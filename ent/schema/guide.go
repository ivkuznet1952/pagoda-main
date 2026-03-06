package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Guide holds the schema definition for the Guide entity.
type Guide struct {
	ent.Schema
}

// Fields of the Guide.
func (Guide) Fields() []ent.Field {

	return []ent.Field{
		field.String("firstname").NotEmpty(),
		field.String("lastname").NotEmpty(),
		field.String("phone"),
		field.String("comment"),
		field.Bool("active").Default(true),
	}
}

// Edges of the Guide.
func (Guide) Edges() []ent.Edge {
	return nil
}
