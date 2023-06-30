package demo

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"net/http"
	"ultra-hand/service/grpcdemo/pb"
)

// 网络传输相关的，包括协议（HTTP、gRPC、thrift...）等

// HTTP JSON
// decode
// 请求来了之后根据 协议(HTTP、HTTP2)和编码(JSON、pb、thrift) 去解析数据

func decodeCallRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CallRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCallResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPServer(svc Service, logger log.Logger, rpcConn *grpc.ClientConn) http.Handler {
	call := makeCallEndpoint(svc, rpcConn)
	call = loggingMiddleware(log.With(logger, "method", "call"))(call)
	callHandler := httptransport.NewServer(
		call,
		decodeCallRequest,
		encodeCallResponse,
	)
	r := mux.NewRouter()
	r.Handle("/call", callHandler).Methods("POST")
	return r
}

// rpc

// 将内部使用的数据编码为proto
// 对外发起gRPC请求
func encodeReplyRequest(_ context.Context, in interface{}) (request interface{}, err error) {
	req := in.(rpcReplyRequest)
	return &pb.ReplyRequest{Name: req.name}, nil
}

// 解析pb消息, 转为 next 的请求
func decodeReplyResponse(_ context.Context, in interface{}) (response interface{}, err error) {
	resp := in.(*pb.ReplyResponse)
	return rpcReplyResponse{reply: resp.Reply}, nil
}
