// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationresource"
	"github.com/google/uuid"
)

// ApplicationResourceCreate is the builder for creating a ApplicationResource entity.
type ApplicationResourceCreate struct {
	config
	mutation *ApplicationResourceMutation
	hooks    []Hook
}

// SetAppID sets the "app_id" field.
func (arc *ApplicationResourceCreate) SetAppID(u uuid.UUID) *ApplicationResourceCreate {
	arc.mutation.SetAppID(u)
	return arc
}

// SetResourceName sets the "resource_name" field.
func (arc *ApplicationResourceCreate) SetResourceName(s string) *ApplicationResourceCreate {
	arc.mutation.SetResourceName(s)
	return arc
}

// SetResourceDescription sets the "resource_description" field.
func (arc *ApplicationResourceCreate) SetResourceDescription(s string) *ApplicationResourceCreate {
	arc.mutation.SetResourceDescription(s)
	return arc
}

// SetNillableResourceDescription sets the "resource_description" field if the given value is not nil.
func (arc *ApplicationResourceCreate) SetNillableResourceDescription(s *string) *ApplicationResourceCreate {
	if s != nil {
		arc.SetResourceDescription(*s)
	}
	return arc
}

// SetType sets the "type" field.
func (arc *ApplicationResourceCreate) SetType(s string) *ApplicationResourceCreate {
	arc.mutation.SetType(s)
	return arc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (arc *ApplicationResourceCreate) SetNillableType(s *string) *ApplicationResourceCreate {
	if s != nil {
		arc.SetType(*s)
	}
	return arc
}

// SetCreator sets the "creator" field.
func (arc *ApplicationResourceCreate) SetCreator(u uuid.UUID) *ApplicationResourceCreate {
	arc.mutation.SetCreator(u)
	return arc
}

// SetCreateAt sets the "create_at" field.
func (arc *ApplicationResourceCreate) SetCreateAt(u uint32) *ApplicationResourceCreate {
	arc.mutation.SetCreateAt(u)
	return arc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (arc *ApplicationResourceCreate) SetNillableCreateAt(u *uint32) *ApplicationResourceCreate {
	if u != nil {
		arc.SetCreateAt(*u)
	}
	return arc
}

// SetUpdateAt sets the "update_at" field.
func (arc *ApplicationResourceCreate) SetUpdateAt(u uint32) *ApplicationResourceCreate {
	arc.mutation.SetUpdateAt(u)
	return arc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (arc *ApplicationResourceCreate) SetNillableUpdateAt(u *uint32) *ApplicationResourceCreate {
	if u != nil {
		arc.SetUpdateAt(*u)
	}
	return arc
}

// SetDeleteAt sets the "delete_at" field.
func (arc *ApplicationResourceCreate) SetDeleteAt(u uint32) *ApplicationResourceCreate {
	arc.mutation.SetDeleteAt(u)
	return arc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (arc *ApplicationResourceCreate) SetNillableDeleteAt(u *uint32) *ApplicationResourceCreate {
	if u != nil {
		arc.SetDeleteAt(*u)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *ApplicationResourceCreate) SetID(u uuid.UUID) *ApplicationResourceCreate {
	arc.mutation.SetID(u)
	return arc
}

// Mutation returns the ApplicationResourceMutation object of the builder.
func (arc *ApplicationResourceCreate) Mutation() *ApplicationResourceMutation {
	return arc.mutation
}

// Save creates the ApplicationResource in the database.
func (arc *ApplicationResourceCreate) Save(ctx context.Context) (*ApplicationResource, error) {
	var (
		err  error
		node *ApplicationResource
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *ApplicationResourceCreate) SaveX(ctx context.Context) *ApplicationResource {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *ApplicationResourceCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *ApplicationResourceCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *ApplicationResourceCreate) defaults() {
	if _, ok := arc.mutation.GetType(); !ok {
		v := applicationresource.DefaultType
		arc.mutation.SetType(v)
	}
	if _, ok := arc.mutation.CreateAt(); !ok {
		v := applicationresource.DefaultCreateAt()
		arc.mutation.SetCreateAt(v)
	}
	if _, ok := arc.mutation.UpdateAt(); !ok {
		v := applicationresource.DefaultUpdateAt()
		arc.mutation.SetUpdateAt(v)
	}
	if _, ok := arc.mutation.DeleteAt(); !ok {
		v := applicationresource.DefaultDeleteAt()
		arc.mutation.SetDeleteAt(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		v := applicationresource.DefaultID()
		arc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *ApplicationResourceCreate) check() error {
	if _, ok := arc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := arc.mutation.ResourceName(); !ok {
		return &ValidationError{Name: "resource_name", err: errors.New(`ent: missing required field "resource_name"`)}
	}
	if _, ok := arc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if _, ok := arc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "creator"`)}
	}
	if _, ok := arc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := arc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := arc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (arc *ApplicationResourceCreate) sqlSave(ctx context.Context) (*ApplicationResource, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
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

func (arc *ApplicationResourceCreate) createSpec() (*ApplicationResource, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicationResource{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: applicationresource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationresource.FieldID,
			},
		}
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationresource.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := arc.mutation.ResourceName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldResourceName,
		})
		_node.ResourceName = value
	}
	if value, ok := arc.mutation.ResourceDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldResourceDescription,
		})
		_node.ResourceDescription = value
	}
	if value, ok := arc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: applicationresource.FieldType,
		})
		_node.Type = value
	}
	if value, ok := arc.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationresource.FieldCreator,
		})
		_node.Creator = value
	}
	if value, ok := arc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := arc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := arc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationresource.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// ApplicationResourceCreateBulk is the builder for creating many ApplicationResource entities in bulk.
type ApplicationResourceCreateBulk struct {
	config
	builders []*ApplicationResourceCreate
}

// Save creates the ApplicationResource entities in the database.
func (arcb *ApplicationResourceCreateBulk) Save(ctx context.Context) ([]*ApplicationResource, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*ApplicationResource, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *ApplicationResourceCreateBulk) SaveX(ctx context.Context) []*ApplicationResource {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *ApplicationResourceCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *ApplicationResourceCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}
