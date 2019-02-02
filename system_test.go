package vehiclemanagement

import (
	"testing"
	"time"
)

// TestFlow1 test flow#1 described in README.md
func TestFlow1(t *testing.T) {
	idGenerator := NewGenerator()
	user := NewUser(idGenerator.GenerateID(), EndUser)
	hunter := NewUser(idGenerator.GenerateID(), Hunter)
	admin := NewUser(idGenerator.GenerateID(), Admin)

	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	// 02/02/2019 20:00
	localTime := time.Date(2019, 2, 2, 20, 0, 0, 0, time.UTC)

	// 1. User ride the vehicle
	err := Ride(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Riding {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	localTime = localTime.Add(time.Hour*1 + time.Minute*31)

	err = EndRide(user, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	// 2. Vehicle goes to the Bounty because now is 21:31

	// 3. Hunter does next chain:
	// -> Collect
	err = Collect(hunter, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Collected {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	// -> ChargeAndDrop
	err = ChargeAndDrop(hunter, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Dropped {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	// -> PrepareDropped
	err = PrepareDropped(hunter, vehicle, localTime)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Ready {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
	// 4. Admin terminate vehicle
	// -> Terminate
	err = AdminConfig(admin, vehicle, Terminated)
	if err != nil {
		t.Error("Failed to ride the vehicle, err:", err)
	}
	if vehicle.GetState(localTime) != Terminated {
		t.Error("Vehicle state transition is broken, vehicle state change is wrong")
	}
}
