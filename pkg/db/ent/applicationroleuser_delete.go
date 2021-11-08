// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/predicate"
)

// ApplicationRoleUserDelete is the builder for deleting a ApplicationRoleUser entity.
type ApplicationRoleUserDelete struct {
	config
	hooks    []Hook
	mutation *ApplicationRoleUserMutation
}

// Where appends a list predicates to the ApplicationRoleUserDelete builder.
func (arud *ApplicationRoleUserDelete) Where(ps ...predicate.ApplicationRoleUser) *ApplicationRoleUserDelete {
	arud.mutation.Where(ps...)
	return arud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (arud *ApplicationRoleUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(arud.hooks) == 0 {
		affected, err = arud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationRoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			arud.mutation = mutation
			affected, err = arud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(arud.hooks) - 1; i >= 0; i-- {
			if arud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (arud *ApplicationRoleUserDelete) ExecX(ctx context.Context) int {
	n, err := arud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (arud *ApplicationRoleUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: applicationroleuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationroleuser.FieldID,
			},
		},
	}
	if ps := arud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, arud.driver, _spec)
}

// ApplicationRoleUserDeleteOne is the builder for deleting a single ApplicationRoleUser entity.
type ApplicationRoleUserDeleteOne struct {
	arud *ApplicationRoleUserDelete
}

// Exec executes the deletion query.
func (arudo *ApplicationRoleUserDeleteOne) Exec(ctx context.Context) error {
	n, err := arudo.arud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{applicationroleuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (arudo *ApplicationRoleUserDeleteOne) ExecX(ctx context.Context) {
	arudo.arud.ExecX(ctx)
}