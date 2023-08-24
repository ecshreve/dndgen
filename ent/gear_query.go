// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// GearQuery is the builder for querying Gear entities.
type GearQuery struct {
	config
	ctx           *QueryContext
	order         []gear.OrderOption
	inters        []Interceptor
	predicates    []predicate.Gear
	withEquipment *EquipmentQuery
	withFKs       bool
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*Gear) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GearQuery builder.
func (gq *GearQuery) Where(ps ...predicate.Gear) *GearQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GearQuery) Limit(limit int) *GearQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GearQuery) Offset(offset int) *GearQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GearQuery) Unique(unique bool) *GearQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GearQuery) Order(o ...gear.OrderOption) *GearQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (gq *GearQuery) QueryEquipment() *EquipmentQuery {
	query := (&EquipmentClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gear.Table, gear.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, gear.EquipmentTable, gear.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Gear entity from the query.
// Returns a *NotFoundError when no Gear was found.
func (gq *GearQuery) First(ctx context.Context) (*Gear, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gear.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GearQuery) FirstX(ctx context.Context) *Gear {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Gear ID from the query.
// Returns a *NotFoundError when no Gear ID was found.
func (gq *GearQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gear.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GearQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Gear entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Gear entity is found.
// Returns a *NotFoundError when no Gear entities are found.
func (gq *GearQuery) Only(ctx context.Context) (*Gear, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gear.Label}
	default:
		return nil, &NotSingularError{gear.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GearQuery) OnlyX(ctx context.Context) *Gear {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Gear ID in the query.
// Returns a *NotSingularError when more than one Gear ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GearQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gear.Label}
	default:
		err = &NotSingularError{gear.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GearQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Gears.
func (gq *GearQuery) All(ctx context.Context) ([]*Gear, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Gear, *GearQuery]()
	return withInterceptors[[]*Gear](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GearQuery) AllX(ctx context.Context) []*Gear {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Gear IDs.
func (gq *GearQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(gear.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GearQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GearQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GearQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GearQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GearQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GearQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GearQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GearQuery) Clone() *GearQuery {
	if gq == nil {
		return nil
	}
	return &GearQuery{
		config:        gq.config,
		ctx:           gq.ctx.Clone(),
		order:         append([]gear.OrderOption{}, gq.order...),
		inters:        append([]Interceptor{}, gq.inters...),
		predicates:    append([]predicate.Gear{}, gq.predicates...),
		withEquipment: gq.withEquipment.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GearQuery) WithEquipment(opts ...func(*EquipmentQuery)) *GearQuery {
	query := (&EquipmentClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withEquipment = query
	return gq
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
//	client.Gear.Query().
//		GroupBy(gear.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GearQuery) GroupBy(field string, fields ...string) *GearGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GearGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = gear.Label
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
//	client.Gear.Query().
//		Select(gear.FieldIndx).
//		Scan(ctx, &v)
func (gq *GearQuery) Select(fields ...string) *GearSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GearSelect{GearQuery: gq}
	sbuild.label = gear.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GearSelect configured with the given aggregations.
func (gq *GearQuery) Aggregate(fns ...AggregateFunc) *GearSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GearQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !gear.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GearQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Gear, error) {
	var (
		nodes       = []*Gear{}
		withFKs     = gq.withFKs
		_spec       = gq.querySpec()
		loadedTypes = [1]bool{
			gq.withEquipment != nil,
		}
	)
	if gq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, gear.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Gear).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Gear{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withEquipment; query != nil {
		if err := gq.loadEquipment(ctx, query, nodes, nil,
			func(n *Gear, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	for i := range gq.loadTotal {
		if err := gq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GearQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Gear, init func(*Gear), assign func(*Gear, *Equipment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Gear)
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
	query.Where(equipment.IDIn(ids...))
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

func (gq *GearQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GearQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gear.Table, gear.Columns, sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gear.FieldID)
		for i := range fields {
			if fields[i] != gear.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GearQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(gear.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = gear.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GearGroupBy is the group-by builder for Gear entities.
type GearGroupBy struct {
	selector
	build *GearQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GearGroupBy) Aggregate(fns ...AggregateFunc) *GearGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GearGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GearQuery, *GearGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GearGroupBy) sqlScan(ctx context.Context, root *GearQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GearSelect is the builder for selecting fields of Gear entities.
type GearSelect struct {
	*GearQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GearSelect) Aggregate(fns ...AggregateFunc) *GearSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GearSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GearQuery, *GearSelect](ctx, gs.GearQuery, gs, gs.inters, v)
}

func (gs *GearSelect) sqlScan(ctx context.Context, root *GearQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
