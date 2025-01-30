package main

import (
	"net/http"

	"github.com/IceMAN2377/todo.git/internal/auth"
)

func main() {
	//_ := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}
