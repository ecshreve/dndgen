package main

import (
	"context"
	"net/http"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/charmbracelet/log"

	"github.com/ecshreve/dndgen/ent"
	dndgen "github.com/ecshreve/dndgen/gqlserver"
	"github.com/ecshreve/dndgen/internal/popper"

	_ "github.com/hedwigz/entviz"
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
	ctx := context.Background()
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.Info("Starting dndgen/gqlserver...")

	DNDGEN_ENV := "prod" // os.Getenv("DNDGEN_ENV")
	if DNDGEN_ENV == "" {
		DNDGEN_ENV = "dev"
	}
	log.Info("Running in", "environment", DNDGEN_ENV)

	db_url := "file:dev.db?_fk=1"
	if DNDGEN_ENV == "prod" {
		db_url = "file:ent?mode=memory&cache=shared&_fk=1"
	}

	log.Info("Connecting to database...")
	client, err := ent.Open(dialect.SQLite, db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	log.Info("Creating schema...")
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	log.Info("Populating database...")
	p := popper.NewPopper(ctx, client)
	if err := p.PopulateAll(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("Creating http handlers...")
	http.Handle("/", playground.Handler("dndgen", "/graphql"))
	http.HandleFunc("/graphql", graphqlHandler(client))
	http.HandleFunc("/viz", ent.ServeEntviz().ServeHTTP)

	log.Info("Starting the server on :8087...")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		log.Fatal(err)
	}
}
