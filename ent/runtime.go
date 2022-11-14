// Code generated by ent, DO NOT EDIT.

package ent

import (
	"carlord/ent/schema"
	"carlord/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[2].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}