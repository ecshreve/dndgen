package main

import (
	"context"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	log, _ := zap.NewDevelopment()
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	srv := handler.NewDefaultServer(ent.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	http.Handle("/",
		playground.Handler("dndgen", "/query"),
	)
	http.Handle("/query", srv)

	log.Info("listening on", zap.String("address", cli.Addr))
	if err := http.ListenAndServe(cli.Addr, nil); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
