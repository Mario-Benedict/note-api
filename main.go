package main

import (
	"log"
	"os"

	"github.com/Mario-Benedict/note-api/conf"
	"github.com/Mario-Benedict/note-api/db"
	"github.com/Mario-Benedict/note-api/server"
	"github.com/joho/godotenv"
)

func main() {
	loadDotEnv()
	db.Init()

	conf.MigrateDB(db.GetDB())

	server := server.NewServer(":8080")

	log.Printf("Server started on port %s", server.Address)
	log.Fatal(server.Run())

}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}
}
