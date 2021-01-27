package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// A holds the schema definition for the A entity.
type A struct {
	ent.Schema
}

// Fields of the A.
func (A) Fields() []ent.Field {
	return []ent.Field{
		field.String("a"),
	}
}

// Edges of the A.
func (A) Edges() []ent.Edge {
	return nil
}

// Indexes of the A.
func (A) Indexes() []ent.Index {
    return []ent.Index{}
}

