package handlers

import (
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	// if user.Password == r.PostFormValue("password") {

	// }
}
