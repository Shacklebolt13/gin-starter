package config

import (
	"context"
	"sync"
)

type ProcessConfig struct {
	Ctx    context.Context
	Cancel context.CancelCauseFunc
}

var processConfig ProcessConfig

func NewProcessConfig(envConfig *EnvConfig) (*ProcessConfig, error) {
	sync.OnceFunc(func() {
		ctx, cancel := context.WithCancelCause(context.Background())
		processConfig = ProcessConfig{
			Ctx:    ctx,
			Cancel: cancel,
		}
	})()
	return &processConfig, nil
}
