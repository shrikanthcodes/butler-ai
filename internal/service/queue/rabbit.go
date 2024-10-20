package queue

import (
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"github.com/shrikanthcodes/butler-ai/pkg/repository"
)

// QsService is a service for handling caching.
type QsService struct {
	Conn   *repository.ConnPool
	Logger *logger.Logger
}

func NewQueueService(cfg config.RMQ, log *logger.Logger) (*QsService, error) {
	return nil, nil
}

func (qs *QsService) Close() error {
	return nil
}
