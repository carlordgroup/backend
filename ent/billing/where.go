// Code generated by ent, DO NOT EDIT.

package billing

import (
	"carlord/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// BasicCost applies equality check predicate on the "basic_cost" field. It's identical to BasicCostEQ.
func BasicCost(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBasicCost), v))
	})
}

// FuelCost applies equality check predicate on the "fuel_cost" field. It's identical to FuelCostEQ.
func FuelCost(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFuelCost), v))
	})
}

// Compensation applies equality check predicate on the "compensation" field. It's identical to CompensationEQ.
func Compensation(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensation), v))
	})
}

// Deposit applies equality check predicate on the "deposit" field. It's identical to DepositEQ.
func Deposit(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeposit), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// BasicCostEQ applies the EQ predicate on the "basic_cost" field.
func BasicCostEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBasicCost), v))
	})
}

// BasicCostNEQ applies the NEQ predicate on the "basic_cost" field.
func BasicCostNEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBasicCost), v))
	})
}

// BasicCostIn applies the In predicate on the "basic_cost" field.
func BasicCostIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBasicCost), v...))
	})
}

// BasicCostNotIn applies the NotIn predicate on the "basic_cost" field.
func BasicCostNotIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBasicCost), v...))
	})
}

// BasicCostGT applies the GT predicate on the "basic_cost" field.
func BasicCostGT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBasicCost), v))
	})
}

// BasicCostGTE applies the GTE predicate on the "basic_cost" field.
func BasicCostGTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBasicCost), v))
	})
}

// BasicCostLT applies the LT predicate on the "basic_cost" field.
func BasicCostLT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBasicCost), v))
	})
}

// BasicCostLTE applies the LTE predicate on the "basic_cost" field.
func BasicCostLTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBasicCost), v))
	})
}

// FuelCostEQ applies the EQ predicate on the "fuel_cost" field.
func FuelCostEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFuelCost), v))
	})
}

// FuelCostNEQ applies the NEQ predicate on the "fuel_cost" field.
func FuelCostNEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFuelCost), v))
	})
}

// FuelCostIn applies the In predicate on the "fuel_cost" field.
func FuelCostIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFuelCost), v...))
	})
}

// FuelCostNotIn applies the NotIn predicate on the "fuel_cost" field.
func FuelCostNotIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFuelCost), v...))
	})
}

// FuelCostGT applies the GT predicate on the "fuel_cost" field.
func FuelCostGT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFuelCost), v))
	})
}

// FuelCostGTE applies the GTE predicate on the "fuel_cost" field.
func FuelCostGTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFuelCost), v))
	})
}

// FuelCostLT applies the LT predicate on the "fuel_cost" field.
func FuelCostLT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFuelCost), v))
	})
}

// FuelCostLTE applies the LTE predicate on the "fuel_cost" field.
func FuelCostLTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFuelCost), v))
	})
}

// CompensationEQ applies the EQ predicate on the "compensation" field.
func CompensationEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensation), v))
	})
}

// CompensationNEQ applies the NEQ predicate on the "compensation" field.
func CompensationNEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompensation), v))
	})
}

// CompensationIn applies the In predicate on the "compensation" field.
func CompensationIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompensation), v...))
	})
}

// CompensationNotIn applies the NotIn predicate on the "compensation" field.
func CompensationNotIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompensation), v...))
	})
}

// CompensationGT applies the GT predicate on the "compensation" field.
func CompensationGT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompensation), v))
	})
}

// CompensationGTE applies the GTE predicate on the "compensation" field.
func CompensationGTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompensation), v))
	})
}

// CompensationLT applies the LT predicate on the "compensation" field.
func CompensationLT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompensation), v))
	})
}

// CompensationLTE applies the LTE predicate on the "compensation" field.
func CompensationLTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompensation), v))
	})
}

// DepositEQ applies the EQ predicate on the "deposit" field.
func DepositEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeposit), v))
	})
}

// DepositNEQ applies the NEQ predicate on the "deposit" field.
func DepositNEQ(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeposit), v))
	})
}

// DepositIn applies the In predicate on the "deposit" field.
func DepositIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeposit), v...))
	})
}

// DepositNotIn applies the NotIn predicate on the "deposit" field.
func DepositNotIn(vs ...float32) predicate.Billing {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeposit), v...))
	})
}

// DepositGT applies the GT predicate on the "deposit" field.
func DepositGT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeposit), v))
	})
}

// DepositGTE applies the GTE predicate on the "deposit" field.
func DepositGTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeposit), v))
	})
}

// DepositLT applies the LT predicate on the "deposit" field.
func DepositLT(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeposit), v))
	})
}

// DepositLTE applies the LTE predicate on the "deposit" field.
func DepositLTE(v float32) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeposit), v))
	})
}

// HasBooking applies the HasEdge predicate on the "booking" edge.
func HasBooking() predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BookingTable, BookingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingWith applies the HasEdge predicate on the "booking" edge with a given conditions (other predicates).
func HasBookingWith(preds ...predicate.Booking) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BookingTable, BookingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...predicate.Card) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Billing) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Billing) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
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
func Not(p predicate.Billing) predicate.Billing {
	return predicate.Billing(func(s *sql.Selector) {
		p(s.Not())
	})
}
