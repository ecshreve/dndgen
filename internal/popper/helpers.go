package popper

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/samsarahq/go/oops"
)

type Base struct {
	ID   string   `json:"index"`
	Name string   `json:"name"`
	Desc []string `json:"desc"`
}

type IDWrapper struct {
	ID string `json:"index"`
}

// LoadJSONFile loads a JSON file from the given path and unmarshals it into the
// given interface.
func LoadJSONFile(fpath string, v interface{}) error {
	// Check that the given file is a JSON file.
	ext := filepath.Ext(fpath)
	if ext != ".json" {
		return oops.Errorf("input file must be a JSON file: %s", fpath)
	}

	// Open the file specified by path.
	file, err := os.Open(fpath)
	if err != nil {
		return oops.Wrapf(err, "unable to open file %s", fpath)
	}
	defer file.Close()

	// Read the file into a byte array.
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return oops.Wrapf(err, "unable to read file %s to byte array", file.Name())
	}

	// Unmarshall the byte array into the given interface.
	if err = json.Unmarshal(byteValue, &v); err != nil {
		return oops.Wrapf(err, "unable to unmarshal byte array to map")
	}

	return nil
}
