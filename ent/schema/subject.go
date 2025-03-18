package schema

import "entgo.io/ent"

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return nil
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return nil
}
