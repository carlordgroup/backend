// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/account"
	"carlord/ent/card"
	"carlord/ent/flaw"
	"carlord/ent/predicate"
	"carlord/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetPostalCode sets the "postal_code" field.
func (uu *UserUpdate) SetPostalCode(s string) *UserUpdate {
	uu.mutation.SetPostalCode(s)
	return uu
}

// SetTel sets the "tel" field.
func (uu *UserUpdate) SetTel(s string) *UserUpdate {
	uu.mutation.SetTel(s)
	return uu
}

// SetDriverLicenseID sets the "driver_license_id" field.
func (uu *UserUpdate) SetDriverLicenseID(s string) *UserUpdate {
	uu.mutation.SetDriverLicenseID(s)
	return uu
}

// SetDriverLicenseCountry sets the "driver_license_country" field.
func (uu *UserUpdate) SetDriverLicenseCountry(s string) *UserUpdate {
	uu.mutation.SetDriverLicenseCountry(s)
	return uu
}

// SetBirthday sets the "birthday" field.
func (uu *UserUpdate) SetBirthday(t time.Time) *UserUpdate {
	uu.mutation.SetBirthday(t)
	return uu
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (uu *UserUpdate) AddCardIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCardIDs(ids...)
	return uu
}

// AddCard adds the "card" edges to the Card entity.
func (uu *UserUpdate) AddCard(c ...*Card) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCardIDs(ids...)
}

// AddNoteFlawIDs adds the "note_flaws" edge to the Flaw entity by IDs.
func (uu *UserUpdate) AddNoteFlawIDs(ids ...int) *UserUpdate {
	uu.mutation.AddNoteFlawIDs(ids...)
	return uu
}

// AddNoteFlaws adds the "note_flaws" edges to the Flaw entity.
func (uu *UserUpdate) AddNoteFlaws(f ...*Flaw) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddNoteFlawIDs(ids...)
}

// AddAccountIDs adds the "account" edge to the Account entity by IDs.
func (uu *UserUpdate) AddAccountIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAccountIDs(ids...)
	return uu
}

// AddAccount adds the "account" edges to the Account entity.
func (uu *UserUpdate) AddAccount(a ...*Account) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCard clears all "card" edges to the Card entity.
func (uu *UserUpdate) ClearCard() *UserUpdate {
	uu.mutation.ClearCard()
	return uu
}

// RemoveCardIDs removes the "card" edge to Card entities by IDs.
func (uu *UserUpdate) RemoveCardIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCardIDs(ids...)
	return uu
}

// RemoveCard removes "card" edges to Card entities.
func (uu *UserUpdate) RemoveCard(c ...*Card) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCardIDs(ids...)
}

// ClearNoteFlaws clears all "note_flaws" edges to the Flaw entity.
func (uu *UserUpdate) ClearNoteFlaws() *UserUpdate {
	uu.mutation.ClearNoteFlaws()
	return uu
}

// RemoveNoteFlawIDs removes the "note_flaws" edge to Flaw entities by IDs.
func (uu *UserUpdate) RemoveNoteFlawIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveNoteFlawIDs(ids...)
	return uu
}

// RemoveNoteFlaws removes "note_flaws" edges to Flaw entities.
func (uu *UserUpdate) RemoveNoteFlaws(f ...*Flaw) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveNoteFlawIDs(ids...)
}

// ClearAccount clears all "account" edges to the Account entity.
func (uu *UserUpdate) ClearAccount() *UserUpdate {
	uu.mutation.ClearAccount()
	return uu
}

// RemoveAccountIDs removes the "account" edge to Account entities by IDs.
func (uu *UserUpdate) RemoveAccountIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAccountIDs(ids...)
	return uu
}

