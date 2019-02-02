package vehiclemanagement

// Ride func allow user to ride a vehicle
func Ride(user *User, vehicle *Vehicle) error {
	// Do not check role of user, everyone allowed
	// Check status of the vehicle
	if vehicle.GetState() != Ready {
		return StateTransitionErr(vehicle.GetState(), Ready)
	}
	// Actual actions
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
	// Actual actions
	vehicle.UseBattery()
	vehicle.SetState(Ready)
	return nil
}

// Collect func allow hunter to collect vehicle
func Collect(user *User, vehicle *Vehicle) error {
	neededLevel := Hunter
	requiredCurrentState := Bounty
	finalState := Collected
	attemptedAction := "collect vehicle"

	// Check user role
	err := AppropriateRoleLevel(attemptedAction, user.GetRole(), neededLevel)
	if err != nil {
		return err
	}
	// Check status of the vehicle
	if vehicle.GetState() != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(), requiredCurrentState)
	}
	// Actual actions
	vehicle.SetState(finalState)
	return nil
}

// ChargeAndDrop func allow hunter to charge collected vehicle and drop it
func ChargeAndDrop(user *User, vehicle *Vehicle) error {
	neededLevel := Hunter
	requiredCurrentState := Collected
	finalState := Dropped
	attemptedAction := "charge and drop vehicle"

	// Check user role
	err := AppropriateRoleLevel(attemptedAction, user.GetRole(), neededLevel)
	if err != nil {
		return err
	}
	// Check status of the vehicle
	if vehicle.GetState() != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(), requiredCurrentState)
	}
	// Actual actions
	vehicle.Charge()
	vehicle.SetState(finalState)
	return nil
}

// PrepareDropped func allow hunter to prepare dropped vehicle to make it ready to use
func PrepareDropped(user *User, vehicle *Vehicle) error {
	neededLevel := Hunter
	requiredCurrentState := Dropped
	finalState := Ready
	attemptedAction := "prepare dropped vehicle to make it ready"

	// Check user role
	err := AppropriateRoleLevel(attemptedAction, user.GetRole(), neededLevel)
	if err != nil {
		return err
	}
	// Check status of the vehicle
	if vehicle.GetState() != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(), requiredCurrentState)
	}
	// Actual actions
	vehicle.SetState(finalState)
	return nil
}
