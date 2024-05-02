// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lenon/gofii/ent/fnetcategory"
	"github.com/lenon/gofii/ent/fnetdocument"
	"github.com/lenon/gofii/ent/predicate"
)

// FnetCategoryQuery is the builder for querying FnetCategory entities.
type FnetCategoryQuery struct {
	config
	ctx           *QueryContext
	order         []fnetcategory.OrderOption
	inters        []Interceptor
	predicates    []predicate.FnetCategory
	withDocuments *FnetDocumentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FnetCategoryQuery builder.
func (fcq *FnetCategoryQuery) Where(ps ...predicate.FnetCategory) *FnetCategoryQuery {
	fcq.predicates = append(fcq.predicates, ps...)
	return fcq
}

// Limit the number of records to be returned by this query.
func (fcq *FnetCategoryQuery) Limit(limit int) *FnetCategoryQuery {
	fcq.ctx.Limit = &limit
	return fcq
}

// Offset to start from.
func (fcq *FnetCategoryQuery) Offset(offset int) *FnetCategoryQuery {
	fcq.ctx.Offset = &offset
	return fcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcq *FnetCategoryQuery) Unique(unique bool) *FnetCategoryQuery {
	fcq.ctx.Unique = &unique
	return fcq
}

// Order specifies how the records should be ordered.
func (fcq *FnetCategoryQuery) Order(o ...fnetcategory.OrderOption) *FnetCategoryQuery {
	fcq.order = append(fcq.order, o...)
	return fcq
}

// QueryDocuments chains the current query on the "documents" edge.
func (fcq *FnetCategoryQuery) QueryDocuments() *FnetDocumentQuery {
	query := (&FnetDocumentClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fnetcategory.Table, fnetcategory.FieldID, selector),
			sqlgraph.To(fnetdocument.Table, fnetdocument.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, fnetcategory.DocumentsTable, fnetcategory.DocumentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FnetCategory entity from the query.
// Returns a *NotFoundError when no FnetCategory was found.
func (fcq *FnetCategoryQuery) First(ctx context.Context) (*FnetCategory, error) {
	nodes, err := fcq.Limit(1).All(setContextOp(ctx, fcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fnetcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcq *FnetCategoryQuery) FirstX(ctx context.Context) *FnetCategory {
	node, err := fcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FnetCategory ID from the query.
// Returns a *NotFoundError when no FnetCategory ID was found.
func (fcq *FnetCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fcq.Limit(1).IDs(setContextOp(ctx, fcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fnetcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcq *FnetCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := fcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FnetCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FnetCategory entity is found.
// Returns a *NotFoundError when no FnetCategory entities are found.
func (fcq *FnetCategoryQuery) Only(ctx context.Context) (*FnetCategory, error) {
	nodes, err := fcq.Limit(2).All(setContextOp(ctx, fcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fnetcategory.Label}
	default:
		return nil, &NotSingularError{fnetcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcq *FnetCategoryQuery) OnlyX(ctx context.Context) *FnetCategory {
	node, err := fcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FnetCategory ID in the query.
// Returns a *NotSingularError when more than one FnetCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcq *FnetCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fcq.Limit(2).IDs(setContextOp(ctx, fcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fnetcategory.Label}
	default:
		err = &NotSingularError{fnetcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcq *FnetCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := fcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FnetCategories.
func (fcq *FnetCategoryQuery) All(ctx context.Context) ([]*FnetCategory, error) {
	ctx = setContextOp(ctx, fcq.ctx, "All")
	if err := fcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FnetCategory, *FnetCategoryQuery]()
	return withInterceptors[[]*FnetCategory](ctx, fcq, qr, fcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fcq *FnetCategoryQuery) AllX(ctx context.Context) []*FnetCategory {
	nodes, err := fcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FnetCategory IDs.
func (fcq *FnetCategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fcq.ctx.Unique == nil && fcq.path != nil {
		fcq.Unique(true)
	}
	ctx = setContextOp(ctx, fcq.ctx, "IDs")
	if err = fcq.Select(fnetcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcq *FnetCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := fcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcq *FnetCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fcq.ctx, "Count")
	if err := fcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fcq, querierCount[*FnetCategoryQuery](), fcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fcq *FnetCategoryQuery) CountX(ctx context.Context) int {
	count, err := fcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcq *FnetCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fcq.ctx, "Exist")
	switch _, err := fcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fcq *FnetCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := fcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FnetCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcq *FnetCategoryQuery) Clone() *FnetCategoryQuery {
	if fcq == nil {
		return nil
	}
	return &FnetCategoryQuery{
		config:        fcq.config,
		ctx:           fcq.ctx.Clone(),
		order:         append([]fnetcategory.OrderOption{}, fcq.order...),
		inters:        append([]Interceptor{}, fcq.inters...),
		predicates:    append([]predicate.FnetCategory{}, fcq.predicates...),
		withDocuments: fcq.withDocuments.Clone(),
		// clone intermediate query.
		sql:  fcq.sql.Clone(),
		path: fcq.path,
	}
}

// WithDocuments tells the query-builder to eager-load the nodes that are connected to
// the "documents" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FnetCategoryQuery) WithDocuments(opts ...func(*FnetDocumentQuery)) *FnetCategoryQuery {
	query := (&FnetDocumentClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withDocuments = query
	return fcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FnetCategory.Query().
//		GroupBy(fnetcategory.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fcq *FnetCategoryQuery) GroupBy(field string, fields ...string) *FnetCategoryGroupBy {
	fcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FnetCategoryGroupBy{build: fcq}
	grbuild.flds = &fcq.ctx.Fields
	grbuild.label = fnetcategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.FnetCategory.Query().
//		Select(fnetcategory.FieldName).
//		Scan(ctx, &v)
func (fcq *FnetCategoryQuery) Select(fields ...string) *FnetCategorySelect {
	fcq.ctx.Fields = append(fcq.ctx.Fields, fields...)
	sbuild := &FnetCategorySelect{FnetCategoryQuery: fcq}
	sbuild.label = fnetcategory.Label
	sbuild.flds, sbuild.scan = &fcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FnetCategorySelect configured with the given aggregations.
func (fcq *FnetCategoryQuery) Aggregate(fns ...AggregateFunc) *FnetCategorySelect {
	return fcq.Select().Aggregate(fns...)
}

func (fcq *FnetCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fcq); err != nil {
				return err
			}
		}
	}
	for _, f := range fcq.ctx.Fields {
		if !fnetcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fcq.path != nil {
		prev, err := fcq.path(ctx)
		if err != nil {
			return err
		}
		fcq.sql = prev
	}
	return nil
}

func (fcq *FnetCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FnetCategory, error) {
	var (
		nodes       = []*FnetCategory{}
		_spec       = fcq.querySpec()
		loadedTypes = [1]bool{
			fcq.withDocuments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FnetCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FnetCategory{config: fcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fcq.withDocuments; query != nil {
		if err := fcq.loadDocuments(ctx, query, nodes,
			func(n *FnetCategory) { n.Edges.Documents = []*FnetDocument{} },
			func(n *FnetCategory, e *FnetDocument) { n.Edges.Documents = append(n.Edges.Documents, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fcq *FnetCategoryQuery) loadDocuments(ctx context.Context, query *FnetDocumentQuery, nodes []*FnetCategory, init func(*FnetCategory), assign func(*FnetCategory, *FnetDocument)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*FnetCategory)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.FnetDocument(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(fnetcategory.DocumentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.category_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "category_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "category_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fcq *FnetCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcq.querySpec()
	_spec.Node.Columns = fcq.ctx.Fields
	if len(fcq.ctx.Fields) > 0 {
		_spec.Unique = fcq.ctx.Unique != nil && *fcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fcq.driver, _spec)
}

func (fcq *FnetCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fnetcategory.Table, fnetcategory.Columns, sqlgraph.NewFieldSpec(fnetcategory.FieldID, field.TypeInt))
	_spec.From = fcq.sql
	if unique := fcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fcq.path != nil {
		_spec.Unique = true
	}
	if fields := fcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fnetcategory.FieldID)
		for i := range fields {
			if fields[i] != fnetcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcq *FnetCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcq.driver.Dialect())
	t1 := builder.Table(fnetcategory.Table)
	columns := fcq.ctx.Fields
	if len(columns) == 0 {
		columns = fnetcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcq.sql != nil {
		selector = fcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcq.ctx.Unique != nil && *fcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fcq.predicates {
		p(selector)
	}
	for _, p := range fcq.order {
		p(selector)
	}
	if offset := fcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FnetCategoryGroupBy is the group-by builder for FnetCategory entities.
type FnetCategoryGroupBy struct {
	selector
	build *FnetCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcgb *FnetCategoryGroupBy) Aggregate(fns ...AggregateFunc) *FnetCategoryGroupBy {
	fcgb.fns = append(fcgb.fns, fns...)
	return fcgb
}

// Scan applies the selector query and scans the result into the given value.
func (fcgb *FnetCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcgb.build.ctx, "GroupBy")
	if err := fcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FnetCategoryQuery, *FnetCategoryGroupBy](ctx, fcgb.build, fcgb, fcgb.build.inters, v)
}

func (fcgb *FnetCategoryGroupBy) sqlScan(ctx context.Context, root *FnetCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fcgb.fns))
	for _, fn := range fcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fcgb.flds)+len(fcgb.fns))
		for _, f := range *fcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FnetCategorySelect is the builder for selecting fields of FnetCategory entities.
type FnetCategorySelect struct {
	*FnetCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fcs *FnetCategorySelect) Aggregate(fns ...AggregateFunc) *FnetCategorySelect {
	fcs.fns = append(fcs.fns, fns...)
	return fcs
}

// Scan applies the selector query and scans the result into the given value.
func (fcs *FnetCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcs.ctx, "Select")
	if err := fcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FnetCategoryQuery, *FnetCategorySelect](ctx, fcs.FnetCategoryQuery, fcs, fcs.inters, v)
}

func (fcs *FnetCategorySelect) sqlScan(ctx context.Context, root *FnetCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fcs.fns))
	for _, fn := range fcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
