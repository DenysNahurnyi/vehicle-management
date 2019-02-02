# vehicle-management
Solution that validates and handles statetransitions for an abstract vehicle. Valid transitions are defined by various business rules (details below) which are continuously evolving. Some state changes can only be achieved by certain user roles.

# Technical requirements

 1. If the state transition is not valid, the function should return a descriptive error.
 2. If the state transition is valid, the function should return a nil error.
 3. The library needs to have a reasonable performance to be used in a soft realtime API solution.
 4. The solution should include the git history.
 5. The solution must be stateless. Assume that any required state will be provided to the library.

# User roles

 - EndUser: regular app user
 - Hunter: EndUser who have signed up to be chargers of vehicles and are responsible for picking up low battery vehicles.
 - Admin: Super user who can do everything

#Valid states

 1) Ready - The vehicle is operational and can be claimed by an enduser
 2) BatteryLow -  The vehicle is low on battery but otherwise operational. The vehicle cannot be claimed by an EndUser but can be claimed by a Hunter.
 3) Bountu - Only available for Hunter to be picked up for charging.
 4) Riding - An EndUser is currently using this vehicle; it can not be claimed by another user or Hunter.
 5) Collected - A Hunter has picked up a vehicle for charging.
 6) Dropped - A hunter has returned a vehicle after being charged.
*Not commisiooned for service, not claimable by either EndUsers nor Hunters.
 7) ServiceMode
 8) Terminated
 9) Unknown