package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

// threads の一覧を取得する
func Threads(w http.ResponseWriter, r *http.Request) {
	var files []string

	_, err := session(w, r)
	if err == nil { // sessionが存在しているならば
		files = []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/threads.html",
		}
	} else { // sessionが存在していなければ
		files = []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/threads.html",
		}
	}

	t := template.Must(template.ParseFiles(files...))
	threads, err := data.GetThreads()
	if err != nil {
		fmt.Println("Threads():", err)
	}
	t.ExecuteTemplate(w, "layout", threads)

}

// 選択したThreadのPostの一覧を表示する
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
				"templates/private.posts.html",
			}

			templates := template.Must(template.ParseFiles(files...))
			templates.ExecuteTemplate(w, "layout", &thread) // ポインタを渡さないと行けないので注意
		} else { // sessionが取得できなかったら
			files := []string{
				"templates/layout.html",
				"templates/public.navbar.html",
				"templates/public.posts.html",
			}

			templates := template.Must(template.ParseFiles(files...))
			templates.ExecuteTemplate(w, "layout", &thread) // ポインタを渡さないと行けないので注意
		}
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		fmt.Println("Post():", err)
	} else {
		err = r.ParseForm()
		if err != nil {
			fmt.Println("can't parse form:", err)
		}

		u, err := s.GetUser()
		if err != nil {
			fmt.Println("can't get user from form", err)
		}

		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid") // スレッドのUUID
		t, err := data.GetThreadByUUID(uuid)
		if err != nil {
			fmt.Println("can't get thread by uuid", err)
		}

		_, err = u.CreatePost(t, body)
		if err != nil {
			fmt.Println("can't create post", err)
		}

		url := fmt.Sprint("/threads/read?id=", uuid)
		http.Redirect(w, r, url, http.StatusFound)
	}
}
