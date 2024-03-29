// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/echo-stream/database/ent/slackmessage"
	"github.com/echo-stream/database/ent/slackmessageattachment"
	"github.com/google/uuid"
)

// SlackMessageCreate is the builder for creating a SlackMessage entity.
type SlackMessageCreate struct {
	config
	mutation *SlackMessageMutation
	hooks    []Hook
}

// SetApplicationId sets the "ApplicationId" field.
func (smc *SlackMessageCreate) SetApplicationId(u uuid.UUID) *SlackMessageCreate {
	smc.mutation.SetApplicationId(u)
	return smc
}

// SetUserName sets the "UserName" field.
func (smc *SlackMessageCreate) SetUserName(s string) *SlackMessageCreate {
	smc.mutation.SetUserName(s)
	return smc
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableUserName(s *string) *SlackMessageCreate {
	if s != nil {
		smc.SetUserName(*s)
	}
	return smc
}

// SetIconEmoji sets the "IconEmoji" field.
func (smc *SlackMessageCreate) SetIconEmoji(s string) *SlackMessageCreate {
	smc.mutation.SetIconEmoji(s)
	return smc
}

// SetNillableIconEmoji sets the "IconEmoji" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableIconEmoji(s *string) *SlackMessageCreate {
	if s != nil {
		smc.SetIconEmoji(*s)
	}
	return smc
}

// SetChannel sets the "Channel" field.
func (smc *SlackMessageCreate) SetChannel(s string) *SlackMessageCreate {
	smc.mutation.SetChannel(s)
	return smc
}

// SetText sets the "Text" field.
func (smc *SlackMessageCreate) SetText(s string) *SlackMessageCreate {
	smc.mutation.SetText(s)
	return smc
}

// SetNillableText sets the "Text" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableText(s *string) *SlackMessageCreate {
	if s != nil {
		smc.SetText(*s)
	}
	return smc
}

// SetThreadTs sets the "ThreadTs" field.
func (smc *SlackMessageCreate) SetThreadTs(s string) *SlackMessageCreate {
	smc.mutation.SetThreadTs(s)
	return smc
}

// SetNillableThreadTs sets the "ThreadTs" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableThreadTs(s *string) *SlackMessageCreate {
	if s != nil {
		smc.SetThreadTs(*s)
	}
	return smc
}

// SetState sets the "State" field.
func (smc *SlackMessageCreate) SetState(s string) *SlackMessageCreate {
	smc.mutation.SetState(s)
	return smc
}

// SetNillableState sets the "State" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableState(s *string) *SlackMessageCreate {
	if s != nil {
		smc.SetState(*s)
	}
	return smc
}

// SetRetryCount sets the "RetryCount" field.
func (smc *SlackMessageCreate) SetRetryCount(i int16) *SlackMessageCreate {
	smc.mutation.SetRetryCount(i)
	return smc
}

// SetNillableRetryCount sets the "RetryCount" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableRetryCount(i *int16) *SlackMessageCreate {
	if i != nil {
		smc.SetRetryCount(*i)
	}
	return smc
}

