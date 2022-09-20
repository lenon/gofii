package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FnetSubCategory1 holds the schema definition for the FnetSubCategory1 entity.
type FnetSubCategory1 struct {
	ent.Schema
}

// Annotations of the User.
func (FnetSubCategory1) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "fnet_sub_categories1"},
	}
}

// Fields of the FnetSubCategory1.
func (FnetSubCategory1) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().Immutable(),
	}
}

// Edges of the FnetSubCategory1.
func (FnetSubCategory1) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("documents", FnetDocument.Type).StorageKey(edge.Column("sub_category1_id")),
	}
}
