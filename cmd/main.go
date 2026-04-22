package main

import (
	"context"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
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

}