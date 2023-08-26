package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ecshreve/dndgen/ent"
	dndgen "github.com/ecshreve/dndgen/gqlserver"
	_ "github.com/mattn/go-sqlite3"
)

// Defining the Graphql handler
func graphqlHandler(cc *ent.Client) http.HandlerFunc {
	srv := handler.NewDefaultServer(dndgen.NewSchema(cc))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		srv.ServeHTTP(w, r)
	}
}

func main() {
	client, err := ent.Open(
		dialect.SQLite,
		"file:dnd.db?_fk=1",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	http.Handle("/", playground.Handler("dndgen", "/graphql"))
	http.HandleFunc("/graphql", graphqlHandler(client))
	http.HandleFunc("/viz", ent.ServeEntviz().ServeHTTP)

	if err := http.ListenAndServe(":8087", nil); err != nil {
		log.Fatal(err)
	}
}
