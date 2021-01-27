// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/degenerat3/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/core/ent/host"
	"github.com/degenerat3/meteor/core/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// BotUpdate is the builder for updating Bot entities.
type BotUpdate struct {
	config
	hooks    []Hook
	mutation *BotMutation
}

// Where adds a new predicate for the BotUpdate builder.
func (bu *BotUpdate) Where(ps ...predicate.Bot) *BotUpdate {
	bu.mutation.predicates = append(bu.mutation.predicates, ps...)
	return bu
}

// SetUUID sets the "uuid" field.
func (bu *BotUpdate) SetUUID(s string) *BotUpdate {
	bu.mutation.SetUUID(s)
	return bu
}

// SetInterval sets the "interval" field.
func (bu *BotUpdate) SetInterval(i int) *BotUpdate {
	bu.mutation.ResetInterval()
	bu.mutation.SetInterval(i)
	return bu
}

// AddInterval adds i to the "interval" field.
func (bu *BotUpdate) AddInterval(i int) *BotUpdate {
	bu.mutation.AddInterval(i)
	return bu
}

// SetDelta sets the "delta" field.
func (bu *BotUpdate) SetDelta(i int) *BotUpdate {
	bu.mutation.ResetDelta()
	bu.mutation.SetDelta(i)
	return bu
}

// AddDelta adds i to the "delta" field.
func (bu *BotUpdate) AddDelta(i int) *BotUpdate {
	bu.mutation.AddDelta(i)
	return bu
}

// SetLastSeen sets the "lastSeen" field.
func (bu *BotUpdate) SetLastSeen(i int) *BotUpdate {
	bu.mutation.ResetLastSeen()
	bu.mutation.SetLastSeen(i)
	return bu
}

// SetNillableLastSeen sets the "lastSeen" field if the given value is not nil.
func (bu *BotUpdate) SetNillableLastSeen(i *int) *BotUpdate {
	if i != nil {
		bu.SetLastSeen(*i)
	}
	return bu
}

// AddLastSeen adds i to the "lastSeen" field.
func (bu *BotUpdate) AddLastSeen(i int) *BotUpdate {
	bu.mutation.AddLastSeen(i)
	return bu
}

// SetInfectingID sets the "infecting" edge to the Host entity by ID.
func (bu *BotUpdate) SetInfectingID(id int) *BotUpdate {
	bu.mutation.SetInfectingID(id)
	return bu
}

// SetNillableInfectingID sets the "infecting" edge to the Host entity by ID if the given value is not nil.
func (bu *BotUpdate) SetNillableInfectingID(id *int) *BotUpdate {
	if id != nil {
		bu = bu.SetInfectingID(*id)
	}
	return bu
}

// SetInfecting sets the "infecting" edge to the Host entity.
func (bu *BotUpdate) SetInfecting(h *Host) *BotUpdate {
	return bu.SetInfectingID(h.ID)
}

// Mutation returns the BotMutation object of the builder.
func (bu *BotUpdate) Mutation() *BotMutation {
	return bu.mutation
}

