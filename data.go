package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func loadData() *Resume {
	f, err := os.Open("_data/_config.yml")
	chk(err)
	defer f.Close()
	var config Resume
	chk(yaml.NewDecoder(f).Decode(&config))

	f, err = os.Open("_data/sections.yml")
	chk(err)
	defer f.Close()

	chk(yaml.NewDecoder(f).Decode(&config.Sections))
	for i := 0; i < len(config.Sections); i++ {
		s := &config.Sections[i]
		log.Println("decoding", s.ID, "...")
		// log.Println("##", s.Name, s.Layout)
		if s.Layout == "block" {
			f, err := os.Open(fmt.Sprintf("_data/%s.yml", s.ID))
			chk(err)
			var blocks Blocks
			chk(yaml.NewDecoder(f).Decode(&blocks))
			s.Data = blocks
			f.Close()
		} else if s.Layout == "list" {
			f, err := os.Open(fmt.Sprintf("_data/%s.yml", s.ID))
			chk(err)
			var list List
			chk(yaml.NewDecoder(f).Decode(&list))
			s.Data = list
			f.Close()
		}
	}

	return &config
}
