// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ShortLinksColumns holds the columns for the "short_links" table.
	ShortLinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "link", Type: field.TypeString, Unique: true},
	}
	// ShortLinksTable holds the schema information for the "short_links" table.
	ShortLinksTable = &schema.Table{
		Name:       "short_links",
		Columns:    ShortLinksColumns,
		PrimaryKey: []*schema.Column{ShortLinksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "shortlink_key",
				Unique:  true,
				Columns: []*schema.Column{ShortLinksColumns[3]},
			},
			{
				Name:    "shortlink_link",
				Unique:  true,
				Columns: []*schema.Column{ShortLinksColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ShortLinksTable,
	}
)

func init() {
}
