package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"github.com/shrikanthcodes/butler-ai/pkg/repository"
	"sync"
	"time"
)

// CsCache defines the in-memory cache backed by Redis.
type CsCache struct {
	ActiveConversations sync.Map // map[string]*entity.Conversation
	RecentDialogues     sync.Map // map[string][]entity.Dialogue
	Prompt              sync.Map // map[string][]string
	ConvLocks           sync.Map // map[string]*sync.RWMutex
	LastUpdated         sync.Map // map[string]time.Time
}

func initializeCsCache() *CsCache {
	return &CsCache{
		ActiveConversations: sync.Map{},
		RecentDialogues:     sync.Map{},
		Prompt:              sync.Map{},
		ConvLocks:           sync.Map{},
		LastUpdated:         sync.Map{},
	}
}

// UsCache defines the in-memory cache backed by Redis.
type UsCache struct {
	UserBasic      sync.Map // map[string]entity.User
	UserMedical    sync.Map // map[string]entity.Medical
	UserProfile    sync.Map // map[string]entity.Profile
	UserDiet       sync.Map // map[string]entity.Diet
	UserInventory  sync.Map // map[string]entity.Inventory
	UserShopping   sync.Map // map[string]entity.Shopping
	UserGoal       sync.Map // map[string]entity.Goal
	UserMealChoice sync.Map // map[string]entity.MealChoice
	UserLocks      sync.Map // map[string]*sync.RWMutex
	LastUpdated    sync.Map // map[string]time.Time
}

func initializeUsCache() *UsCache {
	return &UsCache{
		UserBasic:      sync.Map{},
		UserMedical:    sync.Map{},
		UserProfile:    sync.Map{},
		UserDiet:       sync.Map{},
		UserInventory:  sync.Map{},
		UserShopping:   sync.Map{},
		UserGoal:       sync.Map{},
		UserMealChoice: sync.Map{},
		UserLocks:      sync.Map{},
		LastUpdated:    sync.Map{},
	}
}

type CcService struct {
	CsCache     *CsCache
	UsCache     *UsCache
	redisClient *repository.ConnPool
	ctx         context.Context
}

