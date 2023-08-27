// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
	"github.com/ecshreve/dndgen/ent/weaponproperty"
)

// WeaponQuery is the builder for querying Weapon entities.
type WeaponQuery struct {
	config
	ctx                       *QueryContext
	order                     []weapon.OrderOption
	inters                    []Interceptor
	predicates                []predicate.Weapon
	withEquipment             *EquipmentQuery
	withDamageType            *DamageTypeQuery
	withWeaponProperties      *WeaponPropertyQuery
	withWeaponDamage          *WeaponDamageQuery
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*Weapon) error
	withNamedDamageType       map[string]*DamageTypeQuery
	withNamedWeaponProperties map[string]*WeaponPropertyQuery
	withNamedWeaponDamage     map[string]*WeaponDamageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WeaponQuery builder.
func (wq *WeaponQuery) Where(ps ...predicate.Weapon) *WeaponQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WeaponQuery) Limit(limit int) *WeaponQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WeaponQuery) Offset(offset int) *WeaponQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WeaponQuery) Unique(unique bool) *WeaponQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WeaponQuery) Order(o ...weapon.OrderOption) *WeaponQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (wq *WeaponQuery) QueryEquipment() *EquipmentQuery {
	query := (&EquipmentClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weapon.Table, weapon.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, weapon.EquipmentTable, weapon.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDamageType chains the current query on the "damage_type" edge.
func (wq *WeaponQuery) QueryDamageType() *DamageTypeQuery {
	query := (&DamageTypeClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weapon.Table, weapon.FieldID, selector),
			sqlgraph.To(damagetype.Table, damagetype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, weapon.DamageTypeTable, weapon.DamageTypePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWeaponProperties chains the current query on the "weapon_properties" edge.
func (wq *WeaponQuery) QueryWeaponProperties() *WeaponPropertyQuery {
	query := (&WeaponPropertyClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weapon.Table, weapon.FieldID, selector),
			sqlgraph.To(weaponproperty.Table, weaponproperty.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, weapon.WeaponPropertiesTable, weapon.WeaponPropertiesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWeaponDamage chains the current query on the "weapon_damage" edge.
func (wq *WeaponQuery) QueryWeaponDamage() *WeaponDamageQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weapon.Table, weapon.FieldID, selector),
			sqlgraph.To(weapondamage.Table, weapondamage.WeaponColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, weapon.WeaponDamageTable, weapon.WeaponDamageColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Weapon entity from the query.
// Returns a *NotFoundError when no Weapon was found.
func (wq *WeaponQuery) First(ctx context.Context) (*Weapon, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{weapon.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WeaponQuery) FirstX(ctx context.Context) *Weapon {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Weapon ID from the query.
// Returns a *NotFoundError when no Weapon ID was found.
func (wq *WeaponQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{weapon.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WeaponQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Weapon entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Weapon entity is found.
// Returns a *NotFoundError when no Weapon entities are found.
func (wq *WeaponQuery) Only(ctx context.Context) (*Weapon, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{weapon.Label}
	default:
		return nil, &NotSingularError{weapon.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WeaponQuery) OnlyX(ctx context.Context) *Weapon {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Weapon ID in the query.
// Returns a *NotSingularError when more than one Weapon ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WeaponQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{weapon.Label}
	default:
		err = &NotSingularError{weapon.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WeaponQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Weapons.
func (wq *WeaponQuery) All(ctx context.Context) ([]*Weapon, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Weapon, *WeaponQuery]()
	return withInterceptors[[]*Weapon](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WeaponQuery) AllX(ctx context.Context) []*Weapon {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Weapon IDs.
func (wq *WeaponQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(weapon.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WeaponQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WeaponQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WeaponQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WeaponQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WeaponQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WeaponQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WeaponQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WeaponQuery) Clone() *WeaponQuery {
	if wq == nil {
		return nil
	}
	return &WeaponQuery{
		config:               wq.config,
		ctx:                  wq.ctx.Clone(),
		order:                append([]weapon.OrderOption{}, wq.order...),
		inters:               append([]Interceptor{}, wq.inters...),
		predicates:           append([]predicate.Weapon{}, wq.predicates...),
		withEquipment:        wq.withEquipment.Clone(),
		withDamageType:       wq.withDamageType.Clone(),
		withWeaponProperties: wq.withWeaponProperties.Clone(),
		withWeaponDamage:     wq.withWeaponDamage.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithEquipment(opts ...func(*EquipmentQuery)) *WeaponQuery {
	query := (&EquipmentClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withEquipment = query
	return wq
}

// WithDamageType tells the query-builder to eager-load the nodes that are connected to
// the "damage_type" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithDamageType(opts ...func(*DamageTypeQuery)) *WeaponQuery {
	query := (&DamageTypeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withDamageType = query
	return wq
}

// WithWeaponProperties tells the query-builder to eager-load the nodes that are connected to
// the "weapon_properties" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithWeaponProperties(opts ...func(*WeaponPropertyQuery)) *WeaponQuery {
	query := (&WeaponPropertyClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWeaponProperties = query
	return wq
}

// WithWeaponDamage tells the query-builder to eager-load the nodes that are connected to
// the "weapon_damage" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithWeaponDamage(opts ...func(*WeaponDamageQuery)) *WeaponQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withWeaponDamage = query
	return wq
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
//	client.Weapon.Query().
//		GroupBy(weapon.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WeaponQuery) GroupBy(field string, fields ...string) *WeaponGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WeaponGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = weapon.Label
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
//	client.Weapon.Query().
//		Select(weapon.FieldIndx).
//		Scan(ctx, &v)
func (wq *WeaponQuery) Select(fields ...string) *WeaponSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WeaponSelect{WeaponQuery: wq}
	sbuild.label = weapon.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WeaponSelect configured with the given aggregations.
func (wq *WeaponQuery) Aggregate(fns ...AggregateFunc) *WeaponSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WeaponQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !weapon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WeaponQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Weapon, error) {
	var (
		nodes       = []*Weapon{}
		_spec       = wq.querySpec()
		loadedTypes = [4]bool{
			wq.withEquipment != nil,
			wq.withDamageType != nil,
			wq.withWeaponProperties != nil,
			wq.withWeaponDamage != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Weapon).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Weapon{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withEquipment; query != nil {
		if err := wq.loadEquipment(ctx, query, nodes, nil,
			func(n *Weapon, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withDamageType; query != nil {
		if err := wq.loadDamageType(ctx, query, nodes,
			func(n *Weapon) { n.Edges.DamageType = []*DamageType{} },
			func(n *Weapon, e *DamageType) { n.Edges.DamageType = append(n.Edges.DamageType, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWeaponProperties; query != nil {
		if err := wq.loadWeaponProperties(ctx, query, nodes,
			func(n *Weapon) { n.Edges.WeaponProperties = []*WeaponProperty{} },
			func(n *Weapon, e *WeaponProperty) { n.Edges.WeaponProperties = append(n.Edges.WeaponProperties, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withWeaponDamage; query != nil {
		if err := wq.loadWeaponDamage(ctx, query, nodes,
			func(n *Weapon) { n.Edges.WeaponDamage = []*WeaponDamage{} },
			func(n *Weapon, e *WeaponDamage) { n.Edges.WeaponDamage = append(n.Edges.WeaponDamage, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedDamageType {
		if err := wq.loadDamageType(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedDamageType(name) },
			func(n *Weapon, e *DamageType) { n.appendNamedDamageType(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedWeaponProperties {
		if err := wq.loadWeaponProperties(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedWeaponProperties(name) },
			func(n *Weapon, e *WeaponProperty) { n.appendNamedWeaponProperties(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedWeaponDamage {
		if err := wq.loadWeaponDamage(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedWeaponDamage(name) },
			func(n *Weapon, e *WeaponDamage) { n.appendNamedWeaponDamage(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range wq.loadTotal {
		if err := wq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WeaponQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *Equipment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Weapon)
	for i := range nodes {
		fk := nodes[i].EquipmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(equipment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WeaponQuery) loadDamageType(ctx context.Context, query *DamageTypeQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *DamageType)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Weapon)
	nids := make(map[int]map[*Weapon]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(weapon.DamageTypeTable)
		s.Join(joinT).On(s.C(damagetype.FieldID), joinT.C(weapon.DamageTypePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(weapon.DamageTypePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(weapon.DamageTypePrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Weapon]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*DamageType](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "damage_type" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WeaponQuery) loadWeaponProperties(ctx context.Context, query *WeaponPropertyQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *WeaponProperty)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Weapon)
	nids := make(map[int]map[*Weapon]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(weapon.WeaponPropertiesTable)
		s.Join(joinT).On(s.C(weaponproperty.FieldID), joinT.C(weapon.WeaponPropertiesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(weapon.WeaponPropertiesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(weapon.WeaponPropertiesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Weapon]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*WeaponProperty](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "weapon_properties" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WeaponQuery) loadWeaponDamage(ctx context.Context, query *WeaponDamageQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *WeaponDamage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Weapon)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(weapondamage.FieldWeaponID)
	}
	query.Where(predicate.WeaponDamage(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(weapon.WeaponDamageColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WeaponID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "weapon_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WeaponQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WeaponQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(weapon.Table, weapon.Columns, sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, weapon.FieldID)
		for i := range fields {
			if fields[i] != weapon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wq.withEquipment != nil {
			_spec.Node.AddColumnOnce(weapon.FieldEquipmentID)
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WeaponQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(weapon.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = weapon.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedDamageType tells the query-builder to eager-load the nodes that are connected to the "damage_type"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedDamageType(name string, opts ...func(*DamageTypeQuery)) *WeaponQuery {
	query := (&DamageTypeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedDamageType == nil {
		wq.withNamedDamageType = make(map[string]*DamageTypeQuery)
	}
	wq.withNamedDamageType[name] = query
	return wq
}

// WithNamedWeaponProperties tells the query-builder to eager-load the nodes that are connected to the "weapon_properties"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedWeaponProperties(name string, opts ...func(*WeaponPropertyQuery)) *WeaponQuery {
	query := (&WeaponPropertyClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedWeaponProperties == nil {
		wq.withNamedWeaponProperties = make(map[string]*WeaponPropertyQuery)
	}
	wq.withNamedWeaponProperties[name] = query
	return wq
}

// WithNamedWeaponDamage tells the query-builder to eager-load the nodes that are connected to the "weapon_damage"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedWeaponDamage(name string, opts ...func(*WeaponDamageQuery)) *WeaponQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedWeaponDamage == nil {
		wq.withNamedWeaponDamage = make(map[string]*WeaponDamageQuery)
	}
	wq.withNamedWeaponDamage[name] = query
	return wq
}

// WeaponGroupBy is the group-by builder for Weapon entities.
type WeaponGroupBy struct {
	selector
	build *WeaponQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WeaponGroupBy) Aggregate(fns ...AggregateFunc) *WeaponGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WeaponGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WeaponQuery, *WeaponGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WeaponGroupBy) sqlScan(ctx context.Context, root *WeaponQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WeaponSelect is the builder for selecting fields of Weapon entities.
type WeaponSelect struct {
	*WeaponQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WeaponSelect) Aggregate(fns ...AggregateFunc) *WeaponSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WeaponSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WeaponQuery, *WeaponSelect](ctx, ws.WeaponQuery, ws, ws.inters, v)
}

func (ws *WeaponSelect) sqlScan(ctx context.Context, root *WeaponQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
