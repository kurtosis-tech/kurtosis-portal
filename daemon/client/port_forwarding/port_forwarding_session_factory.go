package port_forwarding

import (
	"context"

	"github.com/google/uuid"
	portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"
)

type PortForwardingSessionFactory interface {
	NewSession(params *PortForwardingParams, remoteEndpointType portal_api.RemoteEndpointType) (PortForwardingSession, error)

	GetSessions() map[uuid.UUID]PortForwardingSession

	IsHealthy(ctx context.Context) error
}
