package port_forwarding

import "github.com/google/uuid"

type PortForwardingSessionFactory interface {
	NewSession(params *PortForwardingParams) (PortForwardingSession, error)

	GetSessions() map[uuid.UUID]PortForwardingSession
}
