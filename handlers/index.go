package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func Threads(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/private.navbar.html",
		"templates/threads.html",
	}

	t := template.Must(template.ParseFiles(files...))
	threads, err := data.GetThreads()
	if err != nil {
		fmt.Println("Threads():", err)
	}
	t.ExecuteTemplate(w, "layout", threads)

}

func ReadThreads(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")

	thread, err := data.GetThreadByUUID(uuid)
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
			templates.ExecuteTemplate(w, "layout", &thread) // ポインタを渡さないと行けないので注意
		}
	}
}
