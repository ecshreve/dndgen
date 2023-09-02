package popper_test

import (
	"context"
	"testing"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/enttest"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

// TestPopulate tests the Populate methods.
func TestPopulate(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	testClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer testClient.Close()

	p := popper.NewPopper(ctx, testClient)
	err := p.PopulateAll(ctx)
	assert.NoError(t, err)

	testcases := []struct {
		desc  string
		query []interface{}
	}{
		{
			desc: "ability_scores",
			query: []interface{}{
				testClient.AbilityScore.Query().AllX(ctx),
			},
		},
		{
			desc: "skills",
			query: []interface{}{
				testClient.Skill.Query().WithAbilityScore().AllX(ctx),
			},
		},
		{
			desc: "languages",
			query: []interface{}{
				testClient.Language.Query().AllX(ctx),
			},
		},
		{
			desc: "classes",
			query: []interface{}{
				testClient.Class.Query().
					WithProficiencyChoices(
						func(q *ent.ChoiceQuery) {
							q.WithProficiencyOptions()
							q.WithChoices()
						},
					).
					WithEquipmentChoice().AllX(ctx),
			},
		},
		{
			desc: "races",
			query: []interface{}{
				testClient.Race.Query().
					WithSubrace().
					WithStartingProficiencyOptions(
						func(q *ent.ChoiceQuery) {
							q.WithProficiencyOptions().
								WithChoices(func(cq *ent.ChoiceQuery) {
									cq.WithProficiencyOptions()
								})
						},
					).AllX(ctx),
			},
		},
		{
			desc: "proficiencies",
			query: []interface{}{
				testClient.Proficiency.Query().
					WithClasses().
					WithRaces().AllX(ctx),
			},
		},
		{
			desc: "equipment",
			query: []interface{}{
				testClient.Equipment.Query().AllX(ctx),
			},
		},
		{
			desc: "weapons",
			query: []interface{}{
				testClient.Weapon.Query().AllX(ctx),
			},
		},
		{
			desc: "armor",
			query: []interface{}{
				testClient.Armor.Query().AllX(ctx),
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			snap.Snapshot(tc.desc, tc.query)
		})
	}

}
