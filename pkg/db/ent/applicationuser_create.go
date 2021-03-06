// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationuser"
	"github.com/google/uuid"
)

// ApplicationUserCreate is the builder for creating a ApplicationUser entity.
type ApplicationUserCreate struct {
	config
	mutation *ApplicationUserMutation
	hooks    []Hook
}

// SetAppID sets the "app_id" field.
func (auc *ApplicationUserCreate) SetAppID(u uuid.UUID) *ApplicationUserCreate {
	auc.mutation.SetAppID(u)
	return auc
}

// SetUserID sets the "user_id" field.
func (auc *ApplicationUserCreate) SetUserID(u uuid.UUID) *ApplicationUserCreate {
	auc.mutation.SetUserID(u)
	return auc
}

// SetOriginal sets the "original" field.
func (auc *ApplicationUserCreate) SetOriginal(b bool) *ApplicationUserCreate {
	auc.mutation.SetOriginal(b)
	return auc
}

// SetNillableOriginal sets the "original" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableOriginal(b *bool) *ApplicationUserCreate {
	if b != nil {
		auc.SetOriginal(*b)
	}
	return auc
}

// SetKycVerify sets the "kyc_verify" field.
func (auc *ApplicationUserCreate) SetKycVerify(b bool) *ApplicationUserCreate {
	auc.mutation.SetKycVerify(b)
	return auc
}

// SetNillableKycVerify sets the "kyc_verify" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableKycVerify(b *bool) *ApplicationUserCreate {
	if b != nil {
		auc.SetKycVerify(*b)
	}
	return auc
}

// SetGaVerify sets the "ga_verify" field.
func (auc *ApplicationUserCreate) SetGaVerify(b bool) *ApplicationUserCreate {
	auc.mutation.SetGaVerify(b)
	return auc
}

// SetNillableGaVerify sets the "ga_verify" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableGaVerify(b *bool) *ApplicationUserCreate {
	if b != nil {
		auc.SetGaVerify(*b)
	}
	return auc
}

// SetGaLogin sets the "ga_login" field.
func (auc *ApplicationUserCreate) SetGaLogin(b bool) *ApplicationUserCreate {
	auc.mutation.SetGaLogin(b)
	return auc
}

// SetNillableGaLogin sets the "ga_login" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableGaLogin(b *bool) *ApplicationUserCreate {
	if b != nil {
		auc.SetGaLogin(*b)
	}
	return auc
}

// SetSmsLogin sets the "sms_login" field.
func (auc *ApplicationUserCreate) SetSmsLogin(b bool) *ApplicationUserCreate {
	auc.mutation.SetSmsLogin(b)
	return auc
}

// SetNillableSmsLogin sets the "sms_login" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableSmsLogin(b *bool) *ApplicationUserCreate {
	if b != nil {
		auc.SetSmsLogin(*b)
	}
	return auc
}

// SetLoginNumber sets the "Login_number" field.
func (auc *ApplicationUserCreate) SetLoginNumber(u uint32) *ApplicationUserCreate {
	auc.mutation.SetLoginNumber(u)
	return auc
}

// SetNillableLoginNumber sets the "Login_number" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableLoginNumber(u *uint32) *ApplicationUserCreate {
	if u != nil {
		auc.SetLoginNumber(*u)
	}
	return auc
}

