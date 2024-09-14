package main

import (
	"builder/ent"
	"builder/ent/migrate"
	generated "builder/graph"

	"context"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)

	log.Info("listening on", "address", cli.Addr)
	server := &http.Server{
		Addr:              cli.Addr,
		ReadHeaderTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error("http server terminated", "error", err)
	}
}
