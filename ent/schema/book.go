package schema

import (
	"fmt"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Checks if ISBN is valid.
func validateIsbn(isbn string) error {
	matched, err := regexp.MatchString(`^[0-9X]{10,13}$`, isbn)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid isbn: %s", isbn)
	}
	return nil
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("isbn").
			MinLen(10).
			MaxLen(13).
			Validate(validateIsbn),

		field.String("name").
			NotEmpty(),

		field.Int("page_count"),

		field.Time("publication_date"),

		field.Bool("has_hardcover"),

		field.String("cover_path").
			Optional(),

		field.String("description").
			Optional(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("authors", Author.Type),
		edge.To("genres", Genre.Type),
	}
}
