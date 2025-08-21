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
var processOnce sync.Once

func NewProcessConfig(envConfig *EnvConfig) (*ProcessConfig, error) {
	processOnce.Do(func() {
		ctx, cancel := context.WithCancelCause(context.Background())
		processConfig = ProcessConfig{
			Ctx:    ctx,
			Cancel: cancel,
		}
	})
	return &processConfig, nil
}
