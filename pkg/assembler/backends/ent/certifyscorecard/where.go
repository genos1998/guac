// Code generated by ent, DO NOT EDIT.

package certifyscorecard

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldLTE(FieldID, id))
}

// SourceID applies equality check predicate on the "source_id" field. It's identical to SourceIDEQ.
func SourceID(v int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldEQ(FieldSourceID, v))
}

// ScorecardID applies equality check predicate on the "scorecard_id" field. It's identical to ScorecardIDEQ.
func ScorecardID(v int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldEQ(FieldScorecardID, v))
}

// SourceIDEQ applies the EQ predicate on the "source_id" field.
func SourceIDEQ(v int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldEQ(FieldSourceID, v))
}

// SourceIDNEQ applies the NEQ predicate on the "source_id" field.
func SourceIDNEQ(v int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldNEQ(FieldSourceID, v))
}

// SourceIDIn applies the In predicate on the "source_id" field.
func SourceIDIn(vs ...int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldIn(FieldSourceID, vs...))
}

// SourceIDNotIn applies the NotIn predicate on the "source_id" field.
func SourceIDNotIn(vs ...int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldNotIn(FieldSourceID, vs...))
}

// ScorecardIDEQ applies the EQ predicate on the "scorecard_id" field.
func ScorecardIDEQ(v int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldEQ(FieldScorecardID, v))
}

// ScorecardIDNEQ applies the NEQ predicate on the "scorecard_id" field.
func ScorecardIDNEQ(v int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldNEQ(FieldScorecardID, v))
}

// ScorecardIDIn applies the In predicate on the "scorecard_id" field.
func ScorecardIDIn(vs ...int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldIn(FieldScorecardID, vs...))
}

// ScorecardIDNotIn applies the NotIn predicate on the "scorecard_id" field.
func ScorecardIDNotIn(vs ...int) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.FieldNotIn(FieldScorecardID, vs...))
}

// HasScorecard applies the HasEdge predicate on the "scorecard" edge.
func HasScorecard() predicate.CertifyScorecard {
	return predicate.CertifyScorecard(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScorecardTable, ScorecardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScorecardWith applies the HasEdge predicate on the "scorecard" edge with a given conditions (other predicates).
func HasScorecardWith(preds ...predicate.Scorecard) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(func(s *sql.Selector) {
		step := newScorecardStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSource applies the HasEdge predicate on the "source" edge.
func HasSource() predicate.CertifyScorecard {
	return predicate.CertifyScorecard(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SourceTable, SourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourceWith applies the HasEdge predicate on the "source" edge with a given conditions (other predicates).
func HasSourceWith(preds ...predicate.SourceName) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(func(s *sql.Selector) {
		step := newSourceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CertifyScorecard) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CertifyScorecard) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CertifyScorecard) predicate.CertifyScorecard {
	return predicate.CertifyScorecard(sql.NotPredicates(p))
}
