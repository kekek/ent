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
	return []ent.Field{field.Int64("id"), field.String("account"), field.String("password").Immutable(), field.String("name")}

}
func (User) Edges() []ent.Edge {
	return []ent.Edge{edge.To("card", Card.Type), edge.To("groups", Group.Type).StorageKey(edge.Table("user_group"), edge.Columns("uid", "gid"))}
}
func (User) Annotations() []schema.Annotation {
	return nil
}
