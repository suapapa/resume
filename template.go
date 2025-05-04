package main

import (
	"html/template"
	"os"

	"github.com/gomarkdown/markdown"
)

func genResume(config *Resume) {
	tFuncMap := template.FuncMap{
		"markdown": func(str string) template.HTML {
			return template.HTML(markdown.ToHTML([]byte(str), nil, nil))
		},
	}

	tmpl := template.New("index.tmpl.html")
	tmpl = tmpl.Funcs(tFuncMap)
	tmpl, err := tmpl.ParseFiles("index.tmpl.html")
	chk(err)

	f, err := os.Create("_deploy/index.html")
	chk(err)
	defer f.Close()

	err = tmpl.Execute(f, config)
	chk(err)
}
