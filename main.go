package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	. "github.com/shurlz/bank-backend/handlers"
	. "github.com/shurlz/bank-backend/models"
)

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	server := &http.Server{
		Handler: HandleRouter(),
		Addr:    ":" + port,
	}

	// run migrations
	InitialMigrations(
		&Accounts{},
		&CreditCards{},
		&TransactionsHistory{},
		&TransferHistory{},
		&Users{},
	)

	fmt.Printf("server running on port : %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
