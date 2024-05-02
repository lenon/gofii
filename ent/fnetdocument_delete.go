// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lenon/gofii/ent/fnetdocument"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetDocumentDelete is the builder for deleting a FnetDocument entity.
type FnetDocumentDelete struct {
	config
	hooks    []Hook
	mutation *FnetDocumentMutation
}

// Where appends a list predicates to the FnetDocumentDelete builder.
func (fdd *FnetDocumentDelete) Where(ps ...predicate.FnetDocument) *FnetDocumentDelete {
	fdd.mutation.Where(ps...)
	return fdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fdd *FnetDocumentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fdd.sqlExec, fdd.mutation, fdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fdd *FnetDocumentDelete) ExecX(ctx context.Context) int {
	n, err := fdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fdd *FnetDocumentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(fnetdocument.Table, sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt))
	if ps := fdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fdd.mutation.done = true
	return affected, err
}

// FnetDocumentDeleteOne is the builder for deleting a single FnetDocument entity.
type FnetDocumentDeleteOne struct {
	fdd *FnetDocumentDelete
}

// Where appends a list predicates to the FnetDocumentDelete builder.
func (fddo *FnetDocumentDeleteOne) Where(ps ...predicate.FnetDocument) *FnetDocumentDeleteOne {
	fddo.fdd.mutation.Where(ps...)
	return fddo
}

// Exec executes the deletion query.
func (fddo *FnetDocumentDeleteOne) Exec(ctx context.Context) error {
	n, err := fddo.fdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fnetdocument.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fddo *FnetDocumentDeleteOne) ExecX(ctx context.Context) {
	if err := fddo.Exec(ctx); err != nil {
		panic(err)
	}
}
