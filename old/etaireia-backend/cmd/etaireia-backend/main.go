package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mastrogiovanni/etaireia/etaireia-backend/internal/server"
)

func main() {

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	engine, err := server.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	engine.Run(os.ExpandEnv(":${PORT}"))

}
