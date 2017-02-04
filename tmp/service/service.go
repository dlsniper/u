package service

import (
	"github.com/dlsniper/u/tmp/logger"
)

type SkyService interface {
	Logger() logger.Logger
	Metrics() metrics.Metrics
	Logger() logger.Logger
	Logger() logger.Logger
}

type Service struct {
	logger logger.Logger
}

func (s *Service) Logger() logger.Logger {
	return s.logger
}

func NewService() *Service {
	return &Service{}
}

