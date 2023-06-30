package demo

import (
	"context"
)

type Service interface {
	Call(ctx context.Context, name string) (reply string, err error)
}

type service struct {
}

func (s service) Call(_ context.Context, name string) (reply string, err error) {
	reply = "Hello " + name
	return reply, nil
}

func NewService() Service {
	return &service{}
}