// RemoveAccount removes "account" edges to Account entities.
func (uu *UserUpdate) RemoveAccount(a ...*Account) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if value, ok := uu.mutation.PostalCode(); ok {
		_spec.SetField(user.FieldPostalCode, field.TypeString, value)
	}
	if value, ok := uu.mutation.Tel(); ok {
		_spec.SetField(user.FieldTel, field.TypeString, value)
	}
	if value, ok := uu.mutation.DriverLicenseID(); ok {
		_spec.SetField(user.FieldDriverLicenseID, field.TypeString, value)
	}
	if value, ok := uu.mutation.DriverLicenseCountry(); ok {
		_spec.SetField(user.FieldDriverLicenseCountry, field.TypeString, value)
	}
	if value, ok := uu.mutation.Birthday(); ok {
		_spec.SetField(user.FieldBirthday, field.TypeTime, value)
	}
	if uu.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCardIDs(); len(nodes) > 0 && !uu.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.NoteFlawsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NoteFlawsTable,
			Columns: []string{user.NoteFlawsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flaw.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedNoteFlawsIDs(); len(nodes) > 0 && !uu.mutation.NoteFlawsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NoteFlawsTable,
			Columns: []string{user.NoteFlawsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flaw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.NoteFlawsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NoteFlawsTable,
			Columns: []string{user.NoteFlawsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flaw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.AccountTable,
			Columns: user.AccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAccountIDs(); len(nodes) > 0 && !uu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.AccountTable,
			Columns: user.AccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.AccountTable,
			Columns: user.AccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetPostalCode sets the "postal_code" field.
func (uuo *UserUpdateOne) SetPostalCode(s string) *UserUpdateOne {
	uuo.mutation.SetPostalCode(s)
	return uuo
}

// SetTel sets the "tel" field.
func (uuo *UserUpdateOne) SetTel(s string) *UserUpdateOne {
	uuo.mutation.SetTel(s)
	return uuo
}

// SetDriverLicenseID sets the "driver_license_id" field.
func (uuo *UserUpdateOne) SetDriverLicenseID(s string) *UserUpdateOne {
	uuo.mutation.SetDriverLicenseID(s)
	return uuo
}

// SetDriverLicenseCountry sets the "driver_license_country" field.
func (uuo *UserUpdateOne) SetDriverLicenseCountry(s string) *UserUpdateOne {
	uuo.mutation.SetDriverLicenseCountry(s)
	return uuo
}

// SetBirthday sets the "birthday" field.
func (uuo *UserUpdateOne) SetBirthday(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirthday(t)
	return uuo
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (uuo *UserUpdateOne) AddCardIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCardIDs(ids...)
	return uuo
}

// AddCard adds the "card" edges to the Card entity.
func (uuo *UserUpdateOne) AddCard(c ...*Card) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCardIDs(ids...)
}

// AddNoteFlawIDs adds the "note_flaws" edge to the Flaw entity by IDs.
func (uuo *UserUpdateOne) AddNoteFlawIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddNoteFlawIDs(ids...)
	return uuo
}

// AddNoteFlaws adds the "note_flaws" edges to the Flaw entity.
func (uuo *UserUpdateOne) AddNoteFlaws(f ...*Flaw) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddNoteFlawIDs(ids...)
}

// AddAccountIDs adds the "account" edge to the Account entity by IDs.
func (uuo *UserUpdateOne) AddAccountIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAccountIDs(ids...)
	return uuo
}

// AddAccount adds the "account" edges to the Account entity.
func (uuo *UserUpdateOne) AddAccount(a ...*Account) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCard clears all "card" edges to the Card entity.
func (uuo *UserUpdateOne) ClearCard() *UserUpdateOne {
	uuo.mutation.ClearCard()
	return uuo
}

// RemoveCardIDs removes the "card" edge to Card entities by IDs.
func (uuo *UserUpdateOne) RemoveCardIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCardIDs(ids...)
	return uuo
}

// RemoveCard removes "card" edges to Card entities.
func (uuo *UserUpdateOne) RemoveCard(c ...*Card) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCardIDs(ids...)
}

// ClearNoteFlaws clears all "note_flaws" edges to the Flaw entity.
func (uuo *UserUpdateOne) ClearNoteFlaws() *UserUpdateOne {
	uuo.mutation.ClearNoteFlaws()
	return uuo
}

// RemoveNoteFlawIDs removes the "note_flaws" edge to Flaw entities by IDs.
func (uuo *UserUpdateOne) RemoveNoteFlawIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveNoteFlawIDs(ids...)
	return uuo
}

// RemoveNoteFlaws removes "note_flaws" edges to Flaw entities.
func (uuo *UserUpdateOne) RemoveNoteFlaws(f ...*Flaw) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveNoteFlawIDs(ids...)
}

// ClearAccount clears all "account" edges to the Account entity.
func (uuo *UserUpdateOne) ClearAccount() *UserUpdateOne {
	uuo.mutation.ClearAccount()
	return uuo
}

// RemoveAccountIDs removes the "account" edge to Account entities by IDs.
func (uuo *UserUpdateOne) RemoveAccountIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAccountIDs(ids...)
	return uuo
}

// RemoveAccount removes "account" edges to Account entities.
func (uuo *UserUpdateOne) RemoveAccount(a ...*Account) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAccountIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PostalCode(); ok {
		_spec.SetField(user.FieldPostalCode, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Tel(); ok {
		_spec.SetField(user.FieldTel, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DriverLicenseID(); ok {
		_spec.SetField(user.FieldDriverLicenseID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DriverLicenseCountry(); ok {
		_spec.SetField(user.FieldDriverLicenseCountry, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Birthday(); ok {
		_spec.SetField(user.FieldBirthday, field.TypeTime, value)
	}
	if uuo.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCardIDs(); len(nodes) > 0 && !uuo.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.NoteFlawsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NoteFlawsTable,
			Columns: []string{user.NoteFlawsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flaw.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedNoteFlawsIDs(); len(nodes) > 0 && !uuo.mutation.NoteFlawsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NoteFlawsTable,
			Columns: []string{user.NoteFlawsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flaw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.NoteFlawsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NoteFlawsTable,
			Columns: []string{user.NoteFlawsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: flaw.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.AccountTable,
			Columns: user.AccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAccountIDs(); len(nodes) > 0 && !uuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.AccountTable,
			Columns: user.AccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.AccountTable,
			Columns: user.AccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
