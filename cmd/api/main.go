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

func main() {
	// db conn
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// dependency injection
	postRepo := repository.NewPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepo)

	// init router
	r := mux.NewRouter()

	// post handler
	internalHttp.NewPostHandler(r, postUseCase)

	// start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
