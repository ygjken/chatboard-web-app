package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var templates *template.Template

	s, err := session(w, r) // セッションが存在しなければエラーが返される

	public := []string{
		"templates/layout.html",
		"templates/public.index.html",
		"templates/public.navbar.html",
	}
	private := []string{
		"templates/layout.html",
		"templates/private.index.html",
		"templates/private.navbar.html",
	}

	if err != nil { // セッションが取得できた場合
		templates = template.Must(template.ParseFiles(public...))
		templates.ExecuteTemplate(w, "layout", nil)

	} else { // セッションが取得できなかった場合
		templates = template.Must(template.ParseFiles(private...))
		u, _ := s.GetUser()
		templates.ExecuteTemplate(w, "layout", u)
	}

}

func LogIn(w http.ResponseWriter, r *http.Request) {
	t := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/login.html",
	}

	templates := template.Must(template.ParseFiles(t...))
	templates.ExecuteTemplate(w, "layout", nil)

	// メモ, GETから取得してきた情報を取り出す
	// fmt.Println(r.FormValue("email"))
}

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

func session(w http.ResponseWriter, r *http.Request) (s data.Session, err error) {
	cookie, err := r.Cookie(("_cookie")) // リクエストからクッキーを取り出す
	if err == nil {
		s = data.Session{Uuid: cookie.Value}
		ok, _ := s.Check()
		if !ok {
			err = errors.New("invalid session") // セッションが存在しなければエラーを返す
		}
	}

	return
}
