// Code generated by ent, DO NOT EDIT.

package booking

const (
	// Label holds the string label denoting the booking type in the database.
	Label = "booking"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldReturnCarAt holds the string denoting the return_car_at field in the database.
	FieldReturnCarAt = "return_car_at"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// FieldExceedRate holds the string denoting the exceed_rate field in the database.
	FieldExceedRate = "exceed_rate"
	// FieldDeposit holds the string denoting the deposit field in the database.
	FieldDeposit = "deposit"
	// FieldFuelLevelAtBegin holds the string denoting the fuel_level_at_begin field in the database.
	FieldFuelLevelAtBegin = "fuel_level_at_begin"
	// FieldFuelLevelAtEnd holds the string denoting the fuel_level_at_end field in the database.
	FieldFuelLevelAtEnd = "fuel_level_at_end"
	// FieldMileageBegin holds the string denoting the mileage_begin field in the database.
	FieldMileageBegin = "mileage_begin"
	// FieldMileageEnd holds the string denoting the mileage_end field in the database.
	FieldMileageEnd = "mileage_end"
	// FieldBookingStatus holds the string denoting the booking_status field in the database.
	FieldBookingStatus = "booking_status"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"
	// EdgeBilling holds the string denoting the billing edge name in mutations.
	EdgeBilling = "billing"
	// Table holds the table name of the booking in the database.
	Table = "bookings"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "bookings"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "booking_user"
	// CarTable is the table that holds the car relation/edge.
	CarTable = "bookings"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "booking_car"
	// BillingTable is the table that holds the billing relation/edge.
	BillingTable = "bookings"
	// BillingInverseTable is the table name for the Billing entity.
	// It exists in this package in order to avoid circular dependency with the "billing" package.
	BillingInverseTable = "billings"
	// BillingColumn is the table column denoting the billing relation/edge.
	BillingColumn = "billing_booking"
)

// Columns holds all SQL columns for booking fields.
var Columns = []string{
	FieldID,
	FieldStartAt,
	FieldEndAt,
	FieldReturnCarAt,
	FieldRate,
	FieldExceedRate,
	FieldDeposit,
	FieldFuelLevelAtBegin,
	FieldFuelLevelAtEnd,
	FieldMileageBegin,
	FieldMileageEnd,
	FieldBookingStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "bookings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"billing_booking",
	"booking_user",
	"booking_car",
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

var (
	// DefaultRate holds the default value on creation for the "rate" field.
	DefaultRate float32
	// DefaultExceedRate holds the default value on creation for the "exceed_rate" field.
	DefaultExceedRate float32
	// DefaultDeposit holds the default value on creation for the "deposit" field.
	DefaultDeposit float32
	// DefaultBookingStatus holds the default value on creation for the "booking_status" field.
	DefaultBookingStatus string
)
