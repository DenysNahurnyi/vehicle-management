package vehiclemanagement

// Role type describe possible role of the user
type Role int

// Constants that describe role of the user
const (
	EndUser Role = iota
	Hunter
	Admin
)

func (s Role) String() string {
	roleStr := "Role type is invalid"
	switch s {
	case EndUser:
		roleStr = "EndUser"
	case Hunter:
		roleStr = "Hunter"
	case Admin:
		roleStr = "Admin"
	}
	return roleStr
}

// User that has some role
type User struct {
	ID   string
	role Role
}

// GetRole func return current role of the user
func (u *User) GetRole() Role {
	return u.role
}

// NewUser func creates new user with specified id and role
func NewUser(id string, role Role) *User {
	return &User{
		ID:   id,
		role: role,
	}
}

// Since we have hierarchical structure of roles, we can create levels of roles where the highest can do anything the lower can
// AppropriateRoleLevel check wheter user can do actions appropriate to some neededLevel role
func AppropriateRoleLevel(action string, currentRole Role, neededLevel Role) error {
	if currentRole < neededLevel {
		RolePermissionErr(action, currentRole, neededLevel)
	}
	return nil
}
