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
	"github.com/ecshreve/dndgen/ent/ammunition"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// AmmunitionQuery is the builder for querying Ammunition entities.
type AmmunitionQuery struct {
	config
	ctx           *QueryContext
	order         []ammunition.OrderOption
	inters        []Interceptor
	predicates    []predicate.Ammunition
	withEquipment *EquipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AmmunitionQuery builder.
func (aq *AmmunitionQuery) Where(ps ...predicate.Ammunition) *AmmunitionQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AmmunitionQuery) Limit(limit int) *AmmunitionQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AmmunitionQuery) Offset(offset int) *AmmunitionQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AmmunitionQuery) Unique(unique bool) *AmmunitionQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AmmunitionQuery) Order(o ...ammunition.OrderOption) *AmmunitionQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (aq *AmmunitionQuery) QueryEquipment() *EquipmentQuery {
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
			sqlgraph.From(ammunition.Table, ammunition.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ammunition.EquipmentTable, ammunition.EquipmentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Ammunition entity from the query.
// Returns a *NotFoundError when no Ammunition was found.
func (aq *AmmunitionQuery) First(ctx context.Context) (*Ammunition, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ammunition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AmmunitionQuery) FirstX(ctx context.Context) *Ammunition {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Ammunition ID from the query.
// Returns a *NotFoundError when no Ammunition ID was found.
func (aq *AmmunitionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ammunition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AmmunitionQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Ammunition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Ammunition entity is found.
// Returns a *NotFoundError when no Ammunition entities are found.
func (aq *AmmunitionQuery) Only(ctx context.Context) (*Ammunition, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ammunition.Label}
	default:
		return nil, &NotSingularError{ammunition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AmmunitionQuery) OnlyX(ctx context.Context) *Ammunition {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Ammunition ID in the query.
// Returns a *NotSingularError when more than one Ammunition ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AmmunitionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ammunition.Label}
	default:
		err = &NotSingularError{ammunition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AmmunitionQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ammunitions.
func (aq *AmmunitionQuery) All(ctx context.Context) ([]*Ammunition, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Ammunition, *AmmunitionQuery]()
	return withInterceptors[[]*Ammunition](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AmmunitionQuery) AllX(ctx context.Context) []*Ammunition {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Ammunition IDs.
func (aq *AmmunitionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(ammunition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AmmunitionQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AmmunitionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AmmunitionQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AmmunitionQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AmmunitionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
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
func (aq *AmmunitionQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AmmunitionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AmmunitionQuery) Clone() *AmmunitionQuery {
	if aq == nil {
		return nil
	}
	return &AmmunitionQuery{
		config:        aq.config,
		ctx:           aq.ctx.Clone(),
		order:         append([]ammunition.OrderOption{}, aq.order...),
		inters:        append([]Interceptor{}, aq.inters...),
		predicates:    append([]predicate.Ammunition{}, aq.predicates...),
		withEquipment: aq.withEquipment.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AmmunitionQuery) WithEquipment(opts ...func(*EquipmentQuery)) *AmmunitionQuery {
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
//		Quantity int `json:"quantity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ammunition.Query().
//		GroupBy(ammunition.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AmmunitionQuery) GroupBy(field string, fields ...string) *AmmunitionGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AmmunitionGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = ammunition.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//	}
//
//	client.Ammunition.Query().
//		Select(ammunition.FieldQuantity).
//		Scan(ctx, &v)
func (aq *AmmunitionQuery) Select(fields ...string) *AmmunitionSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AmmunitionSelect{AmmunitionQuery: aq}
	sbuild.label = ammunition.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AmmunitionSelect configured with the given aggregations.
func (aq *AmmunitionQuery) Aggregate(fns ...AggregateFunc) *AmmunitionSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AmmunitionQuery) prepareQuery(ctx context.Context) error {
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
		if !ammunition.ValidColumn(f) {
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

func (aq *AmmunitionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Ammunition, error) {
	var (
		nodes       = []*Ammunition{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withEquipment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Ammunition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Ammunition{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := aq.withEquipment; query != nil {
		if err := aq.loadEquipment(ctx, query, nodes,
			func(n *Ammunition) { n.Edges.Equipment = []*Equipment{} },
			func(n *Ammunition, e *Equipment) { n.Edges.Equipment = append(n.Edges.Equipment, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AmmunitionQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Ammunition, init func(*Ammunition), assign func(*Ammunition, *Equipment)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Ammunition)
	nids := make(map[int]map[*Ammunition]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(ammunition.EquipmentTable)
		s.Join(joinT).On(s.C(equipment.FieldID), joinT.C(ammunition.EquipmentPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(ammunition.EquipmentPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(ammunition.EquipmentPrimaryKey[1]))
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
					nids[inValue] = map[*Ammunition]struct{}{byID[outValue]: {}}
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

func (aq *AmmunitionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AmmunitionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ammunition.Table, ammunition.Columns, sqlgraph.NewFieldSpec(ammunition.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ammunition.FieldID)
		for i := range fields {
			if fields[i] != ammunition.FieldID {
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

func (aq *AmmunitionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(ammunition.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = ammunition.Columns
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

// AmmunitionGroupBy is the group-by builder for Ammunition entities.
type AmmunitionGroupBy struct {
	selector
	build *AmmunitionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AmmunitionGroupBy) Aggregate(fns ...AggregateFunc) *AmmunitionGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AmmunitionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AmmunitionQuery, *AmmunitionGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AmmunitionGroupBy) sqlScan(ctx context.Context, root *AmmunitionQuery, v any) error {
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

// AmmunitionSelect is the builder for selecting fields of Ammunition entities.
type AmmunitionSelect struct {
	*AmmunitionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AmmunitionSelect) Aggregate(fns ...AggregateFunc) *AmmunitionSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AmmunitionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AmmunitionQuery, *AmmunitionSelect](ctx, as.AmmunitionQuery, as, as.inters, v)
}

func (as *AmmunitionSelect) sqlScan(ctx context.Context, root *AmmunitionQuery, v any) error {
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
