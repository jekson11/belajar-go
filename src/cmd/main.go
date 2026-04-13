package main

import (
	"fmt"
	"log"

	"belajar-go/src/config/database"
	"belajar-go/src/config/query"
	"belajar-go/src/handler/rest"
	"belajar-go/src/repository"
	"belajar-go/src/service"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	// Load DB config
	cfg := database.LoadDBConfig()

	// Open PostgreSQL connection
	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	fmt.Println("Connected to PostgreSQL")

	ql, err := query.NewLoadQuery("etc/query/user.sql")
	if err != nil {
		log.Fatalf("failed to load query: %v", err)
	}

	userRepo := repository.InitRepository(db, ql)
	userService := service.InitService(userRepo)
	appCfg, err := LoadAppConfig()

	if err != nil {
		log.Fatalf("failed to load app config: %v", err)
	}

	rest.InitRestHandler(userService, appCfg.Port)

	fmt.Println("Server running on : ", appCfg.Port)
}
