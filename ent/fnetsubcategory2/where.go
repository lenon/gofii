// Code generated by ent, DO NOT EDIT.

package fnetsubcategory2

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lenon/gofii/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.FnetSubCategory2 {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.FnetSubCategory2 {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasDocuments applies the HasEdge predicate on the "documents" edge.
func HasDocuments() predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DocumentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DocumentsTable, DocumentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentsWith applies the HasEdge predicate on the "documents" edge with a given conditions (other predicates).
func HasDocumentsWith(preds ...predicate.FnetDocument) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DocumentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DocumentsTable, DocumentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FnetSubCategory2) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FnetSubCategory2) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
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
func Not(p predicate.FnetSubCategory2) predicate.FnetSubCategory2 {
	return predicate.FnetSubCategory2(func(s *sql.Selector) {
		p(s.Not())
	})
}
