// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"carlord/ent/migrate"

	"carlord/ent/account"
	"carlord/ent/billing"
	"carlord/ent/booking"
	"carlord/ent/car"
	"carlord/ent/card"
	"carlord/ent/flaw"
	"carlord/ent/location"
	"carlord/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Billing is the client for interacting with the Billing builders.
	Billing *BillingClient
	// Booking is the client for interacting with the Booking builders.
	Booking *BookingClient
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// Flaw is the client for interacting with the Flaw builders.
	Flaw *FlawClient
	// Location is the client for interacting with the Location builders.
	Location *LocationClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Billing = NewBillingClient(c.config)
	c.Booking = NewBookingClient(c.config)
	c.Car = NewCarClient(c.config)
	c.Card = NewCardClient(c.config)
	c.Flaw = NewFlawClient(c.config)
	c.Location = NewLocationClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Account:  NewAccountClient(cfg),
		Billing:  NewBillingClient(cfg),
		Booking:  NewBookingClient(cfg),
		Car:      NewCarClient(cfg),
		Card:     NewCardClient(cfg),
		Flaw:     NewFlawClient(cfg),
		Location: NewLocationClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Account:  NewAccountClient(cfg),
		Billing:  NewBillingClient(cfg),
		Booking:  NewBookingClient(cfg),
		Car:      NewCarClient(cfg),
		Card:     NewCardClient(cfg),
		Flaw:     NewFlawClient(cfg),
		Location: NewLocationClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Billing.Use(hooks...)
	c.Booking.Use(hooks...)
	c.Car.Use(hooks...)
	c.Card.Use(hooks...)
	c.Flaw.Use(hooks...)
	c.Location.Use(hooks...)
	c.User.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id int) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id int) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id int) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id int) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Account.
func (c *AccountClient) QueryUser(a *Account) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, account.UserTable, account.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// BillingClient is a client for the Billing schema.
type BillingClient struct {
	config
}

