// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/billing"
	"carlord/ent/booking"
	"carlord/ent/card"
	"carlord/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Billing is the model entity for the Billing schema.
type Billing struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// BasicCost holds the value of the "basic_cost" field.
	BasicCost float32 `json:"basic_cost,omitempty"`
	// FuelCost holds the value of the "fuel_cost" field.
	FuelCost float32 `json:"fuel_cost,omitempty"`
	// Compensation holds the value of the "compensation" field.
	Compensation float32 `json:"compensation,omitempty"`
	// Deposit holds the value of the "deposit" field.
	Deposit float32 `json:"deposit,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillingQuery when eager-loading is set.
	Edges        BillingEdges `json:"edges"`
	billing_card *int
	billing_user *int
}

// BillingEdges holds the relations/edges for other nodes in the graph.
type BillingEdges struct {
	// Booking holds the value of the booking edge.
	Booking *Booking `json:"booking,omitempty"`
	// Card holds the value of the card edge.
	Card *Card `json:"card,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BookingOrErr returns the Booking value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingEdges) BookingOrErr() (*Booking, error) {
	if e.loadedTypes[0] {
		if e.Booking == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: booking.Label}
		}
		return e.Booking, nil
	}
	return nil, &NotLoadedError{edge: "booking"}
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingEdges) CardOrErr() (*Card, error) {
	if e.loadedTypes[1] {
		if e.Card == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: card.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Billing) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billing.FieldBasicCost, billing.FieldFuelCost, billing.FieldCompensation, billing.FieldDeposit:
			values[i] = new(sql.NullFloat64)
		case billing.FieldID:
			values[i] = new(sql.NullInt64)
		case billing.FieldStatus:
			values[i] = new(sql.NullString)
		case billing.ForeignKeys[0]: // billing_card
			values[i] = new(sql.NullInt64)
		case billing.ForeignKeys[1]: // billing_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Billing", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Billing fields.
func (b *Billing) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billing.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case billing.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = value.String
			}
		case billing.FieldBasicCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field basic_cost", values[i])
			} else if value.Valid {
				b.BasicCost = float32(value.Float64)
			}
		case billing.FieldFuelCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fuel_cost", values[i])
			} else if value.Valid {
				b.FuelCost = float32(value.Float64)
			}
		case billing.FieldCompensation:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field compensation", values[i])
			} else if value.Valid {
				b.Compensation = float32(value.Float64)
			}
		case billing.FieldDeposit:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field deposit", values[i])
			} else if value.Valid {
				b.Deposit = float32(value.Float64)
			}
		case billing.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field billing_card", value)
			} else if value.Valid {
				b.billing_card = new(int)
				*b.billing_card = int(value.Int64)
			}
		case billing.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field billing_user", value)
			} else if value.Valid {
				b.billing_user = new(int)
				*b.billing_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBooking queries the "booking" edge of the Billing entity.
func (b *Billing) QueryBooking() *BookingQuery {
	return (&BillingClient{config: b.config}).QueryBooking(b)
}

// QueryCard queries the "card" edge of the Billing entity.
func (b *Billing) QueryCard() *CardQuery {
	return (&BillingClient{config: b.config}).QueryCard(b)
}

// QueryUser queries the "user" edge of the Billing entity.
func (b *Billing) QueryUser() *UserQuery {
	return (&BillingClient{config: b.config}).QueryUser(b)
}

// Update returns a builder for updating this Billing.
// Note that you need to call Billing.Unwrap() before calling this method if this Billing
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Billing) Update() *BillingUpdateOne {
	return (&BillingClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Billing entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Billing) Unwrap() *Billing {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Billing is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Billing) String() string {
	var builder strings.Builder
	builder.WriteString("Billing(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("status=")
	builder.WriteString(b.Status)
	builder.WriteString(", ")
	builder.WriteString("basic_cost=")
	builder.WriteString(fmt.Sprintf("%v", b.BasicCost))
	builder.WriteString(", ")
	builder.WriteString("fuel_cost=")
	builder.WriteString(fmt.Sprintf("%v", b.FuelCost))
	builder.WriteString(", ")
	builder.WriteString("compensation=")
	builder.WriteString(fmt.Sprintf("%v", b.Compensation))
	builder.WriteString(", ")
	builder.WriteString("deposit=")
	builder.WriteString(fmt.Sprintf("%v", b.Deposit))
	builder.WriteByte(')')
	return builder.String()
}

// Billings is a parsable slice of Billing.
type Billings []*Billing

func (b Billings) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
