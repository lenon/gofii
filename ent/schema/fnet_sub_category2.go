package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FnetSubCategory2 holds the schema definition for the FnetSubCategory2 entity.
type FnetSubCategory2 struct {
	ent.Schema
}

// Annotations of the User.
func (FnetSubCategory2) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "fnet_sub_categories2"},
	}
}

// Fields of the FnetSubCategory2.
func (FnetSubCategory2) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().Immutable(),
	}
}

// Edges of the FnetSubCategory2.
func (FnetSubCategory2) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("documents", FnetDocument.Type).StorageKey(edge.Column("sub_category2_id")),
	}
}
