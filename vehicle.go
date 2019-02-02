package vehiclemanagement

import "time"

// State type describe possible state of the vehicle
type State int

// Constants that describe state of the vehicle
const (
	Ready State = iota
	BatteryLow
	Bountu
	Riding
	Collected
	Dropped
	ServiceMode
	Terminated
	Unknown
)

// Vehicle that has some state
type Vehicle struct {
	ID        string
	updatedAt time.Time
	state     State
}

// GetState func return current state of the vehicle
func (v *Vehicle) GetState() State {
	return v.state
}
