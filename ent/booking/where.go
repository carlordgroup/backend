// Code generated by ent, DO NOT EDIT.

package booking

import (
	"carlord/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// ReturnCarAt applies equality check predicate on the "return_car_at" field. It's identical to ReturnCarAtEQ.
func ReturnCarAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReturnCarAt), v))
	})
}

// Rate applies equality check predicate on the "rate" field. It's identical to RateEQ.
func Rate(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRate), v))
	})
}

// ExceedRate applies equality check predicate on the "exceed_rate" field. It's identical to ExceedRateEQ.
func ExceedRate(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExceedRate), v))
	})
}

// Deposit applies equality check predicate on the "deposit" field. It's identical to DepositEQ.
func Deposit(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeposit), v))
	})
}

// FuelLevelAtBegin applies equality check predicate on the "fuel_level_at_begin" field. It's identical to FuelLevelAtBeginEQ.
func FuelLevelAtBegin(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtEnd applies equality check predicate on the "fuel_level_at_end" field. It's identical to FuelLevelAtEndEQ.
func FuelLevelAtEnd(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFuelLevelAtEnd), v))
	})
}

// MileageBegin applies equality check predicate on the "mileage_begin" field. It's identical to MileageBeginEQ.
func MileageBegin(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMileageBegin), v))
	})
}

// MileageEnd applies equality check predicate on the "mileage_end" field. It's identical to MileageEndEQ.
func MileageEnd(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMileageEnd), v))
	})
}

// BookingStatus applies equality check predicate on the "booking_status" field. It's identical to BookingStatusEQ.
func BookingStatus(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBookingStatus), v))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// ReturnCarAtEQ applies the EQ predicate on the "return_car_at" field.
func ReturnCarAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReturnCarAt), v))
	})
}

// ReturnCarAtNEQ applies the NEQ predicate on the "return_car_at" field.
func ReturnCarAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReturnCarAt), v))
	})
}

// ReturnCarAtIn applies the In predicate on the "return_car_at" field.
func ReturnCarAtIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReturnCarAt), v...))
	})
}

// ReturnCarAtNotIn applies the NotIn predicate on the "return_car_at" field.
func ReturnCarAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReturnCarAt), v...))
	})
}

// ReturnCarAtGT applies the GT predicate on the "return_car_at" field.
func ReturnCarAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReturnCarAt), v))
	})
}

// ReturnCarAtGTE applies the GTE predicate on the "return_car_at" field.
func ReturnCarAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReturnCarAt), v))
	})
}

// ReturnCarAtLT applies the LT predicate on the "return_car_at" field.
func ReturnCarAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReturnCarAt), v))
	})
}

// ReturnCarAtLTE applies the LTE predicate on the "return_car_at" field.
func ReturnCarAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReturnCarAt), v))
	})
}

// RateEQ applies the EQ predicate on the "rate" field.
func RateEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRate), v))
	})
}

// RateNEQ applies the NEQ predicate on the "rate" field.
func RateNEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRate), v))
	})
}

// RateIn applies the In predicate on the "rate" field.
func RateIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRate), v...))
	})
}

// RateNotIn applies the NotIn predicate on the "rate" field.
func RateNotIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRate), v...))
	})
}

// RateGT applies the GT predicate on the "rate" field.
func RateGT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRate), v))
	})
}

// RateGTE applies the GTE predicate on the "rate" field.
func RateGTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRate), v))
	})
}

// RateLT applies the LT predicate on the "rate" field.
func RateLT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRate), v))
	})
}

// RateLTE applies the LTE predicate on the "rate" field.
func RateLTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRate), v))
	})
}

// ExceedRateEQ applies the EQ predicate on the "exceed_rate" field.
func ExceedRateEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExceedRate), v))
	})
}

// ExceedRateNEQ applies the NEQ predicate on the "exceed_rate" field.
func ExceedRateNEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExceedRate), v))
	})
}

// ExceedRateIn applies the In predicate on the "exceed_rate" field.
func ExceedRateIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExceedRate), v...))
	})
}

// ExceedRateNotIn applies the NotIn predicate on the "exceed_rate" field.
func ExceedRateNotIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExceedRate), v...))
	})
}

// ExceedRateGT applies the GT predicate on the "exceed_rate" field.
func ExceedRateGT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExceedRate), v))
	})
}

// ExceedRateGTE applies the GTE predicate on the "exceed_rate" field.
func ExceedRateGTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExceedRate), v))
	})
}

// ExceedRateLT applies the LT predicate on the "exceed_rate" field.
func ExceedRateLT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExceedRate), v))
	})
}

// ExceedRateLTE applies the LTE predicate on the "exceed_rate" field.
func ExceedRateLTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExceedRate), v))
	})
}

// DepositEQ applies the EQ predicate on the "deposit" field.
func DepositEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeposit), v))
	})
}

// DepositNEQ applies the NEQ predicate on the "deposit" field.
func DepositNEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeposit), v))
	})
}

// DepositIn applies the In predicate on the "deposit" field.
func DepositIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeposit), v...))
	})
}

// DepositNotIn applies the NotIn predicate on the "deposit" field.
func DepositNotIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeposit), v...))
	})
}

// DepositGT applies the GT predicate on the "deposit" field.
func DepositGT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeposit), v))
	})
}

// DepositGTE applies the GTE predicate on the "deposit" field.
func DepositGTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeposit), v))
	})
}

// DepositLT applies the LT predicate on the "deposit" field.
func DepositLT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeposit), v))
	})
}

// DepositLTE applies the LTE predicate on the "deposit" field.
func DepositLTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeposit), v))
	})
}

// FuelLevelAtBeginEQ applies the EQ predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtBeginNEQ applies the NEQ predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginNEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtBeginIn applies the In predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFuelLevelAtBegin), v...))
	})
}

