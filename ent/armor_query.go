// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/armorclass"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ArmorQuery is the builder for querying Armor entities.
type ArmorQuery struct {
	config
	ctx            *QueryContext
	order          []armor.OrderOption
	inters         []Interceptor
	predicates     []predicate.Armor
	withArmorClass *ArmorClassQuery
	withEquipment  *EquipmentQuery
	withFKs        bool
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*Armor) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArmorQuery builder.
func (aq *ArmorQuery) Where(ps ...predicate.Armor) *ArmorQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ArmorQuery) Limit(limit int) *ArmorQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ArmorQuery) Offset(offset int) *ArmorQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ArmorQuery) Unique(unique bool) *ArmorQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ArmorQuery) Order(o ...armor.OrderOption) *ArmorQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryArmorClass chains the current query on the "armor_class" edge.
func (aq *ArmorQuery) QueryArmorClass() *ArmorClassQuery {
	query := (&ArmorClassClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(armor.Table, armor.FieldID, selector),
			sqlgraph.To(armorclass.Table, armorclass.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, armor.ArmorClassTable, armor.ArmorClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEquipment chains the current query on the "equipment" edge.
func (aq *ArmorQuery) QueryEquipment() *EquipmentQuery {
	query := (&EquipmentClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(armor.Table, armor.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, armor.EquipmentTable, armor.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Armor entity from the query.
// Returns a *NotFoundError when no Armor was found.
func (aq *ArmorQuery) First(ctx context.Context) (*Armor, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{armor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ArmorQuery) FirstX(ctx context.Context) *Armor {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Armor ID from the query.
// Returns a *NotFoundError when no Armor ID was found.
func (aq *ArmorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{armor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ArmorQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Armor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Armor entity is found.
// Returns a *NotFoundError when no Armor entities are found.
func (aq *ArmorQuery) Only(ctx context.Context) (*Armor, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{armor.Label}
	default:
		return nil, &NotSingularError{armor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ArmorQuery) OnlyX(ctx context.Context) *Armor {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Armor ID in the query.
// Returns a *NotSingularError when more than one Armor ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ArmorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{armor.Label}
	default:
		err = &NotSingularError{armor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ArmorQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Armors.
func (aq *ArmorQuery) All(ctx context.Context) ([]*Armor, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Armor, *ArmorQuery]()
	return withInterceptors[[]*Armor](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ArmorQuery) AllX(ctx context.Context) []*Armor {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Armor IDs.
func (aq *ArmorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(armor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ArmorQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ArmorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ArmorQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ArmorQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ArmorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ArmorQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArmorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ArmorQuery) Clone() *ArmorQuery {
	if aq == nil {
		return nil
	}
	return &ArmorQuery{
		config:         aq.config,
		ctx:            aq.ctx.Clone(),
		order:          append([]armor.OrderOption{}, aq.order...),
		inters:         append([]Interceptor{}, aq.inters...),
		predicates:     append([]predicate.Armor{}, aq.predicates...),
		withArmorClass: aq.withArmorClass.Clone(),
		withEquipment:  aq.withEquipment.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithArmorClass tells the query-builder to eager-load the nodes that are connected to
// the "armor_class" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArmorQuery) WithArmorClass(opts ...func(*ArmorClassQuery)) *ArmorQuery {
	query := (&ArmorClassClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withArmorClass = query
	return aq
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArmorQuery) WithEquipment(opts ...func(*EquipmentQuery)) *ArmorQuery {
	query := (&EquipmentClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withEquipment = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ArmorCategory armor.ArmorCategory `json:"armor_category,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Armor.Query().
//		GroupBy(armor.FieldArmorCategory).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ArmorQuery) GroupBy(field string, fields ...string) *ArmorGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArmorGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = armor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ArmorCategory armor.ArmorCategory `json:"armor_category,omitempty"`
//	}
//
//	client.Armor.Query().
//		Select(armor.FieldArmorCategory).
//		Scan(ctx, &v)
func (aq *ArmorQuery) Select(fields ...string) *ArmorSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ArmorSelect{ArmorQuery: aq}
	sbuild.label = armor.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArmorSelect configured with the given aggregations.
func (aq *ArmorQuery) Aggregate(fns ...AggregateFunc) *ArmorSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ArmorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !armor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ArmorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Armor, error) {
	var (
		nodes       = []*Armor{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withArmorClass != nil,
			aq.withEquipment != nil,
		}
	)
	if aq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, armor.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Armor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Armor{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withArmorClass; query != nil {
		if err := aq.loadArmorClass(ctx, query, nodes, nil,
			func(n *Armor, e *ArmorClass) { n.Edges.ArmorClass = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withEquipment; query != nil {
		if err := aq.loadEquipment(ctx, query, nodes, nil,
			func(n *Armor, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *ArmorQuery) loadArmorClass(ctx context.Context, query *ArmorClassQuery, nodes []*Armor, init func(*Armor), assign func(*Armor, *ArmorClass)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Armor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.ArmorClass(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(armor.ArmorClassColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.armor_armor_class
		if fk == nil {
			return fmt.Errorf(`foreign-key "armor_armor_class" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "armor_armor_class" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *ArmorQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Armor, init func(*Armor), assign func(*Armor, *Equipment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Armor)
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
	query.Where(equipment.IDIn(ids...))
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

func (aq *ArmorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ArmorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(armor.Table, armor.Columns, sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, armor.FieldID)
		for i := range fields {
			if fields[i] != armor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ArmorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(armor.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = armor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArmorGroupBy is the group-by builder for Armor entities.
type ArmorGroupBy struct {
	selector
	build *ArmorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ArmorGroupBy) Aggregate(fns ...AggregateFunc) *ArmorGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ArmorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArmorQuery, *ArmorGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ArmorGroupBy) sqlScan(ctx context.Context, root *ArmorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArmorSelect is the builder for selecting fields of Armor entities.
type ArmorSelect struct {
	*ArmorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ArmorSelect) Aggregate(fns ...AggregateFunc) *ArmorSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ArmorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArmorQuery, *ArmorSelect](ctx, as.ArmorQuery, as, as.inters, v)
}

func (as *ArmorSelect) sqlScan(ctx context.Context, root *ArmorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
