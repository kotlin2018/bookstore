// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "surplus_catch", Type: field.TypeInt},
		{Name: "author", Type: field.TypeString},
		{Name: "describe", Type: field.TypeString},
		{Name: "ebook", Type: field.TypeString},
		{Name: "cover", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "order_book", Type: field.TypeInt, Nullable: true},
		{Name: "shopping_cart_book", Type: field.TypeInt, Nullable: true},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "books_orders_book",
				Columns:    []*schema.Column{BooksColumns[9]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "books_shopping_carts_book",
				Columns:    []*schema.Column{BooksColumns[10]},
				RefColumns: []*schema.Column{ShoppingCartsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "book_category", Type: field.TypeInt, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_books_category",
				Columns:    []*schema.Column{CategoriesColumns[2]},
				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"completed", "to_be_paid", "transferring"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// ShoppingCartsColumns holds the columns for the "shopping_carts" table.
	ShoppingCartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ShoppingCartsTable holds the schema information for the "shopping_carts" table.
	ShoppingCartsTable = &schema.Table{
		Name:       "shopping_carts",
		Columns:    ShoppingCartsColumns,
		PrimaryKey: []*schema.Column{ShoppingCartsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"normal", "vip", "admin"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "order_user", Type: field.TypeInt, Nullable: true},
		{Name: "shopping_cart_user", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_orders_user",
				Columns:    []*schema.Column{UsersColumns[5]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_shopping_carts_user",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{ShoppingCartsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BooksTable,
		CategoriesTable,
		OrdersTable,
		ShoppingCartsTable,
		UsersTable,
	}
)

func init() {
	BooksTable.ForeignKeys[0].RefTable = OrdersTable
	BooksTable.ForeignKeys[1].RefTable = ShoppingCartsTable
	CategoriesTable.ForeignKeys[0].RefTable = BooksTable
	UsersTable.ForeignKeys[0].RefTable = OrdersTable
	UsersTable.ForeignKeys[1].RefTable = ShoppingCartsTable
}