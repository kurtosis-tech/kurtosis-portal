package server

import (
	"context"
	chserver "github.com/jpillora/chisel/server"
	"github.com/kurtosis-tech/kurtosis-cloud/portal/api/golang/constructors"
	api "github.com/kurtosis-tech/kurtosis-cloud/portal/api/golang/generated"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

const (
	PortalServerGrpcPort            = 9720
	PortalServerTunnelListeningHost = "0.0.0.0"
	PortalServerTunnelPort          = 9721
)

type KurtosisPortalServer struct {
	tunnelMutex *sync.Mutex

	tlsCaFilePath         string
	tlsServerKeyFilePath  string
	tlsServerCertFilePath string

	killTunnelFunc func() error
}

func NewKurtosisPortalServer(tlsCaFilePath string, tlsServerKeyFilePath string, tlsServerCertFilePath string) *KurtosisPortalServer {
	return &KurtosisPortalServer{
		tunnelMutex:           &sync.Mutex{},
		tlsCaFilePath:         tlsCaFilePath,
		tlsServerKeyFilePath:  tlsServerKeyFilePath,
		tlsServerCertFilePath: tlsServerCertFilePath,
		killTunnelFunc:        nil,
	}
}

func (portalServer *KurtosisPortalServer) Ping(ctx context.Context, args *api.PortalPing) (*api.PortalPong, error) {
	return constructors.NewPortalPong(), nil
}

func (portalServer *KurtosisPortalServer) StartTunnelServer(ctx context.Context, host string, listeningPort uint32) error {
	portalServer.tunnelMutex.Lock()
	defer portalServer.tunnelMutex.Unlock()

	if portalServer.killTunnelFunc != nil {
		logrus.Warn("Trying to start a server-side tunnel while one seem to already be running. Killing the current one first")
		if err := portalServer.killTunnelFunc(); err != nil {
			logrus.Errorf("An error occurred trying to kill the current tunnel. This might prevent the new tunnel from starting. Error was: \n%v", err.Error())
		}
	}

	server, err := chserver.NewServer(&chserver.Config{
		KeySeed:   "",
		AuthFile:  "",
		Auth:      "",
		Proxy:     "",
		Socks5:    false,
		Reverse:   false, // reverse tunnelling is not exposed through the API yet, turn it off here
		KeepAlive: 25 * time.Second,
		TLS: chserver.TLSConfig{
			CA:   portalServer.tlsCaFilePath,
			Cert: portalServer.tlsServerCertFilePath,
			Key:  portalServer.tlsServerKeyFilePath,
		},
	})
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating chisel server")
	}

	chiselStartedSuccessfully := false
	cancellableContext, cancelFunc := context.WithCancel(ctx)
	defer func() {
		if chiselStartedSuccessfully {
			return
		}
		cancelFunc()
	}()

	listeningPortStr := strconv.Itoa(int(listeningPort))
	if err := server.StartContext(cancellableContext, host, listeningPortStr); err != nil {
		return stacktrace.Propagate(err, "error running chisel server")
	}
	portalServer.killTunnelFunc = func() error {
		cancelFunc() // cancelling the context will stop Chisel
		return nil
	}
	chiselStartedSuccessfully = true
	return nil
}

func (portalServer *KurtosisPortalServer) Close() error {
	portalServer.tunnelMutex.Lock()
	defer portalServer.tunnelMutex.Unlock()

	errorOccurred := false
	if portalServer.killTunnelFunc == nil {
		logrus.Debug("Tunnel not running, nothing to stop")
	} else {
		if err := portalServer.killTunnelFunc(); err != nil {
			logrus.Errorf("Unable to stop tunnel. It might still be running in the background. Error was: \n%v", err.Error())
			errorOccurred = true
		}
	}
	if errorOccurred {
		return stacktrace.NewError("An error occurred stopping Portal server. See ERROR logs above for more details")
	}
	return nil
}
