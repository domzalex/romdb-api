package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Entry map[string]any

type Dataset struct {
	Version string           `json:"version"`
	Count   int              `json:"count"`
	Entries map[string]Entry `json:"entries"`
}

func main() {
	entries := make(map[string]Entry)

	files, err := filepath.Glob("../dataset/*.json")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		var entry Entry
		if err := json.Unmarshal(data, &entry); err != nil {
			panic(err)
		}

		id, ok := entry["id"].(string)
		if !ok || id == "" {
			panic("missing id in " + file)
		}

		delete(entry, "id")
		entries[id] = entry
	}

	dataset := Dataset{
		Version: time.Now().Format("2006-01-02"),
		Count:   len(entries),
		Entries: entries,
	}

	os.MkdirAll("../generated", 0755)

	out, _ := json.MarshalIndent(dataset, "", "  ")
	os.WriteFile("../generated/dataset.json", out, 0644)
}
