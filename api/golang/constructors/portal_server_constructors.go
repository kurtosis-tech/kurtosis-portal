package constructors

import portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"

func NewGetRemoteEndpointsArgs() *portal_api.GetRemoteEndpointsArgs {
	return &portal_api.GetRemoteEndpointsArgs{}
}

func NewGetRemoteEndpointsResponse(endpointTypes []portal_api.RemoteEndpointType, remoteHost string) *portal_api.GetRemoteEndpointsResponse {
	endpoints := []*portal_api.RemoteEndpoint{}
	for endpointType := range endpointTypes {
		endpoint := portal_api.RemoteEndpoint{
			RemoteEndpointType: portal_api.RemoteEndpointType(endpointType),
			RemoteHost:         remoteHost,
		}
		endpoints = append(endpoints, &endpoint)
	}
	return &portal_api.GetRemoteEndpointsResponse{
		RemoteEndpoints: endpoints,
	}
}
