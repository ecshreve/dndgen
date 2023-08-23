package popper

import (
	"context"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

//go:generate go run gen.go

// Popper is a populator for the ent database.
type Popper struct {
	Client   *ent.Client
	IdToIndx map[int]string
	IndxToId map[string]int
}

// NewPopper creates a new Popper.
func NewPopper(ctx context.Context, c *ent.Client) *Popper {
	var client = c
	if c == nil {
		client, _ = ent.Open(
			dialect.SQLite,
			// "file:dev?mode=memory&cache=shared&_fk=1",
			"file:file?mode=memory&_fk=1",
		)
		if client == nil {
			log.Fatalln("unable to open ent client")
		}

		if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
			log.Fatalln(err)
		}
	}

	return &Popper{
		Client:   client,
		IdToIndx: map[int]string{},
		IndxToId: map[string]int{},
	}
}
