package vehiclemanagement

import "time"

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

// GetState func return current state of the vehicle
func (v *Vehicle) GetState() State {
	return v.state
}

// SetState func set incomming state to the vehicle
func (v *Vehicle) SetState(state State) {
	v.state = state
}
