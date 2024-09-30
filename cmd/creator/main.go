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
	"github.com/ecshreve/dndgen/ent/character"
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
	cc := HandleNewCharacterCreation(ctx, client, data)
	log.Info("Character created", "character", cc)

}

// 	character := client.Character.Create().
// 		SetName(data.Name).
// 		SetAge(data.Age).
// 		SetLevel(data.Level).
// 		SetRace(client.Race.Query().
// 			Where(race.Indx(data.Race)).
// 			OnlyX(ctx)).
// 		SetClass(client.Class.Query().
// 			Where(class.Indx(data.Class)).
// 			OnlyX(ctx)).
// 		SetAlignment(client.Alignment.Query().
// 			Where(alignment.Indx(data.Alignment)).
// 			OnlyX(ctx)).
// 		SetProficiencyBonus(utils.LevelProficiencyBonus(data.Level)).
// 		SaveX(ctx)

// 	for as, initScore := range data.AbilityScores {
// 		client.CharacterAbilityScore.Create().
// 			SetCharacter(character).
// 			SetAbilityScore(
// 				client.AbilityScore.Query().
// 					Where(abilityscore.Indx(as)).
// 					OnlyX(ctx)).
// 			SetScore(initScore).
// 			SaveX(ctx)
// 	}

// 	skills := client.Skill.Query().AllX(ctx)
// 	for _, skill := range skills {
// 		csk := client.CharacterSkill.Create().
// 			SetCharacter(character).
// 			SetSkill(skill).
// 			SetModifier(
// 				client.CharacterAbilityScore.Query().
// 					Where(characterabilityscore.And(
// 						characterabilityscore.CharacterID(character.ID),
// 						characterabilityscore.AbilityScoreID(skill.QueryAbilityScore().FirstIDX(ctx)),
// 					)).
// 					OnlyX(ctx).
// 					Modifier,
// 			).
// 			SaveX(ctx)
// 		log.Info("Character skill created", "character_skill", csk)
// 	}

// 	proficiencies := character.QueryRace().QueryStartingProficiencies().AllX(ctx)
// 	proficiencies = append(proficiencies, character.QueryClass().QueryProficiencies().AllX(ctx)...)
// 	for _, proficiency := range proficiencies {
// 		cpr := client.CharacterProficiency.Create().
// 			SetCharacter(character).
// 			SetProficiency(proficiency).
// 			SaveX(ctx)
// 		log.Info("Character proficiency created", "character_proficiency", cpr)

// 		skSplits := strings.Split(proficiency.Indx, "-")
// 		if skSplits[0] == "skill" {
// 			log.Debug("Proficiency is a skill", "proficiency", proficiency)
// 			profSkill := client.Skill.Query().
// 				Where(skill.Indx(skSplits[1])).
// 				WithAbilityScore().
// 				FirstX(ctx)
// 			if profSkill == nil {
// 				log.Warn("Proficiency skill not found", "proficiency", proficiency)
// 				continue
// 			}

// 			log.Info("Proficiency skill", "skill", profSkill)

// 			charAs := character.QueryCharacterAbilityScores().
// 				Where(characterabilityscore.And(
// 					characterabilityscore.CharacterID(character.ID),
// 					characterabilityscore.AbilityScoreID(profSkill.QueryAbilityScore().FirstIDX(ctx)),
// 				)).
// 				OnlyX(ctx)
// 			if charAs == nil {
// 				log.Warn("Character ability score not found", "character", character, "ability_score", profSkill.AbilityScore)
// 				continue
// 			}

// 			mod := charAs.Modifier + utils.LevelProficiencyBonus(character.Level)

// 			characterSkillUpdate := client.CharacterSkill.Update().
// 				Where(characterskill.And(
// 					characterskill.CharacterID(character.ID),
// 					characterskill.SkillID(profSkill.ID),
// 				)).
// 				SetModifier(mod).
// 				SetProficient(true).
// 				SaveX(ctx)
// 			log.Info("Character skill updated", "character_skill", characterSkillUpdate)
// 		}
// 	}
// }

func CreateCharacter(ctx context.Context, client *ent.Client, charJSON CharacterJSON) *ent.Character {
	character := HandleNewCharacterCreation(ctx, client, charJSON)
	HandleCharacterAbilityScores(ctx, client, character, charJSON)
	HandleCharacterSkills(ctx, client, character, charJSON)
	HandleCharacterProficiencies(ctx, client, character, charJSON)
	return character
}

