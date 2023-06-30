package demo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeCallEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CallRequest)
		reply, err := svc.Call(ctx, req.Name)
		if err != nil {
			return CallResponse{Reply: reply}, nil
		}
		return CallResponse{Reply: reply}, nil
	}
}
