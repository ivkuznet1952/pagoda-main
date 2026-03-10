package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GLog holds the schema definition for the GLog entity.
type Cost struct {
	ent.Schema
}

// Fields of the GLog.
func (Cost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("cost").Default(0.0),
		field.Int("trip_id").Default(0.0),
		field.Int("transport_id").Default(0.0),
	}
}

// Edges of the GLog.
func (Cost) Edges() []ent.Edge {
	return nil
}
