// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/degenerat3/meteor/core/ent/action"
	"github.com/degenerat3/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/core/ent/group"
	"github.com/degenerat3/meteor/core/ent/host"
	"github.com/degenerat3/meteor/core/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostQuery is the builder for querying Host entities.
type HostQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Host
	// eager-loading edges.
	withBots    *BotQuery
	withActions *ActionQuery
	withMember  *GroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HostQuery builder.
func (hq *HostQuery) Where(ps ...predicate.Host) *HostQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HostQuery) Limit(limit int) *HostQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HostQuery) Offset(offset int) *HostQuery {
	hq.offset = &offset
	return hq
}

// Order adds an order step to the query.
func (hq *HostQuery) Order(o ...OrderFunc) *HostQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryBots chains the current query on the "bots" edge.
func (hq *HostQuery) QueryBots() *BotQuery {
	query := &BotQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(bot.Table, bot.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.BotsTable, host.BotsColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryActions chains the current query on the "actions" edge.
func (hq *HostQuery) QueryActions() *ActionQuery {
	query := &ActionQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(action.Table, action.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.ActionsTable, host.ActionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMember chains the current query on the "member" edge.
func (hq *HostQuery) QueryMember() *GroupQuery {
	query := &GroupQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, host.MemberTable, host.MemberPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Host entity from the query.
// Returns a *NotFoundError when no Host was found.
func (hq *HostQuery) First(ctx context.Context) (*Host, error) {
	nodes, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{host.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HostQuery) FirstX(ctx context.Context) *Host {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Host ID from the query.
// Returns a *NotFoundError when no Host ID was found.
func (hq *HostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{host.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HostQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Host entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Host entity is not found.
// Returns a *NotFoundError when no Host entities are found.
func (hq *HostQuery) Only(ctx context.Context) (*Host, error) {
	nodes, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{host.Label}
	default:
		return nil, &NotSingularError{host.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HostQuery) OnlyX(ctx context.Context) *Host {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Host ID in the query.
// Returns a *NotSingularError when exactly one Host ID is not found.
// Returns a *NotFoundError when no entities are found.
func (hq *HostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = &NotSingularError{host.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HostQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Hosts.
func (hq *HostQuery) All(ctx context.Context) ([]*Host, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HostQuery) AllX(ctx context.Context) []*Host {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Host IDs.
func (hq *HostQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hq.Select(host.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HostQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HostQuery) Count(ctx context.Context) (int, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HostQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HostQuery) Exist(ctx context.Context) (bool, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HostQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HostQuery) Clone() *HostQuery {
	if hq == nil {
		return nil
	}
	return &HostQuery{
		config:      hq.config,
		limit:       hq.limit,
		offset:      hq.offset,
		order:       append([]OrderFunc{}, hq.order...),
		predicates:  append([]predicate.Host{}, hq.predicates...),
		withBots:    hq.withBots.Clone(),
		withActions: hq.withActions.Clone(),
		withMember:  hq.withMember.Clone(),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// WithBots tells the query-builder to eager-load the nodes that are connected to
// the "bots" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithBots(opts ...func(*BotQuery)) *HostQuery {
	query := &BotQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withBots = query
	return hq
}

// WithActions tells the query-builder to eager-load the nodes that are connected to
// the "actions" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithActions(opts ...func(*ActionQuery)) *HostQuery {
	query := &ActionQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withActions = query
	return hq
}

// WithMember tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithMember(opts ...func(*GroupQuery)) *HostQuery {
	query := &GroupQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withMember = query
	return hq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hostname string `json:"hostname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Host.Query().
//		GroupBy(host.FieldHostname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hq *HostQuery) GroupBy(field string, fields ...string) *HostGroupBy {
	group := &HostGroupBy{config: hq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Hostname string `json:"hostname,omitempty"`
//	}
//
//	client.Host.Query().
//		Select(host.FieldHostname).
//		Scan(ctx, &v)
//
func (hq *HostQuery) Select(field string, fields ...string) *HostSelect {
	hq.fields = append([]string{field}, fields...)
	return &HostSelect{HostQuery: hq}
}

func (hq *HostQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hq.fields {
		if !host.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HostQuery) sqlAll(ctx context.Context) ([]*Host, error) {
	var (
		nodes       = []*Host{}
		_spec       = hq.querySpec()
		loadedTypes = [3]bool{
			hq.withBots != nil,
			hq.withActions != nil,
			hq.withMember != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Host{config: hq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := hq.withBots; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Host)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Bots = []*Bot{}
		}
		query.withFKs = true
		query.Where(predicate.Bot(func(s *sql.Selector) {
			s.Where(sql.InValues(host.BotsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.host_bots
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "host_bots" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "host_bots" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Bots = append(node.Edges.Bots, n)
		}
	}

	if query := hq.withActions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Host)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Actions = []*Action{}
		}
		query.withFKs = true
		query.Where(predicate.Action(func(s *sql.Selector) {
			s.Where(sql.InValues(host.ActionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.host_actions
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "host_actions" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "host_actions" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Actions = append(node.Edges.Actions, n)
		}
	}

	if query := hq.withMember; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Host, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Member = []*Group{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Host)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   host.MemberTable,
				Columns: host.MemberPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(host.MemberPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, hq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "member": %v`, err)
		}
		query.Where(group.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "member" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Member = append(nodes[i].Edges.Member, n)
			}
		}
	}

	return nodes, nil
}

func (hq *HostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HostQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (hq *HostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   host.Table,
			Columns: host.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: host.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if fields := hq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, host.FieldID)
		for i := range fields {
			if fields[i] != host.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, host.ValidColumn)
			}
		}
	}
	return _spec
}

func (hq *HostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(host.Table)
	selector := builder.Select(t1.Columns(host.Columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(host.Columns...)...)
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector, host.ValidColumn)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HostGroupBy is the group-by builder for Host entities.
type HostGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HostGroupBy) Aggregate(fns ...AggregateFunc) *HostGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hgb *HostGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hgb.path(ctx)
	if err != nil {
		return err
	}
	hgb.sql = query
	return hgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hgb *HostGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HostGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hgb *HostGroupBy) StringsX(ctx context.Context) []string {
	v, err := hgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hgb *HostGroupBy) StringX(ctx context.Context) string {
	v, err := hgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HostGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hgb *HostGroupBy) IntsX(ctx context.Context) []int {
	v, err := hgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hgb *HostGroupBy) IntX(ctx context.Context) int {
	v, err := hgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HostGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hgb *HostGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hgb *HostGroupBy) Float64X(ctx context.Context) float64 {
	v, err := hgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HostGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hgb *HostGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HostGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hgb *HostGroupBy) BoolX(ctx context.Context) bool {
	v, err := hgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hgb *HostGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hgb.fields {
		if !host.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HostGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql
	columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
	columns = append(columns, hgb.fields...)
	for _, fn := range hgb.fns {
		columns = append(columns, fn(selector, host.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(hgb.fields...)
}

// HostSelect is the builder for selecting fields of Host entities.
type HostSelect struct {
	*HostQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HostSelect) Scan(ctx context.Context, v interface{}) error {
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	hs.sql = hs.HostQuery.sqlQuery(ctx)
	return hs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hs *HostSelect) ScanX(ctx context.Context, v interface{}) {
	if err := hs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Strings(ctx context.Context) ([]string, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HostSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hs *HostSelect) StringsX(ctx context.Context) []string {
	v, err := hs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hs *HostSelect) StringX(ctx context.Context) string {
	v, err := hs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Ints(ctx context.Context) ([]int, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HostSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hs *HostSelect) IntsX(ctx context.Context) []int {
	v, err := hs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hs *HostSelect) IntX(ctx context.Context) int {
	v, err := hs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HostSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hs *HostSelect) Float64sX(ctx context.Context) []float64 {
	v, err := hs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hs *HostSelect) Float64X(ctx context.Context) float64 {
	v, err := hs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HostSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hs *HostSelect) BoolsX(ctx context.Context) []bool {
	v, err := hs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (hs *HostSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = fmt.Errorf("ent: HostSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hs *HostSelect) BoolX(ctx context.Context) bool {
	v, err := hs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hs *HostSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sqlQuery().Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hs *HostSelect) sqlQuery() sql.Querier {
	selector := hs.sql
	selector.Select(selector.Columns(hs.fields...)...)
	return selector
}
