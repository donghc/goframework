package config

import (
	"go.uber.org/zap"
	"goframework/internal/pkg/core"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	Email() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Handler {
	return &handler{logger: logger}
}

func (h *handler) i() {

}
