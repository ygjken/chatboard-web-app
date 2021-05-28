package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/ygjken/chatboard-web-app/data"
)

func main() {
	// マルチプレクサ
	mux := http.NewServeMux()

	// 静的ファイルへのリダイレクト
	// localhost/static/にアクセスした場合
	// ./publicを見に行く
	files := http.FileServer(http.Dir("public"))
	fmt.Println(files)
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ハンドルファンクションにリダイレクト
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html", "templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	thread, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", thread)
	}
}
