package popper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

func intOrDef(i *int, def int) int {
	if i == nil {
		return def
	}

	return *i
}

type indxwrapper struct {
	Indx string `json:"index"`
}

func cleanString(s string) string {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, ",", "")
	return s
}

// GetIDsFromIndxs gets the IDs from the given indxs.
func (p *Popper) GetIDsFromIndxs(v []byte) []int {
	var indxs []IndxWrapper
	if err := json.Unmarshal(v, &indxs); err != nil {
		log.Fatal(err)
	}

	ids := make([]int, len(indxs))
	for i, indx := range indxs {
		ids[i] = p.IndxToId[indx.Indx]
	}

	return ids
}

// LoadJSONFile loads a JSON file from the given path and unmarshals it into the
// provided interface.
func LoadJSONFile(fpath string, v interface{}) error {
	if filepath.Ext(fpath) != ".json" {
		return fmt.Errorf("file %s is not a JSON file", fpath)
	}

	file, err := os.Open(fpath)
	if err != nil {
		return fmt.Errorf("unable to open file %s", fpath)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("unable to read file %s", fpath)
	}

	if err = json.Unmarshal(byteValue, v); err != nil {
		return fmt.Errorf("unable to unmarshal JSON from file %s", fpath)
	}

	return nil
}
