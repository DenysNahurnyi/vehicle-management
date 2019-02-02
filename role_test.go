package vehiclemanagement

import (
	"testing"
)

func TestAppropriateRoleLevel(t *testing.T) {
	hunterLevel := Hunter
	adminLevel := Admin

	idGenerator := NewGenerator()

	user := NewUser(idGenerator.GenerateID(), EndUser)

	err := AppropriateRoleLevel("do some hunter action", user.GetRole(), hunterLevel)
	if err == nil {
		t.Error("Role permission system is wrong, required role: ", hunterLevel, " user role is: ", user.GetRole())
	}

	err = AppropriateRoleLevel("do some admin action", user.GetRole(), adminLevel)
	if err == nil {
		t.Error("Role permission system is wrong, required role: ", adminLevel, " user role is: ", user.GetRole())
	}

	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	err = AppropriateRoleLevel("do some admin action", hunter.GetRole(), adminLevel)
	if err == nil {
		t.Error("Role permission system is wrong, required role: ", adminLevel, " user role is: ", user.GetRole())
	}
}
