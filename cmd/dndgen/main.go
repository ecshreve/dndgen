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

func main() {
	client, err := ent.Open(
		dialect.SQLite,
		"file:file.db?cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(dndgen.NewSchema(client))

	http.Handle("/",
		playground.Handler("dndgen", "/query"),
	)
	http.Handle("/query", srv)

	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}
