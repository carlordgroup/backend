// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/account"
	"carlord/ent/booking"
	"carlord/ent/card"
	"carlord/ent/flaw"
	"carlord/ent/predicate"
	"carlord/ent/user"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.User
	withCard      *CardQuery
	withNoteFlaws *FlawQuery
	withAccount   *AccountQuery
	withBooking   *BookingQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.unique = &unique
	return uq
}

// Order adds an order step to the query.
func (uq *UserQuery) Order(o ...OrderFunc) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryCard chains the current query on the "card" edge.
func (uq *UserQuery) QueryCard() *CardQuery {
	query := &CardQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CardTable, user.CardColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNoteFlaws chains the current query on the "note_flaws" edge.
func (uq *UserQuery) QueryNoteFlaws() *FlawQuery {
	query := &FlawQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(flaw.Table, flaw.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.NoteFlawsTable, user.NoteFlawsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccount chains the current query on the "account" edge.
func (uq *UserQuery) QueryAccount() *AccountQuery {
	query := &AccountQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, user.AccountTable, user.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBooking chains the current query on the "booking" edge.
func (uq *UserQuery) QueryBooking() *BookingQuery {
	query := &BookingQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.BookingTable, user.BookingColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:        uq.config,
		limit:         uq.limit,
		offset:        uq.offset,
		order:         append([]OrderFunc{}, uq.order...),
		predicates:    append([]predicate.User{}, uq.predicates...),
		withCard:      uq.withCard.Clone(),
		withNoteFlaws: uq.withNoteFlaws.Clone(),
		withAccount:   uq.withAccount.Clone(),
		withBooking:   uq.withBooking.Clone(),
		// clone intermediate query.
		sql:    uq.sql.Clone(),
		path:   uq.path,
		unique: uq.unique,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCard(opts ...func(*CardQuery)) *UserQuery {
	query := &CardQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withCard = query
	return uq
}

// WithNoteFlaws tells the query-builder to eager-load the nodes that are connected to
// the "note_flaws" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNoteFlaws(opts ...func(*FlawQuery)) *UserQuery {
	query := &FlawQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withNoteFlaws = query
	return uq
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithAccount(opts ...func(*AccountQuery)) *UserQuery {
	query := &AccountQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withAccount = query
	return uq
}

// WithBooking tells the query-builder to eager-load the nodes that are connected to
// the "booking" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithBooking(opts ...func(*BookingQuery)) *UserQuery {
	query := &BookingQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withBooking = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"first_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldFirstName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	grbuild := &UserGroupBy{config: uq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(ctx), nil
	}
	grbuild.label = user.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"first_name,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldFirstName).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.fields = append(uq.fields, fields...)
	selbuild := &UserSelect{UserQuery: uq}
	selbuild.label = user.Label
	selbuild.flds, selbuild.scan = &uq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uq.fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		withFKs     = uq.withFKs
		_spec       = uq.querySpec()
		loadedTypes = [4]bool{
			uq.withCard != nil,
			uq.withNoteFlaws != nil,
			uq.withAccount != nil,
			uq.withBooking != nil,
		}
	)
	if uq.withAccount != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, user.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withCard; query != nil {
		if err := uq.loadCard(ctx, query, nodes,
			func(n *User) { n.Edges.Card = []*Card{} },
			func(n *User, e *Card) { n.Edges.Card = append(n.Edges.Card, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withNoteFlaws; query != nil {
		if err := uq.loadNoteFlaws(ctx, query, nodes,
			func(n *User) { n.Edges.NoteFlaws = []*Flaw{} },
			func(n *User, e *Flaw) { n.Edges.NoteFlaws = append(n.Edges.NoteFlaws, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withAccount; query != nil {
		if err := uq.loadAccount(ctx, query, nodes, nil,
			func(n *User, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withBooking; query != nil {
		if err := uq.loadBooking(ctx, query, nodes,
			func(n *User) { n.Edges.Booking = []*Booking{} },
			func(n *User, e *Booking) { n.Edges.Booking = append(n.Edges.Booking, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadCard(ctx context.Context, query *CardQuery, nodes []*User, init func(*User), assign func(*User, *Card)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Card(func(s *sql.Selector) {
		s.Where(sql.InValues(user.CardColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_card
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_card" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_card" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadNoteFlaws(ctx context.Context, query *FlawQuery, nodes []*User, init func(*User), assign func(*User, *Flaw)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Flaw(func(s *sql.Selector) {
		s.Where(sql.InValues(user.NoteFlawsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_note_flaws
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_note_flaws" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_note_flaws" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*User, init func(*User), assign func(*User, *Account)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*User)
	for i := range nodes {
		if nodes[i].account_user == nil {
			continue
		}
		fk := *nodes[i].account_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UserQuery) loadBooking(ctx context.Context, query *BookingQuery, nodes []*User, init func(*User), assign func(*User, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(user.BookingColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.booking_user
		if fk == nil {
			return fmt.Errorf(`foreign-key "booking_user" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "booking_user" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.fields
	if len(uq.fields) > 0 {
		_spec.Unique = uq.unique != nil && *uq.unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if unique := uq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.unique != nil && *uq.unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ugb.path(ctx)
	if err != nil {
		return err
	}
	ugb.sql = query
	return ugb.sqlScan(ctx, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ugb.fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UserGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql.Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
		for _, f := range ugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ugb.fields...)...)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	us.sql = us.UserQuery.sqlQuery(ctx)
	return us.sqlScan(ctx, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(us.sql))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		us.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		us.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := us.sql.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
