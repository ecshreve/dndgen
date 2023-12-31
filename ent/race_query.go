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
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/trait"
)

// RaceQuery is the builder for querying Race entities.
type RaceQuery struct {
	config
	ctx                        *QueryContext
	order                      []race.OrderOption
	inters                     []Interceptor
	predicates                 []predicate.Race
	withProficiencies          *ProficiencyQuery
	withProficiencyChoice      *ProficiencyChoiceQuery
	withLanguages              *LanguageQuery
	withSubrace                *SubraceQuery
	withTraits                 *TraitQuery
	withAbilityBonuses         *AbilityBonusQuery
	modifiers                  []func(*sql.Selector)
	loadTotal                  []func(context.Context, []*Race) error
	withNamedProficiencies     map[string]*ProficiencyQuery
	withNamedProficiencyChoice map[string]*ProficiencyChoiceQuery
	withNamedLanguages         map[string]*LanguageQuery
	withNamedSubrace           map[string]*SubraceQuery
	withNamedTraits            map[string]*TraitQuery
	withNamedAbilityBonuses    map[string]*AbilityBonusQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RaceQuery builder.
func (rq *RaceQuery) Where(ps ...predicate.Race) *RaceQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RaceQuery) Limit(limit int) *RaceQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RaceQuery) Offset(offset int) *RaceQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RaceQuery) Unique(unique bool) *RaceQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RaceQuery) Order(o ...race.OrderOption) *RaceQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryProficiencies chains the current query on the "proficiencies" edge.
func (rq *RaceQuery) QueryProficiencies() *ProficiencyQuery {
	query := (&ProficiencyClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(race.Table, race.FieldID, selector),
			sqlgraph.To(proficiency.Table, proficiency.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, race.ProficienciesTable, race.ProficienciesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProficiencyChoice chains the current query on the "proficiency_choice" edge.
func (rq *RaceQuery) QueryProficiencyChoice() *ProficiencyChoiceQuery {
	query := (&ProficiencyChoiceClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(race.Table, race.FieldID, selector),
			sqlgraph.To(proficiencychoice.Table, proficiencychoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, race.ProficiencyChoiceTable, race.ProficiencyChoicePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLanguages chains the current query on the "languages" edge.
func (rq *RaceQuery) QueryLanguages() *LanguageQuery {
	query := (&LanguageClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(race.Table, race.FieldID, selector),
			sqlgraph.To(language.Table, language.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, race.LanguagesTable, race.LanguagesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubrace chains the current query on the "subrace" edge.
func (rq *RaceQuery) QuerySubrace() *SubraceQuery {
	query := (&SubraceClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(race.Table, race.FieldID, selector),
			sqlgraph.To(subrace.Table, subrace.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, race.SubraceTable, race.SubraceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTraits chains the current query on the "traits" edge.
func (rq *RaceQuery) QueryTraits() *TraitQuery {
	query := (&TraitClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(race.Table, race.FieldID, selector),
			sqlgraph.To(trait.Table, trait.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, race.TraitsTable, race.TraitsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAbilityBonuses chains the current query on the "ability_bonuses" edge.
func (rq *RaceQuery) QueryAbilityBonuses() *AbilityBonusQuery {
	query := (&AbilityBonusClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(race.Table, race.FieldID, selector),
			sqlgraph.To(abilitybonus.Table, abilitybonus.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, race.AbilityBonusesTable, race.AbilityBonusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Race entity from the query.
// Returns a *NotFoundError when no Race was found.
func (rq *RaceQuery) First(ctx context.Context) (*Race, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{race.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RaceQuery) FirstX(ctx context.Context) *Race {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Race ID from the query.
// Returns a *NotFoundError when no Race ID was found.
func (rq *RaceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{race.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RaceQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Race entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Race entity is found.
// Returns a *NotFoundError when no Race entities are found.
func (rq *RaceQuery) Only(ctx context.Context) (*Race, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{race.Label}
	default:
		return nil, &NotSingularError{race.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RaceQuery) OnlyX(ctx context.Context) *Race {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Race ID in the query.
// Returns a *NotSingularError when more than one Race ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RaceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{race.Label}
	default:
		err = &NotSingularError{race.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RaceQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Races.
func (rq *RaceQuery) All(ctx context.Context) ([]*Race, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Race, *RaceQuery]()
	return withInterceptors[[]*Race](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RaceQuery) AllX(ctx context.Context) []*Race {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Race IDs.
func (rq *RaceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(race.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RaceQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RaceQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RaceQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RaceQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RaceQuery) Clone() *RaceQuery {
	if rq == nil {
		return nil
	}
	return &RaceQuery{
		config:                rq.config,
		ctx:                   rq.ctx.Clone(),
		order:                 append([]race.OrderOption{}, rq.order...),
		inters:                append([]Interceptor{}, rq.inters...),
		predicates:            append([]predicate.Race{}, rq.predicates...),
		withProficiencies:     rq.withProficiencies.Clone(),
		withProficiencyChoice: rq.withProficiencyChoice.Clone(),
		withLanguages:         rq.withLanguages.Clone(),
		withSubrace:           rq.withSubrace.Clone(),
		withTraits:            rq.withTraits.Clone(),
		withAbilityBonuses:    rq.withAbilityBonuses.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithProficiencies tells the query-builder to eager-load the nodes that are connected to
// the "proficiencies" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithProficiencies(opts ...func(*ProficiencyQuery)) *RaceQuery {
	query := (&ProficiencyClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withProficiencies = query
	return rq
}

// WithProficiencyChoice tells the query-builder to eager-load the nodes that are connected to
// the "proficiency_choice" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithProficiencyChoice(opts ...func(*ProficiencyChoiceQuery)) *RaceQuery {
	query := (&ProficiencyChoiceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withProficiencyChoice = query
	return rq
}

// WithLanguages tells the query-builder to eager-load the nodes that are connected to
// the "languages" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithLanguages(opts ...func(*LanguageQuery)) *RaceQuery {
	query := (&LanguageClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withLanguages = query
	return rq
}

// WithSubrace tells the query-builder to eager-load the nodes that are connected to
// the "subrace" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithSubrace(opts ...func(*SubraceQuery)) *RaceQuery {
	query := (&SubraceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withSubrace = query
	return rq
}

// WithTraits tells the query-builder to eager-load the nodes that are connected to
// the "traits" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithTraits(opts ...func(*TraitQuery)) *RaceQuery {
	query := (&TraitClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withTraits = query
	return rq
}

// WithAbilityBonuses tells the query-builder to eager-load the nodes that are connected to
// the "ability_bonuses" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithAbilityBonuses(opts ...func(*AbilityBonusQuery)) *RaceQuery {
	query := (&AbilityBonusClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withAbilityBonuses = query
	return rq
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
//	client.Race.Query().
//		GroupBy(race.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RaceQuery) GroupBy(field string, fields ...string) *RaceGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RaceGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = race.Label
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
//	client.Race.Query().
//		Select(race.FieldIndx).
//		Scan(ctx, &v)
func (rq *RaceQuery) Select(fields ...string) *RaceSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RaceSelect{RaceQuery: rq}
	sbuild.label = race.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RaceSelect configured with the given aggregations.
func (rq *RaceQuery) Aggregate(fns ...AggregateFunc) *RaceSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !race.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Race, error) {
	var (
		nodes       = []*Race{}
		_spec       = rq.querySpec()
		loadedTypes = [6]bool{
			rq.withProficiencies != nil,
			rq.withProficiencyChoice != nil,
			rq.withLanguages != nil,
			rq.withSubrace != nil,
			rq.withTraits != nil,
			rq.withAbilityBonuses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Race).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Race{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withProficiencies; query != nil {
		if err := rq.loadProficiencies(ctx, query, nodes,
			func(n *Race) { n.Edges.Proficiencies = []*Proficiency{} },
			func(n *Race, e *Proficiency) { n.Edges.Proficiencies = append(n.Edges.Proficiencies, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withProficiencyChoice; query != nil {
		if err := rq.loadProficiencyChoice(ctx, query, nodes,
			func(n *Race) { n.Edges.ProficiencyChoice = []*ProficiencyChoice{} },
			func(n *Race, e *ProficiencyChoice) { n.Edges.ProficiencyChoice = append(n.Edges.ProficiencyChoice, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withLanguages; query != nil {
		if err := rq.loadLanguages(ctx, query, nodes,
			func(n *Race) { n.Edges.Languages = []*Language{} },
			func(n *Race, e *Language) { n.Edges.Languages = append(n.Edges.Languages, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withSubrace; query != nil {
		if err := rq.loadSubrace(ctx, query, nodes,
			func(n *Race) { n.Edges.Subrace = []*Subrace{} },
			func(n *Race, e *Subrace) { n.Edges.Subrace = append(n.Edges.Subrace, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withTraits; query != nil {
		if err := rq.loadTraits(ctx, query, nodes,
			func(n *Race) { n.Edges.Traits = []*Trait{} },
			func(n *Race, e *Trait) { n.Edges.Traits = append(n.Edges.Traits, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withAbilityBonuses; query != nil {
		if err := rq.loadAbilityBonuses(ctx, query, nodes,
			func(n *Race) { n.Edges.AbilityBonuses = []*AbilityBonus{} },
			func(n *Race, e *AbilityBonus) { n.Edges.AbilityBonuses = append(n.Edges.AbilityBonuses, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedProficiencies {
		if err := rq.loadProficiencies(ctx, query, nodes,
			func(n *Race) { n.appendNamedProficiencies(name) },
			func(n *Race, e *Proficiency) { n.appendNamedProficiencies(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedProficiencyChoice {
		if err := rq.loadProficiencyChoice(ctx, query, nodes,
			func(n *Race) { n.appendNamedProficiencyChoice(name) },
			func(n *Race, e *ProficiencyChoice) { n.appendNamedProficiencyChoice(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedLanguages {
		if err := rq.loadLanguages(ctx, query, nodes,
			func(n *Race) { n.appendNamedLanguages(name) },
			func(n *Race, e *Language) { n.appendNamedLanguages(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedSubrace {
		if err := rq.loadSubrace(ctx, query, nodes,
			func(n *Race) { n.appendNamedSubrace(name) },
			func(n *Race, e *Subrace) { n.appendNamedSubrace(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedTraits {
		if err := rq.loadTraits(ctx, query, nodes,
			func(n *Race) { n.appendNamedTraits(name) },
			func(n *Race, e *Trait) { n.appendNamedTraits(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedAbilityBonuses {
		if err := rq.loadAbilityBonuses(ctx, query, nodes,
			func(n *Race) { n.appendNamedAbilityBonuses(name) },
			func(n *Race, e *AbilityBonus) { n.appendNamedAbilityBonuses(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range rq.loadTotal {
		if err := rq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RaceQuery) loadProficiencies(ctx context.Context, query *ProficiencyQuery, nodes []*Race, init func(*Race), assign func(*Race, *Proficiency)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Race)
	nids := make(map[int]map[*Race]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(race.ProficienciesTable)
		s.Join(joinT).On(s.C(proficiency.FieldID), joinT.C(race.ProficienciesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(race.ProficienciesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(race.ProficienciesPrimaryKey[0]))
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
					nids[inValue] = map[*Race]struct{}{byID[outValue]: {}}
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
func (rq *RaceQuery) loadProficiencyChoice(ctx context.Context, query *ProficiencyChoiceQuery, nodes []*Race, init func(*Race), assign func(*Race, *ProficiencyChoice)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Race)
	nids := make(map[int]map[*Race]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(race.ProficiencyChoiceTable)
		s.Join(joinT).On(s.C(proficiencychoice.FieldID), joinT.C(race.ProficiencyChoicePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(race.ProficiencyChoicePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(race.ProficiencyChoicePrimaryKey[0]))
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
					nids[inValue] = map[*Race]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ProficiencyChoice](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "proficiency_choice" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RaceQuery) loadLanguages(ctx context.Context, query *LanguageQuery, nodes []*Race, init func(*Race), assign func(*Race, *Language)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Race)
	nids := make(map[int]map[*Race]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(race.LanguagesTable)
		s.Join(joinT).On(s.C(language.FieldID), joinT.C(race.LanguagesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(race.LanguagesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(race.LanguagesPrimaryKey[0]))
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
					nids[inValue] = map[*Race]struct{}{byID[outValue]: {}}
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
func (rq *RaceQuery) loadSubrace(ctx context.Context, query *SubraceQuery, nodes []*Race, init func(*Race), assign func(*Race, *Subrace)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Race)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Subrace(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(race.SubraceColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.subrace_race
		if fk == nil {
			return fmt.Errorf(`foreign-key "subrace_race" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "subrace_race" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *RaceQuery) loadTraits(ctx context.Context, query *TraitQuery, nodes []*Race, init func(*Race), assign func(*Race, *Trait)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Race)
	nids := make(map[int]map[*Race]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(race.TraitsTable)
		s.Join(joinT).On(s.C(trait.FieldID), joinT.C(race.TraitsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(race.TraitsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(race.TraitsPrimaryKey[0]))
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
					nids[inValue] = map[*Race]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Trait](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "traits" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RaceQuery) loadAbilityBonuses(ctx context.Context, query *AbilityBonusQuery, nodes []*Race, init func(*Race), assign func(*Race, *AbilityBonus)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Race)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AbilityBonus(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(race.AbilityBonusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.race_ability_bonuses
		if fk == nil {
			return fmt.Errorf(`foreign-key "race_ability_bonuses" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "race_ability_bonuses" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(race.Table, race.Columns, sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, race.FieldID)
		for i := range fields {
			if fields[i] != race.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(race.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = race.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProficiencies tells the query-builder to eager-load the nodes that are connected to the "proficiencies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithNamedProficiencies(name string, opts ...func(*ProficiencyQuery)) *RaceQuery {
	query := (&ProficiencyClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedProficiencies == nil {
		rq.withNamedProficiencies = make(map[string]*ProficiencyQuery)
	}
	rq.withNamedProficiencies[name] = query
	return rq
}

// WithNamedProficiencyChoice tells the query-builder to eager-load the nodes that are connected to the "proficiency_choice"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithNamedProficiencyChoice(name string, opts ...func(*ProficiencyChoiceQuery)) *RaceQuery {
	query := (&ProficiencyChoiceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedProficiencyChoice == nil {
		rq.withNamedProficiencyChoice = make(map[string]*ProficiencyChoiceQuery)
	}
	rq.withNamedProficiencyChoice[name] = query
	return rq
}

// WithNamedLanguages tells the query-builder to eager-load the nodes that are connected to the "languages"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithNamedLanguages(name string, opts ...func(*LanguageQuery)) *RaceQuery {
	query := (&LanguageClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedLanguages == nil {
		rq.withNamedLanguages = make(map[string]*LanguageQuery)
	}
	rq.withNamedLanguages[name] = query
	return rq
}

// WithNamedSubrace tells the query-builder to eager-load the nodes that are connected to the "subrace"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithNamedSubrace(name string, opts ...func(*SubraceQuery)) *RaceQuery {
	query := (&SubraceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedSubrace == nil {
		rq.withNamedSubrace = make(map[string]*SubraceQuery)
	}
	rq.withNamedSubrace[name] = query
	return rq
}

// WithNamedTraits tells the query-builder to eager-load the nodes that are connected to the "traits"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithNamedTraits(name string, opts ...func(*TraitQuery)) *RaceQuery {
	query := (&TraitClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedTraits == nil {
		rq.withNamedTraits = make(map[string]*TraitQuery)
	}
	rq.withNamedTraits[name] = query
	return rq
}

// WithNamedAbilityBonuses tells the query-builder to eager-load the nodes that are connected to the "ability_bonuses"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RaceQuery) WithNamedAbilityBonuses(name string, opts ...func(*AbilityBonusQuery)) *RaceQuery {
	query := (&AbilityBonusClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedAbilityBonuses == nil {
		rq.withNamedAbilityBonuses = make(map[string]*AbilityBonusQuery)
	}
	rq.withNamedAbilityBonuses[name] = query
	return rq
}

// RaceGroupBy is the group-by builder for Race entities.
type RaceGroupBy struct {
	selector
	build *RaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RaceGroupBy) Aggregate(fns ...AggregateFunc) *RaceGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RaceQuery, *RaceGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RaceGroupBy) sqlScan(ctx context.Context, root *RaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RaceSelect is the builder for selecting fields of Race entities.
type RaceSelect struct {
	*RaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RaceSelect) Aggregate(fns ...AggregateFunc) *RaceSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RaceQuery, *RaceSelect](ctx, rs.RaceQuery, rs, rs.inters, v)
}

func (rs *RaceSelect) sqlScan(ctx context.Context, root *RaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
