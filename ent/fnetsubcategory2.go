// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/lenon/gofii/ent/fnetsubcategory2"
)

// FnetSubCategory2 is the model entity for the FnetSubCategory2 schema.
type FnetSubCategory2 struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FnetSubCategory2Query when eager-loading is set.
	Edges FnetSubCategory2Edges `json:"edges"`
}

// FnetSubCategory2Edges holds the relations/edges for other nodes in the graph.
type FnetSubCategory2Edges struct {
	// Documents holds the value of the documents edge.
	Documents []*FnetDocument `json:"documents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DocumentsOrErr returns the Documents value or an error if the edge
// was not loaded in eager-loading.
func (e FnetSubCategory2Edges) DocumentsOrErr() ([]*FnetDocument, error) {
	if e.loadedTypes[0] {
		return e.Documents, nil
	}
	return nil, &NotLoadedError{edge: "documents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FnetSubCategory2) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fnetsubcategory2.FieldID:
			values[i] = new(sql.NullInt64)
		case fnetsubcategory2.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FnetSubCategory2", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FnetSubCategory2 fields.
func (fsc *FnetSubCategory2) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fnetsubcategory2.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fsc.ID = int(value.Int64)
		case fnetsubcategory2.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fsc.Name = value.String
			}
		}
	}
	return nil
}

// QueryDocuments queries the "documents" edge of the FnetSubCategory2 entity.
func (fsc *FnetSubCategory2) QueryDocuments() *FnetDocumentQuery {
	return (&FnetSubCategory2Client{config: fsc.config}).QueryDocuments(fsc)
}

// Update returns a builder for updating this FnetSubCategory2.
// Note that you need to call FnetSubCategory2.Unwrap() before calling this method if this FnetSubCategory2
// was returned from a transaction, and the transaction was committed or rolled back.
func (fsc *FnetSubCategory2) Update() *FnetSubCategory2UpdateOne {
	return (&FnetSubCategory2Client{config: fsc.config}).UpdateOne(fsc)
}

// Unwrap unwraps the FnetSubCategory2 entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fsc *FnetSubCategory2) Unwrap() *FnetSubCategory2 {
	_tx, ok := fsc.config.driver.(*txDriver)
	if !ok {
		panic("ent: FnetSubCategory2 is not a transactional entity")
	}
	fsc.config.driver = _tx.drv
	return fsc
}

// String implements the fmt.Stringer.
func (fsc *FnetSubCategory2) String() string {
	var builder strings.Builder
	builder.WriteString("FnetSubCategory2(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fsc.ID))
	builder.WriteString("name=")
	builder.WriteString(fsc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// FnetSubCategory2s is a parsable slice of FnetSubCategory2.
type FnetSubCategory2s []*FnetSubCategory2

func (fsc FnetSubCategory2s) config(cfg config) {
	for _i := range fsc {
		fsc[_i].config = cfg
	}
}