// NewCacheService initializes the CsCache with Redis as the backend.
func NewCacheService(conf config.Redis, log logger.Logger) *CcService {
	redisClient, err := repository.New(conf, log)
	if err != nil {
		log.Error("Cache service could not be initialized")
	}
	return &CcService{
		CsCache:     initializeCsCache(),
		UsCache:     initializeUsCache(),
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

// getConvLock returns the mutex for a specific convID, creating it if necessary.
func (cache *CcService) getConvLock(convID string) *sync.RWMutex {
	if val, ok := cache.CsCache.ConvLocks.Load(convID); ok {
		return val.(*sync.RWMutex)
	}
	// If the mutex doesn't exist, create it
	newLock := &sync.RWMutex{}
	cache.CsCache.ConvLocks.Store(convID, newLock)
	return newLock
}

// getUserLock returns the mutex for a specific convID, creating it if necessary.
func (cache *CcService) getUserLock(userID string) *sync.RWMutex {
	if val, ok := cache.UsCache.UserLocks.Load(userID); ok {
		return val.(*sync.RWMutex)
	}
	// If the mutex doesn't exist, create it
	newLock := &sync.RWMutex{}
	cache.UsCache.UserLocks.Store(userID, newLock)
	return newLock
}

// AddDialogue adds a new dialogue and stores it both in-memory and in Redis.
func (cache *CcService) AddDialogue(convID string, newDialogue entity.Dialogue, newConversation *entity.Conversation) error {
	convLock := cache.CsCache.getConvLock(convID)
	convLock.Lock()
	defer convLock.Unlock()

	// Update in-memory active conversations
	cache.activeConversations.Store(convID, newConversation)

	// Update recent dialogues (limit to 10)
	var dialogues []entity.Dialogue
	if val, ok := cache.recentDialogues.Load(convID); ok {
		dialogues = val.([]entity.Dialogue)
		dialogues = append(dialogues, newDialogue)
		if len(dialogues) > 10 {
			dialogues = dialogues[len(dialogues)-10:]
		}
	} else {
		dialogues = []entity.Dialogue{newDialogue}
	}
	cache.recentDialogues.Store(convID, dialogues)

	// Update Redis
	err := cache.storeInRedis(convID, newConversation, dialogues)
	if err != nil {
		return err
	}

	// Update lastUpdated in-memory
	cache.lastUpdated.Store(convID, time.Now())

	return nil
}

// storeInRedis persists both active conversation and recent dialogues in Redis.
func (cache *CsCache) storeInRedis(convID string, conversation *entity.Conversation, dialogues []entity.Dialogue) error {
	// Convert data to JSON
	conversationJSON, err := json.Marshal(conversation)
	if err != nil {
		return err
	}
	dialogueJSON, err := json.Marshal(dialogues)
	if err != nil {
		return err
	}

	// Store in Redis
	err = cache.redisClient.Set(cache.ctx, convID+":conversation", conversationJSON, 0).Err()
	if err != nil {
		return err
	}
	err = cache.redisClient.Set(cache.ctx, convID+":dialogues", dialogueJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetDialogues retrieves recent dialogues from either in-memory or Redis.
func (cache *CsCache) GetDialogues(convID string) ([]entity.Dialogue, error) {
	convLock := cache.getConvLock(convID)
	convLock.RLock()
	defer convLock.RUnlock()

	// Check in-memory
	if val, ok := cache.recentDialogues.Load(convID); ok {
		return val.([]entity.Dialogue), nil
	}

	// Retrieve from Redis
	data, err := cache.redisClient.Get(cache.ctx, convID+":dialogues").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var dialogues []entity.Dialogue
	err = json.Unmarshal([]byte(data), &dialogues)
	if err != nil {
		return nil, err
	}

	// Cache in memory
	cache.recentDialogues.Store(convID, dialogues)

	return dialogues, nil
}

// GetActiveConversation retrieves the active conversation from in-memory or Redis.
func (cache *CsCache) GetActiveConversation(convID string) (*entity.Conversation, error) {
	convLock := cache.getConvLock(convID)
	convLock.RLock()
	defer convLock.RUnlock()

	// Check in-memory
	if val, ok := cache.activeConversations.Load(convID); ok {
		return val.(*entity.Conversation), nil
	}

	// Retrieve from Redis
	data, err := cache.redisClient.Get(cache.ctx, convID+":conversation").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var conversation entity.Conversation
	err = json.Unmarshal([]byte(data), &conversation)
	if err != nil {
		return nil, err
	}

	// Cache in memory
	cache.activeConversations.Store(convID, &conversation)

	return &conversation, nil
}

// UpdatePrompt updates the prompt in-memory and Redis.
func (cache *CsCache) UpdatePrompt(convID string, newPrompt string) error {
	convLock := cache.getConvLock(convID)
	convLock.Lock()
	defer convLock.Unlock()

	// Update in-memory prompt
	var prompts []string
	if val, ok := cache.prompt.Load(convID); ok {
		prompts = val.([]string)
	}
	prompts = append(prompts, newPrompt)
	cache.prompt.Store(convID, prompts)

	// Store in Redis
	promptsJSON, err := json.Marshal(prompts)
	if err != nil {
		return err
	}
	err = cache.redisClient.Set(cache.ctx, convID+":prompts", promptsJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetPrompt retrieves the prompt for a specific conversation from either in-memory or Redis.
func (cache *CsCache) GetPrompt(convID string) ([]string, error) {
	convLock := cache.getConvLock(convID)
	convLock.RLock()
	defer convLock.RUnlock()

	// Check in-memory
	if val, ok := cache.prompt.Load(convID); ok {
		return val.([]string), nil
	}

	// Retrieve from Redis
	data, err := cache.redisClient.Get(cache.ctx, convID+":prompts").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var prompts []string
	err = json.Unmarshal([]byte(data), &prompts)
	if err != nil {
		return nil, err
	}

	// Cache in memory
	cache.prompt.Store(convID, prompts)

	return prompts, nil
}

// GetLastUpdated retrieves the last update time for a specific conversation (in-memory only).
func (cache *CsCache) GetLastUpdated(convID string) *time.Time {
	if val, ok := cache.lastUpdated.Load(convID); ok {
		t := val.(time.Time)
		return &t
	}
	return nil
}