// NewBillingClient returns a client for the Billing from the given config.
func NewBillingClient(c config) *BillingClient {
	return &BillingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `billing.Hooks(f(g(h())))`.
func (c *BillingClient) Use(hooks ...Hook) {
	c.hooks.Billing = append(c.hooks.Billing, hooks...)
}

// Create returns a builder for creating a Billing entity.
func (c *BillingClient) Create() *BillingCreate {
	mutation := newBillingMutation(c.config, OpCreate)
	return &BillingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Billing entities.
func (c *BillingClient) CreateBulk(builders ...*BillingCreate) *BillingCreateBulk {
	return &BillingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Billing.
func (c *BillingClient) Update() *BillingUpdate {
	mutation := newBillingMutation(c.config, OpUpdate)
	return &BillingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BillingClient) UpdateOne(b *Billing) *BillingUpdateOne {
	mutation := newBillingMutation(c.config, OpUpdateOne, withBilling(b))
	return &BillingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BillingClient) UpdateOneID(id int) *BillingUpdateOne {
	mutation := newBillingMutation(c.config, OpUpdateOne, withBillingID(id))
	return &BillingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Billing.
func (c *BillingClient) Delete() *BillingDelete {
	mutation := newBillingMutation(c.config, OpDelete)
	return &BillingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BillingClient) DeleteOne(b *Billing) *BillingDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BillingClient) DeleteOneID(id int) *BillingDeleteOne {
	builder := c.Delete().Where(billing.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BillingDeleteOne{builder}
}

// Query returns a query builder for Billing.
func (c *BillingClient) Query() *BillingQuery {
	return &BillingQuery{
		config: c.config,
	}
}

// Get returns a Billing entity by its id.
func (c *BillingClient) Get(ctx context.Context, id int) (*Billing, error) {
	return c.Query().Where(billing.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BillingClient) GetX(ctx context.Context, id int) *Billing {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BillingClient) Hooks() []Hook {
	return c.hooks.Billing
}

// BookingClient is a client for the Booking schema.
type BookingClient struct {
	config
}

// NewBookingClient returns a client for the Booking from the given config.
func NewBookingClient(c config) *BookingClient {
	return &BookingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `booking.Hooks(f(g(h())))`.
func (c *BookingClient) Use(hooks ...Hook) {
	c.hooks.Booking = append(c.hooks.Booking, hooks...)
}

// Create returns a builder for creating a Booking entity.
func (c *BookingClient) Create() *BookingCreate {
	mutation := newBookingMutation(c.config, OpCreate)
	return &BookingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Booking entities.
func (c *BookingClient) CreateBulk(builders ...*BookingCreate) *BookingCreateBulk {
	return &BookingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Booking.
func (c *BookingClient) Update() *BookingUpdate {
	mutation := newBookingMutation(c.config, OpUpdate)
	return &BookingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookingClient) UpdateOne(b *Booking) *BookingUpdateOne {
	mutation := newBookingMutation(c.config, OpUpdateOne, withBooking(b))
	return &BookingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookingClient) UpdateOneID(id int) *BookingUpdateOne {
	mutation := newBookingMutation(c.config, OpUpdateOne, withBookingID(id))
	return &BookingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Booking.
func (c *BookingClient) Delete() *BookingDelete {
	mutation := newBookingMutation(c.config, OpDelete)
	return &BookingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BookingClient) DeleteOne(b *Booking) *BookingDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BookingClient) DeleteOneID(id int) *BookingDeleteOne {
	builder := c.Delete().Where(booking.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookingDeleteOne{builder}
}

// Query returns a query builder for Booking.
func (c *BookingClient) Query() *BookingQuery {
	return &BookingQuery{
		config: c.config,
	}
}

// Get returns a Booking entity by its id.
func (c *BookingClient) Get(ctx context.Context, id int) (*Booking, error) {
	return c.Query().Where(booking.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookingClient) GetX(ctx context.Context, id int) *Booking {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BookingClient) Hooks() []Hook {
	return c.hooks.Booking
}

// CarClient is a client for the Car schema.
type CarClient struct {
	config
}

// NewCarClient returns a client for the Car from the given config.
func NewCarClient(c config) *CarClient {
	return &CarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `car.Hooks(f(g(h())))`.
func (c *CarClient) Use(hooks ...Hook) {
	c.hooks.Car = append(c.hooks.Car, hooks...)
}

// Create returns a builder for creating a Car entity.
func (c *CarClient) Create() *CarCreate {
	mutation := newCarMutation(c.config, OpCreate)
	return &CarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Car entities.
func (c *CarClient) CreateBulk(builders ...*CarCreate) *CarCreateBulk {
	return &CarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Car.
func (c *CarClient) Update() *CarUpdate {
	mutation := newCarMutation(c.config, OpUpdate)
	return &CarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CarClient) UpdateOne(ca *Car) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCar(ca))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CarClient) UpdateOneID(id int) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCarID(id))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Car.
func (c *CarClient) Delete() *CarDelete {
	mutation := newCarMutation(c.config, OpDelete)
	return &CarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CarClient) DeleteOne(ca *Car) *CarDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CarClient) DeleteOneID(id int) *CarDeleteOne {
	builder := c.Delete().Where(car.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CarDeleteOne{builder}
}

// Query returns a query builder for Car.
func (c *CarClient) Query() *CarQuery {
	return &CarQuery{
		config: c.config,
	}
}

// Get returns a Car entity by its id.
func (c *CarClient) Get(ctx context.Context, id int) (*Car, error) {
	return c.Query().Where(car.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CarClient) GetX(ctx context.Context, id int) *Car {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `card.Hooks(f(g(h())))`.
func (c *CardClient) Use(hooks ...Hook) {
	c.hooks.Card = append(c.hooks.Card, hooks...)
}

// Create returns a builder for creating a Card entity.
func (c *CardClient) Create() *CardCreate {
	mutation := newCardMutation(c.config, OpCreate)
	return &CardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Card entities.
func (c *CardClient) CreateBulk(builders ...*CardCreate) *CardCreateBulk {
	return &CardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	mutation := newCardMutation(c.config, OpUpdate)
	return &CardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCard(ca))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id int) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCardID(id))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	mutation := newCardMutation(c.config, OpDelete)
	return &CardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CardClient) DeleteOneID(id int) *CardDeleteOne {
	builder := c.Delete().Where(card.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardDeleteOne{builder}
}

// Query returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{
		config: c.config,
	}
}

// Get returns a Card entity by its id.
func (c *CardClient) Get(ctx context.Context, id int) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id int) *Card {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Card.
func (c *CardClient) QueryOwner(ca *Card) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, card.OwnerTable, card.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CardClient) Hooks() []Hook {
	return c.hooks.Card
}

// FlawClient is a client for the Flaw schema.
type FlawClient struct {
	config
}

// NewFlawClient returns a client for the Flaw from the given config.
func NewFlawClient(c config) *FlawClient {
	return &FlawClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flaw.Hooks(f(g(h())))`.
func (c *FlawClient) Use(hooks ...Hook) {
	c.hooks.Flaw = append(c.hooks.Flaw, hooks...)
}

// Create returns a builder for creating a Flaw entity.
func (c *FlawClient) Create() *FlawCreate {
	mutation := newFlawMutation(c.config, OpCreate)
	return &FlawCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Flaw entities.
func (c *FlawClient) CreateBulk(builders ...*FlawCreate) *FlawCreateBulk {
	return &FlawCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Flaw.
func (c *FlawClient) Update() *FlawUpdate {
	mutation := newFlawMutation(c.config, OpUpdate)
	return &FlawUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlawClient) UpdateOne(f *Flaw) *FlawUpdateOne {
	mutation := newFlawMutation(c.config, OpUpdateOne, withFlaw(f))
	return &FlawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlawClient) UpdateOneID(id int) *FlawUpdateOne {
	mutation := newFlawMutation(c.config, OpUpdateOne, withFlawID(id))
	return &FlawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Flaw.
func (c *FlawClient) Delete() *FlawDelete {
	mutation := newFlawMutation(c.config, OpDelete)
	return &FlawDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FlawClient) DeleteOne(f *Flaw) *FlawDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FlawClient) DeleteOneID(id int) *FlawDeleteOne {
	builder := c.Delete().Where(flaw.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlawDeleteOne{builder}
}

// Query returns a query builder for Flaw.
func (c *FlawClient) Query() *FlawQuery {
	return &FlawQuery{
		config: c.config,
	}
}

// Get returns a Flaw entity by its id.
func (c *FlawClient) Get(ctx context.Context, id int) (*Flaw, error) {
	return c.Query().Where(flaw.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlawClient) GetX(ctx context.Context, id int) *Flaw {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FlawClient) Hooks() []Hook {
	return c.hooks.Flaw
}

// LocationClient is a client for the Location schema.
type LocationClient struct {
	config
}

// NewLocationClient returns a client for the Location from the given config.
func NewLocationClient(c config) *LocationClient {
	return &LocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `location.Hooks(f(g(h())))`.
func (c *LocationClient) Use(hooks ...Hook) {
	c.hooks.Location = append(c.hooks.Location, hooks...)
}

// Create returns a builder for creating a Location entity.
func (c *LocationClient) Create() *LocationCreate {
	mutation := newLocationMutation(c.config, OpCreate)
	return &LocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Location entities.
func (c *LocationClient) CreateBulk(builders ...*LocationCreate) *LocationCreateBulk {
	return &LocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Location.
func (c *LocationClient) Update() *LocationUpdate {
	mutation := newLocationMutation(c.config, OpUpdate)
	return &LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LocationClient) UpdateOne(l *Location) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocation(l))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LocationClient) UpdateOneID(id int) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocationID(id))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Location.
func (c *LocationClient) Delete() *LocationDelete {
	mutation := newLocationMutation(c.config, OpDelete)
	return &LocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LocationClient) DeleteOne(l *Location) *LocationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LocationClient) DeleteOneID(id int) *LocationDeleteOne {
	builder := c.Delete().Where(location.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LocationDeleteOne{builder}
}

// Query returns a query builder for Location.
func (c *LocationClient) Query() *LocationQuery {
	return &LocationQuery{
		config: c.config,
	}
}

// Get returns a Location entity by its id.
func (c *LocationClient) Get(ctx context.Context, id int) (*Location, error) {
	return c.Query().Where(location.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LocationClient) GetX(ctx context.Context, id int) *Location {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LocationClient) Hooks() []Hook {
	return c.hooks.Location
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCard queries the card edge of a User.
func (c *UserClient) QueryCard(u *User) *CardQuery {
	query := &CardQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CardTable, user.CardColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNoteFlaws queries the note_flaws edge of a User.
func (c *UserClient) QueryNoteFlaws(u *User) *FlawQuery {
	query := &FlawQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(flaw.Table, flaw.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.NoteFlawsTable, user.NoteFlawsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccount queries the account edge of a User.
func (c *UserClient) QueryAccount(u *User) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.AccountTable, user.AccountPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
