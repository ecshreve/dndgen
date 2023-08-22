package popper

import (
	"context"
	"os"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

//go:generate go run gen.go

// Popper is a populator for the ent database.
type Popper struct {
	Client *ent.Client
}

// NewPopper creates a new Popper.
func NewPopper(ctx context.Context) *Popper {
	client, err := ent.Open(
		dialect.SQLite,
		os.Getenv("DNDDB"),
		// "file:dev?mode=memory&cache=shared&_fk=1",
		// "file:file.db?_fk=1",
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatalln(err)
	}

	return &Popper{
		Client: client,
	}
}
