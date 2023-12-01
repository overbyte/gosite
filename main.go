package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

// create template container and load all templates into it
var tpl *template.Template

type book struct {
	Title	string
	Author	string
}

func (b book) Summarize() string {
	return b.Title + ", " + b.Author
}

// create a FuncMap of functions to pass to the templates
// NOTE it is common to use very short function names, presumably because they
// often get pipelined together like {{.MyString | uc | ft}}
var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"dmy": formatToDMY,
}

// instiate the templates and pass the funcmap up front
// NOTE passing the funcMap after like tpl.Funcs(funcMap) means that the
// template will be instantiated and be looking for any funcs used before they
// are passed so will not work
// this * creates a new template, * adds the func map to the result template
// * runs ParseGlob() on result template
func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func formatToDMY(t time.Time) string {
	// Time.Format() is based on "01/02 03:04:05PM '06 -0700"
	return t.Format("02-01-2006")
}

func main() {
	pageData := struct {
		Books []book
		Time time.Time
		Mol	uint
	} {
		Books: []book{
			{ Title: "Harry Potter", Author: "J K Rowling" },
			{ Title: "The Bible", Author: "JHVH" },
			{ Title: "Neuromancer", Author: "William Gibson"},
		},
		Time: time.Now(),
		Mol: 42,
	}

	// execute to stdout
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", pageData)
	if err != nil {
		log.Fatalln(err)
	}
}
