package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/ecshreve/dndgen/ent/characterproficiency"
)

// IntOrDefault returns the value of the given integer pointer if it is not nil,
// otherwise it returns the default value.
func IntOrDefault(i *int, def int) int {
	if i == nil {
		return def
	}

	return *i
}

// cleanString cleans the given string by converting it to lowercase, replacing
// spaces with underscores, removing hyphens, apostrophes, and commas.
func CleanString(s string) string {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, ",", "")
	return s
}

// LoadJSONFile loads a JSON file from the given path and unmarshals it into the
// provided interface.
func LoadJSONFile(fpath string, v interface{}) error {
	// Check if the file is not a JSON file
	if !strings.HasPrefix(filepath.Ext(fpath), ".json") {
		return fmt.Errorf("file %s is not a JSON file", fpath)
	}

	data, err := os.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", fpath, err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return nil
}

// AbilityScoreModifier returns the modifier for the given ability score.
func AbilityScoreModifier(score int) int {
	if score < 1 {
		return -5
	}

	if score > 30 {
		return 10
	}

	return int(math.Floor(float64(score-10) / 2))
}

func LevelProficiencyBonus(level int) int {
	if level < 1 {
		return 0
	}

	if level < 5 {
		return 2
	}

	if level < 9 {
		return 3
	}

	if level < 13 {
		return 4
	}

	if level < 17 {
		return 5
	}

	return 6
}

func GetProficiencyTypeFromReference(reference string) characterproficiency.ProficiencyType {
	profType := strings.TrimPrefix(reference, "/api/")
	profType = strings.ToUpper(strings.ReplaceAll(strings.Split(profType, "/")[0], "-", "_"))
	profType = strings.ReplaceAll(profType, "IES", "Y")
	profType = strings.TrimSuffix(profType, "S")
	profType = strings.ReplaceAll(profType, "ABILITY_SCORE", "SAVING_THROW")

	return characterproficiency.ProficiencyType(profType)
}
