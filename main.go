package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// create template container and load all templates into it
var tpl *template.Template

type book struct {
	Title	string
	Author	string
}

// create a FuncMap of functions to pass to the templates
var funcMap = template.FuncMap{
	"upperCase": strings.ToUpper,
	"firstThree": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	books := []book {
		{ Title: "Harry Potter", Author: "J K Rowling" },
		{ Title: "The Bible", Author: "JHVH" },
		{ Title: "Neuromancer", Author: "William Gibson"},
	}

	// execute to stdout
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", books)
	if err != nil {
		log.Fatalln(err)
	}
}
