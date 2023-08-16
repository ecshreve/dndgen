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
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
	"github.com/ecshreve/dndgen/ent/weaponrange"
)

// WeaponQuery is the builder for querying Weapon entities.
type WeaponQuery struct {
	config
	ctx                      *QueryContext
	order                    []weapon.OrderOption
	inters                   []Interceptor
	predicates               []predicate.Weapon
	withRange                *WeaponRangeQuery
	withDamage               *WeaponDamageQuery
	withTwoHandedDamage      *WeaponDamageQuery
	withEquipment            *EquipmentQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*Weapon) error
	withNamedRange           map[string]*WeaponRangeQuery
	withNamedDamage          map[string]*WeaponDamageQuery
	withNamedTwoHandedDamage map[string]*WeaponDamageQuery
	withNamedEquipment       map[string]*EquipmentQuery
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

// QueryRange chains the current query on the "range" edge.
func (wq *WeaponQuery) QueryRange() *WeaponRangeQuery {
	query := (&WeaponRangeClient{config: wq.config}).Query()
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
			sqlgraph.To(weaponrange.Table, weaponrange.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, weapon.RangeTable, weapon.RangePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDamage chains the current query on the "damage" edge.
func (wq *WeaponQuery) QueryDamage() *WeaponDamageQuery {
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
			sqlgraph.To(weapondamage.Table, weapondamage.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, weapon.DamageTable, weapon.DamagePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTwoHandedDamage chains the current query on the "two_handed_damage" edge.
func (wq *WeaponQuery) QueryTwoHandedDamage() *WeaponDamageQuery {
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
			sqlgraph.To(weapondamage.Table, weapondamage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, weapon.TwoHandedDamageTable, weapon.TwoHandedDamageColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
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
			sqlgraph.Edge(sqlgraph.M2M, true, weapon.EquipmentTable, weapon.EquipmentPrimaryKey...),
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
		config:              wq.config,
		ctx:                 wq.ctx.Clone(),
		order:               append([]weapon.OrderOption{}, wq.order...),
		inters:              append([]Interceptor{}, wq.inters...),
		predicates:          append([]predicate.Weapon{}, wq.predicates...),
		withRange:           wq.withRange.Clone(),
		withDamage:          wq.withDamage.Clone(),
		withTwoHandedDamage: wq.withTwoHandedDamage.Clone(),
		withEquipment:       wq.withEquipment.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithRange tells the query-builder to eager-load the nodes that are connected to
// the "range" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithRange(opts ...func(*WeaponRangeQuery)) *WeaponQuery {
	query := (&WeaponRangeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withRange = query
	return wq
}

// WithDamage tells the query-builder to eager-load the nodes that are connected to
// the "damage" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithDamage(opts ...func(*WeaponDamageQuery)) *WeaponQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withDamage = query
	return wq
}

// WithTwoHandedDamage tells the query-builder to eager-load the nodes that are connected to
// the "two_handed_damage" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithTwoHandedDamage(opts ...func(*WeaponDamageQuery)) *WeaponQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withTwoHandedDamage = query
	return wq
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

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Properties string `json:"properties,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Weapon.Query().
//		GroupBy(weapon.FieldProperties).
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
//		Properties string `json:"properties,omitempty"`
//	}
//
//	client.Weapon.Query().
//		Select(weapon.FieldProperties).
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
			wq.withRange != nil,
			wq.withDamage != nil,
			wq.withTwoHandedDamage != nil,
			wq.withEquipment != nil,
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
	if query := wq.withRange; query != nil {
		if err := wq.loadRange(ctx, query, nodes,
			func(n *Weapon) { n.Edges.Range = []*WeaponRange{} },
			func(n *Weapon, e *WeaponRange) { n.Edges.Range = append(n.Edges.Range, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withDamage; query != nil {
		if err := wq.loadDamage(ctx, query, nodes,
			func(n *Weapon) { n.Edges.Damage = []*WeaponDamage{} },
			func(n *Weapon, e *WeaponDamage) { n.Edges.Damage = append(n.Edges.Damage, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withTwoHandedDamage; query != nil {
		if err := wq.loadTwoHandedDamage(ctx, query, nodes,
			func(n *Weapon) { n.Edges.TwoHandedDamage = []*WeaponDamage{} },
			func(n *Weapon, e *WeaponDamage) { n.Edges.TwoHandedDamage = append(n.Edges.TwoHandedDamage, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withEquipment; query != nil {
		if err := wq.loadEquipment(ctx, query, nodes,
			func(n *Weapon) { n.Edges.Equipment = []*Equipment{} },
			func(n *Weapon, e *Equipment) { n.Edges.Equipment = append(n.Edges.Equipment, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedRange {
		if err := wq.loadRange(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedRange(name) },
			func(n *Weapon, e *WeaponRange) { n.appendNamedRange(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedDamage {
		if err := wq.loadDamage(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedDamage(name) },
			func(n *Weapon, e *WeaponDamage) { n.appendNamedDamage(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedTwoHandedDamage {
		if err := wq.loadTwoHandedDamage(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedTwoHandedDamage(name) },
			func(n *Weapon, e *WeaponDamage) { n.appendNamedTwoHandedDamage(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedEquipment {
		if err := wq.loadEquipment(ctx, query, nodes,
			func(n *Weapon) { n.appendNamedEquipment(name) },
			func(n *Weapon, e *Equipment) { n.appendNamedEquipment(name, e) }); err != nil {
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

func (wq *WeaponQuery) loadRange(ctx context.Context, query *WeaponRangeQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *WeaponRange)) error {
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
		joinT := sql.Table(weapon.RangeTable)
		s.Join(joinT).On(s.C(weaponrange.FieldID), joinT.C(weapon.RangePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(weapon.RangePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(weapon.RangePrimaryKey[0]))
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
	neighbors, err := withInterceptors[[]*WeaponRange](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "range" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WeaponQuery) loadDamage(ctx context.Context, query *WeaponDamageQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *WeaponDamage)) error {
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
		joinT := sql.Table(weapon.DamageTable)
		s.Join(joinT).On(s.C(weapondamage.FieldID), joinT.C(weapon.DamagePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(weapon.DamagePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(weapon.DamagePrimaryKey[0]))
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
	neighbors, err := withInterceptors[[]*WeaponDamage](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "damage" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WeaponQuery) loadTwoHandedDamage(ctx context.Context, query *WeaponDamageQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *WeaponDamage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Weapon)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WeaponDamage(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(weapon.TwoHandedDamageColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.weapon_two_handed_damage
		if fk == nil {
			return fmt.Errorf(`foreign-key "weapon_two_handed_damage" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "weapon_two_handed_damage" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WeaponQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Weapon, init func(*Weapon), assign func(*Weapon, *Equipment)) error {
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
		joinT := sql.Table(weapon.EquipmentTable)
		s.Join(joinT).On(s.C(equipment.FieldID), joinT.C(weapon.EquipmentPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(weapon.EquipmentPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(weapon.EquipmentPrimaryKey[1]))
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
	neighbors, err := withInterceptors[[]*Equipment](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "equipment" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
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

// WithNamedRange tells the query-builder to eager-load the nodes that are connected to the "range"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedRange(name string, opts ...func(*WeaponRangeQuery)) *WeaponQuery {
	query := (&WeaponRangeClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedRange == nil {
		wq.withNamedRange = make(map[string]*WeaponRangeQuery)
	}
	wq.withNamedRange[name] = query
	return wq
}

// WithNamedDamage tells the query-builder to eager-load the nodes that are connected to the "damage"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedDamage(name string, opts ...func(*WeaponDamageQuery)) *WeaponQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedDamage == nil {
		wq.withNamedDamage = make(map[string]*WeaponDamageQuery)
	}
	wq.withNamedDamage[name] = query
	return wq
}

// WithNamedTwoHandedDamage tells the query-builder to eager-load the nodes that are connected to the "two_handed_damage"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedTwoHandedDamage(name string, opts ...func(*WeaponDamageQuery)) *WeaponQuery {
	query := (&WeaponDamageClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedTwoHandedDamage == nil {
		wq.withNamedTwoHandedDamage = make(map[string]*WeaponDamageQuery)
	}
	wq.withNamedTwoHandedDamage[name] = query
	return wq
}

// WithNamedEquipment tells the query-builder to eager-load the nodes that are connected to the "equipment"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WeaponQuery) WithNamedEquipment(name string, opts ...func(*EquipmentQuery)) *WeaponQuery {
	query := (&EquipmentClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedEquipment == nil {
		wq.withNamedEquipment = make(map[string]*EquipmentQuery)
	}
	wq.withNamedEquipment[name] = query
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
