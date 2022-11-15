// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/account"
	"carlord/ent/card"
	"carlord/ent/flaw"
	"carlord/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetFirstName sets the "first_name" field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetAddress sets the "address" field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.mutation.SetAddress(s)
	return uc
}

// SetPostalCode sets the "postal_code" field.
func (uc *UserCreate) SetPostalCode(s string) *UserCreate {
	uc.mutation.SetPostalCode(s)
	return uc
}

// SetTel sets the "tel" field.
func (uc *UserCreate) SetTel(s string) *UserCreate {
	uc.mutation.SetTel(s)
	return uc
}

// SetDriverLicenseID sets the "driver_license_id" field.
func (uc *UserCreate) SetDriverLicenseID(s string) *UserCreate {
	uc.mutation.SetDriverLicenseID(s)
	return uc
}

// SetDriverLicenseCountry sets the "driver_license_country" field.
func (uc *UserCreate) SetDriverLicenseCountry(s string) *UserCreate {
	uc.mutation.SetDriverLicenseCountry(s)
	return uc
}

// SetBirthday sets the "birthday" field.
func (uc *UserCreate) SetBirthday(t time.Time) *UserCreate {
	uc.mutation.SetBirthday(t)
	return uc
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (uc *UserCreate) AddCardIDs(ids ...int) *UserCreate {
	uc.mutation.AddCardIDs(ids...)
	return uc
}

// AddCard adds the "card" edges to the Card entity.
func (uc *UserCreate) AddCard(c ...*Card) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCardIDs(ids...)
}

// AddNoteFlawIDs adds the "note_flaws" edge to the Flaw entity by IDs.
func (uc *UserCreate) AddNoteFlawIDs(ids ...int) *UserCreate {
	uc.mutation.AddNoteFlawIDs(ids...)
	return uc
}

// AddNoteFlaws adds the "note_flaws" edges to the Flaw entity.
func (uc *UserCreate) AddNoteFlaws(f ...*Flaw) *UserCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uc.AddNoteFlawIDs(ids...)
}

// AddAccountIDs adds the "account" edge to the Account entity by IDs.
func (uc *UserCreate) AddAccountIDs(ids ...int) *UserCreate {
	uc.mutation.AddAccountIDs(ids...)
	return uc
}

// AddAccount adds the "account" edges to the Account entity.
func (uc *UserCreate) AddAccount(a ...*Account) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "User.first_name"`)}
	}
	if _, ok := uc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "User.last_name"`)}
	}
	if _, ok := uc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "User.address"`)}
	}
	if _, ok := uc.mutation.PostalCode(); !ok {
		return &ValidationError{Name: "postal_code", err: errors.New(`ent: missing required field "User.postal_code"`)}
	}
	if _, ok := uc.mutation.Tel(); !ok {
		return &ValidationError{Name: "tel", err: errors.New(`ent: missing required field "User.tel"`)}
	}
	if _, ok := uc.mutation.DriverLicenseID(); !ok {
		return &ValidationError{Name: "driver_license_id", err: errors.New(`ent: missing required field "User.driver_license_id"`)}
	}
	if _, ok := uc.mutation.DriverLicenseCountry(); !ok {
		return &ValidationError{Name: "driver_license_country", err: errors.New(`ent: missing required field "User.driver_license_country"`)}
	}
	if _, ok := uc.mutation.Birthday(); !ok {
		return &ValidationError{Name: "birthday", err: errors.New(`ent: missing required field "User.birthday"`)}
	}
	if len(uc.mutation.AccountIDs()) == 0 {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "User.account"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := uc.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := uc.mutation.PostalCode(); ok {
		_spec.SetField(user.FieldPostalCode, field.TypeString, value)
		_node.PostalCode = value
	}
	if value, ok := uc.mutation.Tel(); ok {
		_spec.SetField(user.FieldTel, field.TypeString, value)
		_node.Tel = value
	}
	if value, ok := uc.mutation.DriverLicenseID(); ok {
		_spec.SetField(user.FieldDriverLicenseID, field.TypeString, value)
		_node.DriverLicenseID = value
	}
	if value, ok := uc.mutation.DriverLicenseCountry(); ok {
		_spec.SetField(user.FieldDriverLicenseCountry, field.TypeString, value)
		_node.DriverLicenseCountry = value
	}
	if value, ok := uc.mutation.Birthday(); ok {
		_spec.SetField(user.FieldBirthday, field.TypeTime, value)
		_node.Birthday = value
	}
	if nodes := uc.mutation.CardIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.NoteFlawsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
