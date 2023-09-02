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
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentchoice"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// EquipmentChoiceQuery is the builder for querying EquipmentChoice entities.
type EquipmentChoiceQuery struct {
	config
	ctx                *QueryContext
	order              []equipmentchoice.OrderOption
	inters             []Interceptor
	predicates         []predicate.EquipmentChoice
	withClass          *ClassQuery
	withEquipment      *EquipmentQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*EquipmentChoice) error
	withNamedEquipment map[string]*EquipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentChoiceQuery builder.
func (ecq *EquipmentChoiceQuery) Where(ps ...predicate.EquipmentChoice) *EquipmentChoiceQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *EquipmentChoiceQuery) Limit(limit int) *EquipmentChoiceQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *EquipmentChoiceQuery) Offset(offset int) *EquipmentChoiceQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EquipmentChoiceQuery) Unique(unique bool) *EquipmentChoiceQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *EquipmentChoiceQuery) Order(o ...equipmentchoice.OrderOption) *EquipmentChoiceQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryClass chains the current query on the "class" edge.
func (ecq *EquipmentChoiceQuery) QueryClass() *ClassQuery {
	query := (&ClassClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentchoice.Table, equipmentchoice.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, equipmentchoice.ClassTable, equipmentchoice.ClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEquipment chains the current query on the "equipment" edge.
func (ecq *EquipmentChoiceQuery) QueryEquipment() *EquipmentQuery {
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
			sqlgraph.From(equipmentchoice.Table, equipmentchoice.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, equipmentchoice.EquipmentTable, equipmentchoice.EquipmentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EquipmentChoice entity from the query.
// Returns a *NotFoundError when no EquipmentChoice was found.
func (ecq *EquipmentChoiceQuery) First(ctx context.Context) (*EquipmentChoice, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipmentchoice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) FirstX(ctx context.Context) *EquipmentChoice {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EquipmentChoice ID from the query.
// Returns a *NotFoundError when no EquipmentChoice ID was found.
func (ecq *EquipmentChoiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipmentchoice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) FirstIDX(ctx context.Context) int {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EquipmentChoice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EquipmentChoice entity is found.
// Returns a *NotFoundError when no EquipmentChoice entities are found.
func (ecq *EquipmentChoiceQuery) Only(ctx context.Context) (*EquipmentChoice, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipmentchoice.Label}
	default:
		return nil, &NotSingularError{equipmentchoice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) OnlyX(ctx context.Context) *EquipmentChoice {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EquipmentChoice ID in the query.
// Returns a *NotSingularError when more than one EquipmentChoice ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EquipmentChoiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipmentchoice.Label}
	default:
		err = &NotSingularError{equipmentchoice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentChoices.
func (ecq *EquipmentChoiceQuery) All(ctx context.Context) ([]*EquipmentChoice, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EquipmentChoice, *EquipmentChoiceQuery]()
	return withInterceptors[[]*EquipmentChoice](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) AllX(ctx context.Context) []*EquipmentChoice {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EquipmentChoice IDs.
func (ecq *EquipmentChoiceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(equipmentchoice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) IDsX(ctx context.Context) []int {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EquipmentChoiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*EquipmentChoiceQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EquipmentChoiceQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EquipmentChoiceQuery) Exist(ctx context.Context) (bool, error) {
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
func (ecq *EquipmentChoiceQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentChoiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EquipmentChoiceQuery) Clone() *EquipmentChoiceQuery {
	if ecq == nil {
		return nil
	}
	return &EquipmentChoiceQuery{
		config:        ecq.config,
		ctx:           ecq.ctx.Clone(),
		order:         append([]equipmentchoice.OrderOption{}, ecq.order...),
		inters:        append([]Interceptor{}, ecq.inters...),
		predicates:    append([]predicate.EquipmentChoice{}, ecq.predicates...),
		withClass:     ecq.withClass.Clone(),
		withEquipment: ecq.withEquipment.Clone(),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// WithClass tells the query-builder to eager-load the nodes that are connected to
// the "class" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EquipmentChoiceQuery) WithClass(opts ...func(*ClassQuery)) *EquipmentChoiceQuery {
	query := (&ClassClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withClass = query
	return ecq
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EquipmentChoiceQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EquipmentChoiceQuery {
	query := (&EquipmentClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withEquipment = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClassID int `json:"class_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentChoice.Query().
//		GroupBy(equipmentchoice.FieldClassID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EquipmentChoiceQuery) GroupBy(field string, fields ...string) *EquipmentChoiceGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EquipmentChoiceGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = equipmentchoice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClassID int `json:"class_id,omitempty"`
//	}
//
//	client.EquipmentChoice.Query().
//		Select(equipmentchoice.FieldClassID).
//		Scan(ctx, &v)
func (ecq *EquipmentChoiceQuery) Select(fields ...string) *EquipmentChoiceSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &EquipmentChoiceSelect{EquipmentChoiceQuery: ecq}
	sbuild.label = equipmentchoice.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EquipmentChoiceSelect configured with the given aggregations.
func (ecq *EquipmentChoiceQuery) Aggregate(fns ...AggregateFunc) *EquipmentChoiceSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *EquipmentChoiceQuery) prepareQuery(ctx context.Context) error {
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
		if !equipmentchoice.ValidColumn(f) {
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

func (ecq *EquipmentChoiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EquipmentChoice, error) {
	var (
		nodes       = []*EquipmentChoice{}
		_spec       = ecq.querySpec()
		loadedTypes = [2]bool{
			ecq.withClass != nil,
			ecq.withEquipment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EquipmentChoice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EquipmentChoice{config: ecq.config}
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
	if query := ecq.withClass; query != nil {
		if err := ecq.loadClass(ctx, query, nodes, nil,
			func(n *EquipmentChoice, e *Class) { n.Edges.Class = e }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withEquipment; query != nil {
		if err := ecq.loadEquipment(ctx, query, nodes,
			func(n *EquipmentChoice) { n.Edges.Equipment = []*Equipment{} },
			func(n *EquipmentChoice, e *Equipment) { n.Edges.Equipment = append(n.Edges.Equipment, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ecq.withNamedEquipment {
		if err := ecq.loadEquipment(ctx, query, nodes,
			func(n *EquipmentChoice) { n.appendNamedEquipment(name) },
			func(n *EquipmentChoice, e *Equipment) { n.appendNamedEquipment(name, e) }); err != nil {
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

func (ecq *EquipmentChoiceQuery) loadClass(ctx context.Context, query *ClassQuery, nodes []*EquipmentChoice, init func(*EquipmentChoice), assign func(*EquipmentChoice, *Class)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EquipmentChoice)
	for i := range nodes {
		fk := nodes[i].ClassID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(class.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ecq *EquipmentChoiceQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*EquipmentChoice, init func(*EquipmentChoice), assign func(*EquipmentChoice, *Equipment)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*EquipmentChoice)
	nids := make(map[int]map[*EquipmentChoice]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(equipmentchoice.EquipmentTable)
		s.Join(joinT).On(s.C(equipment.FieldID), joinT.C(equipmentchoice.EquipmentPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(equipmentchoice.EquipmentPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(equipmentchoice.EquipmentPrimaryKey[1]))
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
					nids[inValue] = map[*EquipmentChoice]struct{}{byID[outValue]: {}}
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

func (ecq *EquipmentChoiceQuery) sqlCount(ctx context.Context) (int, error) {
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

func (ecq *EquipmentChoiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(equipmentchoice.Table, equipmentchoice.Columns, sqlgraph.NewFieldSpec(equipmentchoice.FieldID, field.TypeInt))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentchoice.FieldID)
		for i := range fields {
			if fields[i] != equipmentchoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ecq.withClass != nil {
			_spec.Node.AddColumnOnce(equipmentchoice.FieldClassID)
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

func (ecq *EquipmentChoiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(equipmentchoice.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = equipmentchoice.Columns
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

// WithNamedEquipment tells the query-builder to eager-load the nodes that are connected to the "equipment"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ecq *EquipmentChoiceQuery) WithNamedEquipment(name string, opts ...func(*EquipmentQuery)) *EquipmentChoiceQuery {
	query := (&EquipmentClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ecq.withNamedEquipment == nil {
		ecq.withNamedEquipment = make(map[string]*EquipmentQuery)
	}
	ecq.withNamedEquipment[name] = query
	return ecq
}

// EquipmentChoiceGroupBy is the group-by builder for EquipmentChoice entities.
type EquipmentChoiceGroupBy struct {
	selector
	build *EquipmentChoiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EquipmentChoiceGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentChoiceGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *EquipmentChoiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EquipmentChoiceQuery, *EquipmentChoiceGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *EquipmentChoiceGroupBy) sqlScan(ctx context.Context, root *EquipmentChoiceQuery, v any) error {
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

// EquipmentChoiceSelect is the builder for selecting fields of EquipmentChoice entities.
type EquipmentChoiceSelect struct {
	*EquipmentChoiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *EquipmentChoiceSelect) Aggregate(fns ...AggregateFunc) *EquipmentChoiceSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EquipmentChoiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EquipmentChoiceQuery, *EquipmentChoiceSelect](ctx, ecs.EquipmentChoiceQuery, ecs, ecs.inters, v)
}

func (ecs *EquipmentChoiceSelect) sqlScan(ctx context.Context, root *EquipmentChoiceQuery, v any) error {
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
