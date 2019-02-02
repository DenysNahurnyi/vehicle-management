package vehiclemanagement

// Role type describe possible role of the user
type Role int

// Constants that describe role of the user
const (
	EndUser State = iota
	Hunter
	Admin
)

// User that has some role
type User struct {
	ID   string
	role Role
}

// GetRole func return current role of the user
func (u *User) GetRole() Role {
	return u.role
}
