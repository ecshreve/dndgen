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
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// DamageTypeQuery is the builder for querying DamageType entities.
type DamageTypeQuery struct {
	config
	ctx              *QueryContext
	order            []damagetype.OrderOption
	inters           []Interceptor
	predicates       []predicate.DamageType
	withWeapons      *WeaponQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*DamageType) error
	withNamedWeapons map[string]*WeaponQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DamageTypeQuery builder.
func (dtq *DamageTypeQuery) Where(ps ...predicate.DamageType) *DamageTypeQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit the number of records to be returned by this query.
func (dtq *DamageTypeQuery) Limit(limit int) *DamageTypeQuery {
	dtq.ctx.Limit = &limit
	return dtq
}

// Offset to start from.
func (dtq *DamageTypeQuery) Offset(offset int) *DamageTypeQuery {
	dtq.ctx.Offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DamageTypeQuery) Unique(unique bool) *DamageTypeQuery {
	dtq.ctx.Unique = &unique
	return dtq
}

// Order specifies how the records should be ordered.
func (dtq *DamageTypeQuery) Order(o ...damagetype.OrderOption) *DamageTypeQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// QueryWeapons chains the current query on the "weapons" edge.
func (dtq *DamageTypeQuery) QueryWeapons() *WeaponQuery {
	query := (&WeaponClient{config: dtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(damagetype.Table, damagetype.FieldID, selector),
			sqlgraph.To(weapon.Table, weapon.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, damagetype.WeaponsTable, damagetype.WeaponsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DamageType entity from the query.
// Returns a *NotFoundError when no DamageType was found.
func (dtq *DamageTypeQuery) First(ctx context.Context) (*DamageType, error) {
	nodes, err := dtq.Limit(1).All(setContextOp(ctx, dtq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{damagetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DamageTypeQuery) FirstX(ctx context.Context) *DamageType {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DamageType ID from the query.
// Returns a *NotFoundError when no DamageType ID was found.
func (dtq *DamageTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(1).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{damagetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DamageTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DamageType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DamageType entity is found.
// Returns a *NotFoundError when no DamageType entities are found.
func (dtq *DamageTypeQuery) Only(ctx context.Context) (*DamageType, error) {
	nodes, err := dtq.Limit(2).All(setContextOp(ctx, dtq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{damagetype.Label}
	default:
		return nil, &NotSingularError{damagetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DamageTypeQuery) OnlyX(ctx context.Context) *DamageType {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DamageType ID in the query.
// Returns a *NotSingularError when more than one DamageType ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DamageTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(2).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{damagetype.Label}
	default:
		err = &NotSingularError{damagetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DamageTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DamageTypes.
func (dtq *DamageTypeQuery) All(ctx context.Context) ([]*DamageType, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryAll)
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DamageType, *DamageTypeQuery]()
	return withInterceptors[[]*DamageType](ctx, dtq, qr, dtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DamageTypeQuery) AllX(ctx context.Context) []*DamageType {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DamageType IDs.
func (dtq *DamageTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dtq.ctx.Unique == nil && dtq.path != nil {
		dtq.Unique(true)
	}
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryIDs)
	if err = dtq.Select(damagetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DamageTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DamageTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryCount)
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dtq, querierCount[*DamageTypeQuery](), dtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DamageTypeQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DamageTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryExist)
	switch _, err := dtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DamageTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DamageTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DamageTypeQuery) Clone() *DamageTypeQuery {
	if dtq == nil {
		return nil
	}
	return &DamageTypeQuery{
		config:      dtq.config,
		ctx:         dtq.ctx.Clone(),
		order:       append([]damagetype.OrderOption{}, dtq.order...),
		inters:      append([]Interceptor{}, dtq.inters...),
		predicates:  append([]predicate.DamageType{}, dtq.predicates...),
		withWeapons: dtq.withWeapons.Clone(),
		// clone intermediate query.
		sql:  dtq.sql.Clone(),
		path: dtq.path,
	}
}

// WithWeapons tells the query-builder to eager-load the nodes that are connected to
// the "weapons" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DamageTypeQuery) WithWeapons(opts ...func(*WeaponQuery)) *DamageTypeQuery {
	query := (&WeaponClient{config: dtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dtq.withWeapons = query
	return dtq
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
//	client.DamageType.Query().
//		GroupBy(damagetype.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dtq *DamageTypeQuery) GroupBy(field string, fields ...string) *DamageTypeGroupBy {
	dtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DamageTypeGroupBy{build: dtq}
	grbuild.flds = &dtq.ctx.Fields
	grbuild.label = damagetype.Label
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
//	client.DamageType.Query().
//		Select(damagetype.FieldIndx).
//		Scan(ctx, &v)
func (dtq *DamageTypeQuery) Select(fields ...string) *DamageTypeSelect {
	dtq.ctx.Fields = append(dtq.ctx.Fields, fields...)
	sbuild := &DamageTypeSelect{DamageTypeQuery: dtq}
	sbuild.label = damagetype.Label
	sbuild.flds, sbuild.scan = &dtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DamageTypeSelect configured with the given aggregations.
func (dtq *DamageTypeQuery) Aggregate(fns ...AggregateFunc) *DamageTypeSelect {
	return dtq.Select().Aggregate(fns...)
}

func (dtq *DamageTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dtq); err != nil {
				return err
			}
		}
	}
	for _, f := range dtq.ctx.Fields {
		if !damagetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DamageTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DamageType, error) {
	var (
		nodes       = []*DamageType{}
		_spec       = dtq.querySpec()
		loadedTypes = [1]bool{
			dtq.withWeapons != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DamageType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DamageType{config: dtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dtq.modifiers) > 0 {
		_spec.Modifiers = dtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dtq.withWeapons; query != nil {
		if err := dtq.loadWeapons(ctx, query, nodes,
			func(n *DamageType) { n.Edges.Weapons = []*Weapon{} },
			func(n *DamageType, e *Weapon) { n.Edges.Weapons = append(n.Edges.Weapons, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dtq.withNamedWeapons {
		if err := dtq.loadWeapons(ctx, query, nodes,
			func(n *DamageType) { n.appendNamedWeapons(name) },
			func(n *DamageType, e *Weapon) { n.appendNamedWeapons(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range dtq.loadTotal {
		if err := dtq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dtq *DamageTypeQuery) loadWeapons(ctx context.Context, query *WeaponQuery, nodes []*DamageType, init func(*DamageType), assign func(*DamageType, *Weapon)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DamageType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Weapon(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(damagetype.WeaponsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.weapon_damage_type
		if fk == nil {
			return fmt.Errorf(`foreign-key "weapon_damage_type" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "weapon_damage_type" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dtq *DamageTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	if len(dtq.modifiers) > 0 {
		_spec.Modifiers = dtq.modifiers
	}
	_spec.Node.Columns = dtq.ctx.Fields
	if len(dtq.ctx.Fields) > 0 {
		_spec.Unique = dtq.ctx.Unique != nil && *dtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DamageTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(damagetype.Table, damagetype.Columns, sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt))
	_spec.From = dtq.sql
	if unique := dtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dtq.path != nil {
		_spec.Unique = true
	}
	if fields := dtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, damagetype.FieldID)
		for i := range fields {
			if fields[i] != damagetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DamageTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(damagetype.Table)
	columns := dtq.ctx.Fields
	if len(columns) == 0 {
		columns = damagetype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.ctx.Unique != nil && *dtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedWeapons tells the query-builder to eager-load the nodes that are connected to the "weapons"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dtq *DamageTypeQuery) WithNamedWeapons(name string, opts ...func(*WeaponQuery)) *DamageTypeQuery {
	query := (&WeaponClient{config: dtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dtq.withNamedWeapons == nil {
		dtq.withNamedWeapons = make(map[string]*WeaponQuery)
	}
	dtq.withNamedWeapons[name] = query
	return dtq
}

// DamageTypeGroupBy is the group-by builder for DamageType entities.
type DamageTypeGroupBy struct {
	selector
	build *DamageTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DamageTypeGroupBy) Aggregate(fns ...AggregateFunc) *DamageTypeGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the selector query and scans the result into the given value.
func (dtgb *DamageTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dtgb.build.ctx, ent.OpQueryGroupBy)
	if err := dtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DamageTypeQuery, *DamageTypeGroupBy](ctx, dtgb.build, dtgb, dtgb.build.inters, v)
}

func (dtgb *DamageTypeGroupBy) sqlScan(ctx context.Context, root *DamageTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dtgb.flds)+len(dtgb.fns))
		for _, f := range *dtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DamageTypeSelect is the builder for selecting fields of DamageType entities.
type DamageTypeSelect struct {
	*DamageTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dts *DamageTypeSelect) Aggregate(fns ...AggregateFunc) *DamageTypeSelect {
	dts.fns = append(dts.fns, fns...)
	return dts
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DamageTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dts.ctx, ent.OpQuerySelect)
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DamageTypeQuery, *DamageTypeSelect](ctx, dts.DamageTypeQuery, dts, dts.inters, v)
}

func (dts *DamageTypeSelect) sqlScan(ctx context.Context, root *DamageTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dts.fns))
	for _, fn := range dts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
