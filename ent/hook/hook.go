// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/ecshreve/dndgen/ent"
)

// The AbilityScoreFunc type is an adapter to allow the use of ordinary
// function as AbilityScore mutator.
type AbilityScoreFunc func(context.Context, *ent.AbilityScoreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AbilityScoreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AbilityScoreMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AbilityScoreMutation", m)
}

// The ArmorFunc type is an adapter to allow the use of ordinary
// function as Armor mutator.
type ArmorFunc func(context.Context, *ent.ArmorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArmorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArmorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArmorMutation", m)
}

// The ArmorClassFunc type is an adapter to allow the use of ordinary
// function as ArmorClass mutator.
type ArmorClassFunc func(context.Context, *ent.ArmorClassMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArmorClassFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArmorClassMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArmorClassMutation", m)
}

// The ClassFunc type is an adapter to allow the use of ordinary
// function as Class mutator.
type ClassFunc func(context.Context, *ent.ClassMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ClassFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ClassMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ClassMutation", m)
}

// The CostFunc type is an adapter to allow the use of ordinary
// function as Cost mutator.
type CostFunc func(context.Context, *ent.CostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CostMutation", m)
}

// The DamageTypeFunc type is an adapter to allow the use of ordinary
// function as DamageType mutator.
type DamageTypeFunc func(context.Context, *ent.DamageTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DamageTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DamageTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DamageTypeMutation", m)
}

// The EquipmentFunc type is an adapter to allow the use of ordinary
// function as Equipment mutator.
type EquipmentFunc func(context.Context, *ent.EquipmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EquipmentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentMutation", m)
}

// The EquipmentCategoryFunc type is an adapter to allow the use of ordinary
// function as EquipmentCategory mutator.
type EquipmentCategoryFunc func(context.Context, *ent.EquipmentCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EquipmentCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EquipmentCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EquipmentCategoryMutation", m)
}

// The GearFunc type is an adapter to allow the use of ordinary
// function as Gear mutator.
type GearFunc func(context.Context, *ent.GearMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GearFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GearMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GearMutation", m)
}

// The LanguageFunc type is an adapter to allow the use of ordinary
// function as Language mutator.
type LanguageFunc func(context.Context, *ent.LanguageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LanguageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LanguageMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LanguageMutation", m)
}

// The ProficiencyFunc type is an adapter to allow the use of ordinary
// function as Proficiency mutator.
type ProficiencyFunc func(context.Context, *ent.ProficiencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProficiencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProficiencyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProficiencyMutation", m)
}

// The RaceFunc type is an adapter to allow the use of ordinary
// function as Race mutator.
type RaceFunc func(context.Context, *ent.RaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RaceMutation", m)
}

// The SkillFunc type is an adapter to allow the use of ordinary
// function as Skill mutator.
type SkillFunc func(context.Context, *ent.SkillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SkillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SkillMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SkillMutation", m)
}

// The ToolFunc type is an adapter to allow the use of ordinary
// function as Tool mutator.
type ToolFunc func(context.Context, *ent.ToolMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ToolFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ToolMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ToolMutation", m)
}

// The VehicleFunc type is an adapter to allow the use of ordinary
// function as Vehicle mutator.
type VehicleFunc func(context.Context, *ent.VehicleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VehicleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VehicleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VehicleMutation", m)
}

// The WeaponFunc type is an adapter to allow the use of ordinary
// function as Weapon mutator.
type WeaponFunc func(context.Context, *ent.WeaponMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WeaponFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WeaponMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WeaponMutation", m)
}

// The WeaponDamageFunc type is an adapter to allow the use of ordinary
// function as WeaponDamage mutator.
type WeaponDamageFunc func(context.Context, *ent.WeaponDamageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WeaponDamageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WeaponDamageMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WeaponDamageMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
