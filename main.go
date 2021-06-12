package main

import (
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
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ハンドルファンクションにリダイレクト
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/login", handlers.LogIn)
	mux.HandleFunc("/threads", handlers.Threads)
	mux.HandleFunc("/threads/read", handlers.ReadThreads)
	mux.HandleFunc("/threads/newThreads", handlers.NewThread)
	mux.HandleFunc("/threads/createThreads", handlers.CreateThread)

	// login.htmlにアクセスしたときに/authenticateが呼び出すようになっている
	mux.HandleFunc("/authenticate", handlers.Authenticate)
	mux.HandleFunc("/threads/post", handlers.Post)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
