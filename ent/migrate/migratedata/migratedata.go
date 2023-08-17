//go:build ignore

package migratedata

import (
	"context"
	"fmt"
	"strings"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/migrate"
	"github.com/kr/pretty"
)

func SeedAbilityScores(dir *atlas.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.SQLite, w)))

	wrk := migrate.NewClient()
	q := `{"query":"query Data {\n  abilityScores {\n    index\n    name\n    full_name\n    desc\n  }\n}","variables":{}}`

	type vv struct {
		Data struct {
			AbilityScores []struct {
				Indx     string   `json:"index"`
				Name     string   `json:"name"`
				FullName string   `json:"full_name"`
				Desc     []string `json:"desc"`
			} `json:"abilityScores"`
		} `json:"data"`
	}
	var v vv
	wrk.MakeGQLQuery(q, &v)
	pretty.Print(v)

	scores := make([]*ent.AbilityScoreCreate, len(v.Data.AbilityScores))
	for i, score := range v.Data.AbilityScores {
		scores[i] = client.AbilityScore.Create().
			SetIndx(score.Indx).
			SetName(score.Name).
			SetFullName(score.FullName).
			SetDesc(strings.Join(score.Desc, "\n"))
	}

	// The statement that generates the INSERT statement.
	err := client.AbilityScore.CreateBulk(
		scores...,
	).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"seed_ability_scores",
		"Add the initial data to the database.",
	)
}
