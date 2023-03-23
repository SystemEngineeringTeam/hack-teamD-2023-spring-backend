package main

import (
	"os"

	"github.com/SystemEngineeringTeam/hack-teamD-2023-spring-backend/web"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()
	if os.Getenv("GPT_KEY") == "" {
		panic("GPT_KEY is not set")
	}
	router := web.Router(web.Server{})
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
