package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gomarkdown/markdown"
)

func genResume(config *Resume, outputPath string) error {
	tFuncMap := template.FuncMap{
		"markdown": func(str string) template.HTML {
			return template.HTML(markdown.ToHTML([]byte(str), nil, nil))
		},
	}

	tmpl, err := template.New("index.tmpl.html").
		Funcs(tFuncMap).
		ParseFiles("index.tmpl.html")
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, config); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
