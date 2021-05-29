package main

import (
	"fmt"
	"net/http"

	"github.com/ygjken/chatboard-web-app/handlers"
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
	mux.HandleFunc("/", handlers.Index)
	// mux.HandleFunc("/login", handlers.login)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
