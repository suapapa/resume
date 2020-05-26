package main

import "fmt"

func main() {
	config := loadData()
	genResume(config)

	for _, s := range config.Sections {
		fmt.Printf("%+v\n", s)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
