package handlers

import (
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	// TODO: パスワードを暗号化
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
