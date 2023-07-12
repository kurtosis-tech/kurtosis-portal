package port_forwarding

import (
	"context"
	"github.com/google/uuid"
)

type PortForwardingSessionFactory interface {
	NewSession(params *PortForwardingParams) (PortForwardingSession, error)

	GetSessions() map[uuid.UUID]PortForwardingSession

	IsHealthy(ctx context.Context) error
}
