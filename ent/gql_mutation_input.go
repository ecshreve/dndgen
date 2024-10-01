// Code generated by ent, DO NOT EDIT.

package ent

// CreateCharacterInput represents a mutation input for creating characters.
type CreateCharacterInput struct {
	Name                     string
	Age                      *int
	Level                    *int
	ProficiencyBonus         *int
	RaceID                   *int
	ClassID                  *int
	AlignmentID              *int
	CharacterAbilityScoreIDs []int
	CharacterSkillIDs        []int
}

// Mutate applies the CreateCharacterInput on the CharacterMutation builder.
func (i *CreateCharacterInput) Mutate(m *CharacterMutation) {
	m.SetName(i.Name)
	if v := i.Age; v != nil {
		m.SetAge(*v)
	}
	if v := i.Level; v != nil {
		m.SetLevel(*v)
	}
	if v := i.ProficiencyBonus; v != nil {
		m.SetProficiencyBonus(*v)
	}
	if v := i.RaceID; v != nil {
		m.SetRaceID(*v)
	}
	if v := i.ClassID; v != nil {
		m.SetClassID(*v)
	}
	if v := i.AlignmentID; v != nil {
		m.SetAlignmentID(*v)
	}
	if v := i.CharacterAbilityScoreIDs; len(v) > 0 {
		m.AddCharacterAbilityScoreIDs(v...)
	}
	if v := i.CharacterSkillIDs; len(v) > 0 {
		m.AddCharacterSkillIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCharacterInput on the CharacterCreate builder.
func (c *CharacterCreate) SetInput(i CreateCharacterInput) *CharacterCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCharacterInput represents a mutation input for updating characters.
type UpdateCharacterInput struct {
	Name                           *string
	Age                            *int
	Level                          *int
	ProficiencyBonus               *int
	ClearRace                      bool
	RaceID                         *int
	ClearClass                     bool
	ClassID                        *int
	ClearAlignment                 bool
	AlignmentID                    *int
	ClearCharacterAbilityScores    bool
	AddCharacterAbilityScoreIDs    []int
	RemoveCharacterAbilityScoreIDs []int
	ClearCharacterSkills           bool
	AddCharacterSkillIDs           []int
	RemoveCharacterSkillIDs        []int
}

// Mutate applies the UpdateCharacterInput on the CharacterMutation builder.
func (i *UpdateCharacterInput) Mutate(m *CharacterMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Age; v != nil {
		m.SetAge(*v)
	}
	if v := i.Level; v != nil {
		m.SetLevel(*v)
	}
	if v := i.ProficiencyBonus; v != nil {
		m.SetProficiencyBonus(*v)
	}
	if i.ClearRace {
		m.ClearRace()
	}
	if v := i.RaceID; v != nil {
		m.SetRaceID(*v)
	}
	if i.ClearClass {
		m.ClearClass()
	}
	if v := i.ClassID; v != nil {
		m.SetClassID(*v)
	}
	if i.ClearAlignment {
		m.ClearAlignment()
	}
	if v := i.AlignmentID; v != nil {
		m.SetAlignmentID(*v)
	}
	if i.ClearCharacterAbilityScores {
		m.ClearCharacterAbilityScores()
	}
	if v := i.AddCharacterAbilityScoreIDs; len(v) > 0 {
		m.AddCharacterAbilityScoreIDs(v...)
	}
	if v := i.RemoveCharacterAbilityScoreIDs; len(v) > 0 {
		m.RemoveCharacterAbilityScoreIDs(v...)
	}
	if i.ClearCharacterSkills {
		m.ClearCharacterSkills()
	}
	if v := i.AddCharacterSkillIDs; len(v) > 0 {
		m.AddCharacterSkillIDs(v...)
	}
	if v := i.RemoveCharacterSkillIDs; len(v) > 0 {
		m.RemoveCharacterSkillIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCharacterInput on the CharacterUpdate builder.
func (c *CharacterUpdate) SetInput(i UpdateCharacterInput) *CharacterUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCharacterInput on the CharacterUpdateOne builder.
func (c *CharacterUpdateOne) SetInput(i UpdateCharacterInput) *CharacterUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
