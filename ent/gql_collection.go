// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/skill"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (as *AbilityScoreQuery) CollectFields(ctx context.Context, satisfies ...string) (*AbilityScoreQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return as, nil
	}
	if err := as.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return as, nil
}

func (as *AbilityScoreQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(abilityscore.Columns))
		selectedFields = []string{abilityscore.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "skills":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SkillClient{config: as.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			as.WithNamedSkills(alias, func(wq *SkillQuery) {
				*wq = *query
			})
		case "indx":
			if _, ok := fieldSeen[abilityscore.FieldIndx]; !ok {
				selectedFields = append(selectedFields, abilityscore.FieldIndx)
				fieldSeen[abilityscore.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[abilityscore.FieldName]; !ok {
				selectedFields = append(selectedFields, abilityscore.FieldName)
				fieldSeen[abilityscore.FieldName] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[abilityscore.FieldDesc]; !ok {
				selectedFields = append(selectedFields, abilityscore.FieldDesc)
				fieldSeen[abilityscore.FieldDesc] = struct{}{}
			}
		case "fullName":
			if _, ok := fieldSeen[abilityscore.FieldFullName]; !ok {
				selectedFields = append(selectedFields, abilityscore.FieldFullName)
				fieldSeen[abilityscore.FieldFullName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		as.Select(selectedFields...)
	}
	return nil
}

type abilityscorePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AbilityScorePaginateOption
}

func newAbilityScorePaginateArgs(rv map[string]any) *abilityscorePaginateArgs {
	args := &abilityscorePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &AbilityScoreOrder{Field: &AbilityScoreOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAbilityScoreOrder(order))
			}
		case *AbilityScoreOrder:
			if v != nil {
				args.opts = append(args.opts, WithAbilityScoreOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*AbilityScoreWhereInput); ok {
		args.opts = append(args.opts, WithAbilityScoreFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SkillQuery) CollectFields(ctx context.Context, satisfies ...string) (*SkillQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SkillQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(skill.Columns))
		selectedFields = []string{skill.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "abilityScore":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AbilityScoreClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withAbilityScore = query
		case "indx":
			if _, ok := fieldSeen[skill.FieldIndx]; !ok {
				selectedFields = append(selectedFields, skill.FieldIndx)
				fieldSeen[skill.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[skill.FieldName]; !ok {
				selectedFields = append(selectedFields, skill.FieldName)
				fieldSeen[skill.FieldName] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[skill.FieldDesc]; !ok {
				selectedFields = append(selectedFields, skill.FieldDesc)
				fieldSeen[skill.FieldDesc] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type skillPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SkillPaginateOption
}

func newSkillPaginateArgs(rv map[string]any) *skillPaginateArgs {
	args := &skillPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &SkillOrder{Field: &SkillOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithSkillOrder(order))
			}
		case *SkillOrder:
			if v != nil {
				args.opts = append(args.opts, WithSkillOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*SkillWhereInput); ok {
		args.opts = append(args.opts, WithSkillFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
