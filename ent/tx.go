// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// AbilityBonus is the client for interacting with the AbilityBonus builders.
	AbilityBonus *AbilityBonusClient
	// AbilityScore is the client for interacting with the AbilityScore builders.
	AbilityScore *AbilityScoreClient
	// Alignment is the client for interacting with the Alignment builders.
	Alignment *AlignmentClient
	// Armor is the client for interacting with the Armor builders.
	Armor *ArmorClient
	// Character is the client for interacting with the Character builders.
	Character *CharacterClient
	// CharacterAbilityScore is the client for interacting with the CharacterAbilityScore builders.
	CharacterAbilityScore *CharacterAbilityScoreClient
	// CharacterProficiency is the client for interacting with the CharacterProficiency builders.
	CharacterProficiency *CharacterProficiencyClient
	// CharacterSkill is the client for interacting with the CharacterSkill builders.
	CharacterSkill *CharacterSkillClient
	// Class is the client for interacting with the Class builders.
	Class *ClassClient
	// Coin is the client for interacting with the Coin builders.
	Coin *CoinClient
	// Condition is the client for interacting with the Condition builders.
	Condition *ConditionClient
	// Cost is the client for interacting with the Cost builders.
	Cost *CostClient
	// DamageType is the client for interacting with the DamageType builders.
	DamageType *DamageTypeClient
	// Equipment is the client for interacting with the Equipment builders.
	Equipment *EquipmentClient
	// EquipmentEntry is the client for interacting with the EquipmentEntry builders.
	EquipmentEntry *EquipmentEntryClient
	// Feat is the client for interacting with the Feat builders.
	Feat *FeatClient
	// Feature is the client for interacting with the Feature builders.
	Feature *FeatureClient
	// Gear is the client for interacting with the Gear builders.
	Gear *GearClient
	// Language is the client for interacting with the Language builders.
	Language *LanguageClient
	// LanguageChoice is the client for interacting with the LanguageChoice builders.
	LanguageChoice *LanguageChoiceClient
	// MagicSchool is the client for interacting with the MagicSchool builders.
	MagicSchool *MagicSchoolClient
	// Prerequisite is the client for interacting with the Prerequisite builders.
	Prerequisite *PrerequisiteClient
	// Proficiency is the client for interacting with the Proficiency builders.
	Proficiency *ProficiencyClient
	// ProficiencyChoice is the client for interacting with the ProficiencyChoice builders.
	ProficiencyChoice *ProficiencyChoiceClient
	// Property is the client for interacting with the Property builders.
	Property *PropertyClient
	// Race is the client for interacting with the Race builders.
	Race *RaceClient
	// Rule is the client for interacting with the Rule builders.
	Rule *RuleClient
	// RuleSection is the client for interacting with the RuleSection builders.
	RuleSection *RuleSectionClient
	// Skill is the client for interacting with the Skill builders.
	Skill *SkillClient
	// Tool is the client for interacting with the Tool builders.
	Tool *ToolClient
	// Trait is the client for interacting with the Trait builders.
	Trait *TraitClient
	// Vehicle is the client for interacting with the Vehicle builders.
	Vehicle *VehicleClient
	// Weapon is the client for interacting with the Weapon builders.
	Weapon *WeaponClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.AbilityBonus = NewAbilityBonusClient(tx.config)
	tx.AbilityScore = NewAbilityScoreClient(tx.config)
	tx.Alignment = NewAlignmentClient(tx.config)
	tx.Armor = NewArmorClient(tx.config)
	tx.Character = NewCharacterClient(tx.config)
	tx.CharacterAbilityScore = NewCharacterAbilityScoreClient(tx.config)
	tx.CharacterProficiency = NewCharacterProficiencyClient(tx.config)
	tx.CharacterSkill = NewCharacterSkillClient(tx.config)
	tx.Class = NewClassClient(tx.config)
	tx.Coin = NewCoinClient(tx.config)
	tx.Condition = NewConditionClient(tx.config)
	tx.Cost = NewCostClient(tx.config)
	tx.DamageType = NewDamageTypeClient(tx.config)
	tx.Equipment = NewEquipmentClient(tx.config)
	tx.EquipmentEntry = NewEquipmentEntryClient(tx.config)
	tx.Feat = NewFeatClient(tx.config)
	tx.Feature = NewFeatureClient(tx.config)
	tx.Gear = NewGearClient(tx.config)
	tx.Language = NewLanguageClient(tx.config)
	tx.LanguageChoice = NewLanguageChoiceClient(tx.config)
	tx.MagicSchool = NewMagicSchoolClient(tx.config)
	tx.Prerequisite = NewPrerequisiteClient(tx.config)
	tx.Proficiency = NewProficiencyClient(tx.config)
	tx.ProficiencyChoice = NewProficiencyChoiceClient(tx.config)
	tx.Property = NewPropertyClient(tx.config)
	tx.Race = NewRaceClient(tx.config)
	tx.Rule = NewRuleClient(tx.config)
	tx.RuleSection = NewRuleSectionClient(tx.config)
	tx.Skill = NewSkillClient(tx.config)
	tx.Tool = NewToolClient(tx.config)
	tx.Trait = NewTraitClient(tx.config)
	tx.Vehicle = NewVehicleClient(tx.config)
	tx.Weapon = NewWeaponClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: AbilityBonus.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
