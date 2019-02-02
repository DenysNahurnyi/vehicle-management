package vehiclemanagement

import "testing"

// TestRide test whether user can ride a vehicle
func TestRide(t *testing.T) {
	idGenerator := NewGenerator()

	admin := NewUser(idGenerator.GenerateID(), Admin)

	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)

	// Ready -> Riding
	err := Ride(admin, vehicle)
	if err != nil {
		t.Error("Failed to ride a vehicle")
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
