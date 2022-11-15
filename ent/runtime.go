// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/account"
	"carlord/ent/schema"
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
}
