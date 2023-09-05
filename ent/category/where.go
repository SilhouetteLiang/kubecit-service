// Code generated by ent, DO NOT EDIT.

package category

import (
	"kubecit-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldLevel, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldStatus, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldParentID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldName, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldLevel, v))
}

// LevelContains applies the Contains predicate on the "level" field.
func LevelContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldLevel, v))
}

// LevelHasPrefix applies the HasPrefix predicate on the "level" field.
func LevelHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldLevel, v))
}

// LevelHasSuffix applies the HasSuffix predicate on the "level" field.
func LevelHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldLevel, v))
}

// LevelEqualFold applies the EqualFold predicate on the "level" field.
func LevelEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldLevel, v))
}

// LevelContainsFold applies the ContainsFold predicate on the "level" field.
func LevelContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldLevel, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldStatus, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldParentID))
}

// HasCourses applies the HasEdge predicate on the "courses" edge.
func HasCourses() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CoursesTable, CoursesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoursesWith applies the HasEdge predicate on the "courses" edge with a given conditions (other predicates).
func HasCoursesWith(preds ...predicate.Course) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := newCoursesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
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
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		p(s.Not())
	})
}