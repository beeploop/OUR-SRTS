package utils

import (
	"html/template"
)

func HtmlParser(paths ...string) *template.Template {
	args := []string{}

	for _, path := range paths {
		args = append(args, "views/templates/"+path)
	}
	args = append(args, "views/templates/components/head.html")

	html := template.Must(template.ParseFiles(args...))
	return html
}
