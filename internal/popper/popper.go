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
	Client     *ent.Client
	Reader     *ent.Client
	IdToIndx   map[int]string
	IndxToId   map[string]int
	Context    *context.Context
	PathPrefix string
}

// NewPopper creates a new Popper.
func NewPopper(ctx context.Context, client *ent.Client) *Popper {
	if client == nil {
		client, _ = ent.Open(dialect.SQLite, "file:dev.db?_fk=1")
	}
	reader, _ := ent.Open(dialect.SQLite, "file:dev.db?_fk=1")

	return &Popper{
		Client:   client,
		IdToIndx: map[int]string{},
		IndxToId: map[string]int{},
		Context:  &ctx,
		Reader:   reader,
	}
}

// NewTestPopper creates a new Popper for testing.
func NewTestPopper(ctx context.Context) *Popper {
	client, err := ent.Open("sqlite3", "file:testdb?mode=memory&_fk=1")
	if err != nil {
		log.Error(err)
		return nil
	}
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	return &Popper{
		Client:     client,
		Reader:     client,
		IdToIndx:   map[int]string{},
		IndxToId:   map[string]int{},
		PathPrefix: "dndgen/data/",
	}
}
