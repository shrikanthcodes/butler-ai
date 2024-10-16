package repository

import (
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/pkg/cache"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"sync"
	"time"
)

// Repository is an interface for the repository service to `get`, `set`, and `delete`
type Repository interface {
	// Get retrieves a value from the repository.
	Get(key string) (string, error)
	// Set sets a value in the repository.
	Set(key string, value string) error
	// Delete removes a value from the repository.
	Delete(key string) error
}

// RsService is a service for handling caching.
type RsService struct {
	Conn    *cache.ConnPool
	CsCache *CsCache
	Logger  *logger.Logger
}

func NewRepositoryService(cfg config.Redis, log *logger.Logger) (*RsService, error) {
	pool, err := cache.New(cfg, *log)
	if err != nil {
		log.Fatal("Failed to initialize Redis", err)
	}

	csCache, err := NewChatServiceCache()

	return &RsService{pool,
		csCache, log}, nil
}

// CsCache stores active conversations in-memory for batching
type CsCache struct {
	activeConversations map[string]*entity.Conversation
	recentDialogues     map[string]*[]entity.Dialogue
	lastUpdated         map[string]time.Time
	mu                  sync.Mutex
}

func NewChatServiceCache() (*CsCache, error) {
	return &CsCache{
		activeConversations: make(map[string]*entity.Conversation),
		recentDialogues:     make(map[string]*[]entity.Dialogue),
		lastUpdated:         make(map[string]time.Time),
	}, nil
}

// Close closes the Redis connection pool and performs any necessary cleanup.
func (rs *RsService) Close() error {
	if err := rs.Conn.Close(); err != nil {
		rs.Logger.Error("Failed to close Redis connection pool", err)
		return err
	}
	rs.Logger.Info("RsService closed successfully")
	return nil
}
