package popper_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/ecshreve/dndgen/internal/popper"
// 	"github.com/samsarahq/go/snapshotter"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// // TestPopulateSkill tests the Populate methods.
// func TestPopulateAbilityScoreAndSkills(t *testing.T) {
// 	ctx := context.Background()
// 	snap := snapshotter.New(t)
// 	defer snap.Verify()

// 	p := popper.NewTestPopper(ctx)

// 	err := p.PopulateAbilityScore(ctx)
// 	assert.NoError(t, err)

// 	as, err := p.Client.AbilityScore.Query().All(ctx)
// 	assert.NoError(t, err)
// 	snap.Snapshot("ability_scores", as)

// 	err = p.PopulateSkill(ctx)
// 	assert.NoError(t, err)

// 	sk, err := p.Client.Skill.Query().All(ctx)
// 	assert.NoError(t, err)
// 	snap.Snapshot("skills", sk)
// }

// func TestPopulateLanguages(t *testing.T) {
// 	ctx := context.Background()
// 	snap := snapshotter.New(t)
// 	defer snap.Verify()

// 	p := popper.NewTestPopper(ctx)

// 	err := p.PopulateLanguage(ctx)
// 	assert.NoError(t, err)

// 	ll, err := p.Client.Language.Query().All(ctx)
// 	assert.NoError(t, err)
// 	snap.Snapshot("languages", ll)
// }

// func TestPopulateClassesAndRaces(t *testing.T) {
// 	ctx := context.Background()
// 	snap := snapshotter.New(t)
// 	defer snap.Verify()

// 	p := popper.NewTestPopper(ctx)
// 	require.NoError(t, p.PopulateAbilityScore(ctx))
// 	require.NoError(t, p.PopulateSkill(ctx))
// 	require.NoError(t, p.PopulateLanguage(ctx))

// 	err := p.PopulateClass(ctx)
// 	assert.NoError(t, err)

// 	cl, err := p.Client.Class.Query().All(ctx)
// 	assert.NoError(t, err)
// 	snap.Snapshot("classes", cl)

// 	err = p.PopulateRace(ctx)
// 	assert.NoError(t, err)

// 	rr, err := p.Client.Race.Query().WithLanguages().All(ctx)
// 	assert.NoError(t, err)
// 	snap.Snapshot("races", rr)

// 	for _, r := range rr {
// 		ll, err := r.QueryLanguages().WithTypicalSpeakers().All(ctx)
// 		assert.NoError(t, err)
// 		snap.Snapshot("languages_"+r.Indx, ll)
// 	}
// }
