package elasticSearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	redis "targeting-engine/database/redis"
	appInit "targeting-engine/init/prometheous"
	webServiceSchema "targeting-engine/webService/schema"
)

// // Predefined set of campaigns (some active, some inactive for test coverage)
// var Campaigns = []serviceSchema.Campaign{
// 	{
// 		ID:       "spotify",
// 		Name:     "Spotify - Music for everyone",
// 		ImageURL: "https://somelink",
// 		CTA:      "Download",
// 		Status:   "ACTIVE",
// 	},
// 	{
// 		ID:       "duolingo",
// 		Name:     "Duolingo: Best way to learn",
// 		ImageURL: "https://somelink2",
// 		CTA:      "Install",
// 		Status:   "ACTIVE",
// 	},
// 	{
// 		ID:       "subwaysurfer",
// 		Name:     "Subway Surfer",
// 		ImageURL: "https://somelink3",
// 		CTA:      "Play",
// 		Status:   "ACTIVE",
// 	},
// 	{
// 		ID:       "inactive_test",
// 		Name:     "Inactive Campaign",
// 		ImageURL: "https://inactive.link",
// 		CTA:      "Inactive",
// 		Status:   "INACTIVE",
// 	},
// }

// // Targeting rules defining which requests are eligible for which campaigns
// var Rules = map[string]serviceSchema.TargetingRule{
// 	"spotify": {
// 		CampaignID:     "spotify",
// 		IncludeCountry: map[string]bool{"US": true, "Canada": true},
// 	},
// 	"duolingo": {
// 		CampaignID:     "duolingo",
// 		IncludeOS:      map[string]bool{"Android": true, "iOS": true},
// 		ExcludeCountry: map[string]bool{"US": true},
// 	},
// 	"subwaysurfer": {
// 		CampaignID: "subwaysurfer",
// 		IncludeOS:  map[string]bool{"Android": true},
// 		IncludeApp: map[string]bool{"com.gametion.ludokinggame": true},
// 	},
// }

// Example Elasticsearch Document Structure:
// CampaignDocument:{
//   "campaign_id": "spotify",
//   "name": "Spotify - Music for everyone",
//   "image_url": "https://somelink",
//   "cta": "Download",
//   "status": "ACTIVE",
//   "created_at": "2023-01-01T10:00:00Z",
//   "updated_at": "2023-01-01T10:00:00Z",
//   "targeting_rules": [
//     {
//       "dimension": "country",
//       "type": "INCLUDE",
//       "value": "us"
//     },
//     {
//       "dimension": "country",
//       "type": "INCLUDE",
//       "value": "canada"
//     },
//     {
//       "dimension": "os",
//       "type": "EXCLUDE",
//       "value": "web" // Example: Exclude web for Spotify
//     }
//   ]
// }

// This is used for unmarshaling the search results.
type CampaignDocument struct {
	CampaignID string `json:"campaign_id"` // Mapped from "_id" or a specific field in ES
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	CTA        string `json:"cta"`
	Status     string `json:"status"`
	// TargetingRules are not needed here for direct response, but are in the ES document
	// TargetingRules []TargetingRuleDetail `json:"targeting_rules"`
}

// It represents a single hit from an Elasticsearch search result.
type ElasticsearchHit struct {
	Source map[string]interface{} `json:"_source"`
	ID     string                 `json:"_id"`
}

// It represents the overall structure of an Elasticsearch search response.
type ElasticsearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []ElasticsearchHit `json:"hits"`
	} `json:"hits"`
}

