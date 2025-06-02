package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	webServiceSchema "targeting-engine/webService/schema"

	appInit "targeting-engine/init/prometheous"

	"github.com/go-redis/redis/v8"
)

// getCampaignsFromRedis attempts to retrieve matching campaigns from Redis cache.
func (rdb *RedisClient) GetCampaignsFromRedis(ctx context.Context, cacheKey string) ([]webServiceSchema.CampaignResponse, error) {
	cachedData, err := rdb.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Printf("Cache hit for key: %s", cacheKey)
		var matchingCampaigns []webServiceSchema.CampaignResponse
		if err := json.Unmarshal([]byte(cachedData), &matchingCampaigns); err != nil {
			fmt.Printf("Error unmarshaling cached data for key %s: %v", cacheKey, err)
			return nil, err
		}
		appInit.RedisCacheHits.Inc() // Increment cache hit counter
		return matchingCampaigns, nil
	} else if err == redis.Nil { //key not exist
		fmt.Printf("Cache miss for key: %s", cacheKey)
		appInit.RedisCacheMisses.Inc() // Increment cache miss counter
		return nil, redis.Nil
	} else {
		fmt.Printf("Error getting from Redis for key %s: %v", cacheKey, err)
		return nil, err
	}
}

// setCampaignsInRedis stores matching campaigns in Redis cache with a TTL.
func (rdb *RedisClient) SetCampaignsInRedis(ctx context.Context, cacheKey string, campaigns []webServiceSchema.CampaignResponse, ttl time.Duration) {
	jsonData, err := json.Marshal(campaigns)
	if err != nil {
		fmt.Printf("Error marshaling data for cache key %s: %v", cacheKey, err)
		return
	}

	if err := rdb.Client.Set(ctx, cacheKey, jsonData, ttl).Err(); err != nil {
		fmt.Printf("Error setting data in Redis for key %s: %v", cacheKey, err)
	} else {
		fmt.Printf("Successfully cached data for key: %s with TTL %s", cacheKey, ttl)
	}
}
