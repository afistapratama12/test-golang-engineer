package handlers

import (
	"net/http"
	"oauthGoogleGo/config"
	"oauthGoogleGo/user"
)

var (
	DB             = config.Connect()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)

	userHandler = NewHandler(userService)
)

func New() http.Handler {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	// OauthGoogle
	mux.HandleFunc("/login", userHandler.oauthGoogleLogin)
	mux.HandleFunc("/callback", userHandler.oauthGoogleCallback)

	return mux
}
