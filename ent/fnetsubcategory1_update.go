// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lenon/gofii/ent/fnetdocument"
	"github.com/lenon/gofii/ent/fnetsubcategory1"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetSubCategory1Update is the builder for updating FnetSubCategory1 entities.
type FnetSubCategory1Update struct {
	config
	hooks    []Hook
	mutation *FnetSubCategory1Mutation
}

// Where appends a list predicates to the FnetSubCategory1Update builder.
func (fsc *FnetSubCategory1Update) Where(ps ...predicate.FnetSubCategory1) *FnetSubCategory1Update {
	fsc.mutation.Where(ps...)
	return fsc
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fsc *FnetSubCategory1Update) AddDocumentIDs(ids ...int) *FnetSubCategory1Update {
	fsc.mutation.AddDocumentIDs(ids...)
	return fsc
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fsc *FnetSubCategory1Update) AddDocuments(f ...*FnetDocument) *FnetSubCategory1Update {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsc.AddDocumentIDs(ids...)
}

// Mutation returns the FnetSubCategory1Mutation object of the builder.
func (fsc *FnetSubCategory1Update) Mutation() *FnetSubCategory1Mutation {
	return fsc.mutation
}

// ClearDocuments clears all "documents" edges to the FnetDocument entity.
func (fsc *FnetSubCategory1Update) ClearDocuments() *FnetSubCategory1Update {
	fsc.mutation.ClearDocuments()
	return fsc
}

// RemoveDocumentIDs removes the "documents" edge to FnetDocument entities by IDs.
func (fsc *FnetSubCategory1Update) RemoveDocumentIDs(ids ...int) *FnetSubCategory1Update {
	fsc.mutation.RemoveDocumentIDs(ids...)
	return fsc
}

// RemoveDocuments removes "documents" edges to FnetDocument entities.
func (fsc *FnetSubCategory1Update) RemoveDocuments(f ...*FnetDocument) *FnetSubCategory1Update {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsc.RemoveDocumentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fsc *FnetSubCategory1Update) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fsc.sqlSave, fsc.mutation, fsc.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fsc *FnetSubCategory1Update) SaveX(ctx context.Context) int {
	affected, err := fsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fsc *FnetSubCategory1Update) Exec(ctx context.Context) error {
	_, err := fsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FnetSubCategory1Update) ExecX(ctx context.Context) {
	if err := fsc.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fsc *FnetSubCategory1Update) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(fnetsubcategory1.Table, fnetsubcategory1.Columns, sqlgraph.NewFieldSpec(fnetsubcategory1.FieldID, field.TypeInt))
	if ps := fsc.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fsc.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory1.DocumentsTable,
			Columns: []string{fnetsubcategory1.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsc.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !fsc.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory1.DocumentsTable,
			Columns: []string{fnetsubcategory1.DocumentsColumn},
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
	if nodes := fsc.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory1.DocumentsTable,
			Columns: []string{fnetsubcategory1.DocumentsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fsc.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fnetsubcategory1.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fsc.mutation.done = true
	return n, nil
}

// FnetSubCategory1UpdateOne is the builder for updating a single FnetSubCategory1 entity.
type FnetSubCategory1UpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FnetSubCategory1Mutation
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fsco *FnetSubCategory1UpdateOne) AddDocumentIDs(ids ...int) *FnetSubCategory1UpdateOne {
	fsco.mutation.AddDocumentIDs(ids...)
	return fsco
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fsco *FnetSubCategory1UpdateOne) AddDocuments(f ...*FnetDocument) *FnetSubCategory1UpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsco.AddDocumentIDs(ids...)
}

// Mutation returns the FnetSubCategory1Mutation object of the builder.
func (fsco *FnetSubCategory1UpdateOne) Mutation() *FnetSubCategory1Mutation {
	return fsco.mutation
}

// ClearDocuments clears all "documents" edges to the FnetDocument entity.
func (fsco *FnetSubCategory1UpdateOne) ClearDocuments() *FnetSubCategory1UpdateOne {
	fsco.mutation.ClearDocuments()
	return fsco
}

// RemoveDocumentIDs removes the "documents" edge to FnetDocument entities by IDs.
func (fsco *FnetSubCategory1UpdateOne) RemoveDocumentIDs(ids ...int) *FnetSubCategory1UpdateOne {
	fsco.mutation.RemoveDocumentIDs(ids...)
	return fsco
}

// RemoveDocuments removes "documents" edges to FnetDocument entities.
func (fsco *FnetSubCategory1UpdateOne) RemoveDocuments(f ...*FnetDocument) *FnetSubCategory1UpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsco.RemoveDocumentIDs(ids...)
}

// Where appends a list predicates to the FnetSubCategory1Update builder.
func (fsco *FnetSubCategory1UpdateOne) Where(ps ...predicate.FnetSubCategory1) *FnetSubCategory1UpdateOne {
	fsco.mutation.Where(ps...)
	return fsco
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fsco *FnetSubCategory1UpdateOne) Select(field string, fields ...string) *FnetSubCategory1UpdateOne {
	fsco.fields = append([]string{field}, fields...)
	return fsco
}

// Save executes the query and returns the updated FnetSubCategory1 entity.
func (fsco *FnetSubCategory1UpdateOne) Save(ctx context.Context) (*FnetSubCategory1, error) {
	return withHooks(ctx, fsco.sqlSave, fsco.mutation, fsco.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fsco *FnetSubCategory1UpdateOne) SaveX(ctx context.Context) *FnetSubCategory1 {
	node, err := fsco.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fsco *FnetSubCategory1UpdateOne) Exec(ctx context.Context) error {
	_, err := fsco.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsco *FnetSubCategory1UpdateOne) ExecX(ctx context.Context) {
	if err := fsco.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fsco *FnetSubCategory1UpdateOne) sqlSave(ctx context.Context) (_node *FnetSubCategory1, err error) {
	_spec := sqlgraph.NewUpdateSpec(fnetsubcategory1.Table, fnetsubcategory1.Columns, sqlgraph.NewFieldSpec(fnetsubcategory1.FieldID, field.TypeInt))
	id, ok := fsco.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FnetSubCategory1.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fsco.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fnetsubcategory1.FieldID)
		for _, f := range fields {
			if !fnetsubcategory1.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fnetsubcategory1.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fsco.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fsco.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory1.DocumentsTable,
			Columns: []string{fnetsubcategory1.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fnetdocument.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsco.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !fsco.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory1.DocumentsTable,
			Columns: []string{fnetsubcategory1.DocumentsColumn},
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
	if nodes := fsco.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory1.DocumentsTable,
			Columns: []string{fnetsubcategory1.DocumentsColumn},
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
	_node = &FnetSubCategory1{config: fsco.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fsco.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fnetsubcategory1.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fsco.mutation.done = true
	return _node, nil
}
