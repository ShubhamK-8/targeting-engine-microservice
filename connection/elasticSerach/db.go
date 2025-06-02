package elasticSearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

// ESClient holds the Elasticsearch client
type ESClient struct {
	Client *elasticsearch.Client
}

// NewElasticsearchClient initializes and returns a new Elasticsearch client.
func NewElasticsearchClient(addresses []string) (*ESClient, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating Elasticsearch client: %w", err)
	}

	// Test the connection
	res, err := es.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting Elasticsearch client info: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error connecting to Elasticsearch: %s", res.Status())
	}

	fmt.Printf("Successfully connected to Elasticsearch. Cluster Info: %s", res.Status())
	return &ESClient{Client: es}, nil
}

// SearchDocuments performs a search query on an Elasticsearch index.
func (esc *ESClient) SearchDocuments(ctx context.Context, indexName string, query map[string]interface{}) ([]map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	res, err := esc.Client.Search(
		esc.Client.Search.WithContext(ctx),
		esc.Client.Search.WithIndex(indexName),
		esc.Client.Search.WithBody(&buf),
		esc.Client.Search.WithTrackTotalHits(true),
		esc.Client.Search.WithPretty(), // For readable output in logs (optional)
	)
	if err != nil {
		return nil, fmt.Errorf("error performing search: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("error parsing the response body: %w", err)
		}
		return nil, fmt.Errorf("error searching documents %s: %s", res.Status(), e["error"])
	}

	var r struct {
		Hits struct {
			Hits []struct {
				Source map[string]interface{} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	var results []map[string]interface{}
	for _, hit := range r.Hits.Hits {
		results = append(results, hit.Source)
	}

	fmt.Printf("Search successful. Found %d documents.", len(results))
	return results, nil
}
