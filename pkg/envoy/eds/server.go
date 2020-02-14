package eds

import (
	"context"

	xds "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/glog"

	"github.com/deislabs/smc/pkg/catalog"
	"github.com/deislabs/smc/pkg/smi"
)

// NewEDSServer creates a new EDS server
func NewEDSServer(ctx context.Context, catalog catalog.MeshCataloger, meshSpec smi.MeshSpec) *Server {
	glog.Info("[EDS] Create NewEDSServer")
	return &Server{
		ctx:      ctx,
		catalog:  catalog,
		meshSpec: meshSpec,
	}
}

// FetchEndpoints implements envoy.EndpointDiscoveryServiceServer
func (s *Server) FetchEndpoints(context.Context, *xds.DiscoveryRequest) (*xds.DiscoveryResponse, error) {
	panic("NotImplemented")
}

// DeltaEndpoints implements envoy.EndpointDiscoveryServiceServer
func (s *Server) DeltaEndpoints(xds.EndpointDiscoveryService_DeltaEndpointsServer) error {
	panic("NotImplemented")
}
