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
	for i := 0; i < len(sections); i++ {
		s := &sections[i]
		log.Println("##", s.Name, s.Layout)
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

	log.Printf("%+v", sections)
	for _, s := range sections {
		switch s.Data.(type) {
		case Blocks:
			log.Println("block", s.Name)
		case List:
			log.Println("list", s.Name)
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
