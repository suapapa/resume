package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func decodeYAML(path string, v interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(v); err != nil {
		return fmt.Errorf("failed to decode %s: %w", path, err)
	}
	return nil
}

func loadData() (*Resume, error) {
	var config Resume
	if err := decodeYAML("_data/_config.yml", &config); err != nil {
		return nil, err
	}

	if err := decodeYAML("_data/sections.yml", &config.Sections); err != nil {
		return nil, err
	}

	for i := range config.Sections {
		s := &config.Sections[i]
		log.Printf("decoding section %s (%s)...\n", s.ID, s.Layout)

		path := fmt.Sprintf("_data/%s.yml", s.ID)
		switch s.Layout {
		case "block":
			var blocks Blocks
			if err := decodeYAML(path, &blocks); err != nil {
				return nil, err
			}
			s.Data = &blocks
		case "list":
			var list List
			if err := decodeYAML(path, &list); err != nil {
				return nil, err
			}
			s.Data = &list
		case "desc":
			var desc Desc
			if err := decodeYAML(path, &desc); err != nil {
				return nil, err
			}
			s.Data = &desc
		default:
			log.Printf("warning: unknown layout %s for section %s\n", s.Layout, s.ID)
		}
	}

	return &config, nil
}
