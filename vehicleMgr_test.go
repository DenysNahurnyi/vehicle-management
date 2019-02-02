package vehiclemanagement

import (
	"testing"
	"time"
)

// TestRide test whether user can ride a vehicle
func TestRide(t *testing.T) {
	idGenerator := NewGenerator()
	admin := NewUser(idGenerator.GenerateID(), Admin)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	localTime := time.Now()

	// Ready -> Riding
	err := Ride(admin, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle")
	}
	localTime = localTime.Add(time.Hour)
	// Riding -> ???
	err = Ride(admin, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, riding vehicle that is already in riding state should not be possible")
	}
	if vehicle.GetState(localTime) != Riding {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestEndRide test whether user can end it's ride
func TestEndRide(t *testing.T) {
	idGenerator := NewGenerator()
	admin := NewUser(idGenerator.GenerateID(), Admin)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	localTime := time.Now()

	// Riding -> Ready
	err := Ride(admin, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride a vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour)

	err = EndRide(admin, vehicle, localTime)
	if err != nil {
		t.Error("Failed to end the ride of the vehicle, err:", err)
	}
	// Ready -> ???
	err = EndRide(admin, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, end ride for a vehicle that is already in ready state should not be possible")
	}
	if vehicle.GetState(localTime) != Ready {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestCollect(t *testing.T) {
	idGenerator := NewGenerator()
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	vehicle := NewVehicle(idGenerator.GenerateID(), Bounty, 100)
	localTime := time.Now()

	// Bounty -> Collect
	err := Collect(hunter, vehicle, localTime)
	if err != nil {
		t.Error("Failed to collect the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour)

	// Collected -> ???
	err = Collect(hunter, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, collect for a vehicle that is already in collected state should not be possible")
	}
	if vehicle.GetState(localTime) != Collected {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestChargeAndDrop(t *testing.T) {
	idGenerator := NewGenerator()
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	vehicle := NewVehicle(idGenerator.GenerateID(), Collected, 100)
	localTime := time.Now()

	// Collected -> Dropped
	err := ChargeAndDrop(hunter, vehicle, localTime)
	if err != nil {
		t.Error("Failed to charge and drop the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour)

	// Dropped -> ???
	err = ChargeAndDrop(hunter, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, charge and drop for a vehicle that is already in dropped state should not be possible")
	}
	if vehicle.GetState(localTime) != Dropped {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestPrepareDropped(t *testing.T) {
	idGenerator := NewGenerator()
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	vehicle := NewVehicle(idGenerator.GenerateID(), Dropped, 100)
	localTime := time.Now()

	// Dropped -> Ready
	err := PrepareDropped(hunter, vehicle, localTime)
	if err != nil {
		t.Error("Failed to collect the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour)

	// Ready -> ???
	err = PrepareDropped(hunter, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, preapare dropped vehicle that is already in ready state should not be possible")
	}
	if vehicle.GetState(localTime) != Ready {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestBatteryLow test vehicle behavior with low battery level
func TestBatteryLow(t *testing.T) {
	idGenerator := NewGenerator()
	user := NewUser(idGenerator.GenerateID(), EndUser)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 29)
	localTime := time.Now()

	// Ready -> Riding
	err := Ride(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour)

	err = EndRide(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Bounty {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	// Bounty -> Riding
	err = Ride(user, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, ride vehicle that is in bounty state should not be possible")
	}
}

// TestUnknowState test vehicle behavior that was in idle state more than 48 hours
func TestUnknowState(t *testing.T) {
	idGenerator := NewGenerator()
	user := NewUser(idGenerator.GenerateID(), EndUser)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	localTime := time.Now()

	// Ready -> Riding
	err := Ride(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour)

	// Riding -> Ready
	err = EndRide(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour*48 + time.Nanosecond)

	// Ready -> Riding, 48 hours and 1 nanosecond after last action with this vehicle
	err = Ride(user, vehicle, localTime)
	if err == nil || vehicle.GetState(localTime) != Unknown {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	// Bounty -> Riding
	err = Ride(user, vehicle, localTime)
	if err == nil {
		t.Error("Vehicle state transition is broken, ride vehicle that is in bounty state should not be possible")
	}
}

// TestEveningBounty test vehicle behavior in the Evening Bounty transfer time
func TestEveningBounty(t *testing.T) {
	idGenerator := NewGenerator()
	user := NewUser(idGenerator.GenerateID(), EndUser)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	// 02/02/2019 20:00
	localTime := time.Date(2019, 2, 2, 20, 0, 0, 0, time.UTC)

	// Ready -> Riding
	err := Ride(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	localTime = localTime.Add(time.Hour*1 + time.Minute*31)

	// Riding -> Ready
	err = EndRide(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}

	if vehicle.GetState(localTime) != Bounty {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestAdminConfig test vehicle behavior as a response to admin actions
func TestAdminConfig(t *testing.T) {
	idGenerator := NewGenerator()
	user := NewUser(idGenerator.GenerateID(), Admin)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	localTime := time.Now()

	// Ready -> Unknown
	err := AdminConfig(user, vehicle, Unknown)
	if err != nil {
		t.Error("Failed to configure the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Unknown {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}

	// Unknown -> Terminated
	err = AdminConfig(user, vehicle, Terminated)
	if err != nil {
		t.Error("Failed to configure the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Terminated {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}

	// Terminated -> ServiceMode
	err = AdminConfig(user, vehicle, ServiceMode)
	if err != nil {
		t.Error("Failed to configure the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != ServiceMode {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}

	// ServiceMode -> Ready
	err = AdminConfig(user, vehicle, Ready)
	if err != nil {
		t.Error("Failed to configure the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Ready {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}
