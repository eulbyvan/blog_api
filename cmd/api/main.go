package main

import (
	"log"

	"github.com/eulbyvan/blog_api/pkg/database"
)

func main() {
	_, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
}
