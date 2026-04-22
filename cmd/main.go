package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/swastiijain24/psp/internals/handlers"
	"github.com/swastiijain24/psp/internals/httpclient"
	repo "github.com/swastiijain24/psp/internals/repositories"
	"github.com/swastiijain24/psp/internals/routes"
	"github.com/swastiijain24/psp/internals/services"
)

func main() {
	err := godotenv.Load()

	ctx := context.Background()
	dsn := os.Getenv("GOOSE_DBSTRING")

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	log.Printf("connected to database")

	r := gin.New()

	repo := repo.New(pool)
	npciClient := httpclient.NewNpciClient(os.Getenv("BASE_URL"))
	transactionService := services.NewTransactionService(repo)
	accountService := services.NewAccountService(npciClient, transactionService)
	accountHandler := handlers.NewAccountHandler(accountService)
	paymentService := services.NewPaymentService(repo, npciClient)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	routes.RegisterAccountRoutes(r, accountHandler)
	routes.RegisterPaymentRoutes(r, paymentHandler)

}