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
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/languagechoice"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
)

// LanguageChoiceQuery is the builder for querying LanguageChoice entities.
type LanguageChoiceQuery struct {
	config
	ctx                *QueryContext
	order              []languagechoice.OrderOption
	inters             []Interceptor
	predicates         []predicate.LanguageChoice
	withLanguages      *LanguageQuery
	withRace           *RaceQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*LanguageChoice) error
	withNamedLanguages map[string]*LanguageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LanguageChoiceQuery builder.
func (lcq *LanguageChoiceQuery) Where(ps ...predicate.LanguageChoice) *LanguageChoiceQuery {
	lcq.predicates = append(lcq.predicates, ps...)
	return lcq
}

// Limit the number of records to be returned by this query.
func (lcq *LanguageChoiceQuery) Limit(limit int) *LanguageChoiceQuery {
	lcq.ctx.Limit = &limit
	return lcq
}

// Offset to start from.
func (lcq *LanguageChoiceQuery) Offset(offset int) *LanguageChoiceQuery {
	lcq.ctx.Offset = &offset
	return lcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcq *LanguageChoiceQuery) Unique(unique bool) *LanguageChoiceQuery {
	lcq.ctx.Unique = &unique
	return lcq
}

// Order specifies how the records should be ordered.
func (lcq *LanguageChoiceQuery) Order(o ...languagechoice.OrderOption) *LanguageChoiceQuery {
	lcq.order = append(lcq.order, o...)
	return lcq
}

