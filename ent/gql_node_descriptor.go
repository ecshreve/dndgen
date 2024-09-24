// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"

	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/cost"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/property"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/rule"
	"github.com/ecshreve/dndgen/ent/rulesection"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/ent/tool"
	"github.com/ecshreve/dndgen/ent/vehicle"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

// Node implements Noder interface
func (ab *AbilityBonus) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ab.ID,
		Type:   "AbilityBonus",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(ab.Bonus); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "bonus",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "AbilityScore",
		Name: "ability_score",
	}
	err = ab.QueryAbilityScore().
		Select(abilityscore.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Race",
		Name: "race",
	}
	err = ab.QueryRace().
		Select(race.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (as *AbilityScore) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     as.ID,
		Type:   "AbilityScore",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(as.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.FullName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "full_name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Skill",
		Name: "skills",
	}
	err = as.QuerySkills().
		Select(skill.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "AbilityBonus",
		Name: "ability_bonuses",
	}
	err = as.QueryAbilityBonuses().
		Select(abilitybonus.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Class",
		Name: "classes",
	}
	err = as.QueryClasses().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Proficiency",
		Name: "proficiencies",
	}
	err = as.QueryProficiencies().
		Select(proficiency.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (a *Alignment) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Alignment",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(a.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Abbr); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "abbr",
		Value: string(buf),
	}
	return node, nil
}

// Node implements Noder interface
func (a *Armor) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Armor",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(a.ArmorCategory); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "armor.ArmorCategory",
		Name:  "armor_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.StrMinimum); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "str_minimum",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.StealthDisadvantage); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "bool",
		Name:  "stealth_disadvantage",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.AcBase); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "ac_base",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.AcDexBonus); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bool",
		Name:  "ac_dex_bonus",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.AcMaxBonus); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "ac_max_bonus",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = a.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (c *Class) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Class",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.HitDie); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "hit_die",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "AbilityScore",
		Name: "saving_throws",
	}
	err = c.QuerySavingThrows().
		Select(abilityscore.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Proficiency",
		Name: "proficiencies",
	}
	err = c.QueryProficiencies().
		Select(proficiency.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (c *Coin) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Coin",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.GoldConversionRate); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "float64",
		Name:  "gold_conversion_rate",
		Value: string(buf),
	}
	return node, nil
}

// Node implements Noder interface
func (c *Condition) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Condition",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	return node, nil
}

