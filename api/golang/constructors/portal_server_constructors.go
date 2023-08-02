package constructors

import portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"

func NewGetRemoteEndpointsArgs() *portal_api.GetRemoteEndpointsArgs {
	return &portal_api.GetRemoteEndpointsArgs{}
}

func NewGetRemoteEndpointsResponse(remoteEndpointTypes []portal_api.RemoteEndpointType, remoteHost string) *portal_api.GetRemoteEndpointsResponse {
	remoteEndpoints := []*portal_api.RemoteEndpoint{}
	for remoteEndpointType := range remoteEndpointTypes {
		remoteEndpoint := portal_api.RemoteEndpoint{
			RemoteEndpointType: portal_api.RemoteEndpointType(remoteEndpointType),
			RemoteHost:         remoteHost,
		}
		remoteEndpoints = append(remoteEndpoints, &remoteEndpoint)
	}
	return &portal_api.GetRemoteEndpointsResponse{
		RemoteEndpoints: remoteEndpoints,
	}
}
