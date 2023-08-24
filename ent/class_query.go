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
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
)

// ClassQuery is the builder for querying Class entities.
type ClassQuery struct {
	config
	ctx                    *QueryContext
	order                  []class.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Class
	withSavingThrows       *AbilityScoreQuery
	withProficiencies      *ProficiencyQuery
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*Class) error
	withNamedSavingThrows  map[string]*AbilityScoreQuery
	withNamedProficiencies map[string]*ProficiencyQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassQuery builder.
func (cq *ClassQuery) Where(ps ...predicate.Class) *ClassQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ClassQuery) Limit(limit int) *ClassQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ClassQuery) Offset(offset int) *ClassQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ClassQuery) Unique(unique bool) *ClassQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ClassQuery) Order(o ...class.OrderOption) *ClassQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QuerySavingThrows chains the current query on the "saving_throws" edge.
func (cq *ClassQuery) QuerySavingThrows() *AbilityScoreQuery {
	query := (&AbilityScoreClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(abilityscore.Table, abilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, class.SavingThrowsTable, class.SavingThrowsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProficiencies chains the current query on the "proficiencies" edge.
func (cq *ClassQuery) QueryProficiencies() *ProficiencyQuery {
	query := (&ProficiencyClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(proficiency.Table, proficiency.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, class.ProficienciesTable, class.ProficienciesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Class entity from the query.
// Returns a *NotFoundError when no Class was found.
func (cq *ClassQuery) First(ctx context.Context) (*Class, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{class.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ClassQuery) FirstX(ctx context.Context) *Class {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Class ID from the query.
// Returns a *NotFoundError when no Class ID was found.
func (cq *ClassQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{class.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ClassQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Class entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Class entity is found.
// Returns a *NotFoundError when no Class entities are found.
func (cq *ClassQuery) Only(ctx context.Context) (*Class, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{class.Label}
	default:
		return nil, &NotSingularError{class.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ClassQuery) OnlyX(ctx context.Context) *Class {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Class ID in the query.
// Returns a *NotSingularError when more than one Class ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ClassQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{class.Label}
	default:
		err = &NotSingularError{class.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ClassQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Classes.
func (cq *ClassQuery) All(ctx context.Context) ([]*Class, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Class, *ClassQuery]()
	return withInterceptors[[]*Class](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ClassQuery) AllX(ctx context.Context) []*Class {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Class IDs.
func (cq *ClassQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(class.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ClassQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ClassQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ClassQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ClassQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ClassQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
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
func (cq *ClassQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ClassQuery) Clone() *ClassQuery {
	if cq == nil {
		return nil
	}
	return &ClassQuery{
		config:            cq.config,
		ctx:               cq.ctx.Clone(),
		order:             append([]class.OrderOption{}, cq.order...),
		inters:            append([]Interceptor{}, cq.inters...),
		predicates:        append([]predicate.Class{}, cq.predicates...),
		withSavingThrows:  cq.withSavingThrows.Clone(),
		withProficiencies: cq.withProficiencies.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithSavingThrows tells the query-builder to eager-load the nodes that are connected to
// the "saving_throws" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithSavingThrows(opts ...func(*AbilityScoreQuery)) *ClassQuery {
	query := (&AbilityScoreClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withSavingThrows = query
	return cq
}

// WithProficiencies tells the query-builder to eager-load the nodes that are connected to
// the "proficiencies" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithProficiencies(opts ...func(*ProficiencyQuery)) *ClassQuery {
	query := (&ProficiencyClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withProficiencies = query
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
//	client.Class.Query().
//		GroupBy(class.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ClassQuery) GroupBy(field string, fields ...string) *ClassGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClassGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = class.Label
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
//	client.Class.Query().
//		Select(class.FieldIndx).
//		Scan(ctx, &v)
func (cq *ClassQuery) Select(fields ...string) *ClassSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ClassSelect{ClassQuery: cq}
	sbuild.label = class.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClassSelect configured with the given aggregations.
func (cq *ClassQuery) Aggregate(fns ...AggregateFunc) *ClassSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ClassQuery) prepareQuery(ctx context.Context) error {
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
		if !class.ValidColumn(f) {
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

func (cq *ClassQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Class, error) {
	var (
		nodes       = []*Class{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withSavingThrows != nil,
			cq.withProficiencies != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Class).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Class{config: cq.config}
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
	if query := cq.withSavingThrows; query != nil {
		if err := cq.loadSavingThrows(ctx, query, nodes,
			func(n *Class) { n.Edges.SavingThrows = []*AbilityScore{} },
			func(n *Class, e *AbilityScore) { n.Edges.SavingThrows = append(n.Edges.SavingThrows, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withProficiencies; query != nil {
		if err := cq.loadProficiencies(ctx, query, nodes,
			func(n *Class) { n.Edges.Proficiencies = []*Proficiency{} },
			func(n *Class, e *Proficiency) { n.Edges.Proficiencies = append(n.Edges.Proficiencies, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedSavingThrows {
		if err := cq.loadSavingThrows(ctx, query, nodes,
			func(n *Class) { n.appendNamedSavingThrows(name) },
			func(n *Class, e *AbilityScore) { n.appendNamedSavingThrows(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedProficiencies {
		if err := cq.loadProficiencies(ctx, query, nodes,
			func(n *Class) { n.appendNamedProficiencies(name) },
			func(n *Class, e *Proficiency) { n.appendNamedProficiencies(name, e) }); err != nil {
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

func (cq *ClassQuery) loadSavingThrows(ctx context.Context, query *AbilityScoreQuery, nodes []*Class, init func(*Class), assign func(*Class, *AbilityScore)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Class)
	nids := make(map[int]map[*Class]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(class.SavingThrowsTable)
		s.Join(joinT).On(s.C(abilityscore.FieldID), joinT.C(class.SavingThrowsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(class.SavingThrowsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(class.SavingThrowsPrimaryKey[0]))
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
					nids[inValue] = map[*Class]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*AbilityScore](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "saving_throws" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadProficiencies(ctx context.Context, query *ProficiencyQuery, nodes []*Class, init func(*Class), assign func(*Class, *Proficiency)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Class)
	nids := make(map[int]map[*Class]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(class.ProficienciesTable)
		s.Join(joinT).On(s.C(proficiency.FieldID), joinT.C(class.ProficienciesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(class.ProficienciesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(class.ProficienciesPrimaryKey[0]))
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
					nids[inValue] = map[*Class]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Proficiency](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "proficiencies" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *ClassQuery) sqlCount(ctx context.Context) (int, error) {
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

func (cq *ClassQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, class.FieldID)
		for i := range fields {
			if fields[i] != class.FieldID {
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

func (cq *ClassQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(class.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = class.Columns
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

// WithNamedSavingThrows tells the query-builder to eager-load the nodes that are connected to the "saving_throws"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithNamedSavingThrows(name string, opts ...func(*AbilityScoreQuery)) *ClassQuery {
	query := (&AbilityScoreClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedSavingThrows == nil {
		cq.withNamedSavingThrows = make(map[string]*AbilityScoreQuery)
	}
	cq.withNamedSavingThrows[name] = query
	return cq
}

// WithNamedProficiencies tells the query-builder to eager-load the nodes that are connected to the "proficiencies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithNamedProficiencies(name string, opts ...func(*ProficiencyQuery)) *ClassQuery {
	query := (&ProficiencyClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedProficiencies == nil {
		cq.withNamedProficiencies = make(map[string]*ProficiencyQuery)
	}
	cq.withNamedProficiencies[name] = query
	return cq
}

// ClassGroupBy is the group-by builder for Class entities.
type ClassGroupBy struct {
	selector
	build *ClassQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ClassGroupBy) Aggregate(fns ...AggregateFunc) *ClassGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ClassGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassQuery, *ClassGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ClassGroupBy) sqlScan(ctx context.Context, root *ClassQuery, v any) error {
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

// ClassSelect is the builder for selecting fields of Class entities.
type ClassSelect struct {
	*ClassQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ClassSelect) Aggregate(fns ...AggregateFunc) *ClassSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ClassSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassQuery, *ClassSelect](ctx, cs.ClassQuery, cs, cs.inters, v)
}

func (cs *ClassSelect) sqlScan(ctx context.Context, root *ClassQuery, v any) error {
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
