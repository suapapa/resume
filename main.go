package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.Open("_data/sections.yml")
	chk(err)
	defer f.Close()

	var sections Sections
	chk(yaml.NewDecoder(f).Decode(&sections))
	log.Printf("%+v", sections)

	for _, s := range sections {
		log.Println("##", s.Name, s.Layout)
		if s.Layout == "block" {
			f, err := os.Open(fmt.Sprintf("_data/%s.yml", s.ID))
			chk(err)
			var blocks Blocks
			chk(yaml.NewDecoder(f).Decode(&blocks))
			log.Printf("%+v", blocks)
			defer f.Close()
		} else if s.Layout == "list" {
			f, err := os.Open(fmt.Sprintf("_data/%s.yml", s.ID))
			chk(err)
			var list List
			chk(yaml.NewDecoder(f).Decode(&list))
			log.Printf("%+v", list)
			defer f.Close()
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
