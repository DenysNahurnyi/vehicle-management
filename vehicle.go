package vehiclemanagement

import (
	"time"
)

// State type describe possible state of the vehicle
type State int

// Constants that describe state of the vehicle
const (
	Ready State = iota
	BatteryLow
	Bounty
	Riding
	Collected
	Dropped
	ServiceMode
	Terminated
	Unknown
)

const (
	// VehicleIdleStatePossibleHours is amount of hours after which Following the taks conditions
	VehicleIdleStatePossibleHours = time.Hour * 48
)

func (s State) String() string {
	stateStr := "State type is invalid"
	switch s {
	case Ready:
		stateStr = "Ready"
	case BatteryLow:
		stateStr = "BatteryLow"
	case Bounty:
		stateStr = "Bounty"
	case Riding:
		stateStr = "Riding"
	case Collected:
		stateStr = "Collected"
	case Dropped:
		stateStr = "Dropped"
	case ServiceMode:
		stateStr = "ServiceMode"
	case Terminated:
		stateStr = "Terminated"
	case Unknown:
		stateStr = "Unknown"
	}
	return stateStr
}

// Vehicle that has some state
type Vehicle struct {
	ID        string
	updatedAt time.Time
	state     State
	battery   uint8 // Not conventional, but for the sake of saving memory
}

// NewVehicle func creates new vehicle with specified id and state
func NewVehicle(id string, state State, battery uint8) *Vehicle {
	return &Vehicle{
		ID:        id,
		updatedAt: time.Now(),
		state:     state,
		battery:   battery,
	}
}

// AutomaticStateChange checks whether state has to be changed automaticly based on external conditions
func (v *Vehicle) AutomaticStateChange(localTime time.Time) bool {
	if v.state == Riding && v.battery < 20 {
		v.state = BatteryLow
		// Automatic state change because no condition needed...
		v.state = Bounty
		v.updatedAt = time.Now()
		return true
	}
	if v.state == Ready {
		// Check last time state change, possible v.state -> Unknown
		if v.updatedAt.Add(VehicleIdleStatePossibleHours).Before(localTime) {
			v.state = Unknown
			return true
		}
		// Check localTime && v.battery, possible -> Bounty
	}
	return false
}

// GetState func return current state of the vehicle
func (v *Vehicle) GetState(localTime time.Time) State {
	v.AutomaticStateChange(localTime)
	return v.state
}

// SetState func set incomming state to the vehicle
func (v *Vehicle) SetState(state State, localTime time.Time) {
	if !v.AutomaticStateChange(localTime) {
		v.state = state
		v.updatedAt = localTime
	}
	return
}

// Charge func set vehicle battery level to 100%
func (v *Vehicle) Charge() {
	v.battery = 100
}

// UseBattery subtracts 10% of battery level, look in READMR.md: tricky moments
func (v *Vehicle) UseBattery() {
	batteryLevel := v.battery - 10

	// it's not possible in current scheme but anyway it's important
	if batteryLevel < 0 {
		v.battery = 0
	}
	v.battery = batteryLevel
}
