// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpcdemo

import (
	"github.com/google/wire"
	"ultra-hand/pkg/config"
	"ultra-hand/pkg/log"
)

// Injectors from wire.go:

func InitGrpcDemo(cp config.Path) (GrpcDemo, error) {
	viper, err := config.New(cp)
	if err != nil {
		return GrpcDemo{}, err
	}
	option, err := log.NewLoggerOption(viper)
	if err != nil {
		return GrpcDemo{}, err
	}
	sugaredLogger, err := log.NewZapLogger(option)
	if err != nil {
		return GrpcDemo{}, err
	}
	logger, err := log.New(option, sugaredLogger)
	if err != nil {
		return GrpcDemo{}, err
	}
	grpcDemo := New(logger)
	return grpcDemo, nil
}

// wire.go:

var ProvideSet = wire.NewSet(
	GrpcProvideSet, log.ProvideSet, config.ProvideSet,
)
