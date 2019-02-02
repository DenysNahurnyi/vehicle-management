package vehiclemanagement

// Ride func allow user to ride a vehicle
func Ride(user *User, vehicle *Vehicle) error {
	// Do not check role of user, everyone allowed
	// Check status of the vehicle
	if vehicle.GetState() != Ready {
		return StateTransitionErr(vehicle.GetState(), Ready)
	}
	vehicle.SetState(Riding)
	return nil
}

// EndRide func allow user to return vehicle form riding to ready state
func EndRide(user *User, vehicle *Vehicle) error {
	// Do not check role of user, everyone allowed
	// Check status of the vehicle
	if vehicle.GetState() != Riding {
		return StateTransitionErr(vehicle.GetState(), Riding)
	}
	vehicle.SetState(Ready)
	return nil
}
