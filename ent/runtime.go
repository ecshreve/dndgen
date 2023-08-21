// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/schema"
	"github.com/ecshreve/dndgen/ent/skill"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	abilityscoreMixin := schema.AbilityScore{}.Mixin()
	abilityscoreMixinFields0 := abilityscoreMixin[0].Fields()
	_ = abilityscoreMixinFields0
	abilityscoreFields := schema.AbilityScore{}.Fields()
	_ = abilityscoreFields
	// abilityscoreDescID is the schema descriptor for id field.
	abilityscoreDescID := abilityscoreMixinFields0[0].Descriptor()
	// abilityscore.IDValidator is a validator for the "id" field. It is called by the builders before save.
	abilityscore.IDValidator = abilityscoreDescID.Validators[0].(func(string) error)
	skillMixin := schema.Skill{}.Mixin()
	skillMixinFields0 := skillMixin[0].Fields()
	_ = skillMixinFields0
	skillFields := schema.Skill{}.Fields()
	_ = skillFields
	// skillDescID is the schema descriptor for id field.
	skillDescID := skillMixinFields0[0].Descriptor()
	// skill.IDValidator is a validator for the "id" field. It is called by the builders before save.
	skill.IDValidator = skillDescID.Validators[0].(func(string) error)
}
