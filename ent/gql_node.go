// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/armorclass"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentcategory"
	"github.com/ecshreve/dndgen/ent/equipmentchoice"
	"github.com/ecshreve/dndgen/ent/equipmentcost"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/magicschool"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/rule"
	"github.com/ecshreve/dndgen/ent/rulesection"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/tool"
	"github.com/ecshreve/dndgen/ent/trait"
	"github.com/ecshreve/dndgen/ent/vehicle"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
	"github.com/ecshreve/dndgen/ent/weaponproperty"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *AbilityBonus) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *AbilityScore) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Armor) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *ArmorClass) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Class) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Coin) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *DamageType) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Equipment) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *EquipmentCategory) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *EquipmentChoice) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *EquipmentCost) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Gear) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Language) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *MagicSchool) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Proficiency) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *ProficiencyChoice) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Race) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Rule) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *RuleSection) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Skill) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Subrace) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Tool) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Trait) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Vehicle) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Weapon) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *WeaponDamage) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *WeaponProperty) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case abilitybonus.Table:
		query := c.AbilityBonus.Query().
			Where(abilitybonus.ID(id))
		query, err := query.CollectFields(ctx, "AbilityBonus")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case abilityscore.Table:
		query := c.AbilityScore.Query().
			Where(abilityscore.ID(id))
		query, err := query.CollectFields(ctx, "AbilityScore")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case armor.Table:
		query := c.Armor.Query().
			Where(armor.ID(id))
		query, err := query.CollectFields(ctx, "Armor")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case armorclass.Table:
		query := c.ArmorClass.Query().
			Where(armorclass.ID(id))
		query, err := query.CollectFields(ctx, "ArmorClass")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case class.Table:
		query := c.Class.Query().
			Where(class.ID(id))
		query, err := query.CollectFields(ctx, "Class")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case coin.Table:
		query := c.Coin.Query().
			Where(coin.ID(id))
		query, err := query.CollectFields(ctx, "Coin")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case damagetype.Table:
		query := c.DamageType.Query().
			Where(damagetype.ID(id))
		query, err := query.CollectFields(ctx, "DamageType")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case equipment.Table:
		query := c.Equipment.Query().
			Where(equipment.ID(id))
		query, err := query.CollectFields(ctx, "Equipment")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case equipmentcategory.Table:
		query := c.EquipmentCategory.Query().
			Where(equipmentcategory.ID(id))
		query, err := query.CollectFields(ctx, "EquipmentCategory")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case equipmentchoice.Table:
		query := c.EquipmentChoice.Query().
			Where(equipmentchoice.ID(id))
		query, err := query.CollectFields(ctx, "EquipmentChoice")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case equipmentcost.Table:
		query := c.EquipmentCost.Query().
			Where(equipmentcost.ID(id))
		query, err := query.CollectFields(ctx, "EquipmentCost")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case gear.Table:
		query := c.Gear.Query().
			Where(gear.ID(id))
		query, err := query.CollectFields(ctx, "Gear")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case language.Table:
		query := c.Language.Query().
			Where(language.ID(id))
		query, err := query.CollectFields(ctx, "Language")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case magicschool.Table:
		query := c.MagicSchool.Query().
			Where(magicschool.ID(id))
		query, err := query.CollectFields(ctx, "MagicSchool")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case proficiency.Table:
		query := c.Proficiency.Query().
			Where(proficiency.ID(id))
		query, err := query.CollectFields(ctx, "Proficiency")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case proficiencychoice.Table:
		query := c.ProficiencyChoice.Query().
			Where(proficiencychoice.ID(id))
		query, err := query.CollectFields(ctx, "ProficiencyChoice")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case race.Table:
		query := c.Race.Query().
			Where(race.ID(id))
		query, err := query.CollectFields(ctx, "Race")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case rule.Table:
		query := c.Rule.Query().
			Where(rule.ID(id))
		query, err := query.CollectFields(ctx, "Rule")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case rulesection.Table:
		query := c.RuleSection.Query().
			Where(rulesection.ID(id))
		query, err := query.CollectFields(ctx, "RuleSection")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case skill.Table:
		query := c.Skill.Query().
			Where(skill.ID(id))
		query, err := query.CollectFields(ctx, "Skill")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subrace.Table:
		query := c.Subrace.Query().
			Where(subrace.ID(id))
		query, err := query.CollectFields(ctx, "Subrace")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tool.Table:
		query := c.Tool.Query().
			Where(tool.ID(id))
		query, err := query.CollectFields(ctx, "Tool")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case trait.Table:
		query := c.Trait.Query().
			Where(trait.ID(id))
		query, err := query.CollectFields(ctx, "Trait")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case vehicle.Table:
		query := c.Vehicle.Query().
			Where(vehicle.ID(id))
		query, err := query.CollectFields(ctx, "Vehicle")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case weapon.Table:
		query := c.Weapon.Query().
			Where(weapon.ID(id))
		query, err := query.CollectFields(ctx, "Weapon")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case weapondamage.Table:
		query := c.WeaponDamage.Query().
			Where(weapondamage.ID(id))
		query, err := query.CollectFields(ctx, "WeaponDamage")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case weaponproperty.Table:
		query := c.WeaponProperty.Query().
			Where(weaponproperty.ID(id))
		query, err := query.CollectFields(ctx, "WeaponProperty")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case abilitybonus.Table:
		query := c.AbilityBonus.Query().
			Where(abilitybonus.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AbilityBonus")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case abilityscore.Table:
		query := c.AbilityScore.Query().
			Where(abilityscore.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AbilityScore")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case armor.Table:
		query := c.Armor.Query().
			Where(armor.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Armor")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case armorclass.Table:
		query := c.ArmorClass.Query().
			Where(armorclass.IDIn(ids...))
		query, err := query.CollectFields(ctx, "ArmorClass")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case class.Table:
		query := c.Class.Query().
			Where(class.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Class")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case coin.Table:
		query := c.Coin.Query().
			Where(coin.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Coin")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case damagetype.Table:
		query := c.DamageType.Query().
			Where(damagetype.IDIn(ids...))
		query, err := query.CollectFields(ctx, "DamageType")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case equipment.Table:
		query := c.Equipment.Query().
			Where(equipment.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Equipment")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case equipmentcategory.Table:
		query := c.EquipmentCategory.Query().
			Where(equipmentcategory.IDIn(ids...))
		query, err := query.CollectFields(ctx, "EquipmentCategory")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case equipmentchoice.Table:
		query := c.EquipmentChoice.Query().
			Where(equipmentchoice.IDIn(ids...))
		query, err := query.CollectFields(ctx, "EquipmentChoice")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case equipmentcost.Table:
		query := c.EquipmentCost.Query().
			Where(equipmentcost.IDIn(ids...))
		query, err := query.CollectFields(ctx, "EquipmentCost")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case gear.Table:
		query := c.Gear.Query().
			Where(gear.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Gear")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case language.Table:
		query := c.Language.Query().
			Where(language.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Language")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case magicschool.Table:
		query := c.MagicSchool.Query().
			Where(magicschool.IDIn(ids...))
		query, err := query.CollectFields(ctx, "MagicSchool")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case proficiency.Table:
		query := c.Proficiency.Query().
			Where(proficiency.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Proficiency")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case proficiencychoice.Table:
		query := c.ProficiencyChoice.Query().
			Where(proficiencychoice.IDIn(ids...))
		query, err := query.CollectFields(ctx, "ProficiencyChoice")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case race.Table:
		query := c.Race.Query().
			Where(race.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Race")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case rule.Table:
		query := c.Rule.Query().
			Where(rule.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Rule")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case rulesection.Table:
		query := c.RuleSection.Query().
			Where(rulesection.IDIn(ids...))
		query, err := query.CollectFields(ctx, "RuleSection")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case skill.Table:
		query := c.Skill.Query().
			Where(skill.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Skill")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subrace.Table:
		query := c.Subrace.Query().
			Where(subrace.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Subrace")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tool.Table:
		query := c.Tool.Query().
			Where(tool.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Tool")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case trait.Table:
		query := c.Trait.Query().
			Where(trait.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Trait")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case vehicle.Table:
		query := c.Vehicle.Query().
			Where(vehicle.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Vehicle")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case weapon.Table:
		query := c.Weapon.Query().
			Where(weapon.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Weapon")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case weapondamage.Table:
		query := c.WeaponDamage.Query().
			Where(weapondamage.IDIn(ids...))
		query, err := query.CollectFields(ctx, "WeaponDamage")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case weaponproperty.Table:
		query := c.WeaponProperty.Query().
			Where(weaponproperty.IDIn(ids...))
		query, err := query.CollectFields(ctx, "WeaponProperty")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
