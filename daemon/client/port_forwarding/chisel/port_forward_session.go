package chisel

import (
	"context"
	"github.com/google/uuid"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kurtosis-tech/kurtosis-portal/daemon/client/port_forwarding"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"sync"
)

type PortForwardSession struct {
	sync.Mutex

	sessionParams *port_forwarding.PortForwardingParams
	chiselClient  *chclient.Client

	uuid      uuid.UUID
	isRunning bool

	context    context.Context
	cancelFunc context.CancelFunc
}

func NewPortForwardSession(sessionUuid uuid.UUID, params *port_forwarding.PortForwardingParams, chiselClient *chclient.Client) *PortForwardSession {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &PortForwardSession{
		uuid:          sessionUuid,
		isRunning:     false,
		context:       ctx,
		cancelFunc:    cancelFunc,
		sessionParams: params,
		chiselClient:  chiselClient,
	}
}

func NewLocalNoopForwardSession(sessionUuid uuid.UUID) *PortForwardSession {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &PortForwardSession{
		uuid:         sessionUuid,
		isRunning:    false,
		context:      ctx,
		cancelFunc:   cancelFunc,
		chiselClient: nil,
	}
}

func (session *PortForwardSession) Close() error {
	session.Lock()
	defer session.Unlock()

	session.cancelFunc()

	if session.chiselClient != nil {
		if err := session.chiselClient.Close(); err != nil {
			logrus.Warnf("Error encountered closing port tunneling client: \n%v", err.Error())
		}
	}
	session.isRunning = false
	return nil
}

func (session *PortForwardSession) GetParams() *port_forwarding.PortForwardingParams {
	return session.sessionParams
}

func (session *PortForwardSession) IsRunning() bool {
	return session.isRunning
}

func (session *PortForwardSession) RunBlocking() error {
	session.Lock()
	defer session.Unlock() // deferring an unlock to make sure lock gets released if an error happens

	if session.isRunning {
		return stacktrace.NewError("Session already running: '%s'", session.uuid)
	}
	if err := session.chiselClient.Start(session.context); err != nil {
		return stacktrace.Propagate(err, "Unable to start Chisel client for session: '%s'", session.uuid)
	}
	session.isRunning = true
	session.Unlock() // manually releasing lock here as we don't want to hold it for the entire duration of the wait

	if err := session.chiselClient.Wait(); err != nil {
		return stacktrace.Propagate(err, "Error waiting on Chisel client for session: '%s'", session.uuid)
	}
	return nil
}

func (session *PortForwardSession) RunAsync() error {
	session.Lock()
	defer session.Unlock()

	if session.isRunning {
		return stacktrace.NewError("Session already running: '%s'", session.uuid)
	}

	if err := session.chiselClient.Start(session.context); err != nil {
		return stacktrace.Propagate(err, "Unable to start Chisel client for session: '%s'", session.uuid)
	}
	session.isRunning = true
	return nil
}
