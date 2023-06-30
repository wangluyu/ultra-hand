package grpcdemo

import (
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
)

func makeReplyEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {

}
