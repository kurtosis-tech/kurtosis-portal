package constructors

import portal_api "github.com/kurtosis-tech/kurtosis-portal/api/golang/generated"

func NewPortalPing() *portal_api.PortalPing {
	return &portal_api.PortalPing{}
}

func NewPortalPong() *portal_api.PortalPong {
	return &portal_api.PortalPong{}
}
