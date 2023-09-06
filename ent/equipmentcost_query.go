// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentcost"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// EquipmentCostQuery is the builder for querying EquipmentCost entities.
type EquipmentCostQuery struct {
	config
	ctx           *QueryContext
	order         []equipmentcost.OrderOption
	inters        []Interceptor
	predicates    []predicate.EquipmentCost
	withEquipment *EquipmentQuery
	withCoin      *CoinQuery
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*EquipmentCost) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentCostQuery builder.
func (ecq *EquipmentCostQuery) Where(ps ...predicate.EquipmentCost) *EquipmentCostQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *EquipmentCostQuery) Limit(limit int) *EquipmentCostQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *EquipmentCostQuery) Offset(offset int) *EquipmentCostQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EquipmentCostQuery) Unique(unique bool) *EquipmentCostQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *EquipmentCostQuery) Order(o ...equipmentcost.OrderOption) *EquipmentCostQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (ecq *EquipmentCostQuery) QueryEquipment() *EquipmentQuery {
	query := (&EquipmentClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentcost.Table, equipmentcost.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, equipmentcost.EquipmentTable, equipmentcost.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCoin chains the current query on the "coin" edge.
func (ecq *EquipmentCostQuery) QueryCoin() *CoinQuery {
	query := (&CoinClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentcost.Table, equipmentcost.FieldID, selector),
			sqlgraph.To(coin.Table, coin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, equipmentcost.CoinTable, equipmentcost.CoinColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EquipmentCost entity from the query.
// Returns a *NotFoundError when no EquipmentCost was found.
func (ecq *EquipmentCostQuery) First(ctx context.Context) (*EquipmentCost, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipmentcost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EquipmentCostQuery) FirstX(ctx context.Context) *EquipmentCost {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EquipmentCost ID from the query.
// Returns a *NotFoundError when no EquipmentCost ID was found.
func (ecq *EquipmentCostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipmentcost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EquipmentCostQuery) FirstIDX(ctx context.Context) int {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EquipmentCost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EquipmentCost entity is found.
// Returns a *NotFoundError when no EquipmentCost entities are found.
func (ecq *EquipmentCostQuery) Only(ctx context.Context) (*EquipmentCost, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipmentcost.Label}
	default:
		return nil, &NotSingularError{equipmentcost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EquipmentCostQuery) OnlyX(ctx context.Context) *EquipmentCost {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EquipmentCost ID in the query.
// Returns a *NotSingularError when more than one EquipmentCost ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EquipmentCostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipmentcost.Label}
	default:
		err = &NotSingularError{equipmentcost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EquipmentCostQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentCosts.
func (ecq *EquipmentCostQuery) All(ctx context.Context) ([]*EquipmentCost, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EquipmentCost, *EquipmentCostQuery]()
	return withInterceptors[[]*EquipmentCost](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EquipmentCostQuery) AllX(ctx context.Context) []*EquipmentCost {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EquipmentCost IDs.
func (ecq *EquipmentCostQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(equipmentcost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EquipmentCostQuery) IDsX(ctx context.Context) []int {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EquipmentCostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*EquipmentCostQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EquipmentCostQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EquipmentCostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Exist")
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EquipmentCostQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentCostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EquipmentCostQuery) Clone() *EquipmentCostQuery {
	if ecq == nil {
		return nil
	}
	return &EquipmentCostQuery{
		config:        ecq.config,
		ctx:           ecq.ctx.Clone(),
		order:         append([]equipmentcost.OrderOption{}, ecq.order...),
		inters:        append([]Interceptor{}, ecq.inters...),
		predicates:    append([]predicate.EquipmentCost{}, ecq.predicates...),
		withEquipment: ecq.withEquipment.Clone(),
		withCoin:      ecq.withCoin.Clone(),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EquipmentCostQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EquipmentCostQuery {
	query := (&EquipmentClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withEquipment = query
	return ecq
}

// WithCoin tells the query-builder to eager-load the nodes that are connected to
// the "coin" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EquipmentCostQuery) WithCoin(opts ...func(*CoinQuery)) *EquipmentCostQuery {
	query := (&CoinClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withCoin = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EquipmentID int `json:"equipment_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentCost.Query().
//		GroupBy(equipmentcost.FieldEquipmentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EquipmentCostQuery) GroupBy(field string, fields ...string) *EquipmentCostGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EquipmentCostGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = equipmentcost.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EquipmentID int `json:"equipment_id,omitempty"`
//	}
//
//	client.EquipmentCost.Query().
//		Select(equipmentcost.FieldEquipmentID).
//		Scan(ctx, &v)
func (ecq *EquipmentCostQuery) Select(fields ...string) *EquipmentCostSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &EquipmentCostSelect{EquipmentCostQuery: ecq}
	sbuild.label = equipmentcost.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EquipmentCostSelect configured with the given aggregations.
func (ecq *EquipmentCostQuery) Aggregate(fns ...AggregateFunc) *EquipmentCostSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *EquipmentCostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !equipmentcost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *EquipmentCostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EquipmentCost, error) {
	var (
		nodes       = []*EquipmentCost{}
		_spec       = ecq.querySpec()
		loadedTypes = [2]bool{
			ecq.withEquipment != nil,
			ecq.withCoin != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EquipmentCost).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EquipmentCost{config: ecq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ecq.withEquipment; query != nil {
		if err := ecq.loadEquipment(ctx, query, nodes, nil,
			func(n *EquipmentCost, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withCoin; query != nil {
		if err := ecq.loadCoin(ctx, query, nodes, nil,
			func(n *EquipmentCost, e *Coin) { n.Edges.Coin = e }); err != nil {
			return nil, err
		}
	}
	for i := range ecq.loadTotal {
		if err := ecq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *EquipmentCostQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*EquipmentCost, init func(*EquipmentCost), assign func(*EquipmentCost, *Equipment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EquipmentCost)
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
func (ecq *EquipmentCostQuery) loadCoin(ctx context.Context, query *CoinQuery, nodes []*EquipmentCost, init func(*EquipmentCost), assign func(*EquipmentCost, *Coin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EquipmentCost)
	for i := range nodes {
		fk := nodes[i].CoinID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(coin.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "coin_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ecq *EquipmentCostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EquipmentCostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(equipmentcost.Table, equipmentcost.Columns, sqlgraph.NewFieldSpec(equipmentcost.FieldID, field.TypeInt))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentcost.FieldID)
		for i := range fields {
			if fields[i] != equipmentcost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ecq.withEquipment != nil {
			_spec.Node.AddColumnOnce(equipmentcost.FieldEquipmentID)
		}
		if ecq.withCoin != nil {
			_spec.Node.AddColumnOnce(equipmentcost.FieldCoinID)
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *EquipmentCostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(equipmentcost.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = equipmentcost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentCostGroupBy is the group-by builder for EquipmentCost entities.
type EquipmentCostGroupBy struct {
	selector
	build *EquipmentCostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EquipmentCostGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentCostGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *EquipmentCostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EquipmentCostQuery, *EquipmentCostGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *EquipmentCostGroupBy) sqlScan(ctx context.Context, root *EquipmentCostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EquipmentCostSelect is the builder for selecting fields of EquipmentCost entities.
type EquipmentCostSelect struct {
	*EquipmentCostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *EquipmentCostSelect) Aggregate(fns ...AggregateFunc) *EquipmentCostSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EquipmentCostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EquipmentCostQuery, *EquipmentCostSelect](ctx, ecs.EquipmentCostQuery, ecs, ecs.inters, v)
}

func (ecs *EquipmentCostSelect) sqlScan(ctx context.Context, root *EquipmentCostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
