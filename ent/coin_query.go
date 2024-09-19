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
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/equipmentcost"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// CoinQuery is the builder for querying Coin entities.
type CoinQuery struct {
	config
	ctx                     *QueryContext
	order                   []coin.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Coin
	withEquipmentCosts      *EquipmentCostQuery
	modifiers               []func(*sql.Selector)
	loadTotal               []func(context.Context, []*Coin) error
	withNamedEquipmentCosts map[string]*EquipmentCostQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinQuery builder.
func (cq *CoinQuery) Where(ps ...predicate.Coin) *CoinQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CoinQuery) Limit(limit int) *CoinQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CoinQuery) Offset(offset int) *CoinQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CoinQuery) Unique(unique bool) *CoinQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CoinQuery) Order(o ...coin.OrderOption) *CoinQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryEquipmentCosts chains the current query on the "equipment_costs" edge.
func (cq *CoinQuery) QueryEquipmentCosts() *EquipmentCostQuery {
	query := (&EquipmentCostClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(coin.Table, coin.FieldID, selector),
			sqlgraph.To(equipmentcost.Table, equipmentcost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, coin.EquipmentCostsTable, coin.EquipmentCostsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Coin entity from the query.
// Returns a *NotFoundError when no Coin was found.
func (cq *CoinQuery) First(ctx context.Context) (*Coin, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CoinQuery) FirstX(ctx context.Context) *Coin {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Coin ID from the query.
// Returns a *NotFoundError when no Coin ID was found.
func (cq *CoinQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CoinQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Coin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Coin entity is found.
// Returns a *NotFoundError when no Coin entities are found.
func (cq *CoinQuery) Only(ctx context.Context) (*Coin, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coin.Label}
	default:
		return nil, &NotSingularError{coin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CoinQuery) OnlyX(ctx context.Context) *Coin {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Coin ID in the query.
// Returns a *NotSingularError when more than one Coin ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CoinQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coin.Label}
	default:
		err = &NotSingularError{coin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CoinQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Coins.
func (cq *CoinQuery) All(ctx context.Context) ([]*Coin, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Coin, *CoinQuery]()
	return withInterceptors[[]*Coin](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CoinQuery) AllX(ctx context.Context) []*Coin {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Coin IDs.
func (cq *CoinQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(coin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CoinQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CoinQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CoinQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CoinQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CoinQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CoinQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CoinQuery) Clone() *CoinQuery {
	if cq == nil {
		return nil
	}
	return &CoinQuery{
		config:             cq.config,
		ctx:                cq.ctx.Clone(),
		order:              append([]coin.OrderOption{}, cq.order...),
		inters:             append([]Interceptor{}, cq.inters...),
		predicates:         append([]predicate.Coin{}, cq.predicates...),
		withEquipmentCosts: cq.withEquipmentCosts.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithEquipmentCosts tells the query-builder to eager-load the nodes that are connected to
// the "equipment_costs" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CoinQuery) WithEquipmentCosts(opts ...func(*EquipmentCostQuery)) *CoinQuery {
	query := (&EquipmentCostClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withEquipmentCosts = query
	return cq
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
//	client.Coin.Query().
//		GroupBy(coin.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CoinQuery) GroupBy(field string, fields ...string) *CoinGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CoinGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = coin.Label
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
//	client.Coin.Query().
//		Select(coin.FieldIndx).
//		Scan(ctx, &v)
func (cq *CoinQuery) Select(fields ...string) *CoinSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CoinSelect{CoinQuery: cq}
	sbuild.label = coin.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CoinSelect configured with the given aggregations.
func (cq *CoinQuery) Aggregate(fns ...AggregateFunc) *CoinSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CoinQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !coin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CoinQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Coin, error) {
	var (
		nodes       = []*Coin{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withEquipmentCosts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Coin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Coin{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withEquipmentCosts; query != nil {
		if err := cq.loadEquipmentCosts(ctx, query, nodes,
			func(n *Coin) { n.Edges.EquipmentCosts = []*EquipmentCost{} },
			func(n *Coin, e *EquipmentCost) { n.Edges.EquipmentCosts = append(n.Edges.EquipmentCosts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedEquipmentCosts {
		if err := cq.loadEquipmentCosts(ctx, query, nodes,
			func(n *Coin) { n.appendNamedEquipmentCosts(name) },
			func(n *Coin, e *EquipmentCost) { n.appendNamedEquipmentCosts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CoinQuery) loadEquipmentCosts(ctx context.Context, query *EquipmentCostQuery, nodes []*Coin, init func(*Coin), assign func(*Coin, *EquipmentCost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Coin)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EquipmentCost(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(coin.EquipmentCostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.equipment_cost_coin
		if fk == nil {
			return fmt.Errorf(`foreign-key "equipment_cost_coin" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "equipment_cost_coin" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CoinQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CoinQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(coin.Table, coin.Columns, sqlgraph.NewFieldSpec(coin.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coin.FieldID)
		for i := range fields {
			if fields[i] != coin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CoinQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(coin.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = coin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedEquipmentCosts tells the query-builder to eager-load the nodes that are connected to the "equipment_costs"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CoinQuery) WithNamedEquipmentCosts(name string, opts ...func(*EquipmentCostQuery)) *CoinQuery {
	query := (&EquipmentCostClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedEquipmentCosts == nil {
		cq.withNamedEquipmentCosts = make(map[string]*EquipmentCostQuery)
	}
	cq.withNamedEquipmentCosts[name] = query
	return cq
}

// CoinGroupBy is the group-by builder for Coin entities.
type CoinGroupBy struct {
	selector
	build *CoinQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CoinGroupBy) Aggregate(fns ...AggregateFunc) *CoinGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CoinGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CoinQuery, *CoinGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CoinGroupBy) sqlScan(ctx context.Context, root *CoinQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CoinSelect is the builder for selecting fields of Coin entities.
type CoinSelect struct {
	*CoinQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CoinSelect) Aggregate(fns ...AggregateFunc) *CoinSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CoinSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CoinQuery, *CoinSelect](ctx, cs.CoinQuery, cs, cs.inters, v)
}

func (cs *CoinSelect) sqlScan(ctx context.Context, root *CoinQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
