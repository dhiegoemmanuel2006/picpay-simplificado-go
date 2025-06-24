package main

import (
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/controller"
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := godotenv.Load("picpay.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.GetDatabase()

	api := controller.NewApi(db)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      api.NewHandler(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server")
	}
	defer database.CloseDatabase(db)
}
