package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),

		field.Time("created_at").
			Default(time.Now),

		field.String("name").
			NotEmpty().
			Default("Guest"),

		field.String("email").Optional(),

		field.String("password_hash").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("city", City.Type).
			Unique().
			Required(),
		edge.To("roles", Role.Type).
			Required(),
	}
}
