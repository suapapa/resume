package main

import (
	"os"
	"testing"
)

func TestLoadData(t *testing.T) {
	config, err := loadData()
	if err != nil {
		t.Fatalf("loadData failed: %v", err)
	}

	if config.Name == "" {
		t.Error("config.Name is empty")
	}

	if len(config.Sections) == 0 {
		t.Error("config.Sections is empty")
	}

	for _, s := range config.Sections {
		if s.Data == nil {
			t.Errorf("section %s has nil data", s.ID)
		}
	}
}

func TestGenResume(t *testing.T) {
	config, err := loadData()
	if err != nil {
		t.Fatalf("loadData failed: %v", err)
	}

	// Create temp output file
	outPath := "_deploy/index_test.html"
	defer os.Remove(outPath)

	err = genResume(config, outPath)
	if err != nil {
		t.Fatalf("genResume failed: %v", err)
	}

	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		t.Errorf("%s was not created", outPath)
	}
}
