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
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
)

// AbilityBonusChoiceQuery is the builder for querying AbilityBonusChoice entities.
type AbilityBonusChoiceQuery struct {
	config
	ctx                     *QueryContext
	order                   []abilitybonuschoice.OrderOption
	inters                  []Interceptor
	predicates              []predicate.AbilityBonusChoice
	withAbilityBonuses      *AbilityBonusQuery
	withRace                *RaceQuery
	modifiers               []func(*sql.Selector)
	loadTotal               []func(context.Context, []*AbilityBonusChoice) error
	withNamedAbilityBonuses map[string]*AbilityBonusQuery
	withNamedRace           map[string]*RaceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AbilityBonusChoiceQuery builder.
func (abcq *AbilityBonusChoiceQuery) Where(ps ...predicate.AbilityBonusChoice) *AbilityBonusChoiceQuery {
	abcq.predicates = append(abcq.predicates, ps...)
	return abcq
}

// Limit the number of records to be returned by this query.
func (abcq *AbilityBonusChoiceQuery) Limit(limit int) *AbilityBonusChoiceQuery {
	abcq.ctx.Limit = &limit
	return abcq
}

// Offset to start from.
func (abcq *AbilityBonusChoiceQuery) Offset(offset int) *AbilityBonusChoiceQuery {
	abcq.ctx.Offset = &offset
	return abcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (abcq *AbilityBonusChoiceQuery) Unique(unique bool) *AbilityBonusChoiceQuery {
	abcq.ctx.Unique = &unique
	return abcq
}

// Order specifies how the records should be ordered.
func (abcq *AbilityBonusChoiceQuery) Order(o ...abilitybonuschoice.OrderOption) *AbilityBonusChoiceQuery {
	abcq.order = append(abcq.order, o...)
	return abcq
}

// QueryAbilityBonuses chains the current query on the "ability_bonuses" edge.
func (abcq *AbilityBonusChoiceQuery) QueryAbilityBonuses() *AbilityBonusQuery {
	query := (&AbilityBonusClient{config: abcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilitybonuschoice.Table, abilitybonuschoice.FieldID, selector),
			sqlgraph.To(abilitybonus.Table, abilitybonus.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, abilitybonuschoice.AbilityBonusesTable, abilitybonuschoice.AbilityBonusesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(abcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRace chains the current query on the "race" edge.
func (abcq *AbilityBonusChoiceQuery) QueryRace() *RaceQuery {
	query := (&RaceClient{config: abcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(abilitybonuschoice.Table, abilitybonuschoice.FieldID, selector),
			sqlgraph.To(race.Table, race.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, abilitybonuschoice.RaceTable, abilitybonuschoice.RaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(abcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AbilityBonusChoice entity from the query.
// Returns a *NotFoundError when no AbilityBonusChoice was found.
func (abcq *AbilityBonusChoiceQuery) First(ctx context.Context) (*AbilityBonusChoice, error) {
	nodes, err := abcq.Limit(1).All(setContextOp(ctx, abcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{abilitybonuschoice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) FirstX(ctx context.Context) *AbilityBonusChoice {
	node, err := abcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AbilityBonusChoice ID from the query.
// Returns a *NotFoundError when no AbilityBonusChoice ID was found.
func (abcq *AbilityBonusChoiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = abcq.Limit(1).IDs(setContextOp(ctx, abcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{abilitybonuschoice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) FirstIDX(ctx context.Context) int {
	id, err := abcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AbilityBonusChoice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AbilityBonusChoice entity is found.
// Returns a *NotFoundError when no AbilityBonusChoice entities are found.
func (abcq *AbilityBonusChoiceQuery) Only(ctx context.Context) (*AbilityBonusChoice, error) {
	nodes, err := abcq.Limit(2).All(setContextOp(ctx, abcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{abilitybonuschoice.Label}
	default:
		return nil, &NotSingularError{abilitybonuschoice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) OnlyX(ctx context.Context) *AbilityBonusChoice {
	node, err := abcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AbilityBonusChoice ID in the query.
// Returns a *NotSingularError when more than one AbilityBonusChoice ID is found.
// Returns a *NotFoundError when no entities are found.
func (abcq *AbilityBonusChoiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = abcq.Limit(2).IDs(setContextOp(ctx, abcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{abilitybonuschoice.Label}
	default:
		err = &NotSingularError{abilitybonuschoice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := abcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AbilityBonusChoices.
func (abcq *AbilityBonusChoiceQuery) All(ctx context.Context) ([]*AbilityBonusChoice, error) {
	ctx = setContextOp(ctx, abcq.ctx, ent.OpQueryAll)
	if err := abcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AbilityBonusChoice, *AbilityBonusChoiceQuery]()
	return withInterceptors[[]*AbilityBonusChoice](ctx, abcq, qr, abcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) AllX(ctx context.Context) []*AbilityBonusChoice {
	nodes, err := abcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AbilityBonusChoice IDs.
func (abcq *AbilityBonusChoiceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if abcq.ctx.Unique == nil && abcq.path != nil {
		abcq.Unique(true)
	}
	ctx = setContextOp(ctx, abcq.ctx, ent.OpQueryIDs)
	if err = abcq.Select(abilitybonuschoice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) IDsX(ctx context.Context) []int {
	ids, err := abcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (abcq *AbilityBonusChoiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, abcq.ctx, ent.OpQueryCount)
	if err := abcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, abcq, querierCount[*AbilityBonusChoiceQuery](), abcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) CountX(ctx context.Context) int {
	count, err := abcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (abcq *AbilityBonusChoiceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, abcq.ctx, ent.OpQueryExist)
	switch _, err := abcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (abcq *AbilityBonusChoiceQuery) ExistX(ctx context.Context) bool {
	exist, err := abcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AbilityBonusChoiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (abcq *AbilityBonusChoiceQuery) Clone() *AbilityBonusChoiceQuery {
	if abcq == nil {
		return nil
	}
	return &AbilityBonusChoiceQuery{
		config:             abcq.config,
		ctx:                abcq.ctx.Clone(),
		order:              append([]abilitybonuschoice.OrderOption{}, abcq.order...),
		inters:             append([]Interceptor{}, abcq.inters...),
		predicates:         append([]predicate.AbilityBonusChoice{}, abcq.predicates...),
		withAbilityBonuses: abcq.withAbilityBonuses.Clone(),
		withRace:           abcq.withRace.Clone(),
		// clone intermediate query.
		sql:  abcq.sql.Clone(),
		path: abcq.path,
	}
}

// WithAbilityBonuses tells the query-builder to eager-load the nodes that are connected to
// the "ability_bonuses" edge. The optional arguments are used to configure the query builder of the edge.
func (abcq *AbilityBonusChoiceQuery) WithAbilityBonuses(opts ...func(*AbilityBonusQuery)) *AbilityBonusChoiceQuery {
	query := (&AbilityBonusClient{config: abcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	abcq.withAbilityBonuses = query
	return abcq
}

// WithRace tells the query-builder to eager-load the nodes that are connected to
// the "race" edge. The optional arguments are used to configure the query builder of the edge.
func (abcq *AbilityBonusChoiceQuery) WithRace(opts ...func(*RaceQuery)) *AbilityBonusChoiceQuery {
	query := (&RaceClient{config: abcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	abcq.withRace = query
	return abcq
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
//	client.AbilityBonusChoice.Query().
//		GroupBy(abilitybonuschoice.FieldChoose).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (abcq *AbilityBonusChoiceQuery) GroupBy(field string, fields ...string) *AbilityBonusChoiceGroupBy {
	abcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AbilityBonusChoiceGroupBy{build: abcq}
	grbuild.flds = &abcq.ctx.Fields
	grbuild.label = abilitybonuschoice.Label
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
//	client.AbilityBonusChoice.Query().
//		Select(abilitybonuschoice.FieldChoose).
//		Scan(ctx, &v)
func (abcq *AbilityBonusChoiceQuery) Select(fields ...string) *AbilityBonusChoiceSelect {
	abcq.ctx.Fields = append(abcq.ctx.Fields, fields...)
	sbuild := &AbilityBonusChoiceSelect{AbilityBonusChoiceQuery: abcq}
	sbuild.label = abilitybonuschoice.Label
	sbuild.flds, sbuild.scan = &abcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AbilityBonusChoiceSelect configured with the given aggregations.
func (abcq *AbilityBonusChoiceQuery) Aggregate(fns ...AggregateFunc) *AbilityBonusChoiceSelect {
	return abcq.Select().Aggregate(fns...)
}

func (abcq *AbilityBonusChoiceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range abcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, abcq); err != nil {
				return err
			}
		}
	}
	for _, f := range abcq.ctx.Fields {
		if !abilitybonuschoice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if abcq.path != nil {
		prev, err := abcq.path(ctx)
		if err != nil {
			return err
		}
		abcq.sql = prev
	}
	return nil
}

func (abcq *AbilityBonusChoiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AbilityBonusChoice, error) {
	var (
		nodes       = []*AbilityBonusChoice{}
		_spec       = abcq.querySpec()
		loadedTypes = [2]bool{
			abcq.withAbilityBonuses != nil,
			abcq.withRace != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AbilityBonusChoice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AbilityBonusChoice{config: abcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(abcq.modifiers) > 0 {
		_spec.Modifiers = abcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, abcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := abcq.withAbilityBonuses; query != nil {
		if err := abcq.loadAbilityBonuses(ctx, query, nodes,
			func(n *AbilityBonusChoice) { n.Edges.AbilityBonuses = []*AbilityBonus{} },
			func(n *AbilityBonusChoice, e *AbilityBonus) {
				n.Edges.AbilityBonuses = append(n.Edges.AbilityBonuses, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := abcq.withRace; query != nil {
		if err := abcq.loadRace(ctx, query, nodes,
			func(n *AbilityBonusChoice) { n.Edges.Race = []*Race{} },
			func(n *AbilityBonusChoice, e *Race) { n.Edges.Race = append(n.Edges.Race, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range abcq.withNamedAbilityBonuses {
		if err := abcq.loadAbilityBonuses(ctx, query, nodes,
			func(n *AbilityBonusChoice) { n.appendNamedAbilityBonuses(name) },
			func(n *AbilityBonusChoice, e *AbilityBonus) { n.appendNamedAbilityBonuses(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range abcq.withNamedRace {
		if err := abcq.loadRace(ctx, query, nodes,
			func(n *AbilityBonusChoice) { n.appendNamedRace(name) },
			func(n *AbilityBonusChoice, e *Race) { n.appendNamedRace(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range abcq.loadTotal {
		if err := abcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (abcq *AbilityBonusChoiceQuery) loadAbilityBonuses(ctx context.Context, query *AbilityBonusQuery, nodes []*AbilityBonusChoice, init func(*AbilityBonusChoice), assign func(*AbilityBonusChoice, *AbilityBonus)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*AbilityBonusChoice)
	nids := make(map[int]map[*AbilityBonusChoice]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(abilitybonuschoice.AbilityBonusesTable)
		s.Join(joinT).On(s.C(abilitybonus.FieldID), joinT.C(abilitybonuschoice.AbilityBonusesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(abilitybonuschoice.AbilityBonusesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(abilitybonuschoice.AbilityBonusesPrimaryKey[0]))
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
					nids[inValue] = map[*AbilityBonusChoice]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*AbilityBonus](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "ability_bonuses" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (abcq *AbilityBonusChoiceQuery) loadRace(ctx context.Context, query *RaceQuery, nodes []*AbilityBonusChoice, init func(*AbilityBonusChoice), assign func(*AbilityBonusChoice, *Race)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*AbilityBonusChoice)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Race(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(abilitybonuschoice.RaceColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.race_ability_bonus_options
		if fk == nil {
			return fmt.Errorf(`foreign-key "race_ability_bonus_options" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "race_ability_bonus_options" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (abcq *AbilityBonusChoiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := abcq.querySpec()
	if len(abcq.modifiers) > 0 {
		_spec.Modifiers = abcq.modifiers
	}
	_spec.Node.Columns = abcq.ctx.Fields
	if len(abcq.ctx.Fields) > 0 {
		_spec.Unique = abcq.ctx.Unique != nil && *abcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, abcq.driver, _spec)
}

func (abcq *AbilityBonusChoiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(abilitybonuschoice.Table, abilitybonuschoice.Columns, sqlgraph.NewFieldSpec(abilitybonuschoice.FieldID, field.TypeInt))
	_spec.From = abcq.sql
	if unique := abcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if abcq.path != nil {
		_spec.Unique = true
	}
	if fields := abcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, abilitybonuschoice.FieldID)
		for i := range fields {
			if fields[i] != abilitybonuschoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := abcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := abcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := abcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := abcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (abcq *AbilityBonusChoiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(abcq.driver.Dialect())
	t1 := builder.Table(abilitybonuschoice.Table)
	columns := abcq.ctx.Fields
	if len(columns) == 0 {
		columns = abilitybonuschoice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if abcq.sql != nil {
		selector = abcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if abcq.ctx.Unique != nil && *abcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range abcq.predicates {
		p(selector)
	}
	for _, p := range abcq.order {
		p(selector)
	}
	if offset := abcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := abcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAbilityBonuses tells the query-builder to eager-load the nodes that are connected to the "ability_bonuses"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (abcq *AbilityBonusChoiceQuery) WithNamedAbilityBonuses(name string, opts ...func(*AbilityBonusQuery)) *AbilityBonusChoiceQuery {
	query := (&AbilityBonusClient{config: abcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if abcq.withNamedAbilityBonuses == nil {
		abcq.withNamedAbilityBonuses = make(map[string]*AbilityBonusQuery)
	}
	abcq.withNamedAbilityBonuses[name] = query
	return abcq
}

// WithNamedRace tells the query-builder to eager-load the nodes that are connected to the "race"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (abcq *AbilityBonusChoiceQuery) WithNamedRace(name string, opts ...func(*RaceQuery)) *AbilityBonusChoiceQuery {
	query := (&RaceClient{config: abcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if abcq.withNamedRace == nil {
		abcq.withNamedRace = make(map[string]*RaceQuery)
	}
	abcq.withNamedRace[name] = query
	return abcq
}

// AbilityBonusChoiceGroupBy is the group-by builder for AbilityBonusChoice entities.
type AbilityBonusChoiceGroupBy struct {
	selector
	build *AbilityBonusChoiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (abcgb *AbilityBonusChoiceGroupBy) Aggregate(fns ...AggregateFunc) *AbilityBonusChoiceGroupBy {
	abcgb.fns = append(abcgb.fns, fns...)
	return abcgb
}

// Scan applies the selector query and scans the result into the given value.
func (abcgb *AbilityBonusChoiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, abcgb.build.ctx, ent.OpQueryGroupBy)
	if err := abcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AbilityBonusChoiceQuery, *AbilityBonusChoiceGroupBy](ctx, abcgb.build, abcgb, abcgb.build.inters, v)
}

func (abcgb *AbilityBonusChoiceGroupBy) sqlScan(ctx context.Context, root *AbilityBonusChoiceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(abcgb.fns))
	for _, fn := range abcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*abcgb.flds)+len(abcgb.fns))
		for _, f := range *abcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*abcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := abcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AbilityBonusChoiceSelect is the builder for selecting fields of AbilityBonusChoice entities.
type AbilityBonusChoiceSelect struct {
	*AbilityBonusChoiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (abcs *AbilityBonusChoiceSelect) Aggregate(fns ...AggregateFunc) *AbilityBonusChoiceSelect {
	abcs.fns = append(abcs.fns, fns...)
	return abcs
}

// Scan applies the selector query and scans the result into the given value.
func (abcs *AbilityBonusChoiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, abcs.ctx, ent.OpQuerySelect)
	if err := abcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AbilityBonusChoiceQuery, *AbilityBonusChoiceSelect](ctx, abcs.AbilityBonusChoiceQuery, abcs, abcs.inters, v)
}

func (abcs *AbilityBonusChoiceSelect) sqlScan(ctx context.Context, root *AbilityBonusChoiceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(abcs.fns))
	for _, fn := range abcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*abcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := abcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
