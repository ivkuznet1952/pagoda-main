package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Trip holds the schema definition for the Trip entity.
type Trip struct {
	ent.Schema
}

// Fields of the Trip.
func (Trip) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description"),
		field.String("comment"),
		field.Time("duration"),
		field.Bool("active").Default(true),
		field.Time("begin"),
		field.Time("end"),
		field.String("photo"),
		field.Int("type"), // 0 - car 1 - foot
	}
}

// Edges of the Trip.
func (Trip) Edges() []ent.Edge {
	return nil
}
