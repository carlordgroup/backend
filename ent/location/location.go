// Code generated by ent, DO NOT EDIT.

package location

const (
	// Label holds the string label denoting the location type in the database.
	Label = "location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// EdgeCars holds the string denoting the cars edge name in mutations.
	EdgeCars = "cars"
	// Table holds the table name of the location in the database.
	Table = "locations"
	// CarsTable is the table that holds the cars relation/edge.
	CarsTable = "cars"
	// CarsInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarsInverseTable = "cars"
	// CarsColumn is the table column denoting the cars relation/edge.
	CarsColumn = "location_cars"
)

// Columns holds all SQL columns for location fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLatitude,
	FieldLongitude,
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