// SetCreatedAt sets the "CreatedAt" field.
func (smc *SlackMessageCreate) SetCreatedAt(t time.Time) *SlackMessageCreate {
	smc.mutation.SetCreatedAt(t)
	return smc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableCreatedAt(t *time.Time) *SlackMessageCreate {
	if t != nil {
		smc.SetCreatedAt(*t)
	}
	return smc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (smc *SlackMessageCreate) SetUpdatedAt(t time.Time) *SlackMessageCreate {
	smc.mutation.SetUpdatedAt(t)
	return smc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (smc *SlackMessageCreate) SetNillableUpdatedAt(t *time.Time) *SlackMessageCreate {
	if t != nil {
		smc.SetUpdatedAt(*t)
	}
	return smc
}

// AddSlackMessageAttachmentIDs adds the "slack_message_attachments" edge to the SlackMessageAttachment entity by IDs.
func (smc *SlackMessageCreate) AddSlackMessageAttachmentIDs(ids ...int) *SlackMessageCreate {
	smc.mutation.AddSlackMessageAttachmentIDs(ids...)
	return smc
}

// AddSlackMessageAttachments adds the "slack_message_attachments" edges to the SlackMessageAttachment entity.
func (smc *SlackMessageCreate) AddSlackMessageAttachments(s ...*SlackMessageAttachment) *SlackMessageCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return smc.AddSlackMessageAttachmentIDs(ids...)
}

// Mutation returns the SlackMessageMutation object of the builder.
func (smc *SlackMessageCreate) Mutation() *SlackMessageMutation {
	return smc.mutation
}

// Save creates the SlackMessage in the database.
func (smc *SlackMessageCreate) Save(ctx context.Context) (*SlackMessage, error) {
	smc.defaults()
	return withHooks(ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SlackMessageCreate) SaveX(ctx context.Context) *SlackMessage {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SlackMessageCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SlackMessageCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smc *SlackMessageCreate) defaults() {
	if _, ok := smc.mutation.State(); !ok {
		v := slackmessage.DefaultState
		smc.mutation.SetState(v)
	}
	if _, ok := smc.mutation.RetryCount(); !ok {
		v := slackmessage.DefaultRetryCount
		smc.mutation.SetRetryCount(v)
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		v := slackmessage.DefaultCreatedAt()
		smc.mutation.SetCreatedAt(v)
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		v := slackmessage.DefaultUpdatedAt()
		smc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SlackMessageCreate) check() error {
	if _, ok := smc.mutation.ApplicationId(); !ok {
		return &ValidationError{Name: "ApplicationId", err: errors.New(`ent: missing required field "SlackMessage.ApplicationId"`)}
	}
	if _, ok := smc.mutation.Channel(); !ok {
		return &ValidationError{Name: "Channel", err: errors.New(`ent: missing required field "SlackMessage.Channel"`)}
	}
	if v, ok := smc.mutation.Channel(); ok {
		if err := slackmessage.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "Channel", err: fmt.Errorf(`ent: validator failed for field "SlackMessage.Channel": %w`, err)}
		}
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "SlackMessage.CreatedAt"`)}
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "SlackMessage.UpdatedAt"`)}
	}
	return nil
}

func (smc *SlackMessageCreate) sqlSave(ctx context.Context) (*SlackMessage, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *SlackMessageCreate) createSpec() (*SlackMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &SlackMessage{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(slackmessage.Table, sqlgraph.NewFieldSpec(slackmessage.FieldID, field.TypeInt))
	)
	if value, ok := smc.mutation.ApplicationId(); ok {
		_spec.SetField(slackmessage.FieldApplicationId, field.TypeUUID, value)
		_node.ApplicationId = value
	}
	if value, ok := smc.mutation.UserName(); ok {
		_spec.SetField(slackmessage.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := smc.mutation.IconEmoji(); ok {
		_spec.SetField(slackmessage.FieldIconEmoji, field.TypeString, value)
		_node.IconEmoji = value
	}
	if value, ok := smc.mutation.Channel(); ok {
		_spec.SetField(slackmessage.FieldChannel, field.TypeString, value)
		_node.Channel = value
	}
	if value, ok := smc.mutation.Text(); ok {
		_spec.SetField(slackmessage.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := smc.mutation.ThreadTs(); ok {
		_spec.SetField(slackmessage.FieldThreadTs, field.TypeString, value)
		_node.ThreadTs = value
	}
	if value, ok := smc.mutation.State(); ok {
		_spec.SetField(slackmessage.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := smc.mutation.RetryCount(); ok {
		_spec.SetField(slackmessage.FieldRetryCount, field.TypeInt16, value)
		_node.RetryCount = value
	}
	if value, ok := smc.mutation.CreatedAt(); ok {
		_spec.SetField(slackmessage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := smc.mutation.UpdatedAt(); ok {
		_spec.SetField(slackmessage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := smc.mutation.SlackMessageAttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   slackmessage.SlackMessageAttachmentsTable,
			Columns: []string{slackmessage.SlackMessageAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(slackmessageattachment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SlackMessageCreateBulk is the builder for creating many SlackMessage entities in bulk.
type SlackMessageCreateBulk struct {
	config
	err      error
	builders []*SlackMessageCreate
}

// Save creates the SlackMessage entities in the database.
func (smcb *SlackMessageCreateBulk) Save(ctx context.Context) ([]*SlackMessage, error) {
	if smcb.err != nil {
		return nil, smcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SlackMessage, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SlackMessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SlackMessageCreateBulk) SaveX(ctx context.Context) []*SlackMessage {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SlackMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SlackMessageCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
