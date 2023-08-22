// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/cost"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/tool"
	"github.com/ecshreve/dndgen/ent/vehicle"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// EquipmentQuery is the builder for querying Equipment entities.
type EquipmentQuery struct {
	config
	ctx         *QueryContext
	order       []equipment.OrderOption
	inters      []Interceptor
	predicates  []predicate.Equipment
	withWeapon  *WeaponQuery
	withArmor   *ArmorQuery
	withGear    *GearQuery
	withTool    *ToolQuery
	withVehicle *VehicleQuery
	withCost    *CostQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*Equipment) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentQuery builder.
func (eq *EquipmentQuery) Where(ps ...predicate.Equipment) *EquipmentQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EquipmentQuery) Limit(limit int) *EquipmentQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EquipmentQuery) Offset(offset int) *EquipmentQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EquipmentQuery) Unique(unique bool) *EquipmentQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EquipmentQuery) Order(o ...equipment.OrderOption) *EquipmentQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryWeapon chains the current query on the "weapon" edge.
func (eq *EquipmentQuery) QueryWeapon() *WeaponQuery {
	query := (&WeaponClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, selector),
			sqlgraph.To(weapon.Table, weapon.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipment.WeaponTable, equipment.WeaponColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArmor chains the current query on the "armor" edge.
func (eq *EquipmentQuery) QueryArmor() *ArmorQuery {
	query := (&ArmorClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, selector),
			sqlgraph.To(armor.Table, armor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipment.ArmorTable, equipment.ArmorColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGear chains the current query on the "gear" edge.
func (eq *EquipmentQuery) QueryGear() *GearQuery {
	query := (&GearClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, selector),
			sqlgraph.To(gear.Table, gear.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipment.GearTable, equipment.GearColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTool chains the current query on the "tool" edge.
func (eq *EquipmentQuery) QueryTool() *ToolQuery {
	query := (&ToolClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, selector),
			sqlgraph.To(tool.Table, tool.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipment.ToolTable, equipment.ToolColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVehicle chains the current query on the "vehicle" edge.
func (eq *EquipmentQuery) QueryVehicle() *VehicleQuery {
	query := (&VehicleClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, selector),
			sqlgraph.To(vehicle.Table, vehicle.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipment.VehicleTable, equipment.VehicleColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCost chains the current query on the "cost" edge.
func (eq *EquipmentQuery) QueryCost() *CostQuery {
	query := (&CostClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipment.Table, equipment.FieldID, selector),
			sqlgraph.To(cost.Table, cost.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipment.CostTable, equipment.CostColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Equipment entity from the query.
// Returns a *NotFoundError when no Equipment was found.
func (eq *EquipmentQuery) First(ctx context.Context) (*Equipment, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EquipmentQuery) FirstX(ctx context.Context) *Equipment {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Equipment ID from the query.
// Returns a *NotFoundError when no Equipment ID was found.
func (eq *EquipmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EquipmentQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Equipment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Equipment entity is found.
// Returns a *NotFoundError when no Equipment entities are found.
func (eq *EquipmentQuery) Only(ctx context.Context) (*Equipment, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipment.Label}
	default:
		return nil, &NotSingularError{equipment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EquipmentQuery) OnlyX(ctx context.Context) *Equipment {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Equipment ID in the query.
// Returns a *NotSingularError when more than one Equipment ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EquipmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipment.Label}
	default:
		err = &NotSingularError{equipment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EquipmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentSlice.
func (eq *EquipmentQuery) All(ctx context.Context) ([]*Equipment, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Equipment, *EquipmentQuery]()
	return withInterceptors[[]*Equipment](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EquipmentQuery) AllX(ctx context.Context) []*Equipment {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Equipment IDs.
func (eq *EquipmentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(equipment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EquipmentQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EquipmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EquipmentQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EquipmentQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EquipmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EquipmentQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EquipmentQuery) Clone() *EquipmentQuery {
	if eq == nil {
		return nil
	}
	return &EquipmentQuery{
		config:      eq.config,
		ctx:         eq.ctx.Clone(),
		order:       append([]equipment.OrderOption{}, eq.order...),
		inters:      append([]Interceptor{}, eq.inters...),
		predicates:  append([]predicate.Equipment{}, eq.predicates...),
		withWeapon:  eq.withWeapon.Clone(),
		withArmor:   eq.withArmor.Clone(),
		withGear:    eq.withGear.Clone(),
		withTool:    eq.withTool.Clone(),
		withVehicle: eq.withVehicle.Clone(),
		withCost:    eq.withCost.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithWeapon tells the query-builder to eager-load the nodes that are connected to
// the "weapon" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EquipmentQuery) WithWeapon(opts ...func(*WeaponQuery)) *EquipmentQuery {
	query := (&WeaponClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withWeapon = query
	return eq
}

// WithArmor tells the query-builder to eager-load the nodes that are connected to
// the "armor" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EquipmentQuery) WithArmor(opts ...func(*ArmorQuery)) *EquipmentQuery {
	query := (&ArmorClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withArmor = query
	return eq
}

// WithGear tells the query-builder to eager-load the nodes that are connected to
// the "gear" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EquipmentQuery) WithGear(opts ...func(*GearQuery)) *EquipmentQuery {
	query := (&GearClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withGear = query
	return eq
}

// WithTool tells the query-builder to eager-load the nodes that are connected to
// the "tool" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EquipmentQuery) WithTool(opts ...func(*ToolQuery)) *EquipmentQuery {
	query := (&ToolClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withTool = query
	return eq
}

// WithVehicle tells the query-builder to eager-load the nodes that are connected to
// the "vehicle" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EquipmentQuery) WithVehicle(opts ...func(*VehicleQuery)) *EquipmentQuery {
	query := (&VehicleClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withVehicle = query
	return eq
}

// WithCost tells the query-builder to eager-load the nodes that are connected to
// the "cost" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EquipmentQuery) WithCost(opts ...func(*CostQuery)) *EquipmentQuery {
	query := (&CostClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withCost = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Indx string `json:"index"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Equipment.Query().
//		GroupBy(equipment.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EquipmentQuery) GroupBy(field string, fields ...string) *EquipmentGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EquipmentGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = equipment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Indx string `json:"index"`
//	}
//
//	client.Equipment.Query().
//		Select(equipment.FieldIndx).
//		Scan(ctx, &v)
func (eq *EquipmentQuery) Select(fields ...string) *EquipmentSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EquipmentSelect{EquipmentQuery: eq}
	sbuild.label = equipment.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EquipmentSelect configured with the given aggregations.
func (eq *EquipmentQuery) Aggregate(fns ...AggregateFunc) *EquipmentSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EquipmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !equipment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EquipmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Equipment, error) {
	var (
		nodes       = []*Equipment{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [6]bool{
			eq.withWeapon != nil,
			eq.withArmor != nil,
			eq.withGear != nil,
			eq.withTool != nil,
			eq.withVehicle != nil,
			eq.withCost != nil,
		}
	)
	if eq.withWeapon != nil || eq.withArmor != nil || eq.withGear != nil || eq.withTool != nil || eq.withVehicle != nil || eq.withCost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, equipment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Equipment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Equipment{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withWeapon; query != nil {
		if err := eq.loadWeapon(ctx, query, nodes, nil,
			func(n *Equipment, e *Weapon) { n.Edges.Weapon = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withArmor; query != nil {
		if err := eq.loadArmor(ctx, query, nodes, nil,
			func(n *Equipment, e *Armor) { n.Edges.Armor = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withGear; query != nil {
		if err := eq.loadGear(ctx, query, nodes, nil,
			func(n *Equipment, e *Gear) { n.Edges.Gear = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withTool; query != nil {
		if err := eq.loadTool(ctx, query, nodes, nil,
			func(n *Equipment, e *Tool) { n.Edges.Tool = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withVehicle; query != nil {
		if err := eq.loadVehicle(ctx, query, nodes, nil,
			func(n *Equipment, e *Vehicle) { n.Edges.Vehicle = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withCost; query != nil {
		if err := eq.loadCost(ctx, query, nodes, nil,
			func(n *Equipment, e *Cost) { n.Edges.Cost = e }); err != nil {
			return nil, err
		}
	}
	for i := range eq.loadTotal {
		if err := eq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EquipmentQuery) loadWeapon(ctx context.Context, query *WeaponQuery, nodes []*Equipment, init func(*Equipment), assign func(*Equipment, *Weapon)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Equipment)
	for i := range nodes {
		if nodes[i].equipment_weapon == nil {
			continue
		}
		fk := *nodes[i].equipment_weapon
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(weapon.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_weapon" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EquipmentQuery) loadArmor(ctx context.Context, query *ArmorQuery, nodes []*Equipment, init func(*Equipment), assign func(*Equipment, *Armor)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Equipment)
	for i := range nodes {
		if nodes[i].equipment_armor == nil {
			continue
		}
		fk := *nodes[i].equipment_armor
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(armor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_armor" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EquipmentQuery) loadGear(ctx context.Context, query *GearQuery, nodes []*Equipment, init func(*Equipment), assign func(*Equipment, *Gear)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Equipment)
	for i := range nodes {
		if nodes[i].equipment_gear == nil {
			continue
		}
		fk := *nodes[i].equipment_gear
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(gear.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_gear" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EquipmentQuery) loadTool(ctx context.Context, query *ToolQuery, nodes []*Equipment, init func(*Equipment), assign func(*Equipment, *Tool)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Equipment)
	for i := range nodes {
		if nodes[i].equipment_tool == nil {
			continue
		}
		fk := *nodes[i].equipment_tool
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tool.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_tool" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EquipmentQuery) loadVehicle(ctx context.Context, query *VehicleQuery, nodes []*Equipment, init func(*Equipment), assign func(*Equipment, *Vehicle)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Equipment)
	for i := range nodes {
		if nodes[i].equipment_vehicle == nil {
			continue
		}
		fk := *nodes[i].equipment_vehicle
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(vehicle.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_vehicle" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EquipmentQuery) loadCost(ctx context.Context, query *CostQuery, nodes []*Equipment, init func(*Equipment), assign func(*Equipment, *Cost)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Equipment)
	for i := range nodes {
		if nodes[i].equipment_cost == nil {
			continue
		}
		fk := *nodes[i].equipment_cost
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(cost.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_cost" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eq *EquipmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EquipmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(equipment.Table, equipment.Columns, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipment.FieldID)
		for i := range fields {
			if fields[i] != equipment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EquipmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(equipment.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = equipment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentGroupBy is the group-by builder for Equipment entities.
type EquipmentGroupBy struct {
	selector
	build *EquipmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EquipmentGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EquipmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EquipmentQuery, *EquipmentGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EquipmentGroupBy) sqlScan(ctx context.Context, root *EquipmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EquipmentSelect is the builder for selecting fields of Equipment entities.
type EquipmentSelect struct {
	*EquipmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EquipmentSelect) Aggregate(fns ...AggregateFunc) *EquipmentSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EquipmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EquipmentQuery, *EquipmentSelect](ctx, es.EquipmentQuery, es, es.inters, v)
}

func (es *EquipmentSelect) sqlScan(ctx context.Context, root *EquipmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
