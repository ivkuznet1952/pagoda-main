package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Shedule holds the schema definition for the Shedule entity.
type Shedule struct {
	ent.Schema
}

// Fields of the Shedule.
func (Shedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("resource_type"),
		field.Int("resource_id"),
		field.Time("begin").Default(nil),
		field.Time("end").Default(nil),
		field.String("comment").Default(""),
	}
}

// Edges of the Shedule.
func (Shedule) Edges() []ent.Edge {
	return nil
}
