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
	"github.com/ecshreve/dndgen/ent/magicitem"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// MagicItemQuery is the builder for querying MagicItem entities.
type MagicItemQuery struct {
	config
	ctx                *QueryContext
	order              []magicitem.OrderOption
	inters             []Interceptor
	predicates         []predicate.MagicItem
	withEquipment      *EquipmentQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*MagicItem) error
	withNamedEquipment map[string]*EquipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MagicItemQuery builder.
func (miq *MagicItemQuery) Where(ps ...predicate.MagicItem) *MagicItemQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit the number of records to be returned by this query.
func (miq *MagicItemQuery) Limit(limit int) *MagicItemQuery {
	miq.ctx.Limit = &limit
	return miq
}

// Offset to start from.
func (miq *MagicItemQuery) Offset(offset int) *MagicItemQuery {
	miq.ctx.Offset = &offset
	return miq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (miq *MagicItemQuery) Unique(unique bool) *MagicItemQuery {
	miq.ctx.Unique = &unique
	return miq
}

// Order specifies how the records should be ordered.
func (miq *MagicItemQuery) Order(o ...magicitem.OrderOption) *MagicItemQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (miq *MagicItemQuery) QueryEquipment() *EquipmentQuery {
	query := (&EquipmentClient{config: miq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(magicitem.Table, magicitem.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, magicitem.EquipmentTable, magicitem.EquipmentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MagicItem entity from the query.
// Returns a *NotFoundError when no MagicItem was found.
func (miq *MagicItemQuery) First(ctx context.Context) (*MagicItem, error) {
	nodes, err := miq.Limit(1).All(setContextOp(ctx, miq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{magicitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MagicItemQuery) FirstX(ctx context.Context) *MagicItem {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MagicItem ID from the query.
// Returns a *NotFoundError when no MagicItem ID was found.
func (miq *MagicItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(1).IDs(setContextOp(ctx, miq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{magicitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MagicItemQuery) FirstIDX(ctx context.Context) int {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MagicItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MagicItem entity is found.
// Returns a *NotFoundError when no MagicItem entities are found.
func (miq *MagicItemQuery) Only(ctx context.Context) (*MagicItem, error) {
	nodes, err := miq.Limit(2).All(setContextOp(ctx, miq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{magicitem.Label}
	default:
		return nil, &NotSingularError{magicitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MagicItemQuery) OnlyX(ctx context.Context) *MagicItem {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MagicItem ID in the query.
// Returns a *NotSingularError when more than one MagicItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (miq *MagicItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(2).IDs(setContextOp(ctx, miq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{magicitem.Label}
	default:
		err = &NotSingularError{magicitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MagicItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MagicItems.
func (miq *MagicItemQuery) All(ctx context.Context) ([]*MagicItem, error) {
	ctx = setContextOp(ctx, miq.ctx, "All")
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MagicItem, *MagicItemQuery]()
	return withInterceptors[[]*MagicItem](ctx, miq, qr, miq.inters)
}

// AllX is like All, but panics if an error occurs.
func (miq *MagicItemQuery) AllX(ctx context.Context) []*MagicItem {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MagicItem IDs.
func (miq *MagicItemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if miq.ctx.Unique == nil && miq.path != nil {
		miq.Unique(true)
	}
	ctx = setContextOp(ctx, miq.ctx, "IDs")
	if err = miq.Select(magicitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MagicItemQuery) IDsX(ctx context.Context) []int {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MagicItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, miq.ctx, "Count")
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, miq, querierCount[*MagicItemQuery](), miq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MagicItemQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MagicItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, miq.ctx, "Exist")
	switch _, err := miq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MagicItemQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MagicItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MagicItemQuery) Clone() *MagicItemQuery {
	if miq == nil {
		return nil
	}
	return &MagicItemQuery{
		config:        miq.config,
		ctx:           miq.ctx.Clone(),
		order:         append([]magicitem.OrderOption{}, miq.order...),
		inters:        append([]Interceptor{}, miq.inters...),
		predicates:    append([]predicate.MagicItem{}, miq.predicates...),
		withEquipment: miq.withEquipment.Clone(),
		// clone intermediate query.
		sql:  miq.sql.Clone(),
		path: miq.path,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MagicItemQuery) WithEquipment(opts ...func(*EquipmentQuery)) *MagicItemQuery {
	query := (&EquipmentClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	miq.withEquipment = query
	return miq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rarity string `json:"rarity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MagicItem.Query().
//		GroupBy(magicitem.FieldRarity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (miq *MagicItemQuery) GroupBy(field string, fields ...string) *MagicItemGroupBy {
	miq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MagicItemGroupBy{build: miq}
	grbuild.flds = &miq.ctx.Fields
	grbuild.label = magicitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rarity string `json:"rarity,omitempty"`
//	}
//
//	client.MagicItem.Query().
//		Select(magicitem.FieldRarity).
//		Scan(ctx, &v)
func (miq *MagicItemQuery) Select(fields ...string) *MagicItemSelect {
	miq.ctx.Fields = append(miq.ctx.Fields, fields...)
	sbuild := &MagicItemSelect{MagicItemQuery: miq}
	sbuild.label = magicitem.Label
	sbuild.flds, sbuild.scan = &miq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MagicItemSelect configured with the given aggregations.
func (miq *MagicItemQuery) Aggregate(fns ...AggregateFunc) *MagicItemSelect {
	return miq.Select().Aggregate(fns...)
}

func (miq *MagicItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range miq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, miq); err != nil {
				return err
			}
		}
	}
	for _, f := range miq.ctx.Fields {
		if !magicitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MagicItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MagicItem, error) {
	var (
		nodes       = []*MagicItem{}
		_spec       = miq.querySpec()
		loadedTypes = [1]bool{
			miq.withEquipment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MagicItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MagicItem{config: miq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(miq.modifiers) > 0 {
		_spec.Modifiers = miq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := miq.withEquipment; query != nil {
		if err := miq.loadEquipment(ctx, query, nodes,
			func(n *MagicItem) { n.Edges.Equipment = []*Equipment{} },
			func(n *MagicItem, e *Equipment) { n.Edges.Equipment = append(n.Edges.Equipment, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range miq.withNamedEquipment {
		if err := miq.loadEquipment(ctx, query, nodes,
			func(n *MagicItem) { n.appendNamedEquipment(name) },
			func(n *MagicItem, e *Equipment) { n.appendNamedEquipment(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range miq.loadTotal {
		if err := miq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (miq *MagicItemQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*MagicItem, init func(*MagicItem), assign func(*MagicItem, *Equipment)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*MagicItem)
	nids := make(map[int]map[*MagicItem]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(magicitem.EquipmentTable)
		s.Join(joinT).On(s.C(equipment.FieldID), joinT.C(magicitem.EquipmentPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(magicitem.EquipmentPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(magicitem.EquipmentPrimaryKey[1]))
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
					nids[inValue] = map[*MagicItem]struct{}{byID[outValue]: {}}
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

func (miq *MagicItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	if len(miq.modifiers) > 0 {
		_spec.Modifiers = miq.modifiers
	}
	_spec.Node.Columns = miq.ctx.Fields
	if len(miq.ctx.Fields) > 0 {
		_spec.Unique = miq.ctx.Unique != nil && *miq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MagicItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(magicitem.Table, magicitem.Columns, sqlgraph.NewFieldSpec(magicitem.FieldID, field.TypeInt))
	_spec.From = miq.sql
	if unique := miq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if miq.path != nil {
		_spec.Unique = true
	}
	if fields := miq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, magicitem.FieldID)
		for i := range fields {
			if fields[i] != magicitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (miq *MagicItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(magicitem.Table)
	columns := miq.ctx.Fields
	if len(columns) == 0 {
		columns = magicitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if miq.ctx.Unique != nil && *miq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector)
	}
	if offset := miq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedEquipment tells the query-builder to eager-load the nodes that are connected to the "equipment"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (miq *MagicItemQuery) WithNamedEquipment(name string, opts ...func(*EquipmentQuery)) *MagicItemQuery {
	query := (&EquipmentClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if miq.withNamedEquipment == nil {
		miq.withNamedEquipment = make(map[string]*EquipmentQuery)
	}
	miq.withNamedEquipment[name] = query
	return miq
}

// MagicItemGroupBy is the group-by builder for MagicItem entities.
type MagicItemGroupBy struct {
	selector
	build *MagicItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MagicItemGroupBy) Aggregate(fns ...AggregateFunc) *MagicItemGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the selector query and scans the result into the given value.
func (migb *MagicItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, migb.build.ctx, "GroupBy")
	if err := migb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MagicItemQuery, *MagicItemGroupBy](ctx, migb.build, migb, migb.build.inters, v)
}

func (migb *MagicItemGroupBy) sqlScan(ctx context.Context, root *MagicItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(migb.fns))
	for _, fn := range migb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*migb.flds)+len(migb.fns))
		for _, f := range *migb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*migb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MagicItemSelect is the builder for selecting fields of MagicItem entities.
type MagicItemSelect struct {
	*MagicItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mis *MagicItemSelect) Aggregate(fns ...AggregateFunc) *MagicItemSelect {
	mis.fns = append(mis.fns, fns...)
	return mis
}

// Scan applies the selector query and scans the result into the given value.
func (mis *MagicItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mis.ctx, "Select")
	if err := mis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MagicItemQuery, *MagicItemSelect](ctx, mis.MagicItemQuery, mis, mis.inters, v)
}

func (mis *MagicItemSelect) sqlScan(ctx context.Context, root *MagicItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mis.fns))
	for _, fn := range mis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
