package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	. "github.com/shurlz/bank-backend/controllers"
	. "github.com/shurlz/bank-backend/middlewares"
)

func HandleRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	}))

	router.Post("/create-user", CreateUser)
	router.Post("/accounts", ChechAuthStatus(CreateAccount))
	router.Get("/accounts", ChechAuthStatus(GetMyAccounts))
	router.Post("/cards", ChechAuthStatus(CreateCard))
	router.Get("/cards", ChechAuthStatus(GetMyCards))
	router.Post("/transfers/create", ChechAuthStatus(CreateTransfer))
	router.Post("/transfers/get", ChechAuthStatus(GetTransfers))

	return router
}
