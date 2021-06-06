package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func ReadThreads(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")

	thread, err := data.GetThreadByUUID(uuid)
	fmt.Println(thread.GetPosts())
	if err != nil {
		fmt.Println("GetPost():", err)
	} else {
		_, err := session(w, r)

		if err == nil { // session が正しく取得できたら
			files := []string{
				"templates/layout.html",
				"templates/private.navbar.html",
				"templates/posts.html",
			}

			templates := template.Must(template.ParseFiles(files...))
			templates.ExecuteTemplate(w, "layout", thread)
		}
	}
}
