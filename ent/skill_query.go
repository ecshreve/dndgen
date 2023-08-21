// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/skill"
)

// SkillQuery is the builder for querying Skill entities.
type SkillQuery struct {
	config
	ctx              *QueryContext
	order            []skill.OrderOption
	inters           []Interceptor
	predicates       []predicate.Skill
	withAbilityScore *AbilityScoreQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Skill) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SkillQuery builder.
func (sq *SkillQuery) Where(ps ...predicate.Skill) *SkillQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SkillQuery) Limit(limit int) *SkillQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SkillQuery) Offset(offset int) *SkillQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SkillQuery) Unique(unique bool) *SkillQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SkillQuery) Order(o ...skill.OrderOption) *SkillQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryAbilityScore chains the current query on the "ability_score" edge.
func (sq *SkillQuery) QueryAbilityScore() *AbilityScoreQuery {
	query := (&AbilityScoreClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(skill.Table, skill.FieldID, selector),
			sqlgraph.To(abilityscore.Table, abilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, skill.AbilityScoreTable, skill.AbilityScoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Skill entity from the query.
// Returns a *NotFoundError when no Skill was found.
func (sq *SkillQuery) First(ctx context.Context) (*Skill, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SkillQuery) FirstX(ctx context.Context) *Skill {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Skill ID from the query.
// Returns a *NotFoundError when no Skill ID was found.
func (sq *SkillQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SkillQuery) FirstIDX(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Skill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Skill entity is found.
// Returns a *NotFoundError when no Skill entities are found.
func (sq *SkillQuery) Only(ctx context.Context) (*Skill, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skill.Label}
	default:
		return nil, &NotSingularError{skill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SkillQuery) OnlyX(ctx context.Context) *Skill {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Skill ID in the query.
// Returns a *NotSingularError when more than one Skill ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SkillQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skill.Label}
	default:
		err = &NotSingularError{skill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SkillQuery) OnlyIDX(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Skills.
func (sq *SkillQuery) All(ctx context.Context) ([]*Skill, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Skill, *SkillQuery]()
	return withInterceptors[[]*Skill](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SkillQuery) AllX(ctx context.Context) []*Skill {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Skill IDs.
func (sq *SkillQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(skill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SkillQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SkillQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SkillQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SkillQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SkillQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SkillQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SkillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SkillQuery) Clone() *SkillQuery {
	if sq == nil {
		return nil
	}
	return &SkillQuery{
		config:           sq.config,
		ctx:              sq.ctx.Clone(),
		order:            append([]skill.OrderOption{}, sq.order...),
		inters:           append([]Interceptor{}, sq.inters...),
		predicates:       append([]predicate.Skill{}, sq.predicates...),
		withAbilityScore: sq.withAbilityScore.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithAbilityScore tells the query-builder to eager-load the nodes that are connected to
// the "ability_score" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SkillQuery) WithAbilityScore(opts ...func(*AbilityScoreQuery)) *SkillQuery {
	query := (&AbilityScoreClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withAbilityScore = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Skill.Query().
//		GroupBy(skill.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SkillQuery) GroupBy(field string, fields ...string) *SkillGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SkillGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = skill.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Skill.Query().
//		Select(skill.FieldName).
//		Scan(ctx, &v)
func (sq *SkillQuery) Select(fields ...string) *SkillSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SkillSelect{SkillQuery: sq}
	sbuild.label = skill.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SkillSelect configured with the given aggregations.
func (sq *SkillQuery) Aggregate(fns ...AggregateFunc) *SkillSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SkillQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !skill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SkillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Skill, error) {
	var (
		nodes       = []*Skill{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withAbilityScore != nil,
		}
	)
	if sq.withAbilityScore != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, skill.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Skill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Skill{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withAbilityScore; query != nil {
		if err := sq.loadAbilityScore(ctx, query, nodes, nil,
			func(n *Skill, e *AbilityScore) { n.Edges.AbilityScore = e }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SkillQuery) loadAbilityScore(ctx context.Context, query *AbilityScoreQuery, nodes []*Skill, init func(*Skill), assign func(*Skill, *AbilityScore)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Skill)
	for i := range nodes {
		if nodes[i].skill_ability_score == nil {
			continue
		}
		fk := *nodes[i].skill_ability_score
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(abilityscore.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "skill_ability_score" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SkillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SkillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(skill.Table, skill.Columns, sqlgraph.NewFieldSpec(skill.FieldID, field.TypeString))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skill.FieldID)
		for i := range fields {
			if fields[i] != skill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SkillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(skill.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = skill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SkillGroupBy is the group-by builder for Skill entities.
type SkillGroupBy struct {
	selector
	build *SkillQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SkillGroupBy) Aggregate(fns ...AggregateFunc) *SkillGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SkillGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SkillQuery, *SkillGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SkillGroupBy) sqlScan(ctx context.Context, root *SkillQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SkillSelect is the builder for selecting fields of Skill entities.
type SkillSelect struct {
	*SkillQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SkillSelect) Aggregate(fns ...AggregateFunc) *SkillSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SkillSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SkillQuery, *SkillSelect](ctx, ss.SkillQuery, ss, ss.inters, v)
}

func (ss *SkillSelect) sqlScan(ctx context.Context, root *SkillQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
