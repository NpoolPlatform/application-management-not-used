// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroup"
	"github.com/google/uuid"
)

// ApplicationGroupCreate is the builder for creating a ApplicationGroup entity.
type ApplicationGroupCreate struct {
	config
	mutation *ApplicationGroupMutation
	hooks    []Hook
}

// SetAppID sets the "app_id" field.
func (agc *ApplicationGroupCreate) SetAppID(s string) *ApplicationGroupCreate {
	agc.mutation.SetAppID(s)
	return agc
}

// SetGroupName sets the "group_name" field.
func (agc *ApplicationGroupCreate) SetGroupName(s string) *ApplicationGroupCreate {
	agc.mutation.SetGroupName(s)
	return agc
}

// SetGroupLogo sets the "group_logo" field.
func (agc *ApplicationGroupCreate) SetGroupLogo(s string) *ApplicationGroupCreate {
	agc.mutation.SetGroupLogo(s)
	return agc
}

// SetNillableGroupLogo sets the "group_logo" field if the given value is not nil.
func (agc *ApplicationGroupCreate) SetNillableGroupLogo(s *string) *ApplicationGroupCreate {
	if s != nil {
		agc.SetGroupLogo(*s)
	}
	return agc
}

// SetGroupOwner sets the "group_owner" field.
func (agc *ApplicationGroupCreate) SetGroupOwner(u uuid.UUID) *ApplicationGroupCreate {
	agc.mutation.SetGroupOwner(u)
	return agc
}

// SetAnnotation sets the "annotation" field.
func (agc *ApplicationGroupCreate) SetAnnotation(s string) *ApplicationGroupCreate {
	agc.mutation.SetAnnotation(s)
	return agc
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (agc *ApplicationGroupCreate) SetNillableAnnotation(s *string) *ApplicationGroupCreate {
	if s != nil {
		agc.SetAnnotation(*s)
	}
	return agc
}

// SetCreateAt sets the "create_at" field.
func (agc *ApplicationGroupCreate) SetCreateAt(i int64) *ApplicationGroupCreate {
	agc.mutation.SetCreateAt(i)
	return agc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (agc *ApplicationGroupCreate) SetNillableCreateAt(i *int64) *ApplicationGroupCreate {
	if i != nil {
		agc.SetCreateAt(*i)
	}
	return agc
}

// SetUpdateAt sets the "update_at" field.
func (agc *ApplicationGroupCreate) SetUpdateAt(i int64) *ApplicationGroupCreate {
	agc.mutation.SetUpdateAt(i)
	return agc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (agc *ApplicationGroupCreate) SetNillableUpdateAt(i *int64) *ApplicationGroupCreate {
	if i != nil {
		agc.SetUpdateAt(*i)
	}
	return agc
}

// SetDeleteAt sets the "delete_at" field.
func (agc *ApplicationGroupCreate) SetDeleteAt(i int64) *ApplicationGroupCreate {
	agc.mutation.SetDeleteAt(i)
	return agc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (agc *ApplicationGroupCreate) SetNillableDeleteAt(i *int64) *ApplicationGroupCreate {
	if i != nil {
		agc.SetDeleteAt(*i)
	}
	return agc
}

// SetID sets the "id" field.
func (agc *ApplicationGroupCreate) SetID(u uuid.UUID) *ApplicationGroupCreate {
	agc.mutation.SetID(u)
	return agc
}

// Mutation returns the ApplicationGroupMutation object of the builder.
func (agc *ApplicationGroupCreate) Mutation() *ApplicationGroupMutation {
	return agc.mutation
}

// Save creates the ApplicationGroup in the database.
func (agc *ApplicationGroupCreate) Save(ctx context.Context) (*ApplicationGroup, error) {
	var (
		err  error
		node *ApplicationGroup
	)
	agc.defaults()
	if len(agc.hooks) == 0 {
		if err = agc.check(); err != nil {
			return nil, err
		}
		node, err = agc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = agc.check(); err != nil {
				return nil, err
			}
			agc.mutation = mutation
			if node, err = agc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(agc.hooks) - 1; i >= 0; i-- {
			if agc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (agc *ApplicationGroupCreate) SaveX(ctx context.Context) *ApplicationGroup {
	v, err := agc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agc *ApplicationGroupCreate) Exec(ctx context.Context) error {
	_, err := agc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agc *ApplicationGroupCreate) ExecX(ctx context.Context) {
	if err := agc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agc *ApplicationGroupCreate) defaults() {
	if _, ok := agc.mutation.CreateAt(); !ok {
		v := applicationgroup.DefaultCreateAt()
		agc.mutation.SetCreateAt(v)
	}
	if _, ok := agc.mutation.UpdateAt(); !ok {
		v := applicationgroup.DefaultUpdateAt()
		agc.mutation.SetUpdateAt(v)
	}
	if _, ok := agc.mutation.DeleteAt(); !ok {
		v := applicationgroup.DefaultDeleteAt()
		agc.mutation.SetDeleteAt(v)
	}
	if _, ok := agc.mutation.ID(); !ok {
		v := applicationgroup.DefaultID()
		agc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (agc *ApplicationGroupCreate) check() error {
	if _, ok := agc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := agc.mutation.GroupName(); !ok {
		return &ValidationError{Name: "group_name", err: errors.New(`ent: missing required field "group_name"`)}
	}
	if _, ok := agc.mutation.GroupOwner(); !ok {
		return &ValidationError{Name: "group_owner", err: errors.New(`ent: missing required field "group_owner"`)}
	}
	if _, ok := agc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := agc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := agc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (agc *ApplicationGroupCreate) sqlSave(ctx context.Context) (*ApplicationGroup, error) {
	_node, _spec := agc.createSpec()
	if err := sqlgraph.CreateNode(ctx, agc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (agc *ApplicationGroupCreate) createSpec() (*ApplicationGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicationGroup{config: agc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: applicationgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationgroup.FieldID,
			},
		}
	)
	if id, ok := agc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := agc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationgroup.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := agc.mutation.GroupName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationgroup.FieldGroupName,
		})
		_node.GroupName = value
	}
	if value, ok := agc.mutation.GroupLogo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationgroup.FieldGroupLogo,
		})
		_node.GroupLogo = value
	}
	if value, ok := agc.mutation.GroupOwner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationgroup.FieldGroupOwner,
		})
		_node.GroupOwner = value
	}
	if value, ok := agc.mutation.Annotation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationgroup.FieldAnnotation,
		})
		_node.Annotation = value
	}
	if value, ok := agc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: applicationgroup.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := agc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: applicationgroup.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := agc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: applicationgroup.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// ApplicationGroupCreateBulk is the builder for creating many ApplicationGroup entities in bulk.
type ApplicationGroupCreateBulk struct {
	config
	builders []*ApplicationGroupCreate
}

// Save creates the ApplicationGroup entities in the database.
func (agcb *ApplicationGroupCreateBulk) Save(ctx context.Context) ([]*ApplicationGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(agcb.builders))
	nodes := make([]*ApplicationGroup, len(agcb.builders))
	mutators := make([]Mutator, len(agcb.builders))
	for i := range agcb.builders {
		func(i int, root context.Context) {
			builder := agcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, agcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, agcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, agcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (agcb *ApplicationGroupCreateBulk) SaveX(ctx context.Context) []*ApplicationGroup {
	v, err := agcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agcb *ApplicationGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := agcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agcb *ApplicationGroupCreateBulk) ExecX(ctx context.Context) {
	if err := agcb.Exec(ctx); err != nil {
		panic(err)
	}
}