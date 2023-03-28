package port_forwarding

import (
	"fmt"
	portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"
)

const (
	remoteDefaultHost = "localhost"
)

type PortForwardingParams struct {
	LocalPortNumber  uint32
	RemoteHost       string
	RemotePortNumber uint32
	Reverse          bool
	Protocol         portal_api.TransportProtocol
}

func NewPortForwardingParams(localPortNumber uint32, optionalRemoteHost string, remotePortNumber uint32, reverse bool, protocol portal_api.TransportProtocol) *PortForwardingParams {
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
		Protocol:         protocol,
	}
}

func (params *PortForwardingParams) IsIdentify() bool {
	if params.RemoteHost == remoteDefaultHost && params.LocalPortNumber == params.RemotePortNumber {
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
	var protocolSuffix string
	if params.Protocol == portal_api.TransportProtocol_UDP {
		protocolSuffix = "/udp"
	}
	return fmt.Sprintf("%s%d:%s:%d%s", reversePrefix, params.LocalPortNumber, params.RemoteHost, params.RemotePortNumber, protocolSuffix)
}

func (params *PortForwardingParams) Equals(otherParams *PortForwardingParams) bool {
	// TODO: we can resolve the remoteHost here to check whether it effectively points to the server server or not.
	//  Seems not worth it for now
	return params.LocalPortNumber == otherParams.LocalPortNumber &&
		params.RemoteHost == otherParams.RemoteHost &&
		params.RemotePortNumber == otherParams.RemotePortNumber &&
		params.Reverse == otherParams.Reverse
}
