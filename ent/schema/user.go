// File updated by protoc-gen-ent.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.Int64("id"), field.String("account"), field.String("password"), field.String("name")}

}
func (User) Edges() []ent.Edge {
	return []ent.Edge{edge.To("card", Card.Type)}
}
func (User) Annotations() []schema.Annotation {
	return nil
}
