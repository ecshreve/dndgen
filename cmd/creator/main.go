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
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/race"
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
	cc := CreateCharacter(ctx, client, data)
	log.Info("Character created", "character", cc)

}

func CreateCharacter(ctx context.Context, client *ent.Client, charJSON CharacterJSON) *ent.Character {
	character := HandleNewCharacterCreation(ctx, client, charJSON)
	abilityScores, err := HandleCharacterAbilityScores(ctx, client, character, charJSON)
	if err != nil {
		log.Fatal(err)
	}
	_, err = HandleCharacterSkills(ctx, client, character, charJSON, abilityScores)
	if err != nil {
		log.Fatal(err)
	}
	// err = HandleCharacterProficiencies(ctx, client, character, charJSON, skills)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return character
}

// HandleNewCharacterCreation creates a new character in the database
func HandleNewCharacterCreation(ctx context.Context, client *ent.Client, charJSON CharacterJSON) *ent.Character {
	newChar := client.Character.Create().
		SetName(charJSON.Name).
		SetAge(charJSON.Age).
		SetLevel(charJSON.Level).
		SetAlignment(client.Alignment.Query().
			Where(alignment.Indx(charJSON.Alignment)).
			OnlyX(ctx)).
		SetRace(client.Race.Query().
			Where(race.Indx(charJSON.Race)).
			OnlyX(ctx)).
		SetClass(client.Class.Query().
			Where(class.Indx(charJSON.Class)).
			OnlyX(ctx)).
		SetProficiencyBonus(utils.LevelProficiencyBonus(charJSON.Level)).
		SaveX(ctx)
	return newChar
}

// HandleCharacterAbilityScores creates the character's ability scores
func HandleCharacterAbilityScores(ctx context.Context, client *ent.Client, ch *ent.Character, charJSON CharacterJSON) ([]*ent.CharacterAbilityScore, error) {
	asCache := make(map[string]int)
	allAS := client.AbilityScore.Query().AllX(ctx)
	for _, as := range allAS {
		asCache[as.Indx] = as.ID
	}
	for asIndx, asScore := range charJSON.AbilityScores {
		charAbilityScore := client.CharacterAbilityScore.Create().
			SetCharacter(ch).
			SetAbilityScoreID(asCache[asIndx]).
			SetScore(asScore).
			SaveX(ctx)
		log.Info("Character ability score created", "character_ability_score", charAbilityScore.QueryAbilityScore().OnlyX(ctx).Indx)
	}

	// Check if all ability scores are set
	charAbilityScores := ch.QueryCharacterAbilityScores().AllX(ctx)
	if len(charAbilityScores) != 6 {
		return nil, fmt.Errorf("character has wrong number of ability scores, got %d", len(charAbilityScores))
	}

	return charAbilityScores, nil
}

// HandleCharacterSkills creates the character's skills
func HandleCharacterSkills(ctx context.Context, client *ent.Client, ch *ent.Character, charJSON CharacterJSON, abilityScores []*ent.CharacterAbilityScore) ([]*ent.CharacterSkill, error) {
	casCache := make(map[string]*ent.CharacterAbilityScore)
	for _, cas := range abilityScores {
		casCache[cas.QueryAbilityScore().OnlyX(ctx).Indx] = cas
	}

	skills := client.Skill.Query().WithAbilityScore().AllX(ctx)
	for _, sk := range skills {
		charSkill := client.CharacterSkill.Create().
			SetCharacter(ch).
			SetSkill(sk).
			SetCharacterAbilityScore(casCache[sk.QueryAbilityScore().OnlyX(ctx).Indx]).
			SaveX(ctx)
		log.Info("Character skill created", "character_skill", charSkill.QuerySkill().OnlyX(ctx).Indx)
	}

	charSkills := ch.QueryCharacterSkills().WithSkill().AllX(ctx)
	if len(charSkills) != 18 {
		return nil, fmt.Errorf("character has wrong number of skills, got %d", len(charSkills))
	}

	return charSkills, nil
}

func HandleCharacterProficiencies(ctx context.Context, client *ent.Client, ch *ent.Character, charJSON CharacterJSON, skills []*ent.CharacterSkill) error {
	log.Info("Handling character proficiencies")
	proficiencies := ch.QueryRace().QueryStartingProficiencies().AllX(ctx)
	proficiencies = append(proficiencies, ch.QueryClass().QueryProficiencies().AllX(ctx)...)
	skillCache := make(map[string]*ent.CharacterSkill)

	for _, skill := range skills {
		skillCache[skill.Edges.Skill.Indx] = skill
	}

	for _, pp := range proficiencies {
		cpr := client.CharacterProficiency.Create().
			SetCharacter(ch).
			SetProficiency(pp).
			SaveX(ctx)
		log.Info("Character proficiency created", "character_proficiency", pp.Indx, "id", cpr.ProficiencyID)

		// If the proficiency is a skill, set the proficient flag to true for that skill
		if strings.Contains(pp.Indx, "skill") {
			ppIndx := strings.Split(pp.Indx, "-")[1]

			cs, err := skillCache[ppIndx].Update().
				SetProficient(true).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("error updating character skill: %w", err)
			}
			log.Info("Character skill updated", "character_skill", cs)
		}
	}

	return nil
}
