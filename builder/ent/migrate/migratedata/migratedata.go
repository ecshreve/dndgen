package migratedata

import (
	"context"

	"ariga.io/atlas/sql/migrate"
	_ "github.com/mattn/go-sqlite3"
)

var SeedersToRun = []func(ctx context.Context, dir *migrate.LocalDir) error{
	SeedAbilityScoreMigration,
	SeedAlignmentMigration,
	SeedClassMigration,
	SeedLanguageMigration,
	SeedMagicSchoolMigration,
	SeedRaceMigration,
	SeedSkillMigration,
}
