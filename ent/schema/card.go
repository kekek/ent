// File updated by protoc-gen-ent.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{field.Int64("id"), field.String("imei"), field.String("os_version"), field.String("os_type"), field.Int64("user_id")}

}
func (Card) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("card").Required().Unique().Field("user_id")}

}
func (Card) Annotations() []schema.Annotation {
	return nil
}
