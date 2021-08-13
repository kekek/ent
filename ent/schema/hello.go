// File updated by protoc-gen-ent.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Hello struct {
	ent.Schema
}

func (Hello) Fields() []ent.Field {
	return []ent.Field{field.Int64("id").Unique().StructTag("test-id"), field.String("name").Nillable(), field.String("email_address")}
}
func (Hello) Edges() []ent.Edge {
	return nil
}
func (Hello) Annotations() []schema.Annotation {
	return nil
}
