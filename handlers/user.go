package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

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

func Authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		fmt.Println(err)
	}
	// TODO: パスワードを暗号化

	// debug
	fmt.Println("email:", r.PostFormValue("email"))
	fmt.Println("pass:", r.PostFormValue("password"))
	// -----
	if user.Password == r.PostFormValue("password") {
		session, _ := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true, // クッキーへのアクセスはHTTPかHTTPSのみ
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
