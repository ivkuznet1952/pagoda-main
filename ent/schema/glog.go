package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GLog holds the schema definition for the GLog entity.
type GLog struct {
	ent.Schema
}

// Fields of the GLog.
func (GLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("action").NotEmpty(),
		field.String("comment"),
		field.Time("created").Default(time.Now()).Immutable(),
		field.Int("created_by").Default(0.0),
	}
}

// Edges of the GLog.
func (GLog) Edges() []ent.Edge {
	return nil
}
