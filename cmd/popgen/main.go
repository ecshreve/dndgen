package main

import (
	"os"
	"text/template"

	"github.com/charmbracelet/log"

	"github.com/samsarahq/go/oops"
)

func main() {
	log.Info("Starting dndgen/popgen...")

	tmplFile := "template/populator.tmpl"
	tmpl := template.Must(template.ParseFiles(tmplFile))

	var typeNamesToGenerate = []string{
		"AbilityScore",
		"Skill",
		"Language",
		"Alignment",
		// "DamageType",
		"Race",
		// "Class",
		// "WeaponProperty",
		// "MagicSchool",
		// "RuleSection",
		// "Rule",
		// "Subrace",
		// "Trait",
		// "Coin",
	}

	ofile, err := os.Create("internal/popper/generated.go")
	if err != nil {
		log.Fatal(oops.Wrapf(err, "unable to create output file for populators"))
	}
	defer ofile.Close()

	if err := tmpl.Execute(ofile, typeNamesToGenerate); err != nil {
		log.Fatal(oops.Wrapf(err, "unable to execute template"))
	}

	log.Info("done")
}
