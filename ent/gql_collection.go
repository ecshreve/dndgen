// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/magicschool"
	"github.com/ecshreve/dndgen/ent/race"
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
func (a *AlignmentQuery) CollectFields(ctx context.Context, satisfies ...string) (*AlignmentQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AlignmentQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(alignment.Columns))
		selectedFields = []string{alignment.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "indx":
			if _, ok := fieldSeen[alignment.FieldIndx]; !ok {
				selectedFields = append(selectedFields, alignment.FieldIndx)
				fieldSeen[alignment.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[alignment.FieldName]; !ok {
				selectedFields = append(selectedFields, alignment.FieldName)
				fieldSeen[alignment.FieldName] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[alignment.FieldDesc]; !ok {
				selectedFields = append(selectedFields, alignment.FieldDesc)
				fieldSeen[alignment.FieldDesc] = struct{}{}
			}
		case "abbr":
			if _, ok := fieldSeen[alignment.FieldAbbr]; !ok {
				selectedFields = append(selectedFields, alignment.FieldAbbr)
				fieldSeen[alignment.FieldAbbr] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type alignmentPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AlignmentPaginateOption
}

func newAlignmentPaginateArgs(rv map[string]any) *alignmentPaginateArgs {
	args := &alignmentPaginateArgs{}
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
				order      = &AlignmentOrder{Field: &AlignmentOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAlignmentOrder(order))
			}
		case *AlignmentOrder:
			if v != nil {
				args.opts = append(args.opts, WithAlignmentOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*AlignmentWhereInput); ok {
		args.opts = append(args.opts, WithAlignmentFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dt *DamageTypeQuery) CollectFields(ctx context.Context, satisfies ...string) (*DamageTypeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return dt, nil
	}
	if err := dt.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return dt, nil
}

func (dt *DamageTypeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(damagetype.Columns))
		selectedFields = []string{damagetype.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "indx":
			if _, ok := fieldSeen[damagetype.FieldIndx]; !ok {
				selectedFields = append(selectedFields, damagetype.FieldIndx)
				fieldSeen[damagetype.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[damagetype.FieldName]; !ok {
				selectedFields = append(selectedFields, damagetype.FieldName)
				fieldSeen[damagetype.FieldName] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[damagetype.FieldDesc]; !ok {
				selectedFields = append(selectedFields, damagetype.FieldDesc)
				fieldSeen[damagetype.FieldDesc] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		dt.Select(selectedFields...)
	}
	return nil
}

type damagetypePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DamageTypePaginateOption
}

func newDamageTypePaginateArgs(rv map[string]any) *damagetypePaginateArgs {
	args := &damagetypePaginateArgs{}
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
				order      = &DamageTypeOrder{Field: &DamageTypeOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithDamageTypeOrder(order))
			}
		case *DamageTypeOrder:
			if v != nil {
				args.opts = append(args.opts, WithDamageTypeOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*DamageTypeWhereInput); ok {
		args.opts = append(args.opts, WithDamageTypeFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *LanguageQuery) CollectFields(ctx context.Context, satisfies ...string) (*LanguageQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return l, nil
	}
	if err := l.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *LanguageQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(language.Columns))
		selectedFields = []string{language.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "indx":
			if _, ok := fieldSeen[language.FieldIndx]; !ok {
				selectedFields = append(selectedFields, language.FieldIndx)
				fieldSeen[language.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[language.FieldName]; !ok {
				selectedFields = append(selectedFields, language.FieldName)
				fieldSeen[language.FieldName] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[language.FieldDesc]; !ok {
				selectedFields = append(selectedFields, language.FieldDesc)
				fieldSeen[language.FieldDesc] = struct{}{}
			}
		case "languageType":
			if _, ok := fieldSeen[language.FieldLanguageType]; !ok {
				selectedFields = append(selectedFields, language.FieldLanguageType)
				fieldSeen[language.FieldLanguageType] = struct{}{}
			}
		case "script":
			if _, ok := fieldSeen[language.FieldScript]; !ok {
				selectedFields = append(selectedFields, language.FieldScript)
				fieldSeen[language.FieldScript] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		l.Select(selectedFields...)
	}
	return nil
}

type languagePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LanguagePaginateOption
}

func newLanguagePaginateArgs(rv map[string]any) *languagePaginateArgs {
	args := &languagePaginateArgs{}
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
				order      = &LanguageOrder{Field: &LanguageOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithLanguageOrder(order))
			}
		case *LanguageOrder:
			if v != nil {
				args.opts = append(args.opts, WithLanguageOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*LanguageWhereInput); ok {
		args.opts = append(args.opts, WithLanguageFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ms *MagicSchoolQuery) CollectFields(ctx context.Context, satisfies ...string) (*MagicSchoolQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ms, nil
	}
	if err := ms.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ms, nil
}

func (ms *MagicSchoolQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(magicschool.Columns))
		selectedFields = []string{magicschool.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "indx":
			if _, ok := fieldSeen[magicschool.FieldIndx]; !ok {
				selectedFields = append(selectedFields, magicschool.FieldIndx)
				fieldSeen[magicschool.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[magicschool.FieldName]; !ok {
				selectedFields = append(selectedFields, magicschool.FieldName)
				fieldSeen[magicschool.FieldName] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[magicschool.FieldDesc]; !ok {
				selectedFields = append(selectedFields, magicschool.FieldDesc)
				fieldSeen[magicschool.FieldDesc] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ms.Select(selectedFields...)
	}
	return nil
}

type magicschoolPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MagicSchoolPaginateOption
}

func newMagicSchoolPaginateArgs(rv map[string]any) *magicschoolPaginateArgs {
	args := &magicschoolPaginateArgs{}
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
				order      = &MagicSchoolOrder{Field: &MagicSchoolOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMagicSchoolOrder(order))
			}
		case *MagicSchoolOrder:
			if v != nil {
				args.opts = append(args.opts, WithMagicSchoolOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MagicSchoolWhereInput); ok {
		args.opts = append(args.opts, WithMagicSchoolFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RaceQuery) CollectFields(ctx context.Context, satisfies ...string) (*RaceQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RaceQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(race.Columns))
		selectedFields = []string{race.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "indx":
			if _, ok := fieldSeen[race.FieldIndx]; !ok {
				selectedFields = append(selectedFields, race.FieldIndx)
				fieldSeen[race.FieldIndx] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[race.FieldName]; !ok {
				selectedFields = append(selectedFields, race.FieldName)
				fieldSeen[race.FieldName] = struct{}{}
			}
		case "speed":
			if _, ok := fieldSeen[race.FieldSpeed]; !ok {
				selectedFields = append(selectedFields, race.FieldSpeed)
				fieldSeen[race.FieldSpeed] = struct{}{}
			}
		case "size":
			if _, ok := fieldSeen[race.FieldSize]; !ok {
				selectedFields = append(selectedFields, race.FieldSize)
				fieldSeen[race.FieldSize] = struct{}{}
			}
		case "sizeDesc":
			if _, ok := fieldSeen[race.FieldSizeDesc]; !ok {
				selectedFields = append(selectedFields, race.FieldSizeDesc)
				fieldSeen[race.FieldSizeDesc] = struct{}{}
			}
		case "alignmentDesc":
			if _, ok := fieldSeen[race.FieldAlignmentDesc]; !ok {
				selectedFields = append(selectedFields, race.FieldAlignmentDesc)
				fieldSeen[race.FieldAlignmentDesc] = struct{}{}
			}
		case "ageDesc":
			if _, ok := fieldSeen[race.FieldAgeDesc]; !ok {
				selectedFields = append(selectedFields, race.FieldAgeDesc)
				fieldSeen[race.FieldAgeDesc] = struct{}{}
			}
		case "languageDesc":
			if _, ok := fieldSeen[race.FieldLanguageDesc]; !ok {
				selectedFields = append(selectedFields, race.FieldLanguageDesc)
				fieldSeen[race.FieldLanguageDesc] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type racePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RacePaginateOption
}

func newRacePaginateArgs(rv map[string]any) *racePaginateArgs {
	args := &racePaginateArgs{}
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
				order      = &RaceOrder{Field: &RaceOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithRaceOrder(order))
			}
		case *RaceOrder:
			if v != nil {
				args.opts = append(args.opts, WithRaceOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*RaceWhereInput); ok {
		args.opts = append(args.opts, WithRaceFilter(v.Filter))
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
