// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lenon/gofii/ent/fnetcategory"
)

// FnetCategory is the model entity for the FnetCategory schema.
type FnetCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FnetCategoryQuery when eager-loading is set.
	Edges        FnetCategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FnetCategoryEdges holds the relations/edges for other nodes in the graph.
type FnetCategoryEdges struct {
	// Documents holds the value of the documents edge.
	Documents []*FnetDocument `json:"documents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DocumentsOrErr returns the Documents value or an error if the edge
// was not loaded in eager-loading.
func (e FnetCategoryEdges) DocumentsOrErr() ([]*FnetDocument, error) {
	if e.loadedTypes[0] {
		return e.Documents, nil
	}
	return nil, &NotLoadedError{edge: "documents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FnetCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fnetcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case fnetcategory.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FnetCategory fields.
func (fc *FnetCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fnetcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fc.ID = int(value.Int64)
		case fnetcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fc.Name = value.String
			}
		default:
			fc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FnetCategory.
// This includes values selected through modifiers, order, etc.
func (fc *FnetCategory) Value(name string) (ent.Value, error) {
	return fc.selectValues.Get(name)
}

// QueryDocuments queries the "documents" edge of the FnetCategory entity.
func (fc *FnetCategory) QueryDocuments() *FnetDocumentQuery {
	return NewFnetCategoryClient(fc.config).QueryDocuments(fc)
}

// Update returns a builder for updating this FnetCategory.
// Note that you need to call FnetCategory.Unwrap() before calling this method if this FnetCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (fc *FnetCategory) Update() *FnetCategoryUpdateOne {
	return NewFnetCategoryClient(fc.config).UpdateOne(fc)
}

// Unwrap unwraps the FnetCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fc *FnetCategory) Unwrap() *FnetCategory {
	_tx, ok := fc.config.driver.(*txDriver)
	if !ok {
		panic("ent: FnetCategory is not a transactional entity")
	}
	fc.config.driver = _tx.drv
	return fc
}

// String implements the fmt.Stringer.
func (fc *FnetCategory) String() string {
	var builder strings.Builder
	builder.WriteString("FnetCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fc.ID))
	builder.WriteString("name=")
	builder.WriteString(fc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// FnetCategories is a parsable slice of FnetCategory.
type FnetCategories []*FnetCategory
