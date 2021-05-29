package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var templates *template.Template

	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r) // セッションが存在しなければエラーが返される
		public := []string{"templates/layout.html",
			"templates/index.html",
			"templates/public.navbar.html"}
		private := []string{"templates/layout.html",
			"templates/index.html",
			"templates/private.navbar.html"}

		if err != nil { // セッションが取得できた場合
			templates = template.Must(templates.ParseFiles(public...))
		} else { // セッションが取得できなかった場合
			templates = template.Must(template.ParseFiles(private...))
		}

		templates.ExecuteTemplate(w, "layout", threads)
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

func session(w http.ResponseWriter, r *http.Request) (s data.Session, err error) {
	cookie, err := r.Cookie(("_cookie")) // リクエストからクッキーを取り出す
	if err == nil {
		s = data.Session{Uuid: cookie.Value}
		ok, _ := s.Check()
		if !ok {
			err = errors.New("Invalid session") // セッションが存在しなければエラーを返す
		}
	}

	return
}
