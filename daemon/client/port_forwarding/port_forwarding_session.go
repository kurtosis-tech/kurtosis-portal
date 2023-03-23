package port_forwarding

import "io"

type PortForwardingSession interface {
	io.Closer

	GetParams() *PortForwardingParams

	IsRunning() bool

	RunBlocking() error

	RunAsync() error
}
