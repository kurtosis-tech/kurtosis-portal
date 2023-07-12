package chisel

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	chclient "github.com/jpillora/chisel/client"
	portal_constructors "github.com/kurtosis-tech/kurtosis-portal/api/golang/constructors"
	portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/arguments"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/client/port_forwarding"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/server"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"path"
	"sync"
	"time"
)

const (
	locallyRunningServerName = "localhost"

	httpScheme  = "http"
	httpsScheme = "https"

	pingMaxRetries = 5
	pingRetryDelay = 2 * time.Second

	tempDirNamePattern = "kurtosis_backend_tls_*"
	caFileName         = "ca.pem"
	certFileName       = "cert.pem"
	keyFileName        = "key.pem"
	tlsFilesPerm       = 0644
)

type PortForwardSessionFactory struct {
	sync.RWMutex

	isContextLocal bool

	chiselHost string
	chiselPort uint32

	tlsCa         []byte
	tlsClientKey  []byte
	tlsClientCert []byte

	currentSessions map[uuid.UUID]port_forwarding.PortForwardingSession

	portalServerClient portal_api.KurtosisPortalServerClient
}

// NewPortForwardSessionFactory creates a new port tunnelling sessions factory.
func NewPortForwardSessionFactory(portalHost string, portalGrpcPort uint32, portalChiselPort uint32, tlsCa []byte, tlsClientCert []byte, tlsClientKey []byte) (*PortForwardSessionFactory, error) {
	serverClient, err := portalServerClient(portalHost, portalGrpcPort, tlsCa, tlsClientCert, tlsClientKey)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Unable to connect to remote portal server")
	}
	return &PortForwardSessionFactory{
		RWMutex:            sync.RWMutex{},
		isContextLocal:     false,
		tlsCa:              tlsCa,
		tlsClientKey:       tlsClientKey,
		tlsClientCert:      tlsClientCert,
		chiselHost:         portalHost,
		chiselPort:         portalChiselPort,
		currentSessions:    map[uuid.UUID]port_forwarding.PortForwardingSession{},
		portalServerClient: serverClient,
	}, nil
}

// NewPortForwardSessionFactoryForLocalContext creates a new port tunnelling sessions factory.
func NewPortForwardSessionFactoryForLocalContext() (*PortForwardSessionFactory, error) {
	serverClient, err := portalServerClient(locallyRunningServerName, server.PortalServerGrpcPort, nil, nil, nil)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Unable to connect to remote portal server")
	}
	return &PortForwardSessionFactory{
		RWMutex:            sync.RWMutex{},
		isContextLocal:     true,
		tlsCa:              nil,
		tlsClientKey:       nil,
		tlsClientCert:      nil,
		chiselHost:         locallyRunningServerName,
		chiselPort:         server.PortalServerTunnelPort,
		currentSessions:    map[uuid.UUID]port_forwarding.PortForwardingSession{},
		portalServerClient: serverClient,
	}, nil
}

func (factory *PortForwardSessionFactory) NewSession(params *port_forwarding.PortForwardingParams) (port_forwarding.PortForwardingSession, error) {
	factory.Lock()
	defer factory.Unlock()

	if found, existingSessionsUuid := factory.getSimilarExistingSessionsIfAny(params); found {
		logrus.Debugf("Sessions with same params '%s' already exists with UUID '%s'. Returning it.", params.String(), existingSessionsUuid)
		return factory.currentSessions[existingSessionsUuid], nil
	}

	newSessionUuid := uuid.New()
	logrus.Infof("Creating new port forward session '%s' with UUID: '%s'", params.String(), newSessionUuid)

	if factory.isContextLocal && params.IsIdentify() {
		newSession := NewLocalNoopForwardSession(newSessionUuid, params)
		factory.currentSessions[newSessionUuid] = newSession
		return newSession, nil
	}

	tlsFileAbsDirPath, cleanupFunc, err := writeTlsConfigToTempDir(factory.tlsCa, factory.tlsClientCert, factory.tlsClientKey)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Unable to persist TLS information to temporary files")
	}
	defer cleanupFunc()

	var serverUrl string
	var serverName string
	if factory.isContextLocal {
		// local servers do not use HTTPS right now
		serverUrl = fmt.Sprintf("%s://%s:%d", httpScheme, factory.chiselHost, factory.chiselPort)
		serverName = locallyRunningServerName
	} else {
		serverUrl = fmt.Sprintf("%s://%s:%d", httpsScheme, factory.chiselHost, factory.chiselPort)
		serverName = factory.chiselHost
	}

	tunnelString := params.String()

	chiselClientConfig := &chclient.Config{
		Fingerprint:      "",
		Auth:             "",
		KeepAlive:        25 * time.Second,
		MaxRetryCount:    5,
		MaxRetryInterval: 1 * time.Second,
		Server:           serverUrl,
		Proxy:            "",
		Remotes: []string{
			tunnelString,
		},
		Headers: nil,
		TLS: chclient.TLSConfig{
			SkipVerify: false,
			CA:         path.Join(tlsFileAbsDirPath, caFileName),
			Key:        path.Join(tlsFileAbsDirPath, keyFileName),
			Cert:       path.Join(tlsFileAbsDirPath, certFileName),
			ServerName: serverName,
		},
		DialContext: nil,
		Verbose:     false,
	}

	chiselClient, err := chclient.NewClient(chiselClientConfig)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Error creating new Chisel client")
	}

	newSession := NewPortForwardSession(newSessionUuid, params, chiselClient)
	factory.currentSessions[newSessionUuid] = newSession
	return newSession, nil
}

