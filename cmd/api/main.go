package main

import (
	"log"
	"os"

	"lemodie_api_v1/internal/db"
	"lemodie_api_v1/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	envFile := ".env." + appEnv
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️  %s не знайдено, пробуємо .env", envFile)
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  .env теж не знайдено")
		}
	}

	log.Printf("🚀 Середовище: %s", appEnv)

	database := db.Connect(os.Getenv("DB_DSN"))
	defer database.Close()

	r := router.New(database)
	r.Run(":" + os.Getenv("PORT"))
}
