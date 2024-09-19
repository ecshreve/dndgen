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
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/skill"
)

// AbilityScoreQuery is the builder for querying AbilityScore entities.
type AbilityScoreQuery struct {
	config
	ctx             *QueryContext
	order           []abilityscore.OrderOption
	inters          []Interceptor
	predicates      []predicate.AbilityScore
	withSkills      *SkillQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*AbilityScore) error
	withNamedSkills map[string]*SkillQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AbilityScoreQuery builder.
func (asq *AbilityScoreQuery) Where(ps ...predicate.AbilityScore) *AbilityScoreQuery {
	asq.predicates = append(asq.predicates, ps...)
	return asq
}

// Limit the number of records to be returned by this query.
func (asq *AbilityScoreQuery) Limit(limit int) *AbilityScoreQuery {
	asq.ctx.Limit = &limit
	return asq
}

// Offset to start from.
func (asq *AbilityScoreQuery) Offset(offset int) *AbilityScoreQuery {
	asq.ctx.Offset = &offset
	return asq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asq *AbilityScoreQuery) Unique(unique bool) *AbilityScoreQuery {
	asq.ctx.Unique = &unique
	return asq
}

// Order specifies how the records should be ordered.
func (asq *AbilityScoreQuery) Order(o ...abilityscore.OrderOption) *AbilityScoreQuery {
	asq.order = append(asq.order, o...)
	return asq
}

