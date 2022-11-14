// Code generated by ent, DO NOT EDIT.

package booking

const (
	// Label holds the string label denoting the booking type in the database.
	Label = "booking"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the booking in the database.
	Table = "bookings"
)

// Columns holds all SQL columns for booking fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
