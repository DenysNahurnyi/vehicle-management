package vehiclemanagement

import (
	"testing"
	"time"
)

func TestAutomaticStateChangeLowBattery(t *testing.T) {
	expectedState := Bounty

	idGenerator := NewGenerator()
	localTime := time.Now()

	vehicle := NewVehicle(idGenerator.GenerateID(), Riding, 19)
	vState := vehicle.GetState(localTime)

	if vState != expectedState {
		t.Error("Vehicle state change is wrong, expected state: ", expectedState, " current state is: ", vState)
	}
}

func TestAutomaticStateChangeEveningBounty(t *testing.T) {
	expectedState := Bounty

	idGenerator := NewGenerator()
	// 02/02/2019 00:00
	localTime := time.Date(2019, 2, 2, 0, 0, 0, 0, time.UTC)

	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 20)
	localTime = localTime.Add(time.Hour * 22)
	vState := vehicle.GetState(localTime)

	if vState != expectedState {
		t.Error("Vehicle state change is wrong, expected state: ", expectedState, " current state is: ", vState)
	}
}

func TestAutomaticStateChangeToUnknown(t *testing.T) {
	expectedState := Unknown

	idGenerator := NewGenerator()
	localTime := time.Now()

	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	localTime = localTime.Add(time.Hour * 49)
	vState := vehicle.GetState(localTime)

	if vState != expectedState {
		t.Error("Vehicle state change is wrong, expected state: ", expectedState, " current state is: ", vState)
	}
}

func TestAutomaticStateChangeHunterAntiAbuse(t *testing.T) {
	expectedState := Ready

	idGenerator := NewGenerator()
	// 02/02/2019 00:00
	localTime := time.Date(2019, 2, 2, 0, 0, 0, 0, time.UTC)

	vehicle := NewVehicle(idGenerator.GenerateID(), Ready, 100)
	localTime = localTime.Add(time.Hour * 22)
	vState := vehicle.GetState(localTime)

	if vState != expectedState {
		t.Error("Vehicle state change is wrong, expected state: ", expectedState, " current state is: ", vState)
	}
}

func TestUseBattery(t *testing.T) {
	expectedState := Dropped

	idGenerator := NewGenerator()
	localTime := time.Now()

	vehicle := NewVehicle(idGenerator.GenerateID(), Dropped, 10)
	// v.Battery 10% -> 0%
	vehicle.UseBattery()
	// v.Battery 0% -> 0%
	vehicle.UseBattery()
	vState := vehicle.GetState(localTime)

	if vState != expectedState {
		t.Error("Vehicle state change is wrong, expected state: ", expectedState, " current state is: ", vState)
	}
}