// executes an Elasticsearch query based on appID, country, and os
func QueryElasticsearch(esClient *ESClient, appID, country, os string) ([]webServiceSchema.CampaignResponse, error) {
	// Normalize inputs to lowercase for consistent matching with ES index
	appID = strings.ToLower(appID)
	country = strings.ToLower(country)
	os = strings.ToLower(os)

	var matchingCampaigns []webServiceSchema.CampaignResponse

	// Create a cache key based on the query parameters
	cacheKey := fmt.Sprintf("campaigns:%s:%s:%s", appID, country, os)
	ctx := context.Background() // Use a context for Redis operations
	redisClient, err := redis.NewRedisClient()
	if err != nil {
		fmt.Printf("Failed to create connection with redis error: %v", err)
	}
	matchingCampaigns, err = redisClient.GetCampaignsFromRedis(ctx, cacheKey)
	if err == nil {
		return matchingCampaigns, nil // Serve from cache
	} else if err != nil {
		// Log the error but continue to query Elasticsearch if Redis is having issues
		fmt.Printf("Failed to retrieve from cache for key %s, querying Elasticsearch: %v", cacheKey, err)
	} else {
		fmt.Printf("Cache miss for key: %s, querying Elasticsearch.", cacheKey)
	}

	startESQuery := time.Now() // Start timing Elasticsearch query
	//  building the JSON query.
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					// 1. Only consider active campaigns
					{"term": {"status": "ACTIVE"}},
				},
				"must": []map[string]interface{}{
					// 2. Logic for 'app_id' dimension
					buildIncludeLogic("app_id", appID),
					// 3. Logic for 'country' dimension
					buildIncludeLogic("country", country),
					// 4. Logic for 'os' dimension
					buildIncludeLogic("os", os),
				},
				"must_not": []map[string]interface{}{
					// 5. Logic for EXCLUDE rules
					buildExcludeLogic("app_id", appID),
					buildExcludeLogic("country", country),
					buildExcludeLogic("os", os),
				},
			},
		},
		"size": 1000, // Max number of results to return
	}

	searchRawResults, err := esClient.SearchDocuments(context.Background(), "campaigns", query)
	// Observe Elasticsearch query duration
	appInit.EsQueryDuration.Observe(time.Since(startESQuery).Seconds())
	if err != nil {
		return nil, fmt.Errorf("error executing campaign search: %w", err)
	}

	for _, rawHit := range searchRawResults {
		var campaignDoc CampaignDocument
		jsonBytes, _ := json.Marshal(rawHit)
		if err := json.Unmarshal(jsonBytes, &campaignDoc); err != nil {
			fmt.Printf("Warning: Could not unmarshal search hit into CampaignDocument: %v, raw: %+v", err, rawHit)
			continue
		}

		matchingCampaigns = append(matchingCampaigns, webServiceSchema.CampaignResponse{
			CID: campaignDoc.CampaignID,
			Img: campaignDoc.ImageURL,
			CTA: campaignDoc.CTA,
		})
	}

	// Store results in Redis cache for future requests using the new function
	cacheTTL := 15 * time.Minute // Set a TTL (Time-To-Live) for the cache entry, e.g., 5 minutes
	redisClient.SetCampaignsInRedis(ctx, cacheKey, matchingCampaigns, cacheTTL)

	fmt.Printf("Campaign search completed. Total matching campaigns parsed: %d\n", len(matchingCampaigns))
	return matchingCampaigns, nil
}

// It handles both cases: no INCLUDE rule for the dimension, or a matching INCLUDE rule.
func buildIncludeLogic(dimension, value string) map[string]interface{} {
	return map[string]interface{}{
		"bool": map[string]interface{}{
			"should": []map[string]interface{}{
				// Case A: Campaign has NO 'dimension' INCLUDE rule
				{
					"bool": map[string]interface{}{
						"must_not": map[string]interface{}{
							"nested": map[string]interface{}{
								"path": "targeting_rules",
								"query": map[string]interface{}{
									"bool": map[string]interface{}{
										"must": []map[string]interface{}{
											{"term": {"targeting_rules.dimension": dimension}},
											{"term": {"targeting_rules.type": "INCLUDE"}},
										},
									},
								},
							},
						},
					},
				},
				// Case B: Campaign has a 'dimension' INCLUDE rule AND it matches the request value
				{
					"nested": map[string]interface{}{
						"path": "targeting_rules",
						"query": map[string]interface{}{
							"bool": map[string]interface{}{
								"must": []map[string]interface{}{
									{"term": {"targeting_rules.dimension": dimension}},
									{"term": {"targeting_rules.type": "INCLUDE"}},
									{"term": {"targeting_rules.value": value}},
								},
							},
						},
					},
				},
			},
			"minimum_should_match": 1,
		},
	}
}

// It checks if there's an EXCLUDE rule for the dimension that matches the request value.
func buildExcludeLogic(dimension, value string) map[string]interface{} {
	return map[string]interface{}{
		"nested": map[string]interface{}{
			"path": "targeting_rules",
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{"term": {"targeting_rules.type": "EXCLUDE"}},
						{"term": {"targeting_rules.dimension": dimension}},
						{"term": {"targeting_rules.value": value}},
					},
				},
			},
		},
	}
}
