// File updated by protoc-gen-ent.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return []ent.Field{field.Int64("id"), field.Int64("name")}
}
func (Group) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("groups")}
}
func (Group) Annotations() []schema.Annotation {
	return nil
}
