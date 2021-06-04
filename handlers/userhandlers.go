package handlers

import (
	"fmt"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

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
