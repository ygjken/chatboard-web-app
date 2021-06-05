package handlers

import (
	"fmt"
	"net/http"

	"github.com/ygjken/chatboard-web-app/data"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")

	t, err := data.GetThreadByUUID(uuid)
	if err != nil {
		fmt.Println("GetPost():", err)
	} else {
		// TODO: fill here
	}
}
