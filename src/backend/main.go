package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	_, err = pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.SetTrustedProxies(nil)
	r.RunTLS(
		os.Getenv("DOMAIN")+":"+os.Getenv("PORT"),
		"./wark-com.crt",
		"./wark-com-privateKey.key",
	)
}
