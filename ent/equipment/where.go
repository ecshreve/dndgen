// Code generated by ent, DO NOT EDIT.

package equipment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldName, v))
}

// EquipmentSubcategory applies equality check predicate on the "equipment_subcategory" field. It's identical to EquipmentSubcategoryEQ.
func EquipmentSubcategory(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldEquipmentSubcategory, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContainsFold(FieldName, v))
}

// EquipmentCategoryEQ applies the EQ predicate on the "equipment_category" field.
func EquipmentCategoryEQ(v EquipmentCategory) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldEquipmentCategory, v))
}

// EquipmentCategoryNEQ applies the NEQ predicate on the "equipment_category" field.
func EquipmentCategoryNEQ(v EquipmentCategory) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldEquipmentCategory, v))
}

// EquipmentCategoryIn applies the In predicate on the "equipment_category" field.
func EquipmentCategoryIn(vs ...EquipmentCategory) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldEquipmentCategory, vs...))
}

// EquipmentCategoryNotIn applies the NotIn predicate on the "equipment_category" field.
func EquipmentCategoryNotIn(vs ...EquipmentCategory) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldEquipmentCategory, vs...))
}

// EquipmentSubcategoryEQ applies the EQ predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryNEQ applies the NEQ predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryNEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryIn applies the In predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldEquipmentSubcategory, vs...))
}

// EquipmentSubcategoryNotIn applies the NotIn predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryNotIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldEquipmentSubcategory, vs...))
}

// EquipmentSubcategoryGT applies the GT predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryGT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryGTE applies the GTE predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryGTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryLT applies the LT predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryLT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryLTE applies the LTE predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryLTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryContains applies the Contains predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryContains(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContains(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryHasPrefix applies the HasPrefix predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasPrefix(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryHasSuffix applies the HasSuffix predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasSuffix(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryIsNil applies the IsNil predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryIsNil() predicate.Equipment {
	return predicate.Equipment(sql.FieldIsNull(FieldEquipmentSubcategory))
}

// EquipmentSubcategoryNotNil applies the NotNil predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryNotNil() predicate.Equipment {
	return predicate.Equipment(sql.FieldNotNull(FieldEquipmentSubcategory))
}

// EquipmentSubcategoryEqualFold applies the EqualFold predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEqualFold(FieldEquipmentSubcategory, v))
}

// EquipmentSubcategoryContainsFold applies the ContainsFold predicate on the "equipment_subcategory" field.
func EquipmentSubcategoryContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContainsFold(FieldEquipmentSubcategory, v))
}

// HasCost applies the HasEdge predicate on the "cost" edge.
func HasCost() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CostTable, CostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCostWith applies the HasEdge predicate on the "cost" edge with a given conditions (other predicates).
func HasCostWith(preds ...predicate.Cost) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newCostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWeapon applies the HasEdge predicate on the "weapon" edge.
func HasWeapon() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, WeaponTable, WeaponColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeaponWith applies the HasEdge predicate on the "weapon" edge with a given conditions (other predicates).
func HasWeaponWith(preds ...predicate.Weapon) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newWeaponStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArmor applies the HasEdge predicate on the "armor" edge.
func HasArmor() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ArmorTable, ArmorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArmorWith applies the HasEdge predicate on the "armor" edge with a given conditions (other predicates).
func HasArmorWith(preds ...predicate.Armor) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newArmorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGear applies the HasEdge predicate on the "gear" edge.
func HasGear() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, GearTable, GearColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGearWith applies the HasEdge predicate on the "gear" edge with a given conditions (other predicates).
func HasGearWith(preds ...predicate.Gear) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newGearStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTool applies the HasEdge predicate on the "tool" edge.
func HasTool() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ToolTable, ToolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToolWith applies the HasEdge predicate on the "tool" edge with a given conditions (other predicates).
func HasToolWith(preds ...predicate.Tool) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newToolStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVehicle applies the HasEdge predicate on the "vehicle" edge.
func HasVehicle() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, VehicleTable, VehicleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVehicleWith applies the HasEdge predicate on the "vehicle" edge with a given conditions (other predicates).
func HasVehicleWith(preds ...predicate.Vehicle) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newVehicleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassEquipment applies the HasEdge predicate on the "class_equipment" edge.
func HasClassEquipment() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ClassEquipmentTable, ClassEquipmentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassEquipmentWith applies the HasEdge predicate on the "class_equipment" edge with a given conditions (other predicates).
func HasClassEquipmentWith(preds ...predicate.Class) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newClassEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChoice applies the HasEdge predicate on the "choice" edge.
func HasChoice() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChoiceTable, ChoicePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChoiceWith applies the HasEdge predicate on the "choice" edge with a given conditions (other predicates).
func HasChoiceWith(preds ...predicate.EquipmentChoice) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newChoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassStartingEquipment applies the HasEdge predicate on the "class_starting_equipment" edge.
func HasClassStartingEquipment() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ClassStartingEquipmentTable, ClassStartingEquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassStartingEquipmentWith applies the HasEdge predicate on the "class_starting_equipment" edge with a given conditions (other predicates).
func HasClassStartingEquipmentWith(preds ...predicate.StartingEquipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newClassStartingEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		p(s.Not())
	})
}
