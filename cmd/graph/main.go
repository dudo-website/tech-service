package main

import (
	"dudo/tech_service/graph"
	"dudo/tech_service/graph/generated"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// https://gqlgen.com/recipes/cors/
func main() {
	router := chi.NewRouter()

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Fatal("PORT environment variable not set")
	}

	router.Use(corsConfig().Handler)

	router.Handle("/query", newGraphQLServer())

	log.Println("listening on port", serverPort)
	log.Fatal(http.ListenAndServe(":" + serverPort, router))
}

func corsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
}

func newGraphQLServer() *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
}
