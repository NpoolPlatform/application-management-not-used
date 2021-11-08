// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/application"
	"github.com/google/uuid"
)

// ApplicationCreate is the builder for creating a Application entity.
type ApplicationCreate struct {
	config
	mutation *ApplicationMutation
	hooks    []Hook
}

// SetApplicationName sets the "application_name" field.
func (ac *ApplicationCreate) SetApplicationName(s string) *ApplicationCreate {
	ac.mutation.SetApplicationName(s)
	return ac
}

// SetApplicationOwner sets the "application_owner" field.
func (ac *ApplicationCreate) SetApplicationOwner(u uuid.UUID) *ApplicationCreate {
	ac.mutation.SetApplicationOwner(u)
	return ac
}

// SetHomepageURL sets the "homepage_url" field.
func (ac *ApplicationCreate) SetHomepageURL(s string) *ApplicationCreate {
	ac.mutation.SetHomepageURL(s)
	return ac
}

// SetNillableHomepageURL sets the "homepage_url" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableHomepageURL(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetHomepageURL(*s)
	}
	return ac
}

// SetRedirectURL sets the "redirect_url" field.
func (ac *ApplicationCreate) SetRedirectURL(s string) *ApplicationCreate {
	ac.mutation.SetRedirectURL(s)
	return ac
}

// SetNillableRedirectURL sets the "redirect_url" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableRedirectURL(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetRedirectURL(*s)
	}
	return ac
}

// SetClientSecret sets the "client_secret" field.
func (ac *ApplicationCreate) SetClientSecret(s string) *ApplicationCreate {
	ac.mutation.SetClientSecret(s)
	return ac
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableClientSecret(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetClientSecret(*s)
	}
	return ac
}

// SetApplicationLogo sets the "application_logo" field.
func (ac *ApplicationCreate) SetApplicationLogo(s string) *ApplicationCreate {
	ac.mutation.SetApplicationLogo(s)
	return ac
}

// SetNillableApplicationLogo sets the "application_logo" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableApplicationLogo(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetApplicationLogo(*s)
	}
	return ac
}

// SetCreateAt sets the "create_at" field.
func (ac *ApplicationCreate) SetCreateAt(i int64) *ApplicationCreate {
	ac.mutation.SetCreateAt(i)
	return ac
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreateAt(i *int64) *ApplicationCreate {
	if i != nil {
		ac.SetCreateAt(*i)
	}
	return ac
}

// SetUpdateAt sets the "update_at" field.
func (ac *ApplicationCreate) SetUpdateAt(i int64) *ApplicationCreate {
	ac.mutation.SetUpdateAt(i)
	return ac
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableUpdateAt(i *int64) *ApplicationCreate {
	if i != nil {
		ac.SetUpdateAt(*i)
	}
	return ac
}

// SetDeleteAt sets the "delete_at" field.
func (ac *ApplicationCreate) SetDeleteAt(i int64) *ApplicationCreate {
	ac.mutation.SetDeleteAt(i)
	return ac
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableDeleteAt(i *int64) *ApplicationCreate {
	if i != nil {
		ac.SetDeleteAt(*i)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ApplicationCreate) SetID(s string) *ApplicationCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableID(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetID(*s)
	}
	return ac
}

// Mutation returns the ApplicationMutation object of the builder.
func (ac *ApplicationCreate) Mutation() *ApplicationMutation {
	return ac.mutation
}

// Save creates the Application in the database.
func (ac *ApplicationCreate) Save(ctx context.Context) (*Application, error) {
	var (
		err  error
		node *Application
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApplicationCreate) SaveX(ctx context.Context) *Application {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApplicationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApplicationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ApplicationCreate) defaults() {
	if _, ok := ac.mutation.ClientSecret(); !ok {
		v := application.DefaultClientSecret()
		ac.mutation.SetClientSecret(v)
	}
	if _, ok := ac.mutation.CreateAt(); !ok {
		v := application.DefaultCreateAt()
		ac.mutation.SetCreateAt(v)
	}
	if _, ok := ac.mutation.UpdateAt(); !ok {
		v := application.DefaultUpdateAt()
		ac.mutation.SetUpdateAt(v)
	}
	if _, ok := ac.mutation.DeleteAt(); !ok {
		v := application.DefaultDeleteAt()
		ac.mutation.SetDeleteAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := application.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApplicationCreate) check() error {
	if _, ok := ac.mutation.ApplicationName(); !ok {
		return &ValidationError{Name: "application_name", err: errors.New(`ent: missing required field "application_name"`)}
	}
	if _, ok := ac.mutation.ApplicationOwner(); !ok {
		return &ValidationError{Name: "application_owner", err: errors.New(`ent: missing required field "application_owner"`)}
	}
	if _, ok := ac.mutation.ClientSecret(); !ok {
		return &ValidationError{Name: "client_secret", err: errors.New(`ent: missing required field "client_secret"`)}
	}
	if _, ok := ac.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := ac.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := ac.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (ac *ApplicationCreate) sqlSave(ctx context.Context) (*Application, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(string)
	}
	return _node, nil
}

func (ac *ApplicationCreate) createSpec() (*Application, *sqlgraph.CreateSpec) {
	var (
		_node = &Application{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: application.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: application.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.ApplicationName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldApplicationName,
		})
		_node.ApplicationName = value
	}
	if value, ok := ac.mutation.ApplicationOwner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: application.FieldApplicationOwner,
		})
		_node.ApplicationOwner = value
	}
	if value, ok := ac.mutation.HomepageURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldHomepageURL,
		})
		_node.HomepageURL = value
	}
	if value, ok := ac.mutation.RedirectURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldRedirectURL,
		})
		_node.RedirectURL = value
	}
	if value, ok := ac.mutation.ClientSecret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldClientSecret,
		})
		_node.ClientSecret = value
	}
	if value, ok := ac.mutation.ApplicationLogo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldApplicationLogo,
		})
		_node.ApplicationLogo = value
	}
	if value, ok := ac.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: application.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := ac.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: application.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := ac.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: application.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// ApplicationCreateBulk is the builder for creating many Application entities in bulk.
type ApplicationCreateBulk struct {
	config
	builders []*ApplicationCreate
}

// Save creates the Application entities in the database.
func (acb *ApplicationCreateBulk) Save(ctx context.Context) ([]*Application, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Application, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ApplicationCreateBulk) SaveX(ctx context.Context) []*Application {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApplicationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApplicationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}