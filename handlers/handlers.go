package handlers

import (
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html", "templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}

}
