package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GCost holds the schema definition for the GCost entity.
type GCost struct {
	ent.Schema
}

// Fields of the GLog.
func (GCost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("cost").Default(0.0),
		field.Int("trip_id").Default(0.0),
		field.Int("transport_id").Default(0.0),
	}
}

// Edges of the GLog.
func (GCost) Edges() []ent.Edge {
	return nil
}
