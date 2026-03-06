package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OrderNumber holds the schema definition for the OrderNumber entity.
type OrderNumber struct {
	ent.Schema
}

// Fields of the OrderNumber.
func (OrderNumber) Fields() []ent.Field {
	return []ent.Field{
		field.Int("num").Default(0.0),
	}
}

// Edges of the OrderNumber.
func (OrderNumber) Edges() []ent.Edge {
	return nil
}
