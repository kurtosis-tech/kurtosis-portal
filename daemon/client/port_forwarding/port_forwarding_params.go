package port_forwarding

import "fmt"

const (
	remoteDefaultHost = "localhost"
)

type PortForwardingParams struct {
	LocalPortNumber  uint32
	RemoteHost       string
	RemotePortNumber uint32
	Reverse          bool
}

func NewPortForwardingParams(localPortNumber uint32, optionalRemoteHost string, remotePortNumber uint32, reverse bool) *PortForwardingParams {
	var remoteHost string
	if optionalRemoteHost == "" {
		remoteHost = remoteDefaultHost
	} else {
		remoteHost = optionalRemoteHost
	}
	return &PortForwardingParams{
		LocalPortNumber:  localPortNumber,
		RemoteHost:       remoteHost,
		RemotePortNumber: remotePortNumber,
		Reverse:          reverse,
	}
}

func (params *PortForwardingParams) IsDummySession() bool {
	if params.RemoteHost == "" && params.LocalPortNumber == params.RemotePortNumber {
		// This is a case where we need to tunnel a local port to its same port number. This is useless, we don't need
		// to do anything
		return true
	}
	return false
}

func (params *PortForwardingParams) String() string {
	var reversePrefix string
	if params.Reverse {
		reversePrefix = "R:"
	}
	return fmt.Sprintf("%s%d:%s:%d", reversePrefix, params.LocalPortNumber, params.RemoteHost, params.RemotePortNumber)
}

func (params *PortForwardingParams) Equals(otherParams *PortForwardingParams) bool {
	// TODO: we can resolve the remoteHost here to check whether it effectively points to the server server or not.
	//  Seems not worth it for now
	return params.LocalPortNumber == otherParams.LocalPortNumber &&
		params.RemoteHost == otherParams.RemoteHost &&
		params.RemotePortNumber == otherParams.RemotePortNumber &&
		params.Reverse == otherParams.Reverse
}