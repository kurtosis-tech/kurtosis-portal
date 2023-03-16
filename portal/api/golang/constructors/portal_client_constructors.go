package constructors

import generated "github.com/kurtosis-tech/kurtosis-cloud/portal/api/golang/generated"

func NewSwitchContextArgs() *generated.SwitchContextArgs {
	return &generated.SwitchContextArgs{}
}

func NewSwitchContextResponse() *generated.SwitchContextResponse {
	return &generated.SwitchContextResponse{}
}

func NewForwardPortArgs(localPortNumber uint32, remotePortNumber uint32) *generated.ForwardPortArgs {
	return &generated.ForwardPortArgs{
		LocalPortNumber:  localPortNumber,
		RemotePortNumber: remotePortNumber,
	}
}

func NewForwardPortResponse() *generated.ForwardPortResponse {
	return &generated.ForwardPortResponse{}
}
