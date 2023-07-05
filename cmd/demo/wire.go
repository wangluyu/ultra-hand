//go:build wireinject
// +build wireinject

package demo

import (
	"github.com/google/wire"
	"ultra-hand/pkg/config"
	"ultra-hand/pkg/log"
)

var ProvideSet = wire.NewSet(
	DemoProvideSet,
	log.ProvideSet,
	config.ProvideSet)

func InitGrpcDemo(cp config.Path) (Demo, error) {
	wire.Build(ProvideSet)
	return Demo{}, nil
}
