// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "level", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_categories_children",
				Columns:    []*schema.Column{CategoriesColumns[4]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "level", Type: field.TypeInt32},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "detail", Type: field.TypeString},
		{Name: "cover", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat32},
		{Name: "tags", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt32},
		{Name: "category_id", Type: field.TypeInt, Nullable: true},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "courses_categories_courses",
				Columns:    []*schema.Column{CoursesColumns[10]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SlidersColumns holds the columns for the "sliders" table.
	SlidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_by", Type: field.TypeString},
		{Name: "image_name", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString},
	}
	// SlidersTable holds the schema information for the "sliders" table.
	SlidersTable = &schema.Table{
		Name:       "sliders",
		Columns:    SlidersColumns,
		PrimaryKey: []*schema.Column{SlidersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "telephone", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		CoursesTable,
		SlidersTable,
		UsersTable,
	}
)

func init() {
	CategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	CoursesTable.ForeignKeys[0].RefTable = CategoriesTable
}