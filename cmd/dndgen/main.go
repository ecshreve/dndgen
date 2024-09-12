package main

import (
	"context"
	"net/http"
	"text/template"

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

var tmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <title>dndgen</title>
</head>
<body>
    <h1>pages</h1>
    <a type="button" class="btn btn-primary" href="/playground">playground</a>
    <a type="button" class="btn btn-secondary" href="/viz">viz</a>
</body>
</html>
`

func uiHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("webpage")
	t, _ = t.Parse(tmpl)
	t.Execute(w, nil)
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
	http.HandleFunc("/", uiHandler)
	http.Handle("/playground", playground.Handler("dndgen", "/graphql"))
	http.HandleFunc("/graphql", graphqlHandler(client))
	http.HandleFunc("/viz", ent.ServeEntviz().ServeHTTP)

	log.Info("Starting the server on :8087...")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		log.Fatal(err)
	}
}
