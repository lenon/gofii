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
	"github.com/lenon/gofii/ent/fnetsubcategory2"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetSubCategory2Update is the builder for updating FnetSubCategory2 entities.
type FnetSubCategory2Update struct {
	config
	hooks    []Hook
	mutation *FnetSubCategory2Mutation
}

// Where appends a list predicates to the FnetSubCategory2Update builder.
func (fsc *FnetSubCategory2Update) Where(ps ...predicate.FnetSubCategory2) *FnetSubCategory2Update {
	fsc.mutation.Where(ps...)
	return fsc
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fsc *FnetSubCategory2Update) AddDocumentIDs(ids ...int) *FnetSubCategory2Update {
	fsc.mutation.AddDocumentIDs(ids...)
	return fsc
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fsc *FnetSubCategory2Update) AddDocuments(f ...*FnetDocument) *FnetSubCategory2Update {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsc.AddDocumentIDs(ids...)
}

// Mutation returns the FnetSubCategory2Mutation object of the builder.
func (fsc *FnetSubCategory2Update) Mutation() *FnetSubCategory2Mutation {
	return fsc.mutation
}

// ClearDocuments clears all "documents" edges to the FnetDocument entity.
func (fsc *FnetSubCategory2Update) ClearDocuments() *FnetSubCategory2Update {
	fsc.mutation.ClearDocuments()
	return fsc
}

// RemoveDocumentIDs removes the "documents" edge to FnetDocument entities by IDs.
func (fsc *FnetSubCategory2Update) RemoveDocumentIDs(ids ...int) *FnetSubCategory2Update {
	fsc.mutation.RemoveDocumentIDs(ids...)
	return fsc
}

// RemoveDocuments removes "documents" edges to FnetDocument entities.
func (fsc *FnetSubCategory2Update) RemoveDocuments(f ...*FnetDocument) *FnetSubCategory2Update {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsc.RemoveDocumentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fsc *FnetSubCategory2Update) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fsc.hooks) == 0 {
		affected, err = fsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FnetSubCategory2Mutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fsc.mutation = mutation
			affected, err = fsc.sqlSave(ctx)
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

// SaveX is like Save, but panics if an error occurs.
func (fsc *FnetSubCategory2Update) SaveX(ctx context.Context) int {
	affected, err := fsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fsc *FnetSubCategory2Update) Exec(ctx context.Context) error {
	_, err := fsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FnetSubCategory2Update) ExecX(ctx context.Context) {
	if err := fsc.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fsc *FnetSubCategory2Update) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fnetsubcategory2.Table,
			Columns: fnetsubcategory2.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fnetsubcategory2.FieldID,
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
	if fsc.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory2.DocumentsTable,
			Columns: []string{fnetsubcategory2.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fnetdocument.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsc.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !fsc.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory2.DocumentsTable,
			Columns: []string{fnetsubcategory2.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fnetdocument.FieldID,
				},
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
			Table:   fnetsubcategory2.DocumentsTable,
			Columns: []string{fnetsubcategory2.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fnetdocument.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fsc.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fnetsubcategory2.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FnetSubCategory2UpdateOne is the builder for updating a single FnetSubCategory2 entity.
type FnetSubCategory2UpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FnetSubCategory2Mutation
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fsco *FnetSubCategory2UpdateOne) AddDocumentIDs(ids ...int) *FnetSubCategory2UpdateOne {
	fsco.mutation.AddDocumentIDs(ids...)
	return fsco
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fsco *FnetSubCategory2UpdateOne) AddDocuments(f ...*FnetDocument) *FnetSubCategory2UpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsco.AddDocumentIDs(ids...)
}

// Mutation returns the FnetSubCategory2Mutation object of the builder.
func (fsco *FnetSubCategory2UpdateOne) Mutation() *FnetSubCategory2Mutation {
	return fsco.mutation
}

// ClearDocuments clears all "documents" edges to the FnetDocument entity.
func (fsco *FnetSubCategory2UpdateOne) ClearDocuments() *FnetSubCategory2UpdateOne {
	fsco.mutation.ClearDocuments()
	return fsco
}

// RemoveDocumentIDs removes the "documents" edge to FnetDocument entities by IDs.
func (fsco *FnetSubCategory2UpdateOne) RemoveDocumentIDs(ids ...int) *FnetSubCategory2UpdateOne {
	fsco.mutation.RemoveDocumentIDs(ids...)
	return fsco
}

// RemoveDocuments removes "documents" edges to FnetDocument entities.
func (fsco *FnetSubCategory2UpdateOne) RemoveDocuments(f ...*FnetDocument) *FnetSubCategory2UpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsco.RemoveDocumentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fsco *FnetSubCategory2UpdateOne) Select(field string, fields ...string) *FnetSubCategory2UpdateOne {
	fsco.fields = append([]string{field}, fields...)
	return fsco
}

// Save executes the query and returns the updated FnetSubCategory2 entity.
func (fsco *FnetSubCategory2UpdateOne) Save(ctx context.Context) (*FnetSubCategory2, error) {
	var (
		err  error
		node *FnetSubCategory2
	)
	if len(fsco.hooks) == 0 {
		node, err = fsco.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FnetSubCategory2Mutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fsco.mutation = mutation
			node, err = fsco.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fsco.hooks) - 1; i >= 0; i-- {
			if fsco.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fsco.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fsco.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FnetSubCategory2)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FnetSubCategory2Mutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fsco *FnetSubCategory2UpdateOne) SaveX(ctx context.Context) *FnetSubCategory2 {
	node, err := fsco.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fsco *FnetSubCategory2UpdateOne) Exec(ctx context.Context) error {
	_, err := fsco.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsco *FnetSubCategory2UpdateOne) ExecX(ctx context.Context) {
	if err := fsco.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fsco *FnetSubCategory2UpdateOne) sqlSave(ctx context.Context) (_node *FnetSubCategory2, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fnetsubcategory2.Table,
			Columns: fnetsubcategory2.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fnetsubcategory2.FieldID,
			},
		},
	}
	id, ok := fsco.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FnetSubCategory2.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fsco.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fnetsubcategory2.FieldID)
		for _, f := range fields {
			if !fnetsubcategory2.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fnetsubcategory2.FieldID {
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
			Table:   fnetsubcategory2.DocumentsTable,
			Columns: []string{fnetsubcategory2.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fnetdocument.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsco.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !fsco.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fnetsubcategory2.DocumentsTable,
			Columns: []string{fnetsubcategory2.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fnetdocument.FieldID,
				},
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
			Table:   fnetsubcategory2.DocumentsTable,
			Columns: []string{fnetsubcategory2.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fnetdocument.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FnetSubCategory2{config: fsco.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fsco.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fnetsubcategory2.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}