package main

import (
	"net/http"

	"github.com/IceMAN2377/todo.git/configs"
	"github.com/IceMAN2377/todo.git/internal/auth"
	"github.com/IceMAN2377/todo.git/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()
	_ = db.NewDb(cfg)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: cfg,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}
