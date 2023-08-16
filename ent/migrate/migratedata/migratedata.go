//go:build ignore

package migratedata

import (
	"context"
	"fmt"
	"strings"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	"github.com/kr/pretty"
)

func SeedAbilityScores(dir *migrate.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.SQLite, w)))

	// wrk := NewClient()
	q := `{"query":"query RawData {\n  abilityScores {\n    index\n    name\n    full_name\n    desc\n    skills {\n      index\n      name\n      desc\n    }\n  }\n}","variables":{}}`

	type vv struct {
		Data struct {
			AbilityScores []struct {
				Indx     string   `json:"index"`
				Name     string   `json:"name"`
				FullName string   `json:"full_name"`
				Desc     []string `json:"desc"`
				// Skills   []struct {
				// 	Index string `json:"index"`
				// 	Name  string `json:"name"`
				// 	Desc  string `json:"desc"`
				// } `json:"skills"`
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
		"seed_users",
		"Add the initial users to the database.",
	)
}
