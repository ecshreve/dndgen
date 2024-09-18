package seeder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ecshreve/dndgen/builder/templates"
)

type SeederMeta struct {
	ObjName string
	MigName string
	MigComm string
}

func GenerateSeederFns(destDir string, data []SeederMeta) error {
	namesOnly := make([]string, len(data))
	for i, v := range data {
		namesOnly[i] = v.ObjName
	}

	fp := filepath.Join(destDir, "seeders_gen.go")
	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	t := template.Must(template.New("pop").Parse(templates.POP_TEMPLATE))
	if err := t.Execute(f, namesOnly); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func GenerateDataMigrationFns(destDir string, data []SeederMeta) error {
	fp := filepath.Join(destDir, "data_migration_generated.go")
	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	t := template.Must(template.New("data_migration").Parse(templates.DATA_MIGRATION_TEMPLATE))
	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