// Node implements Noder interface
func (c *Cost) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Cost",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Quantity); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "quantity",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Coin",
		Name: "coin",
	}
	err = c.QueryCoin().
		Select(coin.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = c.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (dt *DamageType) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     dt.ID,
		Type:   "DamageType",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(dt.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dt.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dt.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Weapon",
		Name: "weapons",
	}
	err = dt.QueryWeapons().
		Select(weapon.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (e *Equipment) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     e.ID,
		Type:   "Equipment",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 7),
	}
	var buf []byte
	if buf, err = json.Marshal(e.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.EquipmentCategory); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "equipment.EquipmentCategory",
		Name:  "equipment_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.Weight); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "float64",
		Name:  "weight",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Cost",
		Name: "cost",
	}
	err = e.QueryCost().
		Select(cost.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Gear",
		Name: "gear",
	}
	err = e.QueryGear().
		Select(gear.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Tool",
		Name: "tool",
	}
	err = e.QueryTool().
		Select(tool.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Weapon",
		Name: "weapon",
	}
	err = e.QueryWeapon().
		Select(weapon.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "Vehicle",
		Name: "vehicle",
	}
	err = e.QueryVehicle().
		Select(vehicle.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "Armor",
		Name: "armor",
	}
	err = e.QueryArmor().
		Select(armor.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[6] = &Edge{
		Type: "Proficiency",
		Name: "proficiencies",
	}
	err = e.QueryProficiencies().
		Select(proficiency.FieldID).
		Scan(ctx, &node.Edges[6].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (f *Feat) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Feat",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(f.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	return node, nil
}

// Node implements Noder interface
func (ge *Gear) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ge.ID,
		Type:   "Gear",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ge.GearCategory); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "gear_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.Desc); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = ge.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (l *Language) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     l.ID,
		Type:   "Language",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(l.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.LanguageType); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "language.LanguageType",
		Name:  "language_type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.Script); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "language.Script",
		Name:  "script",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Race",
		Name: "races",
	}
	err = l.QueryRaces().
		Select(race.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (ms *MagicSchool) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ms.ID,
		Type:   "MagicSchool",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(ms.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	return node, nil
}

// Node implements Noder interface
func (pr *Proficiency) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pr.ID,
		Type:   "Proficiency",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 5),
	}
	var buf []byte
	if buf, err = json.Marshal(pr.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Category); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "category",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = pr.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Skill",
		Name: "skill",
	}
	err = pr.QuerySkill().
		Select(skill.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "AbilityScore",
		Name: "saving_throw",
	}
	err = pr.QuerySavingThrow().
		Select(abilityscore.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Class",
		Name: "classes",
	}
	err = pr.QueryClasses().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "Race",
		Name: "races",
	}
	err = pr.QueryRaces().
		Select(race.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (pr *Property) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pr.ID,
		Type:   "Property",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(pr.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Weapon",
		Name: "weapons",
	}
	err = pr.QueryWeapons().
		Select(weapon.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (r *Race) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Race",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(r.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Speed); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "speed",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Size); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "race.Size",
		Name:  "size",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.SizeDesc); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "size_desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.AlignmentDesc); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "alignment_desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.AgeDesc); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "age_desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LanguageDesc); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "language_desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "AbilityBonus",
		Name: "ability_bonuses",
	}
	err = r.QueryAbilityBonuses().
		Select(abilitybonus.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Language",
		Name: "languages",
	}
	err = r.QueryLanguages().
		Select(language.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Proficiency",
		Name: "proficiencies",
	}
	err = r.QueryProficiencies().
		Select(proficiency.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (r *Rule) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Rule",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(r.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "RuleSection",
		Name: "sections",
	}
	err = r.QuerySections().
		Select(rulesection.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (rs *RuleSection) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     rs.ID,
		Type:   "RuleSection",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(rs.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(rs.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(rs.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Rule",
		Name: "rule",
	}
	err = rs.QueryRule().
		Select(rule.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (s *Skill) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "Skill",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(s.Indx); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "indx",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "AbilityScore",
		Name: "ability_score",
	}
	err = s.QueryAbilityScore().
		Select(abilityscore.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Proficiency",
		Name: "proficiencies",
	}
	err = s.QueryProficiencies().
		Select(proficiency.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (t *Tool) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Tool",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(t.ToolCategory); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "tool_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Desc); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = t.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (v *Vehicle) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     v.ID,
		Type:   "Vehicle",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(v.VehicleCategory); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "vehicle.VehicleCategory",
		Name:  "vehicle_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(v.Capacity); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "capacity",
		Value: string(buf),
	}
	if buf, err = json.Marshal(v.Desc); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "desc",
		Value: string(buf),
	}
	if buf, err = json.Marshal(v.SpeedQuantity); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "float64",
		Name:  "speed_quantity",
		Value: string(buf),
	}
	if buf, err = json.Marshal(v.SpeedUnits); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "vehicle.SpeedUnits",
		Name:  "speed_units",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = v.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (w *Weapon) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     w.ID,
		Type:   "Weapon",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(w.WeaponCategory); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "weapon.WeaponCategory",
		Name:  "weapon_category",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.WeaponSubcategory); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "weapon.WeaponSubcategory",
		Name:  "weapon_subcategory",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.RangeNormal); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "range_normal",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.RangeLong); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "range_long",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.ThrowRangeNormal); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int",
		Name:  "throw_range_normal",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.ThrowRangeLong); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "throw_range_long",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.DamageDice); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "damage_dice",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Property",
		Name: "properties",
	}
	err = w.QueryProperties().
		Select(property.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "DamageType",
		Name: "damage_type",
	}
	err = w.QueryDamageType().
		Select(damagetype.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Equipment",
		Name: "equipment",
	}
	err = w.QueryEquipment().
		Select(equipment.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node returns the node with given global ID.
//
// This API helpful in case you want to build
// an administrator tool to browser all types in system.
func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}
