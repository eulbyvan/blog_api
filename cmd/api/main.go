package main

import (
	"log"
	"net/http"

	internalHttp "github.com/eulbyvan/blog_api/internal/delivery/http"
	"github.com/eulbyvan/blog_api/internal/repository"
	"github.com/eulbyvan/blog_api/internal/usecase"
	"github.com/eulbyvan/blog_api/pkg/database"
	"github.com/gorilla/mux"
)

// printRoutes logs all registered routes
func printRoutes(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		log.Printf("Endpoint: %s %v\n", path, methods)
		return nil
	})
}

func main() {
	// db conn
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// dependency injection
	postRepo := repository.NewPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepo)

	// set router
	r := mux.NewRouter()

	// register HTTP handlers
	internalHttp.NewPostHandler(r, postUseCase)

	// print available routes
	log.Println("Registered Endpoints:")
	printRoutes(r)

	log.Println("Starting server on port 8080")

	// start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
