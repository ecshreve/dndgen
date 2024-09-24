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
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilitybonuschoice"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
)

// AbilityBonusQuery is the builder for querying AbilityBonus entities.
type AbilityBonusQuery struct {
	config
	ctx              *QueryContext
	order            []abilitybonus.OrderOption
	inters           []Interceptor
	predicates       []predicate.AbilityBonus
	withAbilityScore *AbilityScoreQuery
	withRace         *RaceQuery
	withOptions      *AbilityBonusChoiceQuery
	withSubrace      *SubraceQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*AbilityBonus) error
	withNamedRace    map[string]*RaceQuery
	withNamedOptions map[string]*AbilityBonusChoiceQuery
	withNamedSubrace map[string]*SubraceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AbilityBonusQuery builder.
func (abq *AbilityBonusQuery) Where(ps ...predicate.AbilityBonus) *AbilityBonusQuery {
	abq.predicates = append(abq.predicates, ps...)
	return abq
}

// Limit the number of records to be returned by this query.
func (abq *AbilityBonusQuery) Limit(limit int) *AbilityBonusQuery {
	abq.ctx.Limit = &limit
	return abq
}

// Offset to start from.
func (abq *AbilityBonusQuery) Offset(offset int) *AbilityBonusQuery {
	abq.ctx.Offset = &offset
	return abq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (abq *AbilityBonusQuery) Unique(unique bool) *AbilityBonusQuery {
	abq.ctx.Unique = &unique
	return abq
}

// Order specifies how the records should be ordered.
func (abq *AbilityBonusQuery) Order(o ...abilitybonus.OrderOption) *AbilityBonusQuery {
	abq.order = append(abq.order, o...)
	return abq
}

// QueryAbilityScore chains the current query on the "ability_score" edge.
func (abq *AbilityBonusQuery) QueryAbilityScore() *AbilityScoreQuery {
	query := (&AbilityScoreClient{config: abq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilitybonus.Table, abilitybonus.FieldID, selector),
			sqlgraph.To(abilityscore.Table, abilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, abilitybonus.AbilityScoreTable, abilitybonus.AbilityScoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(abq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRace chains the current query on the "race" edge.
func (abq *AbilityBonusQuery) QueryRace() *RaceQuery {
	query := (&RaceClient{config: abq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilitybonus.Table, abilitybonus.FieldID, selector),
			sqlgraph.To(race.Table, race.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, abilitybonus.RaceTable, abilitybonus.RacePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(abq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOptions chains the current query on the "options" edge.
func (abq *AbilityBonusQuery) QueryOptions() *AbilityBonusChoiceQuery {
	query := (&AbilityBonusChoiceClient{config: abq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilitybonus.Table, abilitybonus.FieldID, selector),
			sqlgraph.To(abilitybonuschoice.Table, abilitybonuschoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, abilitybonus.OptionsTable, abilitybonus.OptionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(abq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubrace chains the current query on the "subrace" edge.
func (abq *AbilityBonusQuery) QuerySubrace() *SubraceQuery {
	query := (&SubraceClient{config: abq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilitybonus.Table, abilitybonus.FieldID, selector),
			sqlgraph.To(subrace.Table, subrace.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, abilitybonus.SubraceTable, abilitybonus.SubracePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(abq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AbilityBonus entity from the query.
// Returns a *NotFoundError when no AbilityBonus was found.
func (abq *AbilityBonusQuery) First(ctx context.Context) (*AbilityBonus, error) {
	nodes, err := abq.Limit(1).All(setContextOp(ctx, abq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{abilitybonus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (abq *AbilityBonusQuery) FirstX(ctx context.Context) *AbilityBonus {
	node, err := abq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AbilityBonus ID from the query.
// Returns a *NotFoundError when no AbilityBonus ID was found.
func (abq *AbilityBonusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = abq.Limit(1).IDs(setContextOp(ctx, abq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{abilitybonus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (abq *AbilityBonusQuery) FirstIDX(ctx context.Context) int {
	id, err := abq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AbilityBonus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AbilityBonus entity is found.
// Returns a *NotFoundError when no AbilityBonus entities are found.
func (abq *AbilityBonusQuery) Only(ctx context.Context) (*AbilityBonus, error) {
	nodes, err := abq.Limit(2).All(setContextOp(ctx, abq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{abilitybonus.Label}
	default:
		return nil, &NotSingularError{abilitybonus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (abq *AbilityBonusQuery) OnlyX(ctx context.Context) *AbilityBonus {
	node, err := abq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AbilityBonus ID in the query.
// Returns a *NotSingularError when more than one AbilityBonus ID is found.
// Returns a *NotFoundError when no entities are found.
func (abq *AbilityBonusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = abq.Limit(2).IDs(setContextOp(ctx, abq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{abilitybonus.Label}
	default:
		err = &NotSingularError{abilitybonus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (abq *AbilityBonusQuery) OnlyIDX(ctx context.Context) int {
	id, err := abq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AbilityBonusSlice.
func (abq *AbilityBonusQuery) All(ctx context.Context) ([]*AbilityBonus, error) {
	ctx = setContextOp(ctx, abq.ctx, ent.OpQueryAll)
	if err := abq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AbilityBonus, *AbilityBonusQuery]()
	return withInterceptors[[]*AbilityBonus](ctx, abq, qr, abq.inters)
}

// AllX is like All, but panics if an error occurs.
func (abq *AbilityBonusQuery) AllX(ctx context.Context) []*AbilityBonus {
	nodes, err := abq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AbilityBonus IDs.
func (abq *AbilityBonusQuery) IDs(ctx context.Context) (ids []int, err error) {
	if abq.ctx.Unique == nil && abq.path != nil {
		abq.Unique(true)
	}
	ctx = setContextOp(ctx, abq.ctx, ent.OpQueryIDs)
	if err = abq.Select(abilitybonus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (abq *AbilityBonusQuery) IDsX(ctx context.Context) []int {
	ids, err := abq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (abq *AbilityBonusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, abq.ctx, ent.OpQueryCount)
	if err := abq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, abq, querierCount[*AbilityBonusQuery](), abq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (abq *AbilityBonusQuery) CountX(ctx context.Context) int {
	count, err := abq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (abq *AbilityBonusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, abq.ctx, ent.OpQueryExist)
	switch _, err := abq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (abq *AbilityBonusQuery) ExistX(ctx context.Context) bool {
	exist, err := abq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AbilityBonusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (abq *AbilityBonusQuery) Clone() *AbilityBonusQuery {
	if abq == nil {
		return nil
	}
	return &AbilityBonusQuery{
		config:           abq.config,
		ctx:              abq.ctx.Clone(),
		order:            append([]abilitybonus.OrderOption{}, abq.order...),
		inters:           append([]Interceptor{}, abq.inters...),
		predicates:       append([]predicate.AbilityBonus{}, abq.predicates...),
		withAbilityScore: abq.withAbilityScore.Clone(),
		withRace:         abq.withRace.Clone(),
		withOptions:      abq.withOptions.Clone(),
		withSubrace:      abq.withSubrace.Clone(),
		// clone intermediate query.
		sql:  abq.sql.Clone(),
		path: abq.path,
	}
}

// WithAbilityScore tells the query-builder to eager-load the nodes that are connected to
// the "ability_score" edge. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithAbilityScore(opts ...func(*AbilityScoreQuery)) *AbilityBonusQuery {
	query := (&AbilityScoreClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	abq.withAbilityScore = query
	return abq
}

// WithRace tells the query-builder to eager-load the nodes that are connected to
// the "race" edge. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithRace(opts ...func(*RaceQuery)) *AbilityBonusQuery {
	query := (&RaceClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	abq.withRace = query
	return abq
}

// WithOptions tells the query-builder to eager-load the nodes that are connected to
// the "options" edge. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithOptions(opts ...func(*AbilityBonusChoiceQuery)) *AbilityBonusQuery {
	query := (&AbilityBonusChoiceClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	abq.withOptions = query
	return abq
}

// WithSubrace tells the query-builder to eager-load the nodes that are connected to
// the "subrace" edge. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithSubrace(opts ...func(*SubraceQuery)) *AbilityBonusQuery {
	query := (&SubraceClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	abq.withSubrace = query
	return abq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Bonus int `json:"bonus,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AbilityBonus.Query().
//		GroupBy(abilitybonus.FieldBonus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (abq *AbilityBonusQuery) GroupBy(field string, fields ...string) *AbilityBonusGroupBy {
	abq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AbilityBonusGroupBy{build: abq}
	grbuild.flds = &abq.ctx.Fields
	grbuild.label = abilitybonus.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Bonus int `json:"bonus,omitempty"`
//	}
//
//	client.AbilityBonus.Query().
//		Select(abilitybonus.FieldBonus).
//		Scan(ctx, &v)
func (abq *AbilityBonusQuery) Select(fields ...string) *AbilityBonusSelect {
	abq.ctx.Fields = append(abq.ctx.Fields, fields...)
	sbuild := &AbilityBonusSelect{AbilityBonusQuery: abq}
	sbuild.label = abilitybonus.Label
	sbuild.flds, sbuild.scan = &abq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AbilityBonusSelect configured with the given aggregations.
func (abq *AbilityBonusQuery) Aggregate(fns ...AggregateFunc) *AbilityBonusSelect {
	return abq.Select().Aggregate(fns...)
}

func (abq *AbilityBonusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range abq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, abq); err != nil {
				return err
			}
		}
	}
	for _, f := range abq.ctx.Fields {
		if !abilitybonus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if abq.path != nil {
		prev, err := abq.path(ctx)
		if err != nil {
			return err
		}
		abq.sql = prev
	}
	return nil
}

func (abq *AbilityBonusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AbilityBonus, error) {
	var (
		nodes       = []*AbilityBonus{}
		withFKs     = abq.withFKs
		_spec       = abq.querySpec()
		loadedTypes = [4]bool{
			abq.withAbilityScore != nil,
			abq.withRace != nil,
			abq.withOptions != nil,
			abq.withSubrace != nil,
		}
	)
	if abq.withAbilityScore != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, abilitybonus.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AbilityBonus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AbilityBonus{config: abq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(abq.modifiers) > 0 {
		_spec.Modifiers = abq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, abq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := abq.withAbilityScore; query != nil {
		if err := abq.loadAbilityScore(ctx, query, nodes, nil,
			func(n *AbilityBonus, e *AbilityScore) { n.Edges.AbilityScore = e }); err != nil {
			return nil, err
		}
	}
	if query := abq.withRace; query != nil {
		if err := abq.loadRace(ctx, query, nodes,
			func(n *AbilityBonus) { n.Edges.Race = []*Race{} },
			func(n *AbilityBonus, e *Race) { n.Edges.Race = append(n.Edges.Race, e) }); err != nil {
			return nil, err
		}
	}
	if query := abq.withOptions; query != nil {
		if err := abq.loadOptions(ctx, query, nodes,
			func(n *AbilityBonus) { n.Edges.Options = []*AbilityBonusChoice{} },
			func(n *AbilityBonus, e *AbilityBonusChoice) { n.Edges.Options = append(n.Edges.Options, e) }); err != nil {
			return nil, err
		}
	}
	if query := abq.withSubrace; query != nil {
		if err := abq.loadSubrace(ctx, query, nodes,
			func(n *AbilityBonus) { n.Edges.Subrace = []*Subrace{} },
			func(n *AbilityBonus, e *Subrace) { n.Edges.Subrace = append(n.Edges.Subrace, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range abq.withNamedRace {
		if err := abq.loadRace(ctx, query, nodes,
			func(n *AbilityBonus) { n.appendNamedRace(name) },
			func(n *AbilityBonus, e *Race) { n.appendNamedRace(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range abq.withNamedOptions {
		if err := abq.loadOptions(ctx, query, nodes,
			func(n *AbilityBonus) { n.appendNamedOptions(name) },
			func(n *AbilityBonus, e *AbilityBonusChoice) { n.appendNamedOptions(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range abq.withNamedSubrace {
		if err := abq.loadSubrace(ctx, query, nodes,
			func(n *AbilityBonus) { n.appendNamedSubrace(name) },
			func(n *AbilityBonus, e *Subrace) { n.appendNamedSubrace(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range abq.loadTotal {
		if err := abq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (abq *AbilityBonusQuery) loadAbilityScore(ctx context.Context, query *AbilityScoreQuery, nodes []*AbilityBonus, init func(*AbilityBonus), assign func(*AbilityBonus, *AbilityScore)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AbilityBonus)
	for i := range nodes {
		if nodes[i].ability_bonus_ability_score == nil {
			continue
		}
		fk := *nodes[i].ability_bonus_ability_score
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
			return fmt.Errorf(`unexpected foreign-key "ability_bonus_ability_score" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (abq *AbilityBonusQuery) loadRace(ctx context.Context, query *RaceQuery, nodes []*AbilityBonus, init func(*AbilityBonus), assign func(*AbilityBonus, *Race)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*AbilityBonus)
	nids := make(map[int]map[*AbilityBonus]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(abilitybonus.RaceTable)
		s.Join(joinT).On(s.C(race.FieldID), joinT.C(abilitybonus.RacePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(abilitybonus.RacePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(abilitybonus.RacePrimaryKey[1]))
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
					nids[inValue] = map[*AbilityBonus]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "race" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (abq *AbilityBonusQuery) loadOptions(ctx context.Context, query *AbilityBonusChoiceQuery, nodes []*AbilityBonus, init func(*AbilityBonus), assign func(*AbilityBonus, *AbilityBonusChoice)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*AbilityBonus)
	nids := make(map[int]map[*AbilityBonus]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(abilitybonus.OptionsTable)
		s.Join(joinT).On(s.C(abilitybonuschoice.FieldID), joinT.C(abilitybonus.OptionsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(abilitybonus.OptionsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(abilitybonus.OptionsPrimaryKey[1]))
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
					nids[inValue] = map[*AbilityBonus]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*AbilityBonusChoice](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "options" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (abq *AbilityBonusQuery) loadSubrace(ctx context.Context, query *SubraceQuery, nodes []*AbilityBonus, init func(*AbilityBonus), assign func(*AbilityBonus, *Subrace)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*AbilityBonus)
	nids := make(map[int]map[*AbilityBonus]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(abilitybonus.SubraceTable)
		s.Join(joinT).On(s.C(subrace.FieldID), joinT.C(abilitybonus.SubracePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(abilitybonus.SubracePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(abilitybonus.SubracePrimaryKey[1]))
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
					nids[inValue] = map[*AbilityBonus]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Subrace](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "subrace" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (abq *AbilityBonusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := abq.querySpec()
	if len(abq.modifiers) > 0 {
		_spec.Modifiers = abq.modifiers
	}
	_spec.Node.Columns = abq.ctx.Fields
	if len(abq.ctx.Fields) > 0 {
		_spec.Unique = abq.ctx.Unique != nil && *abq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, abq.driver, _spec)
}

func (abq *AbilityBonusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(abilitybonus.Table, abilitybonus.Columns, sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt))
	_spec.From = abq.sql
	if unique := abq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if abq.path != nil {
		_spec.Unique = true
	}
	if fields := abq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, abilitybonus.FieldID)
		for i := range fields {
			if fields[i] != abilitybonus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := abq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := abq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := abq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := abq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (abq *AbilityBonusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(abq.driver.Dialect())
	t1 := builder.Table(abilitybonus.Table)
	columns := abq.ctx.Fields
	if len(columns) == 0 {
		columns = abilitybonus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if abq.sql != nil {
		selector = abq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if abq.ctx.Unique != nil && *abq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range abq.predicates {
		p(selector)
	}
	for _, p := range abq.order {
		p(selector)
	}
	if offset := abq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := abq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedRace tells the query-builder to eager-load the nodes that are connected to the "race"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithNamedRace(name string, opts ...func(*RaceQuery)) *AbilityBonusQuery {
	query := (&RaceClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if abq.withNamedRace == nil {
		abq.withNamedRace = make(map[string]*RaceQuery)
	}
	abq.withNamedRace[name] = query
	return abq
}

// WithNamedOptions tells the query-builder to eager-load the nodes that are connected to the "options"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithNamedOptions(name string, opts ...func(*AbilityBonusChoiceQuery)) *AbilityBonusQuery {
	query := (&AbilityBonusChoiceClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if abq.withNamedOptions == nil {
		abq.withNamedOptions = make(map[string]*AbilityBonusChoiceQuery)
	}
	abq.withNamedOptions[name] = query
	return abq
}

// WithNamedSubrace tells the query-builder to eager-load the nodes that are connected to the "subrace"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (abq *AbilityBonusQuery) WithNamedSubrace(name string, opts ...func(*SubraceQuery)) *AbilityBonusQuery {
	query := (&SubraceClient{config: abq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if abq.withNamedSubrace == nil {
		abq.withNamedSubrace = make(map[string]*SubraceQuery)
	}
	abq.withNamedSubrace[name] = query
	return abq
}

// AbilityBonusGroupBy is the group-by builder for AbilityBonus entities.
type AbilityBonusGroupBy struct {
	selector
	build *AbilityBonusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (abgb *AbilityBonusGroupBy) Aggregate(fns ...AggregateFunc) *AbilityBonusGroupBy {
	abgb.fns = append(abgb.fns, fns...)
	return abgb
}

// Scan applies the selector query and scans the result into the given value.
func (abgb *AbilityBonusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, abgb.build.ctx, ent.OpQueryGroupBy)
	if err := abgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AbilityBonusQuery, *AbilityBonusGroupBy](ctx, abgb.build, abgb, abgb.build.inters, v)
}

func (abgb *AbilityBonusGroupBy) sqlScan(ctx context.Context, root *AbilityBonusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(abgb.fns))
	for _, fn := range abgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*abgb.flds)+len(abgb.fns))
		for _, f := range *abgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*abgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := abgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AbilityBonusSelect is the builder for selecting fields of AbilityBonus entities.
type AbilityBonusSelect struct {
	*AbilityBonusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (abs *AbilityBonusSelect) Aggregate(fns ...AggregateFunc) *AbilityBonusSelect {
	abs.fns = append(abs.fns, fns...)
	return abs
}

// Scan applies the selector query and scans the result into the given value.
func (abs *AbilityBonusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, abs.ctx, ent.OpQuerySelect)
	if err := abs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AbilityBonusQuery, *AbilityBonusSelect](ctx, abs.AbilityBonusQuery, abs, abs.inters, v)
}

func (abs *AbilityBonusSelect) sqlScan(ctx context.Context, root *AbilityBonusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(abs.fns))
	for _, fn := range abs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*abs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := abs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
