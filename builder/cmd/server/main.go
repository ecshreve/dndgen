package main

import (
	"net/http"
	"time"

	"github.com/ecshreve/dndgen/builder/ent"
	generated "github.com/ecshreve/dndgen/builder/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

// Defining the Graphql handler
func graphqlHandler(cc *ent.Client) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewSchema(cc))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		srv.ServeHTTP(w, r)
	}
}

func main() {
	var cli struct {
		Addr  string `name:"address" default:"127.0.0.1:8089" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	client, err := ent.Open("sqlite3", "file:prod.db?_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/",
		playground.Handler("builder", "/query"),
	)
	http.Handle("/query", graphqlHandler(client))

	log.Info("listening on", "address", cli.Addr)
	server := &http.Server{
		Addr:              cli.Addr,
		ReadHeaderTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error("http server terminated", "error", err)
	}
}
