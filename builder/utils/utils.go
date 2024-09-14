package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
