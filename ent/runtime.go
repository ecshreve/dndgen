// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilitybonuschoice"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/condition"
	"github.com/ecshreve/dndgen/ent/cost"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/feat"
	"github.com/ecshreve/dndgen/ent/feature"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/languagechoice"
	"github.com/ecshreve/dndgen/ent/magicschool"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
	"github.com/ecshreve/dndgen/ent/property"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/rule"
	"github.com/ecshreve/dndgen/ent/rulesection"
	"github.com/ecshreve/dndgen/ent/schema"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/ent/trait"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	abilitybonusFields := schema.AbilityBonus{}.Fields()
	_ = abilitybonusFields
	// abilitybonusDescBonus is the schema descriptor for bonus field.
	abilitybonusDescBonus := abilitybonusFields[0].Descriptor()
	// abilitybonus.BonusValidator is a validator for the "bonus" field. It is called by the builders before save.
	abilitybonus.BonusValidator = abilitybonusDescBonus.Validators[0].(func(int) error)
	abilitybonuschoiceFields := schema.AbilityBonusChoice{}.Fields()
	_ = abilitybonuschoiceFields
	// abilitybonuschoiceDescChoose is the schema descriptor for choose field.
	abilitybonuschoiceDescChoose := abilitybonuschoiceFields[0].Descriptor()
	// abilitybonuschoice.ChooseValidator is a validator for the "choose" field. It is called by the builders before save.
	abilitybonuschoice.ChooseValidator = abilitybonuschoiceDescChoose.Validators[0].(func(int) error)
	abilityscoreMixin := schema.AbilityScore{}.Mixin()
	abilityscoreMixinFields0 := abilityscoreMixin[0].Fields()
	_ = abilityscoreMixinFields0
	abilityscoreFields := schema.AbilityScore{}.Fields()
	_ = abilityscoreFields
	// abilityscoreDescIndx is the schema descriptor for indx field.
	abilityscoreDescIndx := abilityscoreMixinFields0[0].Descriptor()
	// abilityscore.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	abilityscore.IndxValidator = abilityscoreDescIndx.Validators[0].(func(string) error)
	// abilityscoreDescName is the schema descriptor for name field.
	abilityscoreDescName := abilityscoreMixinFields0[1].Descriptor()
	// abilityscore.NameValidator is a validator for the "name" field. It is called by the builders before save.
	abilityscore.NameValidator = abilityscoreDescName.Validators[0].(func(string) error)
	alignmentMixin := schema.Alignment{}.Mixin()
	alignmentMixinFields0 := alignmentMixin[0].Fields()
	_ = alignmentMixinFields0
	alignmentFields := schema.Alignment{}.Fields()
	_ = alignmentFields
	// alignmentDescIndx is the schema descriptor for indx field.
	alignmentDescIndx := alignmentMixinFields0[0].Descriptor()
	// alignment.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	alignment.IndxValidator = alignmentDescIndx.Validators[0].(func(string) error)
	// alignmentDescName is the schema descriptor for name field.
	alignmentDescName := alignmentMixinFields0[1].Descriptor()
	// alignment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	alignment.NameValidator = alignmentDescName.Validators[0].(func(string) error)
	armorFields := schema.Armor{}.Fields()
	_ = armorFields
	// armorDescAcBase is the schema descriptor for ac_base field.
	armorDescAcBase := armorFields[3].Descriptor()
	// armor.AcBaseValidator is a validator for the "ac_base" field. It is called by the builders before save.
	armor.AcBaseValidator = armorDescAcBase.Validators[0].(func(int) error)
	// armorDescAcDexBonus is the schema descriptor for ac_dex_bonus field.
	armorDescAcDexBonus := armorFields[4].Descriptor()
	// armor.DefaultAcDexBonus holds the default value on creation for the ac_dex_bonus field.
	armor.DefaultAcDexBonus = armorDescAcDexBonus.Default.(bool)
	// armorDescAcMaxBonus is the schema descriptor for ac_max_bonus field.
	armorDescAcMaxBonus := armorFields[5].Descriptor()
	// armor.DefaultAcMaxBonus holds the default value on creation for the ac_max_bonus field.
	armor.DefaultAcMaxBonus = armorDescAcMaxBonus.Default.(int)
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescIndx is the schema descriptor for indx field.
	classDescIndx := classFields[0].Descriptor()
	// class.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	class.IndxValidator = classDescIndx.Validators[0].(func(string) error)
	// classDescName is the schema descriptor for name field.
	classDescName := classFields[1].Descriptor()
	// class.NameValidator is a validator for the "name" field. It is called by the builders before save.
	class.NameValidator = classDescName.Validators[0].(func(string) error)
	// classDescHitDie is the schema descriptor for hit_die field.
	classDescHitDie := classFields[2].Descriptor()
	// class.HitDieValidator is a validator for the "hit_die" field. It is called by the builders before save.
	class.HitDieValidator = classDescHitDie.Validators[0].(func(int) error)
	coinMixin := schema.Coin{}.Mixin()
	coinMixinFields0 := coinMixin[0].Fields()
	_ = coinMixinFields0
	coinFields := schema.Coin{}.Fields()
	_ = coinFields
	// coinDescIndx is the schema descriptor for indx field.
	coinDescIndx := coinMixinFields0[0].Descriptor()
	// coin.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	coin.IndxValidator = coinDescIndx.Validators[0].(func(string) error)
	// coinDescName is the schema descriptor for name field.
	coinDescName := coinMixinFields0[1].Descriptor()
	// coin.NameValidator is a validator for the "name" field. It is called by the builders before save.
	coin.NameValidator = coinDescName.Validators[0].(func(string) error)
	conditionMixin := schema.Condition{}.Mixin()
	conditionMixinFields0 := conditionMixin[0].Fields()
	_ = conditionMixinFields0
	conditionFields := schema.Condition{}.Fields()
	_ = conditionFields
	// conditionDescIndx is the schema descriptor for indx field.
	conditionDescIndx := conditionMixinFields0[0].Descriptor()
	// condition.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	condition.IndxValidator = conditionDescIndx.Validators[0].(func(string) error)
	// conditionDescName is the schema descriptor for name field.
	conditionDescName := conditionMixinFields0[1].Descriptor()
	// condition.NameValidator is a validator for the "name" field. It is called by the builders before save.
	condition.NameValidator = conditionDescName.Validators[0].(func(string) error)
	costFields := schema.Cost{}.Fields()
	_ = costFields
	// costDescQuantity is the schema descriptor for quantity field.
	costDescQuantity := costFields[0].Descriptor()
	// cost.DefaultQuantity holds the default value on creation for the quantity field.
	cost.DefaultQuantity = costDescQuantity.Default.(int)
	damagetypeMixin := schema.DamageType{}.Mixin()
	damagetypeMixinFields0 := damagetypeMixin[0].Fields()
	_ = damagetypeMixinFields0
	damagetypeFields := schema.DamageType{}.Fields()
	_ = damagetypeFields
	// damagetypeDescIndx is the schema descriptor for indx field.
	damagetypeDescIndx := damagetypeMixinFields0[0].Descriptor()
	// damagetype.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	damagetype.IndxValidator = damagetypeDescIndx.Validators[0].(func(string) error)
	// damagetypeDescName is the schema descriptor for name field.
	damagetypeDescName := damagetypeMixinFields0[1].Descriptor()
	// damagetype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	damagetype.NameValidator = damagetypeDescName.Validators[0].(func(string) error)
	equipmentFields := schema.Equipment{}.Fields()
	_ = equipmentFields
	// equipmentDescIndx is the schema descriptor for indx field.
	equipmentDescIndx := equipmentFields[0].Descriptor()
	// equipment.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	equipment.IndxValidator = equipmentDescIndx.Validators[0].(func(string) error)
	// equipmentDescName is the schema descriptor for name field.
	equipmentDescName := equipmentFields[1].Descriptor()
	// equipment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	equipment.NameValidator = equipmentDescName.Validators[0].(func(string) error)
	featMixin := schema.Feat{}.Mixin()
	featMixinFields0 := featMixin[0].Fields()
	_ = featMixinFields0
	featFields := schema.Feat{}.Fields()
	_ = featFields
	// featDescIndx is the schema descriptor for indx field.
	featDescIndx := featMixinFields0[0].Descriptor()
	// feat.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	feat.IndxValidator = featDescIndx.Validators[0].(func(string) error)
	// featDescName is the schema descriptor for name field.
	featDescName := featMixinFields0[1].Descriptor()
	// feat.NameValidator is a validator for the "name" field. It is called by the builders before save.
	feat.NameValidator = featDescName.Validators[0].(func(string) error)
	featureMixin := schema.Feature{}.Mixin()
	featureMixinFields0 := featureMixin[0].Fields()
	_ = featureMixinFields0
	featureFields := schema.Feature{}.Fields()
	_ = featureFields
	// featureDescIndx is the schema descriptor for indx field.
	featureDescIndx := featureMixinFields0[0].Descriptor()
	// feature.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	feature.IndxValidator = featureDescIndx.Validators[0].(func(string) error)
	// featureDescName is the schema descriptor for name field.
	featureDescName := featureMixinFields0[1].Descriptor()
	// feature.NameValidator is a validator for the "name" field. It is called by the builders before save.
	feature.NameValidator = featureDescName.Validators[0].(func(string) error)
	// featureDescLevel is the schema descriptor for level field.
	featureDescLevel := featureFields[0].Descriptor()
	// feature.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	feature.LevelValidator = featureDescLevel.Validators[0].(func(int) error)
	languageMixin := schema.Language{}.Mixin()
	languageMixinFields0 := languageMixin[0].Fields()
	_ = languageMixinFields0
	languageFields := schema.Language{}.Fields()
	_ = languageFields
	// languageDescIndx is the schema descriptor for indx field.
	languageDescIndx := languageMixinFields0[0].Descriptor()
	// language.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	language.IndxValidator = languageDescIndx.Validators[0].(func(string) error)
	// languageDescName is the schema descriptor for name field.
	languageDescName := languageMixinFields0[1].Descriptor()
	// language.NameValidator is a validator for the "name" field. It is called by the builders before save.
	language.NameValidator = languageDescName.Validators[0].(func(string) error)
	languagechoiceFields := schema.LanguageChoice{}.Fields()
	_ = languagechoiceFields
	// languagechoiceDescChoose is the schema descriptor for choose field.
	languagechoiceDescChoose := languagechoiceFields[0].Descriptor()
	// languagechoice.ChooseValidator is a validator for the "choose" field. It is called by the builders before save.
	languagechoice.ChooseValidator = languagechoiceDescChoose.Validators[0].(func(int) error)
	magicschoolMixin := schema.MagicSchool{}.Mixin()
	magicschoolMixinFields0 := magicschoolMixin[0].Fields()
	_ = magicschoolMixinFields0
	magicschoolFields := schema.MagicSchool{}.Fields()
	_ = magicschoolFields
	// magicschoolDescIndx is the schema descriptor for indx field.
	magicschoolDescIndx := magicschoolMixinFields0[0].Descriptor()
	// magicschool.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	magicschool.IndxValidator = magicschoolDescIndx.Validators[0].(func(string) error)
	// magicschoolDescName is the schema descriptor for name field.
	magicschoolDescName := magicschoolMixinFields0[1].Descriptor()
	// magicschool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	magicschool.NameValidator = magicschoolDescName.Validators[0].(func(string) error)
	proficiencyFields := schema.Proficiency{}.Fields()
	_ = proficiencyFields
	// proficiencyDescIndx is the schema descriptor for indx field.
	proficiencyDescIndx := proficiencyFields[0].Descriptor()
	// proficiency.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	proficiency.IndxValidator = proficiencyDescIndx.Validators[0].(func(string) error)
	// proficiencyDescName is the schema descriptor for name field.
	proficiencyDescName := proficiencyFields[1].Descriptor()
	// proficiency.NameValidator is a validator for the "name" field. It is called by the builders before save.
	proficiency.NameValidator = proficiencyDescName.Validators[0].(func(string) error)
	// proficiencyDescReference is the schema descriptor for reference field.
	proficiencyDescReference := proficiencyFields[2].Descriptor()
	// proficiency.ReferenceValidator is a validator for the "reference" field. It is called by the builders before save.
	proficiency.ReferenceValidator = proficiencyDescReference.Validators[0].(func(string) error)
	proficiencychoiceFields := schema.ProficiencyChoice{}.Fields()
	_ = proficiencychoiceFields
	// proficiencychoiceDescChoose is the schema descriptor for choose field.
	proficiencychoiceDescChoose := proficiencychoiceFields[0].Descriptor()
	// proficiencychoice.ChooseValidator is a validator for the "choose" field. It is called by the builders before save.
	proficiencychoice.ChooseValidator = proficiencychoiceDescChoose.Validators[0].(func(int) error)
	propertyMixin := schema.Property{}.Mixin()
	propertyMixinFields0 := propertyMixin[0].Fields()
	_ = propertyMixinFields0
	propertyFields := schema.Property{}.Fields()
	_ = propertyFields
	// propertyDescIndx is the schema descriptor for indx field.
	propertyDescIndx := propertyMixinFields0[0].Descriptor()
	// property.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	property.IndxValidator = propertyDescIndx.Validators[0].(func(string) error)
	// propertyDescName is the schema descriptor for name field.
	propertyDescName := propertyMixinFields0[1].Descriptor()
	// property.NameValidator is a validator for the "name" field. It is called by the builders before save.
	property.NameValidator = propertyDescName.Validators[0].(func(string) error)
	raceFields := schema.Race{}.Fields()
	_ = raceFields
	// raceDescIndx is the schema descriptor for indx field.
	raceDescIndx := raceFields[0].Descriptor()
	// race.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	race.IndxValidator = raceDescIndx.Validators[0].(func(string) error)
	// raceDescName is the schema descriptor for name field.
	raceDescName := raceFields[1].Descriptor()
	// race.NameValidator is a validator for the "name" field. It is called by the builders before save.
	race.NameValidator = raceDescName.Validators[0].(func(string) error)
	// raceDescSpeed is the schema descriptor for speed field.
	raceDescSpeed := raceFields[2].Descriptor()
	// race.SpeedValidator is a validator for the "speed" field. It is called by the builders before save.
	race.SpeedValidator = raceDescSpeed.Validators[0].(func(int) error)
	ruleMixin := schema.Rule{}.Mixin()
	ruleMixinFields0 := ruleMixin[0].Fields()
	_ = ruleMixinFields0
	ruleFields := schema.Rule{}.Fields()
	_ = ruleFields
	// ruleDescIndx is the schema descriptor for indx field.
	ruleDescIndx := ruleMixinFields0[0].Descriptor()
	// rule.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	rule.IndxValidator = ruleDescIndx.Validators[0].(func(string) error)
	// ruleDescName is the schema descriptor for name field.
	ruleDescName := ruleMixinFields0[1].Descriptor()
	// rule.NameValidator is a validator for the "name" field. It is called by the builders before save.
	rule.NameValidator = ruleDescName.Validators[0].(func(string) error)
	rulesectionMixin := schema.RuleSection{}.Mixin()
	rulesectionMixinFields0 := rulesectionMixin[0].Fields()
	_ = rulesectionMixinFields0
	rulesectionFields := schema.RuleSection{}.Fields()
	_ = rulesectionFields
	// rulesectionDescIndx is the schema descriptor for indx field.
	rulesectionDescIndx := rulesectionMixinFields0[0].Descriptor()
	// rulesection.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	rulesection.IndxValidator = rulesectionDescIndx.Validators[0].(func(string) error)
	// rulesectionDescName is the schema descriptor for name field.
	rulesectionDescName := rulesectionMixinFields0[1].Descriptor()
	// rulesection.NameValidator is a validator for the "name" field. It is called by the builders before save.
	rulesection.NameValidator = rulesectionDescName.Validators[0].(func(string) error)
	skillMixin := schema.Skill{}.Mixin()
	skillMixinFields0 := skillMixin[0].Fields()
	_ = skillMixinFields0
	skillFields := schema.Skill{}.Fields()
	_ = skillFields
	// skillDescIndx is the schema descriptor for indx field.
	skillDescIndx := skillMixinFields0[0].Descriptor()
	// skill.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	skill.IndxValidator = skillDescIndx.Validators[0].(func(string) error)
	// skillDescName is the schema descriptor for name field.
	skillDescName := skillMixinFields0[1].Descriptor()
	// skill.NameValidator is a validator for the "name" field. It is called by the builders before save.
	skill.NameValidator = skillDescName.Validators[0].(func(string) error)
	traitMixin := schema.Trait{}.Mixin()
	traitMixinFields0 := traitMixin[0].Fields()
	_ = traitMixinFields0
	traitFields := schema.Trait{}.Fields()
	_ = traitFields
	// traitDescIndx is the schema descriptor for indx field.
	traitDescIndx := traitMixinFields0[0].Descriptor()
	// trait.IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	trait.IndxValidator = traitDescIndx.Validators[0].(func(string) error)
	// traitDescName is the schema descriptor for name field.
	traitDescName := traitMixinFields0[1].Descriptor()
	// trait.NameValidator is a validator for the "name" field. It is called by the builders before save.
	trait.NameValidator = traitDescName.Validators[0].(func(string) error)
}
