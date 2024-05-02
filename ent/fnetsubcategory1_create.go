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
)

// FnetSubCategory1Create is the builder for creating a FnetSubCategory1 entity.
type FnetSubCategory1Create struct {
	config
	mutation *FnetSubCategory1Mutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (fsc *FnetSubCategory1Create) SetName(s string) *FnetSubCategory1Create {
	fsc.mutation.SetName(s)
	return fsc
}

// AddDocumentIDs adds the "documents" edge to the FnetDocument entity by IDs.
func (fsc *FnetSubCategory1Create) AddDocumentIDs(ids ...int) *FnetSubCategory1Create {
	fsc.mutation.AddDocumentIDs(ids...)
	return fsc
}

// AddDocuments adds the "documents" edges to the FnetDocument entity.
func (fsc *FnetSubCategory1Create) AddDocuments(f ...*FnetDocument) *FnetSubCategory1Create {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsc.AddDocumentIDs(ids...)
}

// Mutation returns the FnetSubCategory1Mutation object of the builder.
func (fsc *FnetSubCategory1Create) Mutation() *FnetSubCategory1Mutation {
	return fsc.mutation
}

// Save creates the FnetSubCategory1 in the database.
func (fsc *FnetSubCategory1Create) Save(ctx context.Context) (*FnetSubCategory1, error) {
	return withHooks(ctx, fsc.sqlSave, fsc.mutation, fsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fsc *FnetSubCategory1Create) SaveX(ctx context.Context) *FnetSubCategory1 {
	v, err := fsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fsc *FnetSubCategory1Create) Exec(ctx context.Context) error {
	_, err := fsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FnetSubCategory1Create) ExecX(ctx context.Context) {
	if err := fsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsc *FnetSubCategory1Create) check() error {
	if _, ok := fsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FnetSubCategory1.name"`)}
	}
	if v, ok := fsc.mutation.Name(); ok {
		if err := fnetsubcategory1.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FnetSubCategory1.name": %w`, err)}
		}
	}
	return nil
}

func (fsc *FnetSubCategory1Create) sqlSave(ctx context.Context) (*FnetSubCategory1, error) {
	if err := fsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fsc.mutation.id = &_node.ID
	fsc.mutation.done = true
	return _node, nil
}

func (fsc *FnetSubCategory1Create) createSpec() (*FnetSubCategory1, *sqlgraph.CreateSpec) {
	var (
		_node = &FnetSubCategory1{config: fsc.config}
		_spec = sqlgraph.NewCreateSpec(fnetsubcategory1.Table, sqlgraph.NewFieldSpec(fnetsubcategory1.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fsc.conflict
	if value, ok := fsc.mutation.Name(); ok {
		_spec.SetField(fnetsubcategory1.FieldName, field.TypeString, value)
		_node.Name = value
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FnetSubCategory1.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FnetSubCategory1Upsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fsc *FnetSubCategory1Create) OnConflict(opts ...sql.ConflictOption) *FnetSubCategory1UpsertOne {
	fsc.conflict = opts
	return &FnetSubCategory1UpsertOne{
		create: fsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FnetSubCategory1.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fsc *FnetSubCategory1Create) OnConflictColumns(columns ...string) *FnetSubCategory1UpsertOne {
	fsc.conflict = append(fsc.conflict, sql.ConflictColumns(columns...))
	return &FnetSubCategory1UpsertOne{
		create: fsc,
	}
}

type (
	// FnetSubCategory1UpsertOne is the builder for "upsert"-ing
	//  one FnetSubCategory1 node.
	FnetSubCategory1UpsertOne struct {
		create *FnetSubCategory1Create
	}

	// FnetSubCategory1Upsert is the "OnConflict" setter.
	FnetSubCategory1Upsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.FnetSubCategory1.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FnetSubCategory1UpsertOne) UpdateNewValues() *FnetSubCategory1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(fnetsubcategory1.FieldName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FnetSubCategory1.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FnetSubCategory1UpsertOne) Ignore() *FnetSubCategory1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FnetSubCategory1UpsertOne) DoNothing() *FnetSubCategory1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FnetSubCategory1Create.OnConflict
// documentation for more info.
func (u *FnetSubCategory1UpsertOne) Update(set func(*FnetSubCategory1Upsert)) *FnetSubCategory1UpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FnetSubCategory1Upsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *FnetSubCategory1UpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FnetSubCategory1Create.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FnetSubCategory1UpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FnetSubCategory1UpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FnetSubCategory1UpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FnetSubCategory1CreateBulk is the builder for creating many FnetSubCategory1 entities in bulk.
type FnetSubCategory1CreateBulk struct {
	config
	err      error
	builders []*FnetSubCategory1Create
	conflict []sql.ConflictOption
}

// Save creates the FnetSubCategory1 entities in the database.
func (fscb *FnetSubCategory1CreateBulk) Save(ctx context.Context) ([]*FnetSubCategory1, error) {
	if fscb.err != nil {
		return nil, fscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fscb.builders))
	nodes := make([]*FnetSubCategory1, len(fscb.builders))
	mutators := make([]Mutator, len(fscb.builders))
	for i := range fscb.builders {
		func(i int, root context.Context) {
			builder := fscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FnetSubCategory1Mutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, fscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fscb *FnetSubCategory1CreateBulk) SaveX(ctx context.Context) []*FnetSubCategory1 {
	v, err := fscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fscb *FnetSubCategory1CreateBulk) Exec(ctx context.Context) error {
	_, err := fscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fscb *FnetSubCategory1CreateBulk) ExecX(ctx context.Context) {
	if err := fscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FnetSubCategory1.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FnetSubCategory1Upsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fscb *FnetSubCategory1CreateBulk) OnConflict(opts ...sql.ConflictOption) *FnetSubCategory1UpsertBulk {
	fscb.conflict = opts
	return &FnetSubCategory1UpsertBulk{
		create: fscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FnetSubCategory1.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fscb *FnetSubCategory1CreateBulk) OnConflictColumns(columns ...string) *FnetSubCategory1UpsertBulk {
	fscb.conflict = append(fscb.conflict, sql.ConflictColumns(columns...))
	return &FnetSubCategory1UpsertBulk{
		create: fscb,
	}
}

// FnetSubCategory1UpsertBulk is the builder for "upsert"-ing
// a bulk of FnetSubCategory1 nodes.
type FnetSubCategory1UpsertBulk struct {
	create *FnetSubCategory1CreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FnetSubCategory1.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FnetSubCategory1UpsertBulk) UpdateNewValues() *FnetSubCategory1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(fnetsubcategory1.FieldName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FnetSubCategory1.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FnetSubCategory1UpsertBulk) Ignore() *FnetSubCategory1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FnetSubCategory1UpsertBulk) DoNothing() *FnetSubCategory1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FnetSubCategory1CreateBulk.OnConflict
// documentation for more info.
func (u *FnetSubCategory1UpsertBulk) Update(set func(*FnetSubCategory1Upsert)) *FnetSubCategory1UpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FnetSubCategory1Upsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *FnetSubCategory1UpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FnetSubCategory1CreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FnetSubCategory1CreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FnetSubCategory1UpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
