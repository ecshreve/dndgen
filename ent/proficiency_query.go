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
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
)

// ProficiencyQuery is the builder for querying Proficiency entities.
type ProficiencyQuery struct {
	config
	ctx              *QueryContext
	order            []proficiency.OrderOption
	inters           []Interceptor
	predicates       []predicate.Proficiency
	withRace         *RaceQuery
	withOptions      *ProficiencyChoiceQuery
	withSubrace      *SubraceQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Proficiency) error
	withNamedRace    map[string]*RaceQuery
	withNamedOptions map[string]*ProficiencyChoiceQuery
	withNamedSubrace map[string]*SubraceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProficiencyQuery builder.
func (pq *ProficiencyQuery) Where(ps ...predicate.Proficiency) *ProficiencyQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProficiencyQuery) Limit(limit int) *ProficiencyQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProficiencyQuery) Offset(offset int) *ProficiencyQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProficiencyQuery) Unique(unique bool) *ProficiencyQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProficiencyQuery) Order(o ...proficiency.OrderOption) *ProficiencyQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryRace chains the current query on the "race" edge.
func (pq *ProficiencyQuery) QueryRace() *RaceQuery {
	query := (&RaceClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(proficiency.Table, proficiency.FieldID, selector),
			sqlgraph.To(race.Table, race.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, proficiency.RaceTable, proficiency.RacePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOptions chains the current query on the "options" edge.
func (pq *ProficiencyQuery) QueryOptions() *ProficiencyChoiceQuery {
	query := (&ProficiencyChoiceClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(proficiency.Table, proficiency.FieldID, selector),
			sqlgraph.To(proficiencychoice.Table, proficiencychoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, proficiency.OptionsTable, proficiency.OptionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubrace chains the current query on the "subrace" edge.
func (pq *ProficiencyQuery) QuerySubrace() *SubraceQuery {
	query := (&SubraceClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(proficiency.Table, proficiency.FieldID, selector),
			sqlgraph.To(subrace.Table, subrace.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, proficiency.SubraceTable, proficiency.SubracePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Proficiency entity from the query.
// Returns a *NotFoundError when no Proficiency was found.
func (pq *ProficiencyQuery) First(ctx context.Context) (*Proficiency, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{proficiency.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProficiencyQuery) FirstX(ctx context.Context) *Proficiency {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Proficiency ID from the query.
// Returns a *NotFoundError when no Proficiency ID was found.
func (pq *ProficiencyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{proficiency.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProficiencyQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Proficiency entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Proficiency entity is found.
// Returns a *NotFoundError when no Proficiency entities are found.
func (pq *ProficiencyQuery) Only(ctx context.Context) (*Proficiency, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{proficiency.Label}
	default:
		return nil, &NotSingularError{proficiency.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProficiencyQuery) OnlyX(ctx context.Context) *Proficiency {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Proficiency ID in the query.
// Returns a *NotSingularError when more than one Proficiency ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProficiencyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{proficiency.Label}
	default:
		err = &NotSingularError{proficiency.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProficiencyQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Proficiencies.
func (pq *ProficiencyQuery) All(ctx context.Context) ([]*Proficiency, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Proficiency, *ProficiencyQuery]()
	return withInterceptors[[]*Proficiency](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProficiencyQuery) AllX(ctx context.Context) []*Proficiency {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Proficiency IDs.
func (pq *ProficiencyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(proficiency.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProficiencyQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProficiencyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProficiencyQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProficiencyQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProficiencyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProficiencyQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProficiencyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProficiencyQuery) Clone() *ProficiencyQuery {
	if pq == nil {
		return nil
	}
	return &ProficiencyQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]proficiency.OrderOption{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Proficiency{}, pq.predicates...),
		withRace:    pq.withRace.Clone(),
		withOptions: pq.withOptions.Clone(),
		withSubrace: pq.withSubrace.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithRace tells the query-builder to eager-load the nodes that are connected to
// the "race" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProficiencyQuery) WithRace(opts ...func(*RaceQuery)) *ProficiencyQuery {
	query := (&RaceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withRace = query
	return pq
}

// WithOptions tells the query-builder to eager-load the nodes that are connected to
// the "options" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProficiencyQuery) WithOptions(opts ...func(*ProficiencyChoiceQuery)) *ProficiencyQuery {
	query := (&ProficiencyChoiceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOptions = query
	return pq
}

// WithSubrace tells the query-builder to eager-load the nodes that are connected to
// the "subrace" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProficiencyQuery) WithSubrace(opts ...func(*SubraceQuery)) *ProficiencyQuery {
	query := (&SubraceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withSubrace = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Indx string `json:"indx,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Proficiency.Query().
//		GroupBy(proficiency.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProficiencyQuery) GroupBy(field string, fields ...string) *ProficiencyGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProficiencyGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = proficiency.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Indx string `json:"indx,omitempty"`
//	}
//
//	client.Proficiency.Query().
//		Select(proficiency.FieldIndx).
//		Scan(ctx, &v)
func (pq *ProficiencyQuery) Select(fields ...string) *ProficiencySelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProficiencySelect{ProficiencyQuery: pq}
	sbuild.label = proficiency.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProficiencySelect configured with the given aggregations.
func (pq *ProficiencyQuery) Aggregate(fns ...AggregateFunc) *ProficiencySelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProficiencyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !proficiency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProficiencyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Proficiency, error) {
	var (
		nodes       = []*Proficiency{}
		_spec       = pq.querySpec()
		loadedTypes = [3]bool{
			pq.withRace != nil,
			pq.withOptions != nil,
			pq.withSubrace != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Proficiency).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Proficiency{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withRace; query != nil {
		if err := pq.loadRace(ctx, query, nodes,
			func(n *Proficiency) { n.Edges.Race = []*Race{} },
			func(n *Proficiency, e *Race) { n.Edges.Race = append(n.Edges.Race, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withOptions; query != nil {
		if err := pq.loadOptions(ctx, query, nodes,
			func(n *Proficiency) { n.Edges.Options = []*ProficiencyChoice{} },
			func(n *Proficiency, e *ProficiencyChoice) { n.Edges.Options = append(n.Edges.Options, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withSubrace; query != nil {
		if err := pq.loadSubrace(ctx, query, nodes,
			func(n *Proficiency) { n.Edges.Subrace = []*Subrace{} },
			func(n *Proficiency, e *Subrace) { n.Edges.Subrace = append(n.Edges.Subrace, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedRace {
		if err := pq.loadRace(ctx, query, nodes,
			func(n *Proficiency) { n.appendNamedRace(name) },
			func(n *Proficiency, e *Race) { n.appendNamedRace(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedOptions {
		if err := pq.loadOptions(ctx, query, nodes,
			func(n *Proficiency) { n.appendNamedOptions(name) },
			func(n *Proficiency, e *ProficiencyChoice) { n.appendNamedOptions(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedSubrace {
		if err := pq.loadSubrace(ctx, query, nodes,
			func(n *Proficiency) { n.appendNamedSubrace(name) },
			func(n *Proficiency, e *Subrace) { n.appendNamedSubrace(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProficiencyQuery) loadRace(ctx context.Context, query *RaceQuery, nodes []*Proficiency, init func(*Proficiency), assign func(*Proficiency, *Race)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Proficiency)
	nids := make(map[int]map[*Proficiency]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(proficiency.RaceTable)
		s.Join(joinT).On(s.C(race.FieldID), joinT.C(proficiency.RacePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(proficiency.RacePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(proficiency.RacePrimaryKey[1]))
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
					nids[inValue] = map[*Proficiency]struct{}{byID[outValue]: {}}
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
func (pq *ProficiencyQuery) loadOptions(ctx context.Context, query *ProficiencyChoiceQuery, nodes []*Proficiency, init func(*Proficiency), assign func(*Proficiency, *ProficiencyChoice)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Proficiency)
	nids := make(map[int]map[*Proficiency]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(proficiency.OptionsTable)
		s.Join(joinT).On(s.C(proficiencychoice.FieldID), joinT.C(proficiency.OptionsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(proficiency.OptionsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(proficiency.OptionsPrimaryKey[1]))
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
					nids[inValue] = map[*Proficiency]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "options" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *ProficiencyQuery) loadSubrace(ctx context.Context, query *SubraceQuery, nodes []*Proficiency, init func(*Proficiency), assign func(*Proficiency, *Subrace)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Proficiency)
	nids := make(map[int]map[*Proficiency]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(proficiency.SubraceTable)
		s.Join(joinT).On(s.C(subrace.FieldID), joinT.C(proficiency.SubracePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(proficiency.SubracePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(proficiency.SubracePrimaryKey[1]))
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
					nids[inValue] = map[*Proficiency]struct{}{byID[outValue]: {}}
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

func (pq *ProficiencyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProficiencyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(proficiency.Table, proficiency.Columns, sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, proficiency.FieldID)
		for i := range fields {
			if fields[i] != proficiency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProficiencyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(proficiency.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = proficiency.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedRace tells the query-builder to eager-load the nodes that are connected to the "race"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *ProficiencyQuery) WithNamedRace(name string, opts ...func(*RaceQuery)) *ProficiencyQuery {
	query := (&RaceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedRace == nil {
		pq.withNamedRace = make(map[string]*RaceQuery)
	}
	pq.withNamedRace[name] = query
	return pq
}

// WithNamedOptions tells the query-builder to eager-load the nodes that are connected to the "options"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *ProficiencyQuery) WithNamedOptions(name string, opts ...func(*ProficiencyChoiceQuery)) *ProficiencyQuery {
	query := (&ProficiencyChoiceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedOptions == nil {
		pq.withNamedOptions = make(map[string]*ProficiencyChoiceQuery)
	}
	pq.withNamedOptions[name] = query
	return pq
}

// WithNamedSubrace tells the query-builder to eager-load the nodes that are connected to the "subrace"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *ProficiencyQuery) WithNamedSubrace(name string, opts ...func(*SubraceQuery)) *ProficiencyQuery {
	query := (&SubraceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedSubrace == nil {
		pq.withNamedSubrace = make(map[string]*SubraceQuery)
	}
	pq.withNamedSubrace[name] = query
	return pq
}

// ProficiencyGroupBy is the group-by builder for Proficiency entities.
type ProficiencyGroupBy struct {
	selector
	build *ProficiencyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProficiencyGroupBy) Aggregate(fns ...AggregateFunc) *ProficiencyGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProficiencyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProficiencyQuery, *ProficiencyGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProficiencyGroupBy) sqlScan(ctx context.Context, root *ProficiencyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProficiencySelect is the builder for selecting fields of Proficiency entities.
type ProficiencySelect struct {
	*ProficiencyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProficiencySelect) Aggregate(fns ...AggregateFunc) *ProficiencySelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProficiencySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProficiencyQuery, *ProficiencySelect](ctx, ps.ProficiencyQuery, ps, ps.inters, v)
}

func (ps *ProficiencySelect) sqlScan(ctx context.Context, root *ProficiencyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
