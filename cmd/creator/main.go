package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/race"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetCallerFormatter(func(file string, line int, _ string) string {
		return fmt.Sprintf("%s:%d", strings.Replace(file, "/home/eric/repos/dndgen/", "", 1), line)
	})
	log.SetPrefix("🛠️")
	log.SetTimeFormat("15:04:05")
	log.Info("Starting dndgen/creator...")

	ctx := context.Background()

	dbname := "dev.db"
	log.Info("Using database", "dbname", dbname)

	dburl := "file:" + dbname + "?_fk=1"
	client, err := ent.Open(dialect.SQLite, dburl)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Creating schema...")
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	// CreateCharacter(ctx, client)

	CreateCharacterFromJSON(ctx, client, zekeJSON)
	log.Info("Character created")
}

var zekeJSON = `
{
	"name": "Zeke",
	"race": "human",
	"class": "rogue",
	"age": 30,
	"level": 5,
	"alignment": "chaotic-neutral",
	"ability_scores": {
		"str": 8,
		"dex": 10,
		"con": 12,
		"int": 14,
		"wis": 16,
		"cha": 18
	}
}
`

type CharacterJSON struct {
	Name          string         `json:"name"`
	Race          string         `json:"race"`
	Class         string         `json:"class"`
	Age           int            `json:"age"`
	Level         int            `json:"level"`
	Alignment     string         `json:"alignment"`
	AbilityScores map[string]int `json:"ability_scores"`
}

// CreateCharacter creates a character
func CreateCharacter(ctx context.Context, client *ent.Client) {
	character := client.Character.Create().
		SetName("Zeke").
		SetRace(client.Race.Query().
			Where(race.Name("Human")).
			FirstX(ctx)).
		SetClass(client.Class.Query().
			Where(class.Name("Rogue")).
			FirstX(ctx)).
		SaveX(ctx)

	// Attach alignment
	client.Character.UpdateOne(character).
		SetAlignment(client.Alignment.Query().
			Where(alignment.Indx("CN")).
			OnlyX(ctx)).
		SaveX(ctx)

	charAbilityScores := map[string]int{
		"str": 8,
		"dex": 10,
		"con": 12,
		"int": 14,
		"wis": 16,
		"cha": 18,
	}
	for as, initScore := range charAbilityScores {
		client.CharacterAbilityScore.Create().
			SetCharacter(character).
			SetAbilityScore(
				client.AbilityScore.Query().
					Where(abilityscore.Indx(as)).
					OnlyX(ctx)).
			SetScore(initScore).
			SaveX(ctx)
	}

	log.Info("Character created", "character", character)
}

func CreateCharacterFromJSON(ctx context.Context, client *ent.Client, charJSON string) {
	var data CharacterJSON
	if err := json.Unmarshal([]byte(charJSON), &data); err != nil {
		log.Fatal(err)
	}

	character := client.Character.Create().
		SetName(data.Name).
		SetAge(data.Age).
		SetLevel(data.Level).
		SetRace(client.Race.Query().
			Where(race.Indx(data.Race)).
			OnlyX(ctx)).
		SetClass(client.Class.Query().
			Where(class.Indx(data.Class)).
			OnlyX(ctx)).
		SetAlignment(client.Alignment.Query().
			Where(alignment.Indx(data.Alignment)).
			OnlyX(ctx)).
		SaveX(ctx)

	for as, initScore := range data.AbilityScores {
		client.CharacterAbilityScore.Create().
			SetCharacter(character).
			SetAbilityScore(
				client.AbilityScore.Query().
					Where(abilityscore.Indx(as)).
					OnlyX(ctx)).
			SetScore(initScore).
			SaveX(ctx)
	}

	log.Info("Character created", "character", character)
}
