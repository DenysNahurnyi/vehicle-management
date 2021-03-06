package vehiclemanagement

import "errors"

// ErrorType is type of error specific to the vehicleManagement system
type ErrorType int

const (
	// StateTransition is a type of error that describes the situation when there is an attempt to change vehicle state in a wrong way
	StateTransition ErrorType = iota
	// RolePermission is a type of error that describes the situation when some role attempt to do some actions that are not supported for the current role
	RolePermission
)

// VehicleMgmtError is a error object that describe error specific to vehicleManagement system
type VehicleMgmtError struct {
	err   error
	eType ErrorType
}

// Error implement error interface
func (e *VehicleMgmtError) Error() string {
	return e.err.Error()
}

// EType retutns typo of the error
func (e *VehicleMgmtError) EType() ErrorType {
	return e.eType
}

// StateTransitionErr func creates VehicleMgmtError based on coflict of currState and desiredState
func StateTransitionErr(currState, desiredState State) *VehicleMgmtError {
	msg := "State transition is not possible, vehicle state is " + currState.String() +
		", state " + desiredState.String() + " is needed."
	return &VehicleMgmtError{
		err:   errors.New(msg),
		eType: StateTransition,
	}
}

// RolePermissionErr func creates VehicleMgmtError based on the lack of permissions of the userRole in comparison with the requiredRole
func RolePermissionErr(action string, userRole, requiredRole Role) *VehicleMgmtError {
	msg := "Attempted action " + action + " is not possible, user role " + userRole.String() +
		" is not enough, " + requiredRole.String() + " is needed."
	return &VehicleMgmtError{
		err:   errors.New(msg),
		eType: StateTransition,
	}
}
