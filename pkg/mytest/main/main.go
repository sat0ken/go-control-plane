package main

import (
	"context"
	//cryptotls "crypto/tls"
	"flag"
	"fmt"
	//"io/ioutil"
	"log"
	"net"
	//"os"
	//"time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
    "github.com/envoyproxy/go-control-plane/pkg/server"
    endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
    core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
    discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
)

// MakeEndpoint creates a localhost endpoint on a given port.
func MakeEndpoint(clusterName string, address string, port uint32) *v2.ClusterLoadAssignment {
	return &v2.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  address,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: port,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}



func main() {

    flag.Parse()

    config := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
    srv := server.NewServer(context.Background(), config, nil)

    var endpoint []cache.Resource
    endpoint = append(endpoint, MakeEndpoint("service1", "127.0.0.1", 3000))

    out := cache.NewSnapshot("0", endpoint, nil, nil, nil, nil)
    fmt.Printf("%+v\n", out)

    err := config.SetSnapshot("test-cluster/test-id", out)
	if err != nil {
        log.Fatalf("SnapshotError :%v", err)
	}

    port := ":18000"

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()

    v2.RegisterEndpointDiscoveryServiceServer(s, srv)
    v2.RegisterClusterDiscoveryServiceServer(s, srv)
    v2.RegisterListenerDiscoveryServiceServer(s, srv)
    v2.RegisterRouteDiscoveryServiceServer(s, srv)
    discoverygrpc.RegisterRuntimeDiscoveryServiceServer(s, srv)
    reflection.Register(s)

    log.Printf("management server listening on %s\n", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

    s.GracefulStop()
}

