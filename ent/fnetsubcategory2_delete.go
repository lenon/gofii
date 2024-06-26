// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lenon/gofii/ent/fnetsubcategory2"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetSubCategory2Delete is the builder for deleting a FnetSubCategory2 entity.
type FnetSubCategory2Delete struct {
	config
	hooks    []Hook
	mutation *FnetSubCategory2Mutation
}

// Where appends a list predicates to the FnetSubCategory2Delete builder.
func (fsc *FnetSubCategory2Delete) Where(ps ...predicate.FnetSubCategory2) *FnetSubCategory2Delete {
	fsc.mutation.Where(ps...)
	return fsc
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fsc *FnetSubCategory2Delete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fsc.sqlExec, fsc.mutation, fsc.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FnetSubCategory2Delete) ExecX(ctx context.Context) int {
	n, err := fsc.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fsc *FnetSubCategory2Delete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(fnetsubcategory2.Table, sqlgraph.NewFieldSpec(fnetsubcategory2.FieldID, field.TypeInt))
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
	fsc.mutation.done = true
	return affected, err
}

// FnetSubCategory2DeleteOne is the builder for deleting a single FnetSubCategory2 entity.
type FnetSubCategory2DeleteOne struct {
	fsc *FnetSubCategory2Delete
}

// Where appends a list predicates to the FnetSubCategory2Delete builder.
func (fsco *FnetSubCategory2DeleteOne) Where(ps ...predicate.FnetSubCategory2) *FnetSubCategory2DeleteOne {
	fsco.fsc.mutation.Where(ps...)
	return fsco
}

// Exec executes the deletion query.
func (fsco *FnetSubCategory2DeleteOne) Exec(ctx context.Context) error {
	n, err := fsco.fsc.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fnetsubcategory2.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fsco *FnetSubCategory2DeleteOne) ExecX(ctx context.Context) {
	if err := fsco.Exec(ctx); err != nil {
		panic(err)
	}
}
