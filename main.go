package main

import(
	"log"
	"os"
	"text/template"
)

// create template container and load all templates into it
var tpl *template.Template

type book struct {
	Title	string
	Author	string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	books := []book {
		book{ Title: "Harry Potter", Author: "J K Rowling" },
		book{ Title: "The Bible", Author: "God" },
		book{ Title: "Neuromancer", Author: "William Gibson"},
	}

	// execute to stdout
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", books)
	if err != nil {
		log.Fatalln(err)
	}
}
