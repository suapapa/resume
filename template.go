package main

import (
	"html/template"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
)

func genResume(config *Resume) {
	f, err := os.Create("_deploy/index.html")
	chk(err)
	defer f.Close()

	tFuncMap := template.FuncMap{
		"markdown": func(str string) string {
			md := string(markdown.ToHTML([]byte(str), nil, nil))
			log.Println(str, "->", md)
			return md
		},
	}

	tmpl := template.New("index.tmpl")
	tmpl = tmpl.Funcs(tFuncMap)
	tmpl, err = tmpl.ParseFiles("index.tmpl")
	// tmpl, err := template.ParseFiles("index.tmpl")
	chk(err)
	err = tmpl.Execute(f, config)
	chk(err)
}
