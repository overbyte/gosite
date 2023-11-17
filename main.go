package main

import(
	"log"
	"os"
	"text/template"
)

// create template container and load all templates into it
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	books := map[string]string {
		"Harry Potter": "J K Rowling",
		"The Bible": "God",
		"Neuromancer": "William Gibson",
	}

	// execute to stdout
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", books)
	if err != nil {
		log.Fatalln(err)
	}
}
