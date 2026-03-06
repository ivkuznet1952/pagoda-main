package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transport holds the schema definition for the Transport entity.
type Transport struct {
	ent.Schema
}

// Fields of the Transport.
func (Transport) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description"),
		field.Int("min_count").Default(0.0),
		field.Int("max_count").Default(0.0),
		field.Bool("active").Default(true),
	}
}

// Edges of the Transport.
func (Transport) Edges() []ent.Edge {
	return nil
}
