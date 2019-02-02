package vehiclemanagement

import "time"

// Ride func allow user to ride a vehicle
func Ride(user *User, vehicle *Vehicle, localTime time.Time) error {
	neededLevel := EndUser
	requiredCurrentState := Ready
	finalState := Riding
	attemptedAction := "ride"

	// Check user role
	err := AppropriateRoleLevel(attemptedAction, user.GetRole(), neededLevel)
	if err != nil {
		return err
	}
	// Check status of the vehicle
	if vehicle.GetState(localTime) != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(localTime), requiredCurrentState)
	}
	// Actual actions
	vehicle.SetState(finalState, localTime)
	return nil
}

// EndRide func allow user to return vehicle form riding to ready state
func EndRide(user *User, vehicle *Vehicle, localTime time.Time) error {
	neededLevel := EndUser
	requiredCurrentState := Riding
	finalState := Ready
	attemptedAction := "end the ride"

	// Check user role
	err := AppropriateRoleLevel(attemptedAction, user.GetRole(), neededLevel)
	if err != nil {
		return err
	}
	// Check status of the vehicle
	if vehicle.GetState(localTime) != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(localTime), requiredCurrentState)
	}
	// Actual actions
	vehicle.UseBattery()
	vehicle.SetState(finalState, localTime)
	return nil
}

// Collect func allow hunter to collect vehicle
func Collect(user *User, vehicle *Vehicle, localTime time.Time) error {
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
	if vehicle.GetState(localTime) != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(localTime), requiredCurrentState)
	}
	// Actual actions
	vehicle.SetState(finalState, localTime)
	return nil
}

// ChargeAndDrop func allow hunter to charge collected vehicle and drop it
func ChargeAndDrop(user *User, vehicle *Vehicle, localTime time.Time) error {
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
	if vehicle.GetState(localTime) != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(localTime), requiredCurrentState)
	}
	// Actual actions
	vehicle.Charge()
	vehicle.SetState(finalState, localTime)
	return nil
}

// PrepareDropped func allow hunter to prepare dropped vehicle to make it ready to use
func PrepareDropped(user *User, vehicle *Vehicle, localTime time.Time) error {
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
	if vehicle.GetState(localTime) != requiredCurrentState {
		return StateTransitionErr(vehicle.GetState(localTime), requiredCurrentState)
	}
	// Actual actions
	vehicle.SetState(finalState, localTime)
	return nil
}

// AdminConfig func allow admin manually set any possible state to vehicle with no previous state checks
func AdminConfig(user *User, vehicle *Vehicle, state State) error {
	neededLevel := Admin
	finalState := state
	attemptedAction := "manually set state " + state.String()

	// Check user role
	err := AppropriateRoleLevel(attemptedAction, user.GetRole(), neededLevel)
	if err != nil {
		return err
	}
	// Actual actions
	vehicle.SetStateManually(finalState)
	return nil
}
