// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lenon/gofii/ent/fnetcategory"
	"github.com/lenon/gofii/ent/fnetdocument"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetCategoryUpdate is the builder for updating FnetCategory entities.
type FnetCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *FnetCategoryMutation
}

// Where appends a list predicates to the FnetCategoryUpdate builder.
func (fcu *FnetCategoryUpdate) Where(ps ...predicate.FnetCategory) *FnetCategoryUpdate {
	fcu.mutation.Where(ps...)
	return fcu
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fcu *FnetCategoryUpdate) AddDocumentIDs(ids ...int) *FnetCategoryUpdate {
	fcu.mutation.AddDocumentIDs(ids...)
	return fcu
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fcu *FnetCategoryUpdate) AddDocuments(f ...*FnetDocument) *FnetCategoryUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fcu.AddDocumentIDs(ids...)
}

// Mutation returns the FnetCategoryMutation object of the builder.
func (fcu *FnetCategoryUpdate) Mutation() *FnetCategoryMutation {
	return fcu.mutation
}

// ClearDocuments clears all "documents" edges to the FnetDocument entity.
func (fcu *FnetCategoryUpdate) ClearDocuments() *FnetCategoryUpdate {
	fcu.mutation.ClearDocuments()
	return fcu
}

// RemoveDocumentIDs removes the "documents" edge to FnetDocument entities by IDs.
func (fcu *FnetCategoryUpdate) RemoveDocumentIDs(ids ...int) *FnetCategoryUpdate {
	fcu.mutation.RemoveDocumentIDs(ids...)
	return fcu
}

// RemoveDocuments removes "documents" edges to FnetDocument entities.
func (fcu *FnetCategoryUpdate) RemoveDocuments(f ...*FnetDocument) *FnetCategoryUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fcu.RemoveDocumentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fcu *FnetCategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fcu.sqlSave, fcu.mutation, fcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fcu *FnetCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := fcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fcu *FnetCategoryUpdate) Exec(ctx context.Context) error {
	_, err := fcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcu *FnetCategoryUpdate) ExecX(ctx context.Context) {
	if err := fcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fcu *FnetCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(fnetcategory.Table, fnetcategory.Columns, sqlgraph.NewFieldSpec(fnetcategory.FieldID, field.TypeInt))
	if ps := fcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fcu.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetcategory.DocumentsTable,
			Columns: []string{fnetcategory.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcu.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !fcu.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetcategory.DocumentsTable,
			Columns: []string{fnetcategory.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcu.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetcategory.DocumentsTable,
			Columns: []string{fnetcategory.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fnetcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fcu.mutation.done = true
	return n, nil
}

// FnetCategoryUpdateOne is the builder for updating a single FnetCategory entity.
type FnetCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FnetCategoryMutation
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fcuo *FnetCategoryUpdateOne) AddDocumentIDs(ids ...int) *FnetCategoryUpdateOne {
	fcuo.mutation.AddDocumentIDs(ids...)
	return fcuo
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fcuo *FnetCategoryUpdateOne) AddDocuments(f ...*FnetDocument) *FnetCategoryUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fcuo.AddDocumentIDs(ids...)
}

// Mutation returns the FnetCategoryMutation object of the builder.
func (fcuo *FnetCategoryUpdateOne) Mutation() *FnetCategoryMutation {
	return fcuo.mutation
}

// ClearDocuments clears all "documents" edges to the FnetDocument entity.
func (fcuo *FnetCategoryUpdateOne) ClearDocuments() *FnetCategoryUpdateOne {
	fcuo.mutation.ClearDocuments()
	return fcuo
}

// RemoveDocumentIDs removes the "documents" edge to FnetDocument entities by IDs.
func (fcuo *FnetCategoryUpdateOne) RemoveDocumentIDs(ids ...int) *FnetCategoryUpdateOne {
	fcuo.mutation.RemoveDocumentIDs(ids...)
	return fcuo
}

// RemoveDocuments removes "documents" edges to FnetDocument entities.
func (fcuo *FnetCategoryUpdateOne) RemoveDocuments(f ...*FnetDocument) *FnetCategoryUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fcuo.RemoveDocumentIDs(ids...)
}

// Where appends a list predicates to the FnetCategoryUpdate builder.
func (fcuo *FnetCategoryUpdateOne) Where(ps ...predicate.FnetCategory) *FnetCategoryUpdateOne {
	fcuo.mutation.Where(ps...)
	return fcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fcuo *FnetCategoryUpdateOne) Select(field string, fields ...string) *FnetCategoryUpdateOne {
	fcuo.fields = append([]string{field}, fields...)
	return fcuo
}

// Save executes the query and returns the updated FnetCategory entity.
func (fcuo *FnetCategoryUpdateOne) Save(ctx context.Context) (*FnetCategory, error) {
	return withHooks(ctx, fcuo.sqlSave, fcuo.mutation, fcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fcuo *FnetCategoryUpdateOne) SaveX(ctx context.Context) *FnetCategory {
	node, err := fcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fcuo *FnetCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := fcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcuo *FnetCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := fcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fcuo *FnetCategoryUpdateOne) sqlSave(ctx context.Context) (_node *FnetCategory, err error) {
	_spec := sqlgraph.NewUpdateSpec(fnetcategory.Table, fnetcategory.Columns, sqlgraph.NewFieldSpec(fnetcategory.FieldID, field.TypeInt))
	id, ok := fcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FnetCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fnetcategory.FieldID)
		for _, f := range fields {
			if !fnetcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fnetcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fcuo.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetcategory.DocumentsTable,
			Columns: []string{fnetcategory.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcuo.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !fcuo.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetcategory.DocumentsTable,
			Columns: []string{fnetcategory.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcuo.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetcategory.DocumentsTable,
			Columns: []string{fnetcategory.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FnetCategory{config: fcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fnetcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fcuo.mutation.done = true
	return _node, nil
}
