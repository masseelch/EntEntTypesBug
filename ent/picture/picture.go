// Code generated by ent, DO NOT EDIT.

package picture

const (
	// Label holds the string label denoting the picture type in the database.
	Label = "picture"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the picture in the database.
	Table = "pictures"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_pic"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
)

// Columns holds all SQL columns for picture fields.
var Columns = []string{
	FieldID,
	FieldURL,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "picture_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
