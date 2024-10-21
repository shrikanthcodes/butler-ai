package chat

import (
	"github.com/redis/go-redis/v9"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/internal/service/cache"
	"time"
)

// GetConvLock returns the mutex for a specific convID, creating it if necessary.
func (cs *CsService) GetConvLock(convID string) *cache.RWMutex {
	if val, ok := cs.CacheService.CsCache.ConvLocks.Load(convID); ok {
		return val.(*cache.RWMutex)
	}
	newLock := &cache.RWMutex{}
	cs.CacheService.CsCache.ConvLocks.Store(convID, newLock)
	return newLock
}

// AddDialogue adds a new dialogue and stores it both in-memory and in Redis.
func (cs *CsService) AddDialogue(convID string, newDialogue entity.Dialogue, newConversation *entity.Conversation) error {
	convLock := cs.GetConvLock(convID)
	convLock.Lock()
	defer convLock.Unlock()

	// Update in-memory active conversations
	cs.CacheService.CsCache.ActiveConversations.Store(convID, newConversation)

	// Update recent dialogues (limit to 10)
	var dialogues []entity.Dialogue
	if val, ok := cs.CacheService.CsCache.RecentDialogues.Load(convID); ok {
		dialogues = val.([]entity.Dialogue)
	}
	dialogues = append(dialogues, newDialogue)
	if len(dialogues) > 10 {
		dialogues = dialogues[len(dialogues)-10:]
	}
	cs.CacheService.CsCache.RecentDialogues.Store(convID, dialogues)

	// Update Redis
	err := cs.StoreInRedis(convID, newConversation, dialogues)
	if err != nil {
		return err
	}

	// Update lastUpdated in-memory
	cs.CacheService.CsCache.LastUpdated.Store(convID, time.Now())

	return nil
}

// StoreInRedis persists both active conversation and recent dialogues in Redis.
func (cs *CsService) StoreInRedis(convID string, conversation *entity.Conversation, dialogues []entity.Dialogue) error {
	// Convert data to JSON
	conversationJSON, err := json.Marshal(conversation)
	if err != nil {
		return err
	}
	dialoguesJSON, err := json.Marshal(dialogues)
	if err != nil {
		return err
	}

	// Store in Redis
	err = cache.redisClient.Set(cache.ctx, convID+":conversation", conversationJSON, 0).Err()
	if err != nil {
		return err
	}
	err = cache.redisClient.Set(cache.ctx, convID+":dialogues", dialoguesJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetDialogues retrieves recent dialogues from either in-memory or Redis.
func (cache *CcService) GetDialogues(convID string) ([]entity.Dialogue, error) {
	convLock := cache.GetConvLock(convID) // Correctly refer to GetConvLock method
	convLock.RLock()
	defer convLock.RUnlock()

	// Check in-memory cache
	if val, ok := cache.CsCache.RecentDialogues.Load(convID); ok {
		return val.([]entity.Dialogue), nil
	}

	// Retrieve from Redis if not in memory
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
	cache.CsCache.RecentDialogues.Store(convID, dialogues)

	return dialogues, nil
}

// GetActiveConversation retrieves the active conversation from in-memory or Redis.
func (cache *CcService) GetActiveConversation(convID string) (*entity.Conversation, error) {
	convLock := cache.GetConvLock(convID) // Correctly refer to GetConvLock method
	convLock.RLock()
	defer convLock.RUnlock()

	// Check in-memory cache
	if val, ok := cache.CsCache.ActiveConversations.Load(convID); ok {
		return val.(*entity.Conversation), nil
	}

	// Retrieve from Redis if not in memory
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
	cache.CsCache.ActiveConversations.Store(convID, &conversation)

	return &conversation, nil
}

// UpdatePrompt updates the prompt in-memory and Redis.
func (cache *CcService) UpdatePrompt(convID string, newPrompt string) error {
	convLock := cache.GetConvLock(convID) // Correctly refer to GetConvLock method
	convLock.Lock()
	defer convLock.Unlock()

	// Update in-memory prompt
	var prompts []string
	if val, ok := cache.CsCache.Prompt.Load(convID); ok {
		prompts = val.([]string)
	}
	prompts = append(prompts, newPrompt)
	cache.CsCache.Prompt.Store(convID, prompts)

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
func (cache *CcService) GetPrompt(convID string) ([]string, error) {
	convLock := cache.GetConvLock(convID) // Correctly refer to GetConvLock method
	convLock.RLock()
	defer convLock.RUnlock()

	// Check in-memory cache
	if val, ok := cache.CsCache.Prompt.Load(convID); ok {
		return val.([]string), nil
	}

	// Retrieve from Redis if not in memory
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
	cache.CsCache.Prompt.Store(convID, prompts)

	return prompts, nil
}

// GetLastUpdated retrieves the last update time for a specific conversation (in-memory only).
func (cache *CcService) GetLastUpdated(convID string) *time.Time {
	if val, ok := cache.CsCache.LastUpdated.Load(convID); ok {
		t := val.(time.Time)
		return &t
	}
	return nil
}
