package vehiclemanagement

import "errors"

// Ride func allow user to ride a vehicle
func Ride(user *User, vehicle *Vehicle) error {
	// Do not check role of user, everyone allowed
	// Check status of vehicle
	if vehicle.GetState() != Ready {
		return errors.New("State transition is not possible, vehicle state is " + vehicle.GetState().String() + ", state " + Ready.String() + " is needed.")
	}
	vehicle.SetState(Riding)
	return nil
}
