package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.String("category_str").NotEmpty(),
		field.String("document_status").NotEmpty(),
		field.String("fund_description").NotEmpty(),
		field.Bool("high_priority"),
		field.String("market_name").Optional(),
		field.Time("reference_date"),
		field.String("reference_date_format").NotEmpty(),
		field.String("reference_date_str").NotEmpty(),
		field.String("reviewed").NotEmpty(),
		field.String("status").NotEmpty(),
		field.String("status_description").NotEmpty(),
		field.String("sub_category1_str").Optional(),
		field.String("sub_category2_str").Optional(),
		field.Time("submission_date"),
		field.String("submission_date_str").NotEmpty(),
		field.String("submission_method").NotEmpty(),
		field.String("submission_method_description").NotEmpty(),
		field.Int("version").Positive(),
	}
}

// Edges of the FnetDocument.
func (FnetDocument) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", FnetCategory.Type).Ref("documents").Unique().Required(),
		edge.From("sub_category1", FnetSubCategory1.Type).Ref("documents").Unique(),
		edge.From("sub_category2", FnetSubCategory2.Type).Ref("documents").Unique(),
	}
}
