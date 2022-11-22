// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/account"
	"carlord/ent/billing"
	"carlord/ent/booking"
	"carlord/ent/car"
	"carlord/ent/schema"
	"carlord/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescIsAdmin is the schema descriptor for is_admin field.
	accountDescIsAdmin := accountFields[2].Descriptor()
	// account.DefaultIsAdmin holds the default value on creation for the is_admin field.
	account.DefaultIsAdmin = accountDescIsAdmin.Default.(bool)
	billingFields := schema.Billing{}.Fields()
	_ = billingFields
	// billingDescStatus is the schema descriptor for status field.
	billingDescStatus := billingFields[0].Descriptor()
	// billing.DefaultStatus holds the default value on creation for the status field.
	billing.DefaultStatus = billingDescStatus.Default.(string)
	// billingDescBasicCost is the schema descriptor for basic_cost field.
	billingDescBasicCost := billingFields[1].Descriptor()
	// billing.DefaultBasicCost holds the default value on creation for the basic_cost field.
	billing.DefaultBasicCost = billingDescBasicCost.Default.(float32)
	// billingDescFuelCost is the schema descriptor for fuel_cost field.
	billingDescFuelCost := billingFields[2].Descriptor()
	// billing.DefaultFuelCost holds the default value on creation for the fuel_cost field.
	billing.DefaultFuelCost = billingDescFuelCost.Default.(float32)
	// billingDescCompensation is the schema descriptor for compensation field.
	billingDescCompensation := billingFields[3].Descriptor()
	// billing.DefaultCompensation holds the default value on creation for the compensation field.
	billing.DefaultCompensation = billingDescCompensation.Default.(float32)
	// billingDescDeposit is the schema descriptor for deposit field.
	billingDescDeposit := billingFields[4].Descriptor()
	// billing.DefaultDeposit holds the default value on creation for the deposit field.
	billing.DefaultDeposit = billingDescDeposit.Default.(float32)
	bookingFields := schema.Booking{}.Fields()
	_ = bookingFields
	// bookingDescRate is the schema descriptor for rate field.
	bookingDescRate := bookingFields[3].Descriptor()
	// booking.DefaultRate holds the default value on creation for the rate field.
	booking.DefaultRate = bookingDescRate.Default.(float32)
	// bookingDescExceedRate is the schema descriptor for exceed_rate field.
	bookingDescExceedRate := bookingFields[4].Descriptor()
	// booking.DefaultExceedRate holds the default value on creation for the exceed_rate field.
	booking.DefaultExceedRate = bookingDescExceedRate.Default.(float32)
	// bookingDescDeposit is the schema descriptor for deposit field.
	bookingDescDeposit := bookingFields[5].Descriptor()
	// booking.DefaultDeposit holds the default value on creation for the deposit field.
	booking.DefaultDeposit = bookingDescDeposit.Default.(float32)
	// bookingDescBookingStatus is the schema descriptor for booking_status field.
	bookingDescBookingStatus := bookingFields[10].Descriptor()
	// booking.DefaultBookingStatus holds the default value on creation for the booking_status field.
	booking.DefaultBookingStatus = bookingDescBookingStatus.Default.(string)
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescStatus is the schema descriptor for status field.
	carDescStatus := carFields[4].Descriptor()
	// car.DefaultStatus holds the default value on creation for the status field.
	car.DefaultStatus = carDescStatus.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[0].Descriptor()
	// user.DefaultFirstName holds the default value on creation for the first_name field.
	user.DefaultFirstName = userDescFirstName.Default.(string)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[1].Descriptor()
	// user.DefaultLastName holds the default value on creation for the last_name field.
	user.DefaultLastName = userDescLastName.Default.(string)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[2].Descriptor()
	// user.DefaultAddress holds the default value on creation for the address field.
	user.DefaultAddress = userDescAddress.Default.(string)
	// userDescPostalCode is the schema descriptor for postal_code field.
	userDescPostalCode := userFields[3].Descriptor()
	// user.DefaultPostalCode holds the default value on creation for the postal_code field.
	user.DefaultPostalCode = userDescPostalCode.Default.(string)
	// userDescTel is the schema descriptor for tel field.
	userDescTel := userFields[4].Descriptor()
	// user.DefaultTel holds the default value on creation for the tel field.
	user.DefaultTel = userDescTel.Default.(string)
	// userDescDriverLicenseID is the schema descriptor for driver_license_id field.
	userDescDriverLicenseID := userFields[5].Descriptor()
	// user.DefaultDriverLicenseID holds the default value on creation for the driver_license_id field.
	user.DefaultDriverLicenseID = userDescDriverLicenseID.Default.(string)
	// userDescDriverLicenseCountry is the schema descriptor for driver_license_country field.
	userDescDriverLicenseCountry := userFields[6].Descriptor()
	// user.DefaultDriverLicenseCountry holds the default value on creation for the driver_license_country field.
	user.DefaultDriverLicenseCountry = userDescDriverLicenseCountry.Default.(string)
	// userDescBirthday is the schema descriptor for birthday field.
	userDescBirthday := userFields[7].Descriptor()
	// user.DefaultBirthday holds the default value on creation for the birthday field.
	user.DefaultBirthday = userDescBirthday.Default.(func() time.Time)
}