// SetCreateAt sets the "create_at" field.
func (auc *ApplicationUserCreate) SetCreateAt(u uint32) *ApplicationUserCreate {
	auc.mutation.SetCreateAt(u)
	return auc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableCreateAt(u *uint32) *ApplicationUserCreate {
	if u != nil {
		auc.SetCreateAt(*u)
	}
	return auc
}

// SetDeleteAt sets the "delete_at" field.
func (auc *ApplicationUserCreate) SetDeleteAt(u uint32) *ApplicationUserCreate {
	auc.mutation.SetDeleteAt(u)
	return auc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auc *ApplicationUserCreate) SetNillableDeleteAt(u *uint32) *ApplicationUserCreate {
	if u != nil {
		auc.SetDeleteAt(*u)
	}
	return auc
}

// SetID sets the "id" field.
func (auc *ApplicationUserCreate) SetID(u uuid.UUID) *ApplicationUserCreate {
	auc.mutation.SetID(u)
	return auc
}

// Mutation returns the ApplicationUserMutation object of the builder.
func (auc *ApplicationUserCreate) Mutation() *ApplicationUserMutation {
	return auc.mutation
}

// Save creates the ApplicationUser in the database.
func (auc *ApplicationUserCreate) Save(ctx context.Context) (*ApplicationUser, error) {
	var (
		err  error
		node *ApplicationUser
	)
	auc.defaults()
	if len(auc.hooks) == 0 {
		if err = auc.check(); err != nil {
			return nil, err
		}
		node, err = auc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auc.check(); err != nil {
				return nil, err
			}
			auc.mutation = mutation
			if node, err = auc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(auc.hooks) - 1; i >= 0; i-- {
			if auc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auc *ApplicationUserCreate) SaveX(ctx context.Context) *ApplicationUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auc *ApplicationUserCreate) Exec(ctx context.Context) error {
	_, err := auc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auc *ApplicationUserCreate) ExecX(ctx context.Context) {
	if err := auc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auc *ApplicationUserCreate) defaults() {
	if _, ok := auc.mutation.Original(); !ok {
		v := applicationuser.DefaultOriginal
		auc.mutation.SetOriginal(v)
	}
	if _, ok := auc.mutation.KycVerify(); !ok {
		v := applicationuser.DefaultKycVerify
		auc.mutation.SetKycVerify(v)
	}
	if _, ok := auc.mutation.GaVerify(); !ok {
		v := applicationuser.DefaultGaVerify
		auc.mutation.SetGaVerify(v)
	}
	if _, ok := auc.mutation.GaLogin(); !ok {
		v := applicationuser.DefaultGaLogin
		auc.mutation.SetGaLogin(v)
	}
	if _, ok := auc.mutation.SmsLogin(); !ok {
		v := applicationuser.DefaultSmsLogin
		auc.mutation.SetSmsLogin(v)
	}
	if _, ok := auc.mutation.LoginNumber(); !ok {
		v := applicationuser.DefaultLoginNumber
		auc.mutation.SetLoginNumber(v)
	}
	if _, ok := auc.mutation.CreateAt(); !ok {
		v := applicationuser.DefaultCreateAt()
		auc.mutation.SetCreateAt(v)
	}
	if _, ok := auc.mutation.DeleteAt(); !ok {
		v := applicationuser.DefaultDeleteAt()
		auc.mutation.SetDeleteAt(v)
	}
	if _, ok := auc.mutation.ID(); !ok {
		v := applicationuser.DefaultID()
		auc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *ApplicationUserCreate) check() error {
	if _, ok := auc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := auc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := auc.mutation.Original(); !ok {
		return &ValidationError{Name: "original", err: errors.New(`ent: missing required field "original"`)}
	}
	if _, ok := auc.mutation.KycVerify(); !ok {
		return &ValidationError{Name: "kyc_verify", err: errors.New(`ent: missing required field "kyc_verify"`)}
	}
	if _, ok := auc.mutation.GaVerify(); !ok {
		return &ValidationError{Name: "ga_verify", err: errors.New(`ent: missing required field "ga_verify"`)}
	}
	if _, ok := auc.mutation.GaLogin(); !ok {
		return &ValidationError{Name: "ga_login", err: errors.New(`ent: missing required field "ga_login"`)}
	}
	if _, ok := auc.mutation.SmsLogin(); !ok {
		return &ValidationError{Name: "sms_login", err: errors.New(`ent: missing required field "sms_login"`)}
	}
	if _, ok := auc.mutation.LoginNumber(); !ok {
		return &ValidationError{Name: "Login_number", err: errors.New(`ent: missing required field "Login_number"`)}
	}
	if _, ok := auc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := auc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (auc *ApplicationUserCreate) sqlSave(ctx context.Context) (*ApplicationUser, error) {
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
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

func (auc *ApplicationUserCreate) createSpec() (*ApplicationUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicationUser{config: auc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: applicationuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: applicationuser.FieldID,
			},
		}
	)
	if id, ok := auc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := auc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationuser.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := auc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applicationuser.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := auc.mutation.Original(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldOriginal,
		})
		_node.Original = value
	}
	if value, ok := auc.mutation.KycVerify(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldKycVerify,
		})
		_node.KycVerify = value
	}
	if value, ok := auc.mutation.GaVerify(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldGaVerify,
		})
		_node.GaVerify = value
	}
	if value, ok := auc.mutation.GaLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldGaLogin,
		})
		_node.GaLogin = value
	}
	if value, ok := auc.mutation.SmsLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: applicationuser.FieldSmsLogin,
		})
		_node.SmsLogin = value
	}
	if value, ok := auc.mutation.LoginNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldLoginNumber,
		})
		_node.LoginNumber = value
	}
	if value, ok := auc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := auc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applicationuser.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// ApplicationUserCreateBulk is the builder for creating many ApplicationUser entities in bulk.
type ApplicationUserCreateBulk struct {
	config
	builders []*ApplicationUserCreate
}

// Save creates the ApplicationUser entities in the database.
func (aucb *ApplicationUserCreateBulk) Save(ctx context.Context) ([]*ApplicationUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*ApplicationUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationUserMutation)
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
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *ApplicationUserCreateBulk) SaveX(ctx context.Context) []*ApplicationUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aucb *ApplicationUserCreateBulk) Exec(ctx context.Context) error {
	_, err := aucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucb *ApplicationUserCreateBulk) ExecX(ctx context.Context) {
	if err := aucb.Exec(ctx); err != nil {
		panic(err)
	}
}
