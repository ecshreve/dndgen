// GENERATED BY seeder
package seeder

import (
	"context"
	"fmt"

	"builder/ent"
	"builder/utils"

	"github.com/charmbracelet/log"
)

func SeedAll(ctx context.Context, client *ent.Client) error {
	seedFns := []func(ctx context.Context, client *ent.Client) error{
		SeedAbilityScore,SeedAlignment,SeedClass,SeedRace,SeedSkill,SeedLanguage,SeedMagicSchool,
	}

	for _, fn := range seedFns {
		err := fn(ctx, client)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}


// SeedAbilityScore seeds the AbilityScore entities from the JSON data files.
func SeedAbilityScore(ctx context.Context, client *ent.Client) error {
	fpath := "data/AbilityScore.json"
	var v []ent.AbilityScore

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.AbilityScoreCreate, len(v))
	for i, vv := range v {
		creates[i] = client.AbilityScore.Create().SetAbilityScore(&vv)
	}

	created, err := client.AbilityScore.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("AbilityScore bulk creation success", "created", len(created))

	return nil
}

// SeedAlignment seeds the Alignment entities from the JSON data files.
func SeedAlignment(ctx context.Context, client *ent.Client) error {
	fpath := "data/Alignment.json"
	var v []ent.Alignment

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.AlignmentCreate, len(v))
	for i, vv := range v {
		creates[i] = client.Alignment.Create().SetAlignment(&vv)
	}

	created, err := client.Alignment.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("Alignment bulk creation success", "created", len(created))

	return nil
}

// SeedClass seeds the Class entities from the JSON data files.
func SeedClass(ctx context.Context, client *ent.Client) error {
	fpath := "data/Class.json"
	var v []ent.Class

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.ClassCreate, len(v))
	for i, vv := range v {
		creates[i] = client.Class.Create().SetClass(&vv)
	}

	created, err := client.Class.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("Class bulk creation success", "created", len(created))

	return nil
}

// SeedRace seeds the Race entities from the JSON data files.
func SeedRace(ctx context.Context, client *ent.Client) error {
	fpath := "data/Race.json"
	var v []ent.Race

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.RaceCreate, len(v))
	for i, vv := range v {
		creates[i] = client.Race.Create().SetRace(&vv)
	}

	created, err := client.Race.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("Race bulk creation success", "created", len(created))

	return nil
}

// SeedSkill seeds the Skill entities from the JSON data files.
func SeedSkill(ctx context.Context, client *ent.Client) error {
	fpath := "data/Skill.json"
	var v []ent.Skill

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.SkillCreate, len(v))
	for i, vv := range v {
		creates[i] = client.Skill.Create().SetSkill(&vv)
	}

	created, err := client.Skill.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("Skill bulk creation success", "created", len(created))

	return nil
}

// SeedLanguage seeds the Language entities from the JSON data files.
func SeedLanguage(ctx context.Context, client *ent.Client) error {
	fpath := "data/Language.json"
	var v []ent.Language

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.LanguageCreate, len(v))
	for i, vv := range v {
		creates[i] = client.Language.Create().SetLanguage(&vv)
	}

	created, err := client.Language.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("Language bulk creation success", "created", len(created))

	return nil
}

// SeedMagicSchool seeds the MagicSchool entities from the JSON data files.
func SeedMagicSchool(ctx context.Context, client *ent.Client) error {
	fpath := "data/MagicSchool.json"
	var v []ent.MagicSchool

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("%w", err)
	}

	creates := make([]*ent.MagicSchoolCreate, len(v))
	for i, vv := range v {
		creates[i] = client.MagicSchool.Create().SetMagicSchool(&vv)
	}

	created, err := client.MagicSchool.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Info("MagicSchool bulk creation success", "created", len(created))

	return nil
}

