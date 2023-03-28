package constructors

import portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"

func NewSwitchContextArgs() *portal_api.SwitchContextArgs {
	return &portal_api.SwitchContextArgs{}
}

func NewSwitchContextResponse() *portal_api.SwitchContextResponse {
	return &portal_api.SwitchContextResponse{}
}

func NewForwardPortArgs(localPortNumber uint32, remotePortNumber uint32, protocol *portal_api.TransportProtocol) *portal_api.ForwardPortArgs {
	return &portal_api.ForwardPortArgs{
		LocalPortNumber:  localPortNumber,
		RemotePortNumber: remotePortNumber,
		Protocol:         protocol,
	}
}

func NewForwardPortResponse() *portal_api.ForwardPortResponse {
	return &portal_api.ForwardPortResponse{}
}
