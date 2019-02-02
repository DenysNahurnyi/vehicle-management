package vehiclemanagement

import (
	"testing"
)

// TestRide test whether user can ride a vehicle
func TestRide(t *testing.T) {
	idGenerator := NewGenerator()
	admin := NewUser(idGenerator.GenerateID(), Admin)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)

	// Ready -> Riding
	err := Ride(admin, vehicle)
	if err != nil {
		t.Error("Failed to ride the vehicle")
	}
	// Riding -> ???
	err = Ride(admin, vehicle)
	if err == nil {
		t.Error("Vehicle state transition is broken, riding vehicle that is already in riding state should not be possible")
	}
	if vehicle.GetState() != Riding {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestEndRide test whether user can end it's ride
func TestEndRide(t *testing.T) {
	idGenerator := NewGenerator()
	admin := NewUser(idGenerator.GenerateID(), Admin)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)

	// Riding -> Ready
	err := Ride(admin, vehicle)
	if err != nil {
		t.Error("Failed to ride a vehicle, err:", err)
	}
	err = EndRide(admin, vehicle)
	if err != nil {
		t.Error("Failed to end the ride of the vehicle, err:", err)
	}
	// Ready -> ???
	err = EndRide(admin, vehicle)
	if err == nil {
		t.Error("Vehicle state transition is broken, end ride for a vehicle that is already in ready state should not be possible")
	}
	if vehicle.GetState() != Ready {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestCollect(t *testing.T) {
	idGenerator := NewGenerator()
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	vehicle := NewVehicle(idGenerator.GenerateID(), Bounty, 100)

	// Bounty -> Collect
	err := Collect(hunter, vehicle)
	if err != nil {
		t.Error("Failed to collect the vehicle, err:", err)
	}
	// Collected -> ???
	err = Collect(hunter, vehicle)
	if err == nil {
		t.Error("Vehicle state transition is broken, collect for a vehicle that is already in collected state should not be possible")
	}
	if vehicle.GetState() != Collected {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestChargeAndDrop(t *testing.T) {
	idGenerator := NewGenerator()
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	vehicle := NewVehicle(idGenerator.GenerateID(), Collected, 100)

	// Collected -> Dropped
	err := ChargeAndDrop(hunter, vehicle)
	if err != nil {
		t.Error("Failed to charge and drop the vehicle, err:", err)
	}
	// Dropped -> ???
	err = ChargeAndDrop(hunter, vehicle)
	if err == nil {
		t.Error("Vehicle state transition is broken, charge and drop for a vehicle that is already in dropped state should not be possible")
	}
	if vehicle.GetState() != Dropped {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestPrepareDropped(t *testing.T) {
	idGenerator := NewGenerator()
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	vehicle := NewVehicle(idGenerator.GenerateID(), Dropped, 100)

	// Dropped -> Ready
	err := PrepareDropped(hunter, vehicle)
	if err != nil {
		t.Error("Failed to collect the vehicle, err:", err)
	}
	// Ready -> ???
	err = PrepareDropped(hunter, vehicle)
	if err == nil {
		t.Error("Vehicle state transition is broken, preapare dropped vehicle that is already in ready state should not be possible")
	}
	if vehicle.GetState() != Ready {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}

// TestCollect test whether hunter can collect a vehicle
func TestBatteryLow(t *testing.T) {
	idGenerator := NewGenerator()
	user := NewUser(idGenerator.GenerateID(), EndUser)
	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 29)

	// Dropped -> Ready
	err := Ride(user, vehicle)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	err = EndRide(user, vehicle)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState() != Bounty {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	// Bounty -> Riding
	err = Ride(user, vehicle)
	if err == nil {
		t.Error("Vehicle state transition is broken, ride vehicle that is in bounty state should not be possible")
	}

}
