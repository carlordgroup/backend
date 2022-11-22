// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/billing"
	"carlord/ent/booking"
	"carlord/ent/car"
	"carlord/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookingCreate is the builder for creating a Booking entity.
type BookingCreate struct {
	config
	mutation *BookingMutation
	hooks    []Hook
}

// SetStartAt sets the "start_at" field.
func (bc *BookingCreate) SetStartAt(t time.Time) *BookingCreate {
	bc.mutation.SetStartAt(t)
	return bc
}

// SetEndAt sets the "end_at" field.
func (bc *BookingCreate) SetEndAt(t time.Time) *BookingCreate {
	bc.mutation.SetEndAt(t)
	return bc
}

// SetReturnCarAt sets the "return_car_at" field.
func (bc *BookingCreate) SetReturnCarAt(t time.Time) *BookingCreate {
	bc.mutation.SetReturnCarAt(t)
	return bc
}

// SetNillableReturnCarAt sets the "return_car_at" field if the given value is not nil.
func (bc *BookingCreate) SetNillableReturnCarAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetReturnCarAt(*t)
	}
	return bc
}

// SetRate sets the "rate" field.
func (bc *BookingCreate) SetRate(f float32) *BookingCreate {
	bc.mutation.SetRate(f)
	return bc
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (bc *BookingCreate) SetNillableRate(f *float32) *BookingCreate {
	if f != nil {
		bc.SetRate(*f)
	}
	return bc
}

// SetExceedRate sets the "exceed_rate" field.
func (bc *BookingCreate) SetExceedRate(f float32) *BookingCreate {
	bc.mutation.SetExceedRate(f)
	return bc
}

// SetNillableExceedRate sets the "exceed_rate" field if the given value is not nil.
func (bc *BookingCreate) SetNillableExceedRate(f *float32) *BookingCreate {
	if f != nil {
		bc.SetExceedRate(*f)
	}
	return bc
}

// SetDeposit sets the "deposit" field.
func (bc *BookingCreate) SetDeposit(f float32) *BookingCreate {
	bc.mutation.SetDeposit(f)
	return bc
}

// SetNillableDeposit sets the "deposit" field if the given value is not nil.
func (bc *BookingCreate) SetNillableDeposit(f *float32) *BookingCreate {
	if f != nil {
		bc.SetDeposit(*f)
	}
	return bc
}

// SetFuelLevelAtBegin sets the "fuel_level_at_begin" field.
func (bc *BookingCreate) SetFuelLevelAtBegin(f float32) *BookingCreate {
	bc.mutation.SetFuelLevelAtBegin(f)
	return bc
}

// SetNillableFuelLevelAtBegin sets the "fuel_level_at_begin" field if the given value is not nil.
func (bc *BookingCreate) SetNillableFuelLevelAtBegin(f *float32) *BookingCreate {
	if f != nil {
		bc.SetFuelLevelAtBegin(*f)
	}
	return bc
}

// SetFuelLevelAtEnd sets the "fuel_level_at_end" field.
func (bc *BookingCreate) SetFuelLevelAtEnd(f float32) *BookingCreate {
	bc.mutation.SetFuelLevelAtEnd(f)
	return bc
}

// SetNillableFuelLevelAtEnd sets the "fuel_level_at_end" field if the given value is not nil.
func (bc *BookingCreate) SetNillableFuelLevelAtEnd(f *float32) *BookingCreate {
	if f != nil {
		bc.SetFuelLevelAtEnd(*f)
	}
	return bc
}

// SetMileageBegin sets the "mileage_begin" field.
func (bc *BookingCreate) SetMileageBegin(i int) *BookingCreate {
	bc.mutation.SetMileageBegin(i)
	return bc
}

// SetNillableMileageBegin sets the "mileage_begin" field if the given value is not nil.
func (bc *BookingCreate) SetNillableMileageBegin(i *int) *BookingCreate {
	if i != nil {
		bc.SetMileageBegin(*i)
	}
	return bc
}

// SetMileageEnd sets the "mileage_end" field.
func (bc *BookingCreate) SetMileageEnd(i int) *BookingCreate {
	bc.mutation.SetMileageEnd(i)
	return bc
}

// SetNillableMileageEnd sets the "mileage_end" field if the given value is not nil.
func (bc *BookingCreate) SetNillableMileageEnd(i *int) *BookingCreate {
	if i != nil {
		bc.SetMileageEnd(*i)
	}
	return bc
}

// SetBookingStatus sets the "booking_status" field.
func (bc *BookingCreate) SetBookingStatus(s string) *BookingCreate {
	bc.mutation.SetBookingStatus(s)
	return bc
}

// SetNillableBookingStatus sets the "booking_status" field if the given value is not nil.
func (bc *BookingCreate) SetNillableBookingStatus(s *string) *BookingCreate {
	if s != nil {
		bc.SetBookingStatus(*s)
	}
	return bc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bc *BookingCreate) SetUserID(id int) *BookingCreate {
	bc.mutation.SetUserID(id)
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BookingCreate) SetUser(u *User) *BookingCreate {
	return bc.SetUserID(u.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (bc *BookingCreate) SetCarID(id int) *BookingCreate {
	bc.mutation.SetCarID(id)
	return bc
}

// SetCar sets the "car" edge to the Car entity.
func (bc *BookingCreate) SetCar(c *Car) *BookingCreate {
	return bc.SetCarID(c.ID)
}

// SetBillingID sets the "billing" edge to the Billing entity by ID.
func (bc *BookingCreate) SetBillingID(id int) *BookingCreate {
	bc.mutation.SetBillingID(id)
	return bc
}

// SetNillableBillingID sets the "billing" edge to the Billing entity by ID if the given value is not nil.
func (bc *BookingCreate) SetNillableBillingID(id *int) *BookingCreate {
	if id != nil {
		bc = bc.SetBillingID(*id)
	}
	return bc
}

// SetBilling sets the "billing" edge to the Billing entity.
func (bc *BookingCreate) SetBilling(b *Billing) *BookingCreate {
	return bc.SetBillingID(b.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bc *BookingCreate) Mutation() *BookingMutation {
	return bc.mutation
}

// Save creates the Booking in the database.
func (bc *BookingCreate) Save(ctx context.Context) (*Booking, error) {
	var (
		err  error
		node *Booking
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Booking)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BookingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookingCreate) SaveX(ctx context.Context) *Booking {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookingCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookingCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookingCreate) defaults() {
	if _, ok := bc.mutation.Rate(); !ok {
		v := booking.DefaultRate
		bc.mutation.SetRate(v)
	}
	if _, ok := bc.mutation.ExceedRate(); !ok {
		v := booking.DefaultExceedRate
		bc.mutation.SetExceedRate(v)
	}
	if _, ok := bc.mutation.Deposit(); !ok {
		v := booking.DefaultDeposit
		bc.mutation.SetDeposit(v)
	}
	if _, ok := bc.mutation.BookingStatus(); !ok {
		v := booking.DefaultBookingStatus
		bc.mutation.SetBookingStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookingCreate) check() error {
	if _, ok := bc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "Booking.start_at"`)}
	}
	if _, ok := bc.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New(`ent: missing required field "Booking.end_at"`)}
	}
	if _, ok := bc.mutation.Rate(); !ok {
		return &ValidationError{Name: "rate", err: errors.New(`ent: missing required field "Booking.rate"`)}
	}
	if _, ok := bc.mutation.ExceedRate(); !ok {
		return &ValidationError{Name: "exceed_rate", err: errors.New(`ent: missing required field "Booking.exceed_rate"`)}
	}
	if _, ok := bc.mutation.Deposit(); !ok {
		return &ValidationError{Name: "deposit", err: errors.New(`ent: missing required field "Booking.deposit"`)}
	}
	if _, ok := bc.mutation.BookingStatus(); !ok {
		return &ValidationError{Name: "booking_status", err: errors.New(`ent: missing required field "Booking.booking_status"`)}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Booking.user"`)}
	}
	if _, ok := bc.mutation.CarID(); !ok {
		return &ValidationError{Name: "car", err: errors.New(`ent: missing required edge "Booking.car"`)}
	}
	return nil
}

func (bc *BookingCreate) sqlSave(ctx context.Context) (*Booking, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BookingCreate) createSpec() (*Booking, *sqlgraph.CreateSpec) {
	var (
		_node = &Booking{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: booking.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.StartAt(); ok {
		_spec.SetField(booking.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := bc.mutation.EndAt(); ok {
		_spec.SetField(booking.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := bc.mutation.ReturnCarAt(); ok {
		_spec.SetField(booking.FieldReturnCarAt, field.TypeTime, value)
		_node.ReturnCarAt = value
	}
	if value, ok := bc.mutation.Rate(); ok {
		_spec.SetField(booking.FieldRate, field.TypeFloat32, value)
		_node.Rate = value
	}
	if value, ok := bc.mutation.ExceedRate(); ok {
		_spec.SetField(booking.FieldExceedRate, field.TypeFloat32, value)
		_node.ExceedRate = value
	}
	if value, ok := bc.mutation.Deposit(); ok {
		_spec.SetField(booking.FieldDeposit, field.TypeFloat32, value)
		_node.Deposit = value
	}
	if value, ok := bc.mutation.FuelLevelAtBegin(); ok {
		_spec.SetField(booking.FieldFuelLevelAtBegin, field.TypeFloat32, value)
		_node.FuelLevelAtBegin = value
	}
	if value, ok := bc.mutation.FuelLevelAtEnd(); ok {
		_spec.SetField(booking.FieldFuelLevelAtEnd, field.TypeFloat32, value)
		_node.FuelLevelAtEnd = value
	}
	if value, ok := bc.mutation.MileageBegin(); ok {
		_spec.SetField(booking.FieldMileageBegin, field.TypeInt, value)
		_node.MileageBegin = value
	}
	if value, ok := bc.mutation.MileageEnd(); ok {
		_spec.SetField(booking.FieldMileageEnd, field.TypeInt, value)
		_node.MileageEnd = value
	}
	if value, ok := bc.mutation.BookingStatus(); ok {
		_spec.SetField(booking.FieldBookingStatus, field.TypeString, value)
		_node.BookingStatus = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   booking.UserTable,
			Columns: []string{booking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.booking_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   booking.CarTable,
			Columns: []string{booking.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.booking_car = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BillingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   booking.BillingTable,
			Columns: []string{booking.BillingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: billing.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.billing_booking = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookingCreateBulk is the builder for creating many Booking entities in bulk.
type BookingCreateBulk struct {
	config
	builders []*BookingCreate
}

// Save creates the Booking entities in the database.
func (bcb *BookingCreateBulk) Save(ctx context.Context) ([]*Booking, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Booking, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookingMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookingCreateBulk) SaveX(ctx context.Context) []*Booking {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookingCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookingCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
