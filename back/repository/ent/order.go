// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/repository/ent/order"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// Status holds the value of the "status" field.
	Status order.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges OrderEdges `json:"edges"`
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Book holds the value of the book edge.
	Book []*Book `json:"book,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BookOrErr returns the Book value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) BookOrErr() ([]*Book, error) {
	if e.loadedTypes[1] {
		return e.Book, nil
	}
	return nil, &NotLoadedError{edge: "book"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldAmount:
			values[i] = new(sql.NullInt64)
		case order.FieldStatus:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				o.Amount = int(value.Int64)
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = order.Status(value.String)
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				o.UpdateAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Order entity.
func (o *Order) QueryUser() *UserQuery {
	return (&OrderClient{config: o.config}).QueryUser(o)
}

// QueryBook queries the "book" edge of the Order entity.
func (o *Order) QueryBook() *BookQuery {
	return (&OrderClient{config: o.config}).QueryBook(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", o.Amount))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(o.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}