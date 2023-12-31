package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CreateTimeMixin adds created at time field.
type CreateTimeMixin struct{ mixin.Schema }

// Fields of the creation time mixin.
func (CreateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*CreateTimeMixin)(nil)

// UpdateTimeMixin adds updated at time field.
type UpdateTimeMixin struct{ mixin.Schema }

// Fields of the update time mixin.
func (UpdateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// update time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*UpdateTimeMixin)(nil)

// TimeMixin composes create/update time mixin.
type TimeMixin struct{ mixin.Schema }

// Fields of the time mixin.
func (TimeMixin) Fields() []ent.Field {
	return append(
		CreateTimeMixin{}.Fields(),
		UpdateTimeMixin{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*TimeMixin)(nil)