// QueryLanguages chains the current query on the "languages" edge.
func (lcq *LanguageChoiceQuery) QueryLanguages() *LanguageQuery {
	query := (&LanguageClient{config: lcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(languagechoice.Table, languagechoice.FieldID, selector),
			sqlgraph.To(language.Table, language.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, languagechoice.LanguagesTable, languagechoice.LanguagesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRace chains the current query on the "race" edge.
func (lcq *LanguageChoiceQuery) QueryRace() *RaceQuery {
	query := (&RaceClient{config: lcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(languagechoice.Table, languagechoice.FieldID, selector),
			sqlgraph.To(race.Table, race.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, languagechoice.RaceTable, languagechoice.RaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LanguageChoice entity from the query.
// Returns a *NotFoundError when no LanguageChoice was found.
func (lcq *LanguageChoiceQuery) First(ctx context.Context) (*LanguageChoice, error) {
	nodes, err := lcq.Limit(1).All(setContextOp(ctx, lcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{languagechoice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) FirstX(ctx context.Context) *LanguageChoice {
	node, err := lcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LanguageChoice ID from the query.
// Returns a *NotFoundError when no LanguageChoice ID was found.
func (lcq *LanguageChoiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lcq.Limit(1).IDs(setContextOp(ctx, lcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{languagechoice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) FirstIDX(ctx context.Context) int {
	id, err := lcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LanguageChoice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LanguageChoice entity is found.
// Returns a *NotFoundError when no LanguageChoice entities are found.
func (lcq *LanguageChoiceQuery) Only(ctx context.Context) (*LanguageChoice, error) {
	nodes, err := lcq.Limit(2).All(setContextOp(ctx, lcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{languagechoice.Label}
	default:
		return nil, &NotSingularError{languagechoice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) OnlyX(ctx context.Context) *LanguageChoice {
	node, err := lcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LanguageChoice ID in the query.
// Returns a *NotSingularError when more than one LanguageChoice ID is found.
// Returns a *NotFoundError when no entities are found.
func (lcq *LanguageChoiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lcq.Limit(2).IDs(setContextOp(ctx, lcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{languagechoice.Label}
	default:
		err = &NotSingularError{languagechoice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := lcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LanguageChoices.
func (lcq *LanguageChoiceQuery) All(ctx context.Context) ([]*LanguageChoice, error) {
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryAll)
	if err := lcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LanguageChoice, *LanguageChoiceQuery]()
	return withInterceptors[[]*LanguageChoice](ctx, lcq, qr, lcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) AllX(ctx context.Context) []*LanguageChoice {
	nodes, err := lcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LanguageChoice IDs.
func (lcq *LanguageChoiceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lcq.ctx.Unique == nil && lcq.path != nil {
		lcq.Unique(true)
	}
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryIDs)
	if err = lcq.Select(languagechoice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) IDsX(ctx context.Context) []int {
	ids, err := lcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcq *LanguageChoiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryCount)
	if err := lcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lcq, querierCount[*LanguageChoiceQuery](), lcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) CountX(ctx context.Context) int {
	count, err := lcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcq *LanguageChoiceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lcq.ctx, ent.OpQueryExist)
	switch _, err := lcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lcq *LanguageChoiceQuery) ExistX(ctx context.Context) bool {
	exist, err := lcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LanguageChoiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcq *LanguageChoiceQuery) Clone() *LanguageChoiceQuery {
	if lcq == nil {
		return nil
	}
	return &LanguageChoiceQuery{
		config:        lcq.config,
		ctx:           lcq.ctx.Clone(),
		order:         append([]languagechoice.OrderOption{}, lcq.order...),
		inters:        append([]Interceptor{}, lcq.inters...),
		predicates:    append([]predicate.LanguageChoice{}, lcq.predicates...),
		withLanguages: lcq.withLanguages.Clone(),
		withRace:      lcq.withRace.Clone(),
		// clone intermediate query.
		sql:  lcq.sql.Clone(),
		path: lcq.path,
	}
}

// WithLanguages tells the query-builder to eager-load the nodes that are connected to
// the "languages" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LanguageChoiceQuery) WithLanguages(opts ...func(*LanguageQuery)) *LanguageChoiceQuery {
	query := (&LanguageClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcq.withLanguages = query
	return lcq
}

// WithRace tells the query-builder to eager-load the nodes that are connected to
// the "race" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LanguageChoiceQuery) WithRace(opts ...func(*RaceQuery)) *LanguageChoiceQuery {
	query := (&RaceClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcq.withRace = query
	return lcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Choose int `json:"choose,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LanguageChoice.Query().
//		GroupBy(languagechoice.FieldChoose).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lcq *LanguageChoiceQuery) GroupBy(field string, fields ...string) *LanguageChoiceGroupBy {
	lcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LanguageChoiceGroupBy{build: lcq}
	grbuild.flds = &lcq.ctx.Fields
	grbuild.label = languagechoice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Choose int `json:"choose,omitempty"`
//	}
//
//	client.LanguageChoice.Query().
//		Select(languagechoice.FieldChoose).
//		Scan(ctx, &v)
func (lcq *LanguageChoiceQuery) Select(fields ...string) *LanguageChoiceSelect {
	lcq.ctx.Fields = append(lcq.ctx.Fields, fields...)
	sbuild := &LanguageChoiceSelect{LanguageChoiceQuery: lcq}
	sbuild.label = languagechoice.Label
	sbuild.flds, sbuild.scan = &lcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LanguageChoiceSelect configured with the given aggregations.
func (lcq *LanguageChoiceQuery) Aggregate(fns ...AggregateFunc) *LanguageChoiceSelect {
	return lcq.Select().Aggregate(fns...)
}

func (lcq *LanguageChoiceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lcq); err != nil {
				return err
			}
		}
	}
	for _, f := range lcq.ctx.Fields {
		if !languagechoice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lcq.path != nil {
		prev, err := lcq.path(ctx)
		if err != nil {
			return err
		}
		lcq.sql = prev
	}
	return nil
}

func (lcq *LanguageChoiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LanguageChoice, error) {
	var (
		nodes       = []*LanguageChoice{}
		withFKs     = lcq.withFKs
		_spec       = lcq.querySpec()
		loadedTypes = [2]bool{
			lcq.withLanguages != nil,
			lcq.withRace != nil,
		}
	)
	if lcq.withRace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, languagechoice.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LanguageChoice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LanguageChoice{config: lcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lcq.modifiers) > 0 {
		_spec.Modifiers = lcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lcq.withLanguages; query != nil {
		if err := lcq.loadLanguages(ctx, query, nodes,
			func(n *LanguageChoice) { n.Edges.Languages = []*Language{} },
			func(n *LanguageChoice, e *Language) { n.Edges.Languages = append(n.Edges.Languages, e) }); err != nil {
			return nil, err
		}
	}
	if query := lcq.withRace; query != nil {
		if err := lcq.loadRace(ctx, query, nodes, nil,
			func(n *LanguageChoice, e *Race) { n.Edges.Race = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range lcq.withNamedLanguages {
		if err := lcq.loadLanguages(ctx, query, nodes,
			func(n *LanguageChoice) { n.appendNamedLanguages(name) },
			func(n *LanguageChoice, e *Language) { n.appendNamedLanguages(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range lcq.loadTotal {
		if err := lcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lcq *LanguageChoiceQuery) loadLanguages(ctx context.Context, query *LanguageQuery, nodes []*LanguageChoice, init func(*LanguageChoice), assign func(*LanguageChoice, *Language)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*LanguageChoice)
	nids := make(map[int]map[*LanguageChoice]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(languagechoice.LanguagesTable)
		s.Join(joinT).On(s.C(language.FieldID), joinT.C(languagechoice.LanguagesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(languagechoice.LanguagesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(languagechoice.LanguagesPrimaryKey[0]))
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
					nids[inValue] = map[*LanguageChoice]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Language](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "languages" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lcq *LanguageChoiceQuery) loadRace(ctx context.Context, query *RaceQuery, nodes []*LanguageChoice, init func(*LanguageChoice), assign func(*LanguageChoice, *Race)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*LanguageChoice)
	for i := range nodes {
		if nodes[i].race_language_options == nil {
			continue
		}
		fk := *nodes[i].race_language_options
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(race.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "race_language_options" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lcq *LanguageChoiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcq.querySpec()
	if len(lcq.modifiers) > 0 {
		_spec.Modifiers = lcq.modifiers
	}
	_spec.Node.Columns = lcq.ctx.Fields
	if len(lcq.ctx.Fields) > 0 {
		_spec.Unique = lcq.ctx.Unique != nil && *lcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lcq.driver, _spec)
}

func (lcq *LanguageChoiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(languagechoice.Table, languagechoice.Columns, sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt))
	_spec.From = lcq.sql
	if unique := lcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lcq.path != nil {
		_spec.Unique = true
	}
	if fields := lcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, languagechoice.FieldID)
		for i := range fields {
			if fields[i] != languagechoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcq *LanguageChoiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcq.driver.Dialect())
	t1 := builder.Table(languagechoice.Table)
	columns := lcq.ctx.Fields
	if len(columns) == 0 {
		columns = languagechoice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lcq.sql != nil {
		selector = lcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lcq.ctx.Unique != nil && *lcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lcq.predicates {
		p(selector)
	}
	for _, p := range lcq.order {
		p(selector)
	}
	if offset := lcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedLanguages tells the query-builder to eager-load the nodes that are connected to the "languages"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (lcq *LanguageChoiceQuery) WithNamedLanguages(name string, opts ...func(*LanguageQuery)) *LanguageChoiceQuery {
	query := (&LanguageClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if lcq.withNamedLanguages == nil {
		lcq.withNamedLanguages = make(map[string]*LanguageQuery)
	}
	lcq.withNamedLanguages[name] = query
	return lcq
}

// LanguageChoiceGroupBy is the group-by builder for LanguageChoice entities.
type LanguageChoiceGroupBy struct {
	selector
	build *LanguageChoiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcgb *LanguageChoiceGroupBy) Aggregate(fns ...AggregateFunc) *LanguageChoiceGroupBy {
	lcgb.fns = append(lcgb.fns, fns...)
	return lcgb
}

// Scan applies the selector query and scans the result into the given value.
func (lcgb *LanguageChoiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcgb.build.ctx, ent.OpQueryGroupBy)
	if err := lcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LanguageChoiceQuery, *LanguageChoiceGroupBy](ctx, lcgb.build, lcgb, lcgb.build.inters, v)
}

func (lcgb *LanguageChoiceGroupBy) sqlScan(ctx context.Context, root *LanguageChoiceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lcgb.fns))
	for _, fn := range lcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lcgb.flds)+len(lcgb.fns))
		for _, f := range *lcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LanguageChoiceSelect is the builder for selecting fields of LanguageChoice entities.
type LanguageChoiceSelect struct {
	*LanguageChoiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lcs *LanguageChoiceSelect) Aggregate(fns ...AggregateFunc) *LanguageChoiceSelect {
	lcs.fns = append(lcs.fns, fns...)
	return lcs
}

// Scan applies the selector query and scans the result into the given value.
func (lcs *LanguageChoiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcs.ctx, ent.OpQuerySelect)
	if err := lcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LanguageChoiceQuery, *LanguageChoiceSelect](ctx, lcs.LanguageChoiceQuery, lcs, lcs.inters, v)
}

func (lcs *LanguageChoiceSelect) sqlScan(ctx context.Context, root *LanguageChoiceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lcs.fns))
	for _, fn := range lcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
