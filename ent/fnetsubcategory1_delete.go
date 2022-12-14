// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lenon/gofii/ent/fnetsubcategory1"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetSubCategory1Delete is the builder for deleting a FnetSubCategory1 entity.
type FnetSubCategory1Delete struct {
	config
	hooks    []Hook
	mutation *FnetSubCategory1Mutation
}

// Where appends a list predicates to the FnetSubCategory1Delete builder.
func (fsc *FnetSubCategory1Delete) Where(ps ...predicate.FnetSubCategory1) *FnetSubCategory1Delete {
	fsc.mutation.Where(ps...)
	return fsc
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fsc *FnetSubCategory1Delete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fsc.hooks) == 0 {
		affected, err = fsc.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FnetSubCategory1Mutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fsc.mutation = mutation
			affected, err = fsc.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fsc.hooks) - 1; i >= 0; i-- {
			if fsc.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fsc.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FnetSubCategory1Delete) ExecX(ctx context.Context) int {
	n, err := fsc.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fsc *FnetSubCategory1Delete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fnetsubcategory1.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fnetsubcategory1.FieldID,
			},
		},
	}
	if ps := fsc.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fsc.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FnetSubCategory1DeleteOne is the builder for deleting a single FnetSubCategory1 entity.
type FnetSubCategory1DeleteOne struct {
	fsc *FnetSubCategory1Delete
}

// Exec executes the deletion query.
func (fsco *FnetSubCategory1DeleteOne) Exec(ctx context.Context) error {
	n, err := fsco.fsc.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fnetsubcategory1.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fsco *FnetSubCategory1DeleteOne) ExecX(ctx context.Context) {
	fsco.fsc.ExecX(ctx)
}
