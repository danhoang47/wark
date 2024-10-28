package main

import (
	"log"
	"os"
	appcontext "wark/components/app_context"
	"wark/middlewares"
	"wark/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	appCtx := appcontext.New(db, redisClient, os.Getenv("SECRET_KEY"))

	r := gin.New(func(e *gin.Engine) {
		e.Handlers = append(
			e.Handlers,
			gin.Logger(),
			middlewares.Recovery(),
		)
	})

	v1 := r.Group("/v1")

	routes.ConfigUserRoutes(v1, appCtx)
	routes.ConfigTaskRoutes(v1, appCtx)

	r.SetTrustedProxies(nil)
	r.Run(os.Getenv("DOMAIN") + ":" + os.Getenv("PORT"))
}
