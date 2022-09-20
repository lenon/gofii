package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FnetCategory holds the schema definition for the FnetCategory entity.
type FnetCategory struct {
	ent.Schema
}

// Fields of the FnetCategory.
func (FnetCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Immutable(),
	}
}

// Edges of the FnetCategory.
func (FnetCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("documents", FnetDocument.Type).StorageKey(edge.Column("category_id")),
	}
}