// QuerySkills chains the current query on the "skills" edge.
func (asq *AbilityScoreQuery) QuerySkills() *SkillQuery {
	query := (&SkillClient{config: asq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := asq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilityscore.Table, abilityscore.FieldID, selector),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, abilityscore.SkillsTable, abilityscore.SkillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(asq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AbilityScore entity from the query.
// Returns a *NotFoundError when no AbilityScore was found.
func (asq *AbilityScoreQuery) First(ctx context.Context) (*AbilityScore, error) {
	nodes, err := asq.Limit(1).All(setContextOp(ctx, asq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{abilityscore.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asq *AbilityScoreQuery) FirstX(ctx context.Context) *AbilityScore {
	node, err := asq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AbilityScore ID from the query.
// Returns a *NotFoundError when no AbilityScore ID was found.
func (asq *AbilityScoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = asq.Limit(1).IDs(setContextOp(ctx, asq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{abilityscore.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asq *AbilityScoreQuery) FirstIDX(ctx context.Context) int {
	id, err := asq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AbilityScore entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AbilityScore entity is found.
// Returns a *NotFoundError when no AbilityScore entities are found.
func (asq *AbilityScoreQuery) Only(ctx context.Context) (*AbilityScore, error) {
	nodes, err := asq.Limit(2).All(setContextOp(ctx, asq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{abilityscore.Label}
	default:
		return nil, &NotSingularError{abilityscore.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asq *AbilityScoreQuery) OnlyX(ctx context.Context) *AbilityScore {
	node, err := asq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AbilityScore ID in the query.
// Returns a *NotSingularError when more than one AbilityScore ID is found.
// Returns a *NotFoundError when no entities are found.
func (asq *AbilityScoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = asq.Limit(2).IDs(setContextOp(ctx, asq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{abilityscore.Label}
	default:
		err = &NotSingularError{abilityscore.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asq *AbilityScoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := asq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AbilityScores.
func (asq *AbilityScoreQuery) All(ctx context.Context) ([]*AbilityScore, error) {
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryAll)
	if err := asq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AbilityScore, *AbilityScoreQuery]()
	return withInterceptors[[]*AbilityScore](ctx, asq, qr, asq.inters)
}

// AllX is like All, but panics if an error occurs.
func (asq *AbilityScoreQuery) AllX(ctx context.Context) []*AbilityScore {
	nodes, err := asq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AbilityScore IDs.
func (asq *AbilityScoreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if asq.ctx.Unique == nil && asq.path != nil {
		asq.Unique(true)
	}
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryIDs)
	if err = asq.Select(abilityscore.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asq *AbilityScoreQuery) IDsX(ctx context.Context) []int {
	ids, err := asq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asq *AbilityScoreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryCount)
	if err := asq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, asq, querierCount[*AbilityScoreQuery](), asq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (asq *AbilityScoreQuery) CountX(ctx context.Context) int {
	count, err := asq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asq *AbilityScoreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, asq.ctx, ent.OpQueryExist)
	switch _, err := asq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (asq *AbilityScoreQuery) ExistX(ctx context.Context) bool {
	exist, err := asq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AbilityScoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asq *AbilityScoreQuery) Clone() *AbilityScoreQuery {
	if asq == nil {
		return nil
	}
	return &AbilityScoreQuery{
		config:     asq.config,
		ctx:        asq.ctx.Clone(),
		order:      append([]abilityscore.OrderOption{}, asq.order...),
		inters:     append([]Interceptor{}, asq.inters...),
		predicates: append([]predicate.AbilityScore{}, asq.predicates...),
		withSkills: asq.withSkills.Clone(),
		// clone intermediate query.
		sql:  asq.sql.Clone(),
		path: asq.path,
	}
}

// WithSkills tells the query-builder to eager-load the nodes that are connected to
// the "skills" edge. The optional arguments are used to configure the query builder of the edge.
func (asq *AbilityScoreQuery) WithSkills(opts ...func(*SkillQuery)) *AbilityScoreQuery {
	query := (&SkillClient{config: asq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	asq.withSkills = query
	return asq
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
//	client.AbilityScore.Query().
//		GroupBy(abilityscore.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (asq *AbilityScoreQuery) GroupBy(field string, fields ...string) *AbilityScoreGroupBy {
	asq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AbilityScoreGroupBy{build: asq}
	grbuild.flds = &asq.ctx.Fields
	grbuild.label = abilityscore.Label
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
//	client.AbilityScore.Query().
//		Select(abilityscore.FieldIndx).
//		Scan(ctx, &v)
func (asq *AbilityScoreQuery) Select(fields ...string) *AbilityScoreSelect {
	asq.ctx.Fields = append(asq.ctx.Fields, fields...)
	sbuild := &AbilityScoreSelect{AbilityScoreQuery: asq}
	sbuild.label = abilityscore.Label
	sbuild.flds, sbuild.scan = &asq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AbilityScoreSelect configured with the given aggregations.
func (asq *AbilityScoreQuery) Aggregate(fns ...AggregateFunc) *AbilityScoreSelect {
	return asq.Select().Aggregate(fns...)
}

func (asq *AbilityScoreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range asq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, asq); err != nil {
				return err
			}
		}
	}
	for _, f := range asq.ctx.Fields {
		if !abilityscore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if asq.path != nil {
		prev, err := asq.path(ctx)
		if err != nil {
			return err
		}
		asq.sql = prev
	}
	return nil
}

func (asq *AbilityScoreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AbilityScore, error) {
	var (
		nodes       = []*AbilityScore{}
		_spec       = asq.querySpec()
		loadedTypes = [1]bool{
			asq.withSkills != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AbilityScore).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AbilityScore{config: asq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(asq.modifiers) > 0 {
		_spec.Modifiers = asq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := asq.withSkills; query != nil {
		if err := asq.loadSkills(ctx, query, nodes,
			func(n *AbilityScore) { n.Edges.Skills = []*Skill{} },
			func(n *AbilityScore, e *Skill) { n.Edges.Skills = append(n.Edges.Skills, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range asq.withNamedSkills {
		if err := asq.loadSkills(ctx, query, nodes,
			func(n *AbilityScore) { n.appendNamedSkills(name) },
			func(n *AbilityScore, e *Skill) { n.appendNamedSkills(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range asq.loadTotal {
		if err := asq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (asq *AbilityScoreQuery) loadSkills(ctx context.Context, query *SkillQuery, nodes []*AbilityScore, init func(*AbilityScore), assign func(*AbilityScore, *Skill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*AbilityScore)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Skill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(abilityscore.SkillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ability_score_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "ability_score_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ability_score_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (asq *AbilityScoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asq.querySpec()
	if len(asq.modifiers) > 0 {
		_spec.Modifiers = asq.modifiers
	}
	_spec.Node.Columns = asq.ctx.Fields
	if len(asq.ctx.Fields) > 0 {
		_spec.Unique = asq.ctx.Unique != nil && *asq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, asq.driver, _spec)
}

func (asq *AbilityScoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(abilityscore.Table, abilityscore.Columns, sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt))
	_spec.From = asq.sql
	if unique := asq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if asq.path != nil {
		_spec.Unique = true
	}
	if fields := asq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, abilityscore.FieldID)
		for i := range fields {
			if fields[i] != abilityscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := asq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asq *AbilityScoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asq.driver.Dialect())
	t1 := builder.Table(abilityscore.Table)
	columns := asq.ctx.Fields
	if len(columns) == 0 {
		columns = abilityscore.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asq.sql != nil {
		selector = asq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asq.ctx.Unique != nil && *asq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range asq.predicates {
		p(selector)
	}
	for _, p := range asq.order {
		p(selector)
	}
	if offset := asq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSkills tells the query-builder to eager-load the nodes that are connected to the "skills"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (asq *AbilityScoreQuery) WithNamedSkills(name string, opts ...func(*SkillQuery)) *AbilityScoreQuery {
	query := (&SkillClient{config: asq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if asq.withNamedSkills == nil {
		asq.withNamedSkills = make(map[string]*SkillQuery)
	}
	asq.withNamedSkills[name] = query
	return asq
}

// AbilityScoreGroupBy is the group-by builder for AbilityScore entities.
type AbilityScoreGroupBy struct {
	selector
	build *AbilityScoreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asgb *AbilityScoreGroupBy) Aggregate(fns ...AggregateFunc) *AbilityScoreGroupBy {
	asgb.fns = append(asgb.fns, fns...)
	return asgb
}

// Scan applies the selector query and scans the result into the given value.
func (asgb *AbilityScoreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, asgb.build.ctx, ent.OpQueryGroupBy)
	if err := asgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AbilityScoreQuery, *AbilityScoreGroupBy](ctx, asgb.build, asgb, asgb.build.inters, v)
}

func (asgb *AbilityScoreGroupBy) sqlScan(ctx context.Context, root *AbilityScoreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(asgb.fns))
	for _, fn := range asgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*asgb.flds)+len(asgb.fns))
		for _, f := range *asgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*asgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AbilityScoreSelect is the builder for selecting fields of AbilityScore entities.
type AbilityScoreSelect struct {
	*AbilityScoreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ass *AbilityScoreSelect) Aggregate(fns ...AggregateFunc) *AbilityScoreSelect {
	ass.fns = append(ass.fns, fns...)
	return ass
}

// Scan applies the selector query and scans the result into the given value.
func (ass *AbilityScoreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ass.ctx, ent.OpQuerySelect)
	if err := ass.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AbilityScoreQuery, *AbilityScoreSelect](ctx, ass.AbilityScoreQuery, ass, ass.inters, v)
}

func (ass *AbilityScoreSelect) sqlScan(ctx context.Context, root *AbilityScoreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ass.fns))
	for _, fn := range ass.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ass.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
