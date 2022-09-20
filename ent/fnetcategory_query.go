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
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
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

// Limit adds a limit step to the query.
func (fcq *FnetCategoryQuery) Limit(limit int) *FnetCategoryQuery {
	fcq.limit = &limit
	return fcq
}

// Offset adds an offset step to the query.
func (fcq *FnetCategoryQuery) Offset(offset int) *FnetCategoryQuery {
	fcq.offset = &offset
	return fcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcq *FnetCategoryQuery) Unique(unique bool) *FnetCategoryQuery {
	fcq.unique = &unique
	return fcq
}

// Order adds an order step to the query.
func (fcq *FnetCategoryQuery) Order(o ...OrderFunc) *FnetCategoryQuery {
	fcq.order = append(fcq.order, o...)
	return fcq
}

// QueryDocuments chains the current query on the "documents" edge.
func (fcq *FnetCategoryQuery) QueryDocuments() *FnetDocumentQuery {
	query := &FnetDocumentQuery{config: fcq.config}
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
	nodes, err := fcq.Limit(1).All(ctx)
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
	if ids, err = fcq.Limit(1).IDs(ctx); err != nil {
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
	nodes, err := fcq.Limit(2).All(ctx)
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
	if ids, err = fcq.Limit(2).IDs(ctx); err != nil {
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
	if err := fcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fcq.sqlAll(ctx)
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
func (fcq *FnetCategoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fcq.Select(fnetcategory.FieldID).Scan(ctx, &ids); err != nil {
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
	if err := fcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fcq.sqlCount(ctx)
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
	if err := fcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fcq.sqlExist(ctx)
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
		limit:         fcq.limit,
		offset:        fcq.offset,
		order:         append([]OrderFunc{}, fcq.order...),
		predicates:    append([]predicate.FnetCategory{}, fcq.predicates...),
		withDocuments: fcq.withDocuments.Clone(),
		// clone intermediate query.
		sql:    fcq.sql.Clone(),
		path:   fcq.path,
		unique: fcq.unique,
	}
}

// WithDocuments tells the query-builder to eager-load the nodes that are connected to
// the "documents" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FnetCategoryQuery) WithDocuments(opts ...func(*FnetDocumentQuery)) *FnetCategoryQuery {
	query := &FnetDocumentQuery{config: fcq.config}
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
	grbuild := &FnetCategoryGroupBy{config: fcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fcq.sqlQuery(ctx), nil
	}
	grbuild.label = fnetcategory.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
	fcq.fields = append(fcq.fields, fields...)
	selbuild := &FnetCategorySelect{FnetCategoryQuery: fcq}
	selbuild.label = fnetcategory.Label
	selbuild.flds, selbuild.scan = &fcq.fields, selbuild.Scan
	return selbuild
}

func (fcq *FnetCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fcq.fields {
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
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FnetCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
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
		s.Where(sql.InValues(fnetcategory.DocumentsColumn, fks...))
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
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fcq *FnetCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcq.querySpec()
	_spec.Node.Columns = fcq.fields
	if len(fcq.fields) > 0 {
		_spec.Unique = fcq.unique != nil && *fcq.unique
	}
	return sqlgraph.CountNodes(ctx, fcq.driver, _spec)
}

func (fcq *FnetCategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fcq *FnetCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fnetcategory.Table,
			Columns: fnetcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fnetcategory.FieldID,
			},
		},
		From:   fcq.sql,
		Unique: true,
	}
	if unique := fcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fcq.fields; len(fields) > 0 {
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
	if limit := fcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcq.offset; offset != nil {
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
	columns := fcq.fields
	if len(columns) == 0 {
		columns = fnetcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcq.sql != nil {
		selector = fcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcq.unique != nil && *fcq.unique {
		selector.Distinct()
	}
	for _, p := range fcq.predicates {
		p(selector)
	}
	for _, p := range fcq.order {
		p(selector)
	}
	if offset := fcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FnetCategoryGroupBy is the group-by builder for FnetCategory entities.
type FnetCategoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcgb *FnetCategoryGroupBy) Aggregate(fns ...AggregateFunc) *FnetCategoryGroupBy {
	fcgb.fns = append(fcgb.fns, fns...)
	return fcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fcgb *FnetCategoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fcgb.path(ctx)
	if err != nil {
		return err
	}
	fcgb.sql = query
	return fcgb.sqlScan(ctx, v)
}

func (fcgb *FnetCategoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fcgb.fields {
		if !fnetcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fcgb *FnetCategoryGroupBy) sqlQuery() *sql.Selector {
	selector := fcgb.sql.Select()
	aggregation := make([]string, 0, len(fcgb.fns))
	for _, fn := range fcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fcgb.fields)+len(fcgb.fns))
		for _, f := range fcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fcgb.fields...)...)
}

// FnetCategorySelect is the builder for selecting fields of FnetCategory entities.
type FnetCategorySelect struct {
	*FnetCategoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fcs *FnetCategorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := fcs.prepareQuery(ctx); err != nil {
		return err
	}
	fcs.sql = fcs.FnetCategoryQuery.sqlQuery(ctx)
	return fcs.sqlScan(ctx, v)
}

func (fcs *FnetCategorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fcs.sql.Query()
	if err := fcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
