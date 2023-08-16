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
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
)

// LanguageQuery is the builder for querying Language entities.
type LanguageQuery struct {
	config
	ctx               *QueryContext
	order             []language.OrderOption
	inters            []Interceptor
	predicates        []predicate.Language
	withSpeakers      *RaceQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Language) error
	withNamedSpeakers map[string]*RaceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LanguageQuery builder.
func (lq *LanguageQuery) Where(ps ...predicate.Language) *LanguageQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LanguageQuery) Limit(limit int) *LanguageQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LanguageQuery) Offset(offset int) *LanguageQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LanguageQuery) Unique(unique bool) *LanguageQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LanguageQuery) Order(o ...language.OrderOption) *LanguageQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QuerySpeakers chains the current query on the "speakers" edge.
func (lq *LanguageQuery) QuerySpeakers() *RaceQuery {
	query := (&RaceClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(language.Table, language.FieldID, selector),
			sqlgraph.To(race.Table, race.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, language.SpeakersTable, language.SpeakersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Language entity from the query.
// Returns a *NotFoundError when no Language was found.
func (lq *LanguageQuery) First(ctx context.Context) (*Language, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{language.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LanguageQuery) FirstX(ctx context.Context) *Language {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Language ID from the query.
// Returns a *NotFoundError when no Language ID was found.
func (lq *LanguageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{language.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LanguageQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Language entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Language entity is found.
// Returns a *NotFoundError when no Language entities are found.
func (lq *LanguageQuery) Only(ctx context.Context) (*Language, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{language.Label}
	default:
		return nil, &NotSingularError{language.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LanguageQuery) OnlyX(ctx context.Context) *Language {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Language ID in the query.
// Returns a *NotSingularError when more than one Language ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LanguageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{language.Label}
	default:
		err = &NotSingularError{language.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LanguageQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Languages.
func (lq *LanguageQuery) All(ctx context.Context) ([]*Language, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Language, *LanguageQuery]()
	return withInterceptors[[]*Language](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LanguageQuery) AllX(ctx context.Context) []*Language {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Language IDs.
func (lq *LanguageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(language.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LanguageQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LanguageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LanguageQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LanguageQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LanguageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LanguageQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LanguageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LanguageQuery) Clone() *LanguageQuery {
	if lq == nil {
		return nil
	}
	return &LanguageQuery{
		config:       lq.config,
		ctx:          lq.ctx.Clone(),
		order:        append([]language.OrderOption{}, lq.order...),
		inters:       append([]Interceptor{}, lq.inters...),
		predicates:   append([]predicate.Language{}, lq.predicates...),
		withSpeakers: lq.withSpeakers.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithSpeakers tells the query-builder to eager-load the nodes that are connected to
// the "speakers" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LanguageQuery) WithSpeakers(opts ...func(*RaceQuery)) *LanguageQuery {
	query := (&RaceClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withSpeakers = query
	return lq
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
//	client.Language.Query().
//		GroupBy(language.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LanguageQuery) GroupBy(field string, fields ...string) *LanguageGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LanguageGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = language.Label
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
//	client.Language.Query().
//		Select(language.FieldIndx).
//		Scan(ctx, &v)
func (lq *LanguageQuery) Select(fields ...string) *LanguageSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LanguageSelect{LanguageQuery: lq}
	sbuild.label = language.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LanguageSelect configured with the given aggregations.
func (lq *LanguageQuery) Aggregate(fns ...AggregateFunc) *LanguageSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LanguageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !language.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LanguageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Language, error) {
	var (
		nodes       = []*Language{}
		_spec       = lq.querySpec()
		loadedTypes = [1]bool{
			lq.withSpeakers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Language).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Language{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withSpeakers; query != nil {
		if err := lq.loadSpeakers(ctx, query, nodes,
			func(n *Language) { n.Edges.Speakers = []*Race{} },
			func(n *Language, e *Race) { n.Edges.Speakers = append(n.Edges.Speakers, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range lq.withNamedSpeakers {
		if err := lq.loadSpeakers(ctx, query, nodes,
			func(n *Language) { n.appendNamedSpeakers(name) },
			func(n *Language, e *Race) { n.appendNamedSpeakers(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range lq.loadTotal {
		if err := lq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LanguageQuery) loadSpeakers(ctx context.Context, query *RaceQuery, nodes []*Language, init func(*Language), assign func(*Language, *Race)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Language)
	nids := make(map[int]map[*Language]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(language.SpeakersTable)
		s.Join(joinT).On(s.C(race.FieldID), joinT.C(language.SpeakersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(language.SpeakersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(language.SpeakersPrimaryKey[1]))
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
					nids[inValue] = map[*Language]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Race](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "speakers" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (lq *LanguageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LanguageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, language.FieldID)
		for i := range fields {
			if fields[i] != language.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LanguageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(language.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = language.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSpeakers tells the query-builder to eager-load the nodes that are connected to the "speakers"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (lq *LanguageQuery) WithNamedSpeakers(name string, opts ...func(*RaceQuery)) *LanguageQuery {
	query := (&RaceClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if lq.withNamedSpeakers == nil {
		lq.withNamedSpeakers = make(map[string]*RaceQuery)
	}
	lq.withNamedSpeakers[name] = query
	return lq
}

// LanguageGroupBy is the group-by builder for Language entities.
type LanguageGroupBy struct {
	selector
	build *LanguageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LanguageGroupBy) Aggregate(fns ...AggregateFunc) *LanguageGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LanguageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LanguageQuery, *LanguageGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LanguageGroupBy) sqlScan(ctx context.Context, root *LanguageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LanguageSelect is the builder for selecting fields of Language entities.
type LanguageSelect struct {
	*LanguageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LanguageSelect) Aggregate(fns ...AggregateFunc) *LanguageSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LanguageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LanguageQuery, *LanguageSelect](ctx, ls.LanguageQuery, ls, ls.inters, v)
}

func (ls *LanguageSelect) sqlScan(ctx context.Context, root *LanguageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
