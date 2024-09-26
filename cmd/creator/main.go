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
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterskill"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/internal/utils"

	_ "github.com/ecshreve/dndgen/ent/runtime"

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
	"race": "elf",
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

	skills := client.Skill.Query().AllX(ctx)
	for _, skill := range skills {
		csk := client.CharacterSkill.Create().
			SetCharacter(character).
			SetSkill(skill).
			SetModifier(
				client.CharacterAbilityScore.Query().
					Where(characterabilityscore.And(
						characterabilityscore.CharacterID(character.ID),
						characterabilityscore.AbilityScoreID(skill.QueryAbilityScore().FirstIDX(ctx)),
					)).
					OnlyX(ctx).
					Modifier,
			).
			SaveX(ctx)
		log.Info("Character skill created", "character_skill", csk)
	}

	proficiencies := character.QueryRace().QueryStartingProficiencies().AllX(ctx)
	proficiencies = append(proficiencies, character.QueryClass().QueryProficiencies().AllX(ctx)...)
	for _, proficiency := range proficiencies {
		cpr := client.CharacterProficiency.Create().
			SetCharacter(character).
			SetProficiency(proficiency).
			SaveX(ctx)
		log.Info("Character proficiency created", "character_proficiency", cpr)

		skSplits := strings.Split(proficiency.Indx, "-")
		if skSplits[0] == "skill" {
			log.Debug("Proficiency is a skill", "proficiency", proficiency)
			profSkill := client.Skill.Query().
				Where(skill.Indx(skSplits[1])).
				WithAbilityScore().
				FirstX(ctx)
			if profSkill == nil {
				log.Warn("Proficiency skill not found", "proficiency", proficiency)
				continue
			}

			log.Info("Proficiency skill", "skill", profSkill)

			charAs := character.QueryCharacterAbilityScores().
				Where(characterabilityscore.And(
					characterabilityscore.CharacterID(character.ID),
					characterabilityscore.AbilityScoreID(profSkill.QueryAbilityScore().FirstIDX(ctx)),
				)).
				OnlyX(ctx)
			if charAs == nil {
				log.Warn("Character ability score not found", "character", character, "ability_score", profSkill.AbilityScore)
				continue
			}

			mod := charAs.Modifier + utils.LevelProficiencyBonus(character.Level)

			characterSkillUpdate := client.CharacterSkill.Update().
				Where(characterskill.And(
					characterskill.CharacterID(character.ID),
					characterskill.SkillID(profSkill.ID),
				)).
				SetModifier(mod).
				SetProficient(true).
				SaveX(ctx)
			log.Info("Character skill updated", "character_skill", characterSkillUpdate)
		}
	}
}
