package grpcdemo

import (
	"context"
)

type Service interface {
	Reply(ctx context.Context, name string) (reply string, err error)
}

type service struct {
}

func (s service) Reply(_ context.Context, name string) (reply string, err error) {
	return "Hello " + name, nil
}

func NewService() Service {
	return &service{}
}
