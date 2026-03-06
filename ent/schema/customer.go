package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname").NotEmpty(),
		field.String("lastname").NotEmpty(),
		field.String("phone"),
		field.String("comment"),
		field.Bool("active").Default(true),
		field.String("email"),
		field.String("password"),
		field.Time("created").Default(time.Now()).Immutable(),
		field.Time("updated").Default(time.Now()),
		field.Int("created_by").Default(0.0),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
