package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FnetDocument holds the schema definition for the FnetDocument entity.
type FnetDocument struct {
	ent.Schema
}

// Fields of the FnetDocument.
func (FnetDocument) Fields() []ent.Field {
	return []ent.Field{
		field.Int("fnet_id").Unique().Immutable().Positive(),
		field.String("additional_information").Optional(),
		field.String("document_category").NotEmpty(),
		field.String("document_status").NotEmpty(),
		field.String("document_sub_category1").Optional(),
		field.String("document_sub_category2").Optional(),
		field.String("fund_description").NotEmpty(),
		field.Bool("high_priority"),
		field.String("market_name").Optional(),
		field.String("reference_date_format").NotEmpty(),
		field.String("reference_date").NotEmpty(),
		field.String("reviewed").NotEmpty(),
		field.String("status_description").NotEmpty(),
		field.String("status").NotEmpty(),
		field.String("submission_date").NotEmpty(),
		field.String("submission_method_description").NotEmpty(),
		field.String("submission_method").NotEmpty(),
		field.Int("version").Positive(),
	}
}

// Edges of the FnetDocument.
func (FnetDocument) Edges() []ent.Edge {
	return nil
}
