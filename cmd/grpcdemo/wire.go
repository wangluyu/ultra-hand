//go:build wireinject
// +build wireinject

package grpcdemo

import (
	"github.com/google/wire"
	"ultra-hand/pkg/config"
	"ultra-hand/pkg/log"
)

var ProvideSet = wire.NewSet(
	GrpcProvideSet,
	log.ProvideSet,
	config.ProvideSet)

func InitGrpcDemo(cp config.Path) (GrpcDemo, error) {
	wire.Build(ProvideSet)
	return GrpcDemo{}, nil
}
