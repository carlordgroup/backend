// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/flaw"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Flaw is the model entity for the Flaw schema.
type Flaw struct {
	config
	// ID of the ent.
	ID              int `json:"id,omitempty"`
	user_note_flows *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flaw) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case flaw.FieldID:
			values[i] = new(sql.NullInt64)
		case flaw.ForeignKeys[0]: // user_note_flows
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Flaw", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Flaw fields.
func (f *Flaw) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flaw.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case flaw.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_note_flows", value)
			} else if value.Valid {
				f.user_note_flows = new(int)
				*f.user_note_flows = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Flaw.
// Note that you need to call Flaw.Unwrap() before calling this method if this Flaw
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Flaw) Update() *FlawUpdateOne {
	return (&FlawClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Flaw entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Flaw) Unwrap() *Flaw {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Flaw is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Flaw) String() string {
	var builder strings.Builder
	builder.WriteString("Flaw(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Flaws is a parsable slice of Flaw.
type Flaws []*Flaw

func (f Flaws) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}