// HandleNewCharacterCreation creates a new character in the database
func HandleNewCharacterCreation(ctx context.Context, client *ent.Client, charJSON CharacterJSON) *ent.Character {
	// Check if the character already exists
	existingChar, err := client.Character.Query().
		Where(character.Name(charJSON.Name)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Info("Character not found, creating new character")
		} else {
			log.Error("Error querying character", "error", err)
			return nil
		}
	}

	if existingChar != nil {
		log.Warn("Character already exists", "character", existingChar)
		return existingChar
	}

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
	log.Info("Character created", "character", newChar)

	for as, initScore := range charJSON.AbilityScores {
		client.CharacterAbilityScore.Create().
			SetCharacter(newChar).
			SetAbilityScore(
				client.AbilityScore.Query().
					Where(abilityscore.Indx(as)).
					OnlyX(ctx)).
			SetScore(initScore).
			SetModifier(utils.AbilityScoreModifier(initScore)).
			SaveX(ctx)
	}

	skills := client.Skill.Query().AllX(ctx)
	for _, skill := range skills {
		csk := client.CharacterSkill.Create().
			SetCharacter(newChar).
			SetSkill(skill).
			SetModifier(
				client.CharacterAbilityScore.Query().
					Where(characterabilityscore.And(
						characterabilityscore.CharacterID(newChar.ID),
						characterabilityscore.AbilityScoreID(skill.QueryAbilityScore().FirstIDX(ctx)),
					)).
					OnlyX(ctx).
					Modifier,
			).
			SaveX(ctx)
		log.Info("Character skill created", "character_skill", csk)
	}

	HandleCharacterProficiencies(ctx, client, newChar, charJSON)

	return newChar
}

// HandleCharacterAbilityScores creates the character's ability scores
func HandleCharacterAbilityScores(ctx context.Context, client *ent.Client, ch *ent.Character, charJSON CharacterJSON) {
	charAbilityScores := ch.QueryCharacterAbilityScores().AllX(ctx)
	if len(charAbilityScores) != 6 {
		log.Error("Character has wrong number of ability scores", "character", ch)
		return
	}

	for _, charAS := range charAbilityScores {
		asIndx := charAS.Edges.AbilityScore.Indx
		initScore := charJSON.AbilityScores[asIndx]
		if charAS.Score == initScore {
			log.Info("Ability score unchanged", "ability_score", charAS)
			continue
		}

		charAS.Update().
			SetScore(initScore).
			SetModifier(utils.AbilityScoreModifier(initScore)).
			SaveX(ctx)
		log.Info("Character ability score updated", "character_ability_score", charAS)
	}
}

// HandleCharacterSkills creates the character's skills
func HandleCharacterSkills(ctx context.Context, client *ent.Client, ch *ent.Character, charJSON CharacterJSON) {
	charSkills := ch.QueryCharacterSkills().AllX(ctx)
	if len(charSkills) != 14 {
		log.Error("Character has wrong number of skills", "character", ch)
		return
	}

	for _, charSkill := range charSkills {
		asId := charSkill.Edges.Skill.QueryAbilityScore().FirstIDX(ctx)
		charAs := ch.QueryCharacterAbilityScores().
			Where(characterabilityscore.AbilityScoreID(asId)).
			OnlyX(ctx)
		if charAs == nil {
			log.Error("Character ability score not found", "character", ch, "ability_score", asId)
			continue
		}

		mod := charAs.Modifier

		charSkill.Update().
			SetModifier(mod).
			SetProficient(false).
			SaveX(ctx)
		log.Info("Character skill updated", "character_skill", charSkill)
	}
}

func HandleCharacterProficiencies(ctx context.Context, client *ent.Client, ch *ent.Character, charJSON CharacterJSON) {
	proficiencies := ch.QueryRace().QueryStartingProficiencies().AllX(ctx)
	proficiencies = append(proficiencies, ch.QueryClass().QueryProficiencies().AllX(ctx)...)

	ch.Update().ClearProficiencies().ExecX(ctx)

	for _, proficiency := range proficiencies {
		cpr := client.CharacterProficiency.Create().
			SetCharacter(ch).
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

			charAs := ch.QueryCharacterAbilityScores().
				Where(characterabilityscore.And(
					characterabilityscore.CharacterID(ch.ID),
					characterabilityscore.AbilityScoreID(profSkill.QueryAbilityScore().FirstIDX(ctx)),
				)).
				OnlyX(ctx)
			if charAs == nil {
				log.Warn("Character ability score not found", "character", ch.Name, "ability_score", profSkill.AbilityScore)
				continue
			}

			mod := charAs.Modifier + utils.LevelProficiencyBonus(ch.Level)

			characterSkillUpdate := client.CharacterSkill.Update().
				Where(characterskill.And(
					characterskill.CharacterID(ch.ID),
					characterskill.SkillID(profSkill.ID),
				)).
				SetModifier(mod).
				SetProficient(true).
				SaveX(ctx)
			log.Info("Character skill updated", "character_skill", characterSkillUpdate)
		}
	}

}
