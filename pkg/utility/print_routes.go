package utility

import (
	"log"

	"github.com/gorilla/mux"
)

// printRoutes logs all registered routes
func PrintRoutes(r *mux.Router) {
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
