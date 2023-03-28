package client

import (
	"context"
	portal_constructors "github.com/kurtosis-tech/kurtosis-portal/api/golang/constructors"
	portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/client/port_forwarding"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/client/port_forwarding/chisel"
	contexts_state_store_api "github.com/kurtosis-tech/kurtosis/contexts-config-store/api/golang"
	contexts_state_store_generated "github.com/kurtosis-tech/kurtosis/contexts-config-store/api/golang/generated"
	"github.com/kurtosis-tech/kurtosis/contexts-config-store/store"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	PortalClientGrpcPort = 9731

	DefaultRemoteHost    = ""
	DefaultReverseTunnel = false
)

type KurtosisPortalClient struct {
	sync.RWMutex

	factory port_forwarding.PortForwardingSessionFactory
}

func NewKurtosisClient() *KurtosisPortalClient {
	return &KurtosisPortalClient{
		RWMutex: sync.RWMutex{},
		factory: nil,
	}
}

func (portalClient *KurtosisPortalClient) Ping(ctx context.Context, ping *portal_api.PortalPing) (*portal_api.PortalPong, error) {
	return portal_constructors.NewPortalPong(), nil

}

func (portalClient *KurtosisPortalClient) SwitchContext(_ context.Context, _ *portal_api.SwitchContextArgs) (*portal_api.SwitchContextResponse, error) {
	contextStore := store.GetContextConfigStore()

	portalClient.Lock()
	defer portalClient.Unlock()

	if err := portalClient.closeUnlocked(); err != nil {
		return nil, stacktrace.Propagate(err, "Unable to close current sessions before switching to new context")
	}
	logrus.Infof("Closed all sessions before switching to new context")

	currentContext, err := contextStore.GetCurrentContext()
	if err != nil {
		return nil, stacktrace.Propagate(err, "Unable to load current context")
	}

	return contexts_state_store_api.Visit(currentContext, contexts_state_store_api.KurtosisContextVisitor[portal_api.SwitchContextResponse]{
		VisitLocalOnlyContextV0: func(localOnlyContext *contexts_state_store_generated.LocalOnlyContextV0) (*portal_api.SwitchContextResponse, error) {
			logrus.Infof("Switched to local context '%s'", currentContext.Name)
			newFactory, factoryInitErr := chisel.NewPortForwardSessionFactoryForLocalContext()
			if factoryInitErr != nil {
				return nil, stacktrace.Propagate(factoryInitErr, "Unable to build client to remote portal server")
			}
			portalClient.factory = newFactory
			return portal_constructors.NewSwitchContextResponse(), nil
		},
		VisitRemoteContextV0: func(remoteContext *contexts_state_store_generated.RemoteContextV0) (*portal_api.SwitchContextResponse, error) {
			logrus.Infof("Switched to remote context '%s' running on '%s'",
				currentContext.Name, remoteContext.GetHost())
			var tlsCa []byte
			var tlsKey []byte
			var tlsCert []byte
			if remoteContext.GetTlsConfig() != nil {
				tlsCa = remoteContext.GetTlsConfig().GetCertificateAuthority()
				tlsKey = remoteContext.GetTlsConfig().GetClientKey()
				tlsCert = remoteContext.GetTlsConfig().GetClientCertificate()
			}
			newFactory, factoryInitErr := chisel.NewPortForwardSessionFactory(
				remoteContext.GetHost(),
				remoteContext.GetRemotePortalPort(),
				remoteContext.GetTunnelPort(),
				tlsCa,
				tlsCert,
				tlsKey,
			)
			if factoryInitErr != nil {
				return nil, stacktrace.Propagate(factoryInitErr, "Unable to build client to remote portal server")
			}
			portalClient.factory = newFactory
			return portal_constructors.NewSwitchContextResponse(), nil
		},
	})
}

func (portalClient *KurtosisPortalClient) ForwardPort(_ context.Context, args *portal_api.ForwardPortArgs) (*portal_api.ForwardPortResponse, error) {
	portalClient.RLock()
	defer portalClient.RUnlock()

	if portalClient.factory == nil {
		return nil, stacktrace.NewError("Not connected to any environment at the moment. Connect to an environment first calling the 'ConnectToEnvironment' endpoint")
	}

	localPort := args.LocalPortNumber
	remotePort := args.RemotePortNumber

	// we don't yet use the feature of tunneling to another host. Same for reverse tunneling
	session, err := portalClient.factory.NewSession(port_forwarding.NewPortForwardingParams(localPort, DefaultRemoteHost, remotePort, DefaultReverseTunnel, args.GetProtocol()))
	if err != nil {
		return nil, stacktrace.Propagate(err, "Unable to initiate new session from %d to %d", localPort, remotePort)
	}

	if session.IsRunning() {
		return portal_constructors.NewForwardPortResponse(), nil
	}

	if err = session.RunAsync(); err != nil {
		return nil, stacktrace.Propagate(err, "Error running port forward session")
	}

	return portal_constructors.NewForwardPortResponse(), nil
}

func (portalClient *KurtosisPortalClient) Close() error {
	portalClient.Lock()
	defer portalClient.Unlock()

	errorOccurred := false
	if err := portalClient.closeUnlocked(); err != nil {
		errorOccurred = true
		logrus.Errorf("Unable to close current sessions. Error was: \n%v", err.Error())
	}

	if errorOccurred {
		return stacktrace.NewError("An error occurred stopping Portal server. See ERROR logs above for more details")
	}
	return nil
}

func (portalClient *KurtosisPortalClient) closeUnlocked() error {
	if portalClient.factory == nil {
		// nothing to close
		return nil
	}

	sessionsToClose := portalClient.factory.GetSessions()
	successfullyClosedSessions := 0
	for sessionUuid, session := range sessionsToClose {
		logrus.Debugf("Terminating session '%s'", sessionUuid)
		if err := session.Close(); err != nil {
			logrus.Errorf("Sessions '%s' could not be terminated. It's possible that the port forward is still active",
				sessionUuid)
			continue
		}
		successfullyClosedSessions += 1
	}
	if successfullyClosedSessions != len(sessionsToClose) {
		return stacktrace.NewError("Some sessions could not be terminated. See error logs above for more details")
	}

	portalClient.factory = nil
	return nil
}
