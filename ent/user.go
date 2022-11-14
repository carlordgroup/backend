// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// IsAdmin holds the value of the "is_admin" field.
	IsAdmin bool `json:"is_admin,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode string `json:"postal_code,omitempty"`
	// Tel holds the value of the "tel" field.
	Tel string `json:"tel,omitempty"`
	// DriverLicenseID holds the value of the "driver_license_id" field.
	DriverLicenseID string `json:"driver_license_id,omitempty"`
	// DriverLicenseCountry holds the value of the "driver_license_country" field.
	DriverLicenseCountry string `json:"driver_license_country,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday time.Time `json:"birthday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Card holds the value of the card edge.
	Card []*Card `json:"card,omitempty"`
	// NoteFlows holds the value of the note_flows edge.
	NoteFlows []*Flaw `json:"note_flows,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// NoteFlowsOrErr returns the NoteFlows value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NoteFlowsOrErr() ([]*Flaw, error) {
	if e.loadedTypes[1] {
		return e.NoteFlows, nil
	}
	return nil, &NotLoadedError{edge: "note_flows"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsAdmin:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldPassword, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldAddress, user.FieldPostalCode, user.FieldTel, user.FieldDriverLicenseID, user.FieldDriverLicenseCountry:
			values[i] = new(sql.NullString)
		case user.FieldBirthday:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldIsAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_admin", values[i])
			} else if value.Valid {
				u.IsAdmin = value.Bool
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = value.String
			}
		case user.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				u.PostalCode = value.String
			}
		case user.FieldTel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tel", values[i])
			} else if value.Valid {
				u.Tel = value.String
			}
		case user.FieldDriverLicenseID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field driver_license_id", values[i])
			} else if value.Valid {
				u.DriverLicenseID = value.String
			}
		case user.FieldDriverLicenseCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field driver_license_country", values[i])
			} else if value.Valid {
				u.DriverLicenseCountry = value.String
			}
		case user.FieldBirthday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				u.Birthday = value.Time
			}
		}
	}
	return nil
}

// QueryCard queries the "card" edge of the User entity.
func (u *User) QueryCard() *CardQuery {
	return (&UserClient{config: u.config}).QueryCard(u)
}

// QueryNoteFlows queries the "note_flows" edge of the User entity.
func (u *User) QueryNoteFlows() *FlawQuery {
	return (&UserClient{config: u.config}).QueryNoteFlows(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("is_admin=")
	builder.WriteString(fmt.Sprintf("%v", u.IsAdmin))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(u.Address)
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(u.PostalCode)
	builder.WriteString(", ")
	builder.WriteString("tel=")
	builder.WriteString(u.Tel)
	builder.WriteString(", ")
	builder.WriteString("driver_license_id=")
	builder.WriteString(u.DriverLicenseID)
	builder.WriteString(", ")
	builder.WriteString("driver_license_country=")
	builder.WriteString(u.DriverLicenseCountry)
	builder.WriteString(", ")
	builder.WriteString("birthday=")
	builder.WriteString(u.Birthday.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