// FuelLevelAtBeginNotIn applies the NotIn predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginNotIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFuelLevelAtBegin), v...))
	})
}

// FuelLevelAtBeginGT applies the GT predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginGT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtBeginGTE applies the GTE predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginGTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtBeginLT applies the LT predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginLT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtBeginLTE applies the LTE predicate on the "fuel_level_at_begin" field.
func FuelLevelAtBeginLTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFuelLevelAtBegin), v))
	})
}

// FuelLevelAtEndEQ applies the EQ predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFuelLevelAtEnd), v))
	})
}

// FuelLevelAtEndNEQ applies the NEQ predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndNEQ(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFuelLevelAtEnd), v))
	})
}

// FuelLevelAtEndIn applies the In predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFuelLevelAtEnd), v...))
	})
}

// FuelLevelAtEndNotIn applies the NotIn predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndNotIn(vs ...float32) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFuelLevelAtEnd), v...))
	})
}

// FuelLevelAtEndGT applies the GT predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndGT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFuelLevelAtEnd), v))
	})
}

// FuelLevelAtEndGTE applies the GTE predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndGTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFuelLevelAtEnd), v))
	})
}

// FuelLevelAtEndLT applies the LT predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndLT(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFuelLevelAtEnd), v))
	})
}

// FuelLevelAtEndLTE applies the LTE predicate on the "fuel_level_at_end" field.
func FuelLevelAtEndLTE(v float32) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFuelLevelAtEnd), v))
	})
}

// MileageBeginEQ applies the EQ predicate on the "mileage_begin" field.
func MileageBeginEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMileageBegin), v))
	})
}

// MileageBeginNEQ applies the NEQ predicate on the "mileage_begin" field.
func MileageBeginNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMileageBegin), v))
	})
}

// MileageBeginIn applies the In predicate on the "mileage_begin" field.
func MileageBeginIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMileageBegin), v...))
	})
}

// MileageBeginNotIn applies the NotIn predicate on the "mileage_begin" field.
func MileageBeginNotIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMileageBegin), v...))
	})
}

// MileageBeginGT applies the GT predicate on the "mileage_begin" field.
func MileageBeginGT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMileageBegin), v))
	})
}

// MileageBeginGTE applies the GTE predicate on the "mileage_begin" field.
func MileageBeginGTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMileageBegin), v))
	})
}

// MileageBeginLT applies the LT predicate on the "mileage_begin" field.
func MileageBeginLT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMileageBegin), v))
	})
}

// MileageBeginLTE applies the LTE predicate on the "mileage_begin" field.
func MileageBeginLTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMileageBegin), v))
	})
}

// MileageEndEQ applies the EQ predicate on the "mileage_end" field.
func MileageEndEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMileageEnd), v))
	})
}

// MileageEndNEQ applies the NEQ predicate on the "mileage_end" field.
func MileageEndNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMileageEnd), v))
	})
}

// MileageEndIn applies the In predicate on the "mileage_end" field.
func MileageEndIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMileageEnd), v...))
	})
}

// MileageEndNotIn applies the NotIn predicate on the "mileage_end" field.
func MileageEndNotIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMileageEnd), v...))
	})
}

// MileageEndGT applies the GT predicate on the "mileage_end" field.
func MileageEndGT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMileageEnd), v))
	})
}

// MileageEndGTE applies the GTE predicate on the "mileage_end" field.
func MileageEndGTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMileageEnd), v))
	})
}

// MileageEndLT applies the LT predicate on the "mileage_end" field.
func MileageEndLT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMileageEnd), v))
	})
}

// MileageEndLTE applies the LTE predicate on the "mileage_end" field.
func MileageEndLTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMileageEnd), v))
	})
}

// BookingStatusEQ applies the EQ predicate on the "booking_status" field.
func BookingStatusEQ(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusNEQ applies the NEQ predicate on the "booking_status" field.
func BookingStatusNEQ(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusIn applies the In predicate on the "booking_status" field.
func BookingStatusIn(vs ...string) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBookingStatus), v...))
	})
}

// BookingStatusNotIn applies the NotIn predicate on the "booking_status" field.
func BookingStatusNotIn(vs ...string) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBookingStatus), v...))
	})
}

// BookingStatusGT applies the GT predicate on the "booking_status" field.
func BookingStatusGT(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusGTE applies the GTE predicate on the "booking_status" field.
func BookingStatusGTE(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusLT applies the LT predicate on the "booking_status" field.
func BookingStatusLT(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusLTE applies the LTE predicate on the "booking_status" field.
func BookingStatusLTE(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusContains applies the Contains predicate on the "booking_status" field.
func BookingStatusContains(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusHasPrefix applies the HasPrefix predicate on the "booking_status" field.
func BookingStatusHasPrefix(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusHasSuffix applies the HasSuffix predicate on the "booking_status" field.
func BookingStatusHasSuffix(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusEqualFold applies the EqualFold predicate on the "booking_status" field.
func BookingStatusEqualFold(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBookingStatus), v))
	})
}

// BookingStatusContainsFold applies the ContainsFold predicate on the "booking_status" field.
func BookingStatusContainsFold(v string) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBookingStatus), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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

// HasCar applies the HasEdge predicate on the "car" edge.
func HasCar() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CarTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CarTable, CarColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCarWith applies the HasEdge predicate on the "car" edge with a given conditions (other predicates).
func HasCarWith(preds ...predicate.Car) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CarInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CarTable, CarColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBilling applies the HasEdge predicate on the "billing" edge.
func HasBilling() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillingTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, BillingTable, BillingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingWith applies the HasEdge predicate on the "billing" edge with a given conditions (other predicates).
func HasBillingWith(preds ...predicate.Billing) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, BillingTable, BillingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
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
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
