package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ShortLink holds the schema definition for the ShortLink entity.
type ShortLink struct {
	ent.Schema
}

// Fields of the ShortLink.
func (ShortLink) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Unique(),
		field.String("link").
			Unique(),
	}
}

// Mixin of the ShortLink.
func (ShortLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		UpdateTimeMixin{},
	}
}

// Indexes of the ShortLink
func (ShortLink) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
		index.Fields("link").Unique(),
	}
}
