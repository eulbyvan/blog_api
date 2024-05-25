package main

import (
	"log"
	"net/http"

	internalHttp "github.com/eulbyvan/blog_api/internal/delivery/http"
	"github.com/eulbyvan/blog_api/internal/repository"
	"github.com/eulbyvan/blog_api/internal/usecase"
	"github.com/eulbyvan/blog_api/pkg/database"
	"github.com/eulbyvan/blog_api/pkg/utility"
	"github.com/gorilla/mux"
)

func main() {
	// db conn
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// dependency injection for posts
	postRepo := repository.NewPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepo)

	// dependency injection for tags
	tagRepo := repository.NewTagRepository(db)
	tagUseCase := usecase.NewTagUseCase(tagRepo)

	// set router
	r := mux.NewRouter()

	// register HTTP handlers
	internalHttp.NewPostHandler(r, postUseCase)
	internalHttp.NewTagHandler(r, tagUseCase)

	// print available routes
	log.Println("Registered Endpoints:")
	utility.PrintRoutes(r)

	log.Println("Starting server on port 8080")

	// start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