func (factory *PortForwardSessionFactory) GetSessions() map[uuid.UUID]port_forwarding.PortForwardingSession {
	factory.RLock()
	defer factory.RUnlock()

	return factory.currentSessions
}

func (factory *PortForwardSessionFactory) IsHealthy(ctx context.Context) error {
	if _, err := factory.portalServerClient.Ping(ctx, portal_constructors.NewPortalPing()); err != nil {
		return stacktrace.Propagate(err, "Unable to communicate with Portal Server.")
	}
	return nil
}

func (factory *PortForwardSessionFactory) getSimilarExistingSessionsIfAny(params *port_forwarding.PortForwardingParams) (bool, uuid.UUID) {
	for sessionUuid, otherSession := range factory.currentSessions {
		if otherSession.GetParams().Equals(params) {
			return true, sessionUuid
		}
	}
	return false, uuid.New()
}

func portalServerClient(portalServerHost string, portalServerGrpcPort uint32, tlsCa []byte, tlsCert []byte, tlsKey []byte) (portal_api.KurtosisPortalServerClient, error) {
	url := fmt.Sprintf("%s:%d", portalServerHost, portalServerGrpcPort)

	var err error
	var tlsCredentials credentials.TransportCredentials
	if tlsCert == nil && tlsKey == nil {
		tlsCredentials = insecure.NewCredentials()
	} else {
		tlsCredentials, err = arguments.BuildTlsCredentials(tlsCa, tlsCert, tlsKey)
		if err != nil {
			return nil, stacktrace.Propagate(err, "Error building TLS credentials from CA, cert and key")
		}
	}

	clientConn, err := grpc.Dial(url, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred connecting to the Kurtosis Portal Server on host machine URL '%v'", url)
	}
	client := portal_api.NewKurtosisPortalServerClient(clientConn)

	if err = pingServerWithRetries(client); err != nil {
		return nil, stacktrace.Propagate(err, "Portal server for this context unreachable at '%s'", url)
	}
	return client, nil
}

func pingServerWithRetries(serverClient portal_api.KurtosisPortalServerClient) error {
	var err error
	ticker := time.NewTicker(5 * time.Second)
	for i := 1; i <= pingMaxRetries; i++ {
		_, err = serverClient.Ping(context.Background(), portal_constructors.NewPortalPing())
		if err != nil {
			logrus.Debugf("Error reaching Portal Server for this context (retries %d/%d). Will retry in %v", i, pingMaxRetries, pingRetryDelay)
			<-ticker.C
		} else {
			if i == 1 {
				logrus.Debugf("Succeeded reaching Portal Server")
			} else {
				logrus.Debugf("Finally succeeded reaching Portal Server after %d retries", i)
			}
			return nil
		}
	}
	return stacktrace.Propagate(err, "Unable to reach Portal Server for this context after %d retries", pingMaxRetries)
}

// writeTlsConfigToTempDir writes the different TLS files to a directory, and returns the path to this directory.
// It also returns a function to manually delete those files once they've been used upstream
func writeTlsConfigToTempDir(ca []byte, cert []byte, key []byte) (string, func(), error) {
	tempDirectory, err := os.MkdirTemp("", tempDirNamePattern)
	if err != nil {
		return "", nil, stacktrace.Propagate(err, "Cannot create a temporary directory to store Kurtosis backend TLS files")
	}
	caAbsFileName := path.Join(tempDirectory, caFileName)
	if err = os.WriteFile(caAbsFileName, ca, tlsFilesPerm); err != nil {
		return "", nil, stacktrace.Propagate(err, "Error writing content of CA to temporary file at '%s'", caAbsFileName)
	}
	certAbsFileName := path.Join(tempDirectory, certFileName)
	if err = os.WriteFile(certAbsFileName, cert, tlsFilesPerm); err != nil {
		return "", nil, stacktrace.Propagate(err, "Error writing content of certificate to temporary file at '%s'", certAbsFileName)
	}
	keyAbsFileName := path.Join(tempDirectory, keyFileName)
	if err = os.WriteFile(keyAbsFileName, key, tlsFilesPerm); err != nil {
		return "", nil, stacktrace.Propagate(err, "Error writing content of key to temporary file at '%s'", keyAbsFileName)
	}

	cleanDirectoryFunc := func() {
		if err = os.RemoveAll(tempDirectory); err != nil {
			logrus.Warnf("Error removing TLS config directory at '%s'. Will remain in the OS temporary files folder until the OS removes it", tempDirectory)
		}
	}
	return tempDirectory, cleanDirectoryFunc, nil
}
