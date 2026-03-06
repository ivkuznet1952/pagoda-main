package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GOrder holds the schema definition for the GOrder entity.
type GOrder struct {
	ent.Schema
}

// Fields of the GOrder.
func (GOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int("num").Default(0.0),
		field.Int("trip_id").Default(0.0),
		field.Int("count_person").Default(0.0),
		field.Time("day").Default(nil),
		field.Time("begin").Default(nil),
		field.Int("transport_id").Default(0.0),
		field.Int("guide_id").Default(0.0),
		field.Int("cost").Default(0.0),
		field.Int("status").Default(0.0),
		field.Int("pay_status").Default(0.0),
		field.Int("paid_sum").Default(0.0),
		field.Int("customer_id").Default(0.0),
		field.String("place"),
		field.String("comment"),
		field.Time("created").Default(time.Now()).Immutable(),
		field.Time("updated").Default(time.Now()),
		field.Int("created_by").Default(0.0),
		field.Bool("archived").Default(true),
	}
}

// Edges of the GOrder.
func (GOrder) Edges() []ent.Edge {
	return nil
}
