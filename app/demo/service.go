package demo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"math/rand"
	"time"
	"ultra-hand/pkg/log"
)

type Service interface {
	Call(ctx context.Context, req CallRequest, rpcReplyEndpoint endpoint.Endpoint) (reply string, err error)
}

type service struct {
}

func (s service) Call(ctx context.Context, req CallRequest, rpcReplyEndpoint endpoint.Endpoint) (reply string, err error) {
	rpcResp, err := rpcReplyEndpoint(ctx, rpcReplyRequest{name: req.Name})
	if err != nil {
		return "", err
	}
	return rpcResp.(rpcReplyResponse).reply, nil
}

func NewService() Service {
	return &service{}
}

// 记录入口日志

type logMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw logMiddleware) Call(ctx context.Context, req CallRequest, rpcReplyEndpoint endpoint.Endpoint) (reply string, err error) {
	defer func(begin time.Time) {
		//mw.logger.Info("Test", "randStr", randStr(99999))
		mw.logger.Error("Error")
		mw.logger.Debug("Debug")
		mw.logger.Warn("Warn")
		mw.logger.Info("msg",
			"name", req.Name,
			"output", reply,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	reply, err = mw.next.Call(ctx, req, rpcReplyEndpoint)
	return
}

func NewLogMiddleware(logger log.Logger, svc Service) Service {
	return &logMiddleware{
		logger: logger,
		next:   svc,
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
