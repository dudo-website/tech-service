package main

import (
	"dudo/go_hello_world/graph"
	"dudo/go_hello_world/graph/generated"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// https://gqlgen.com/recipes/cors/
func main() {
	router := chi.NewRouter()
	server_port := os.Getenv("SERVER_PORT")

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/query", srv)

	err := http.ListenAndServe(fmt.Sprintf(":%s", server_port), router)
	if err != nil {
		panic(err)
	}
}
