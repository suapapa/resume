package main

import (
	"html/template"
	"log"
	"os"
)

func genResume(config *Resume) {
	f, err := os.Create("index.html")
	chk(err)
	defer f.Close()

	tFuncMap := map[string]interface{}{
		"List": func(i interface{}) []List {
			log.Println("here?")
			return i.([]List)
		},
	}

	tmpl, err := template.ParseFiles("index.tmpl")
	tmpl = tmpl.Funcs(tFuncMap)
	chk(err)
	err = tmpl.Execute(f, config)
	chk(err)
}
