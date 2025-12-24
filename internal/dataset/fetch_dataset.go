package dataset

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Entry map[string]any

type Dataset struct {
	Version string           `json:"version"`
	Count   int              `json:"count"`
	Entries map[string]Entry `json:"entries"`
}

var (
	cache      *Dataset
	cache_time time.Time
	mu         sync.Mutex
)

const dataset_URL = "https://raw.githubusercontent.com/domzalex/romdb-api/main/generated/dataset.json"
const ttl = 10 * time.Minute

func Load() (*Dataset, error) {
	mu.Lock()
	defer mu.Unlock()

	if cache != nil && time.Since(cache_time) < ttl {
		return cache, nil
	}

	resp, err := http.Get(dataset_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var d Dataset
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}

	cache = &d
	cache_time = time.Now()
	return cache, nil
}
