package main

import (
	"builder/internal/seeder"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// ctx := context.Background()
	// log.Info("Starting seeder")

	// client, err := ent.Open(
	// 	"sqlite3",
	// 	"sqlite://ent?mode=memory&cache=shared&_fk=1",
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := seeder.GenerateSeederFns("internal/seeder", seeder.SeederMetaAll); err != nil {
		log.Fatal(err)
	}

	if err := seeder.GenerateDataMigrationFns("ent/migrate/migratedata", seeder.SeederMetaAll); err != nil {
		log.Fatal(err)
	}

	log.Info("Finished generating seeder files")
}