// ClearInfecting clears the "infecting" edge to the Host entity.
func (bu *BotUpdate) ClearInfecting() *BotUpdate {
	bu.mutation.ClearInfecting()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BotUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BotUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BotUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BotUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bot.Table,
			Columns: bot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bot.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bot.FieldUUID,
		})
	}
	if value, ok := bu.mutation.Interval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldInterval,
		})
	}
	if value, ok := bu.mutation.AddedInterval(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldInterval,
		})
	}
	if value, ok := bu.mutation.Delta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldDelta,
		})
	}
	if value, ok := bu.mutation.AddedDelta(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldDelta,
		})
	}
	if value, ok := bu.mutation.LastSeen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldLastSeen,
		})
	}
	if value, ok := bu.mutation.AddedLastSeen(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldLastSeen,
		})
	}
	if bu.mutation.InfectingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bot.InfectingTable,
			Columns: []string{bot.InfectingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.InfectingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bot.InfectingTable,
			Columns: []string{bot.InfectingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bot.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BotUpdateOne is the builder for updating a single Bot entity.
type BotUpdateOne struct {
	config
	hooks    []Hook
	mutation *BotMutation
}

// SetUUID sets the "uuid" field.
func (buo *BotUpdateOne) SetUUID(s string) *BotUpdateOne {
	buo.mutation.SetUUID(s)
	return buo
}

// SetInterval sets the "interval" field.
func (buo *BotUpdateOne) SetInterval(i int) *BotUpdateOne {
	buo.mutation.ResetInterval()
	buo.mutation.SetInterval(i)
	return buo
}

// AddInterval adds i to the "interval" field.
func (buo *BotUpdateOne) AddInterval(i int) *BotUpdateOne {
	buo.mutation.AddInterval(i)
	return buo
}

// SetDelta sets the "delta" field.
func (buo *BotUpdateOne) SetDelta(i int) *BotUpdateOne {
	buo.mutation.ResetDelta()
	buo.mutation.SetDelta(i)
	return buo
}

// AddDelta adds i to the "delta" field.
func (buo *BotUpdateOne) AddDelta(i int) *BotUpdateOne {
	buo.mutation.AddDelta(i)
	return buo
}

// SetLastSeen sets the "lastSeen" field.
func (buo *BotUpdateOne) SetLastSeen(i int) *BotUpdateOne {
	buo.mutation.ResetLastSeen()
	buo.mutation.SetLastSeen(i)
	return buo
}

// SetNillableLastSeen sets the "lastSeen" field if the given value is not nil.
func (buo *BotUpdateOne) SetNillableLastSeen(i *int) *BotUpdateOne {
	if i != nil {
		buo.SetLastSeen(*i)
	}
	return buo
}

// AddLastSeen adds i to the "lastSeen" field.
func (buo *BotUpdateOne) AddLastSeen(i int) *BotUpdateOne {
	buo.mutation.AddLastSeen(i)
	return buo
}

// SetInfectingID sets the "infecting" edge to the Host entity by ID.
func (buo *BotUpdateOne) SetInfectingID(id int) *BotUpdateOne {
	buo.mutation.SetInfectingID(id)
	return buo
}

// SetNillableInfectingID sets the "infecting" edge to the Host entity by ID if the given value is not nil.
func (buo *BotUpdateOne) SetNillableInfectingID(id *int) *BotUpdateOne {
	if id != nil {
		buo = buo.SetInfectingID(*id)
	}
	return buo
}

// SetInfecting sets the "infecting" edge to the Host entity.
func (buo *BotUpdateOne) SetInfecting(h *Host) *BotUpdateOne {
	return buo.SetInfectingID(h.ID)
}

// Mutation returns the BotMutation object of the builder.
func (buo *BotUpdateOne) Mutation() *BotMutation {
	return buo.mutation
}

// ClearInfecting clears the "infecting" edge to the Host entity.
func (buo *BotUpdateOne) ClearInfecting() *BotUpdateOne {
	buo.mutation.ClearInfecting()
	return buo
}

// Save executes the query and returns the updated Bot entity.
func (buo *BotUpdateOne) Save(ctx context.Context) (*Bot, error) {
	var (
		err  error
		node *Bot
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BotUpdateOne) SaveX(ctx context.Context) *Bot {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BotUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BotUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BotUpdateOne) sqlSave(ctx context.Context) (_node *Bot, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bot.Table,
			Columns: bot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bot.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Bot.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bot.FieldUUID,
		})
	}
	if value, ok := buo.mutation.Interval(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldInterval,
		})
	}
	if value, ok := buo.mutation.AddedInterval(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldInterval,
		})
	}
	if value, ok := buo.mutation.Delta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldDelta,
		})
	}
	if value, ok := buo.mutation.AddedDelta(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldDelta,
		})
	}
	if value, ok := buo.mutation.LastSeen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldLastSeen,
		})
	}
	if value, ok := buo.mutation.AddedLastSeen(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldLastSeen,
		})
	}
	if buo.mutation.InfectingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bot.InfectingTable,
			Columns: []string{bot.InfectingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.InfectingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bot.InfectingTable,
			Columns: []string{bot.InfectingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bot{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bot.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
