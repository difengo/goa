// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc GRPC client
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package client

import (
	"context"

	goa "goa.design/goa"
	calcsvc "goa.design/goa/examples/calc/gen/calc"
	calcpb "goa.design/goa/examples/calc/gen/grpc/calc"
	goagrpc "goa.design/goa/grpc"
	grpc "google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli calcpb.CalcClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the calc service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: calcpb.NewCalcClient(cc),
		opts:    opts,
	}
}

// Add calls the "Add" function in calcpb.CalcClient interface.
func (c *Client) Add() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		p, ok := v.(*calcsvc.AddPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("calc", "add", "*calcsvc.AddPayload", v)
		}
		req := NewAddRequest(p)
		resp, err := c.grpccli.Add(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		res := int(resp.Field)
		return res, nil
	}
}
