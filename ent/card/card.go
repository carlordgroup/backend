// Code generated by ent, DO NOT EDIT.

package card

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldCardholderName holds the string denoting the cardholder_name field in the database.
	FieldCardholderName = "cardholder_name"
	// FieldValidUntil holds the string denoting the valid_until field in the database.
	FieldValidUntil = "valid_until"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_card"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldCardholderName,
	FieldValidUntil,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"billing_card",
	"user_card",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
