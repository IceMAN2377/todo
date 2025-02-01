package main

import (
	"fmt"
	"net/http"

	"github.com/IceMAN2377/todo.git/configs"
	"github.com/IceMAN2377/todo.git/internal/auth"
	"github.com/IceMAN2377/todo.git/internal/link"
	"github.com/IceMAN2377/todo.git/internal/user"
	"github.com/IceMAN2377/todo.git/pkg/db"
	"github.com/IceMAN2377/todo.git/pkg/middleware"
)

func main() {
	cfg := configs.LoadConfig()
	db := db.NewDb(cfg)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	authService := auth.NewAuthService(userRepository)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      cfg,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server is running on port 8080")

	server.ListenAndServe()

}
