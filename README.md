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

# Valid states

 1) Ready - The vehicle is operational and can be claimed by an enduser
 2) BatteryLow -  The vehicle is low on battery but otherwise operational. The vehicle cannot be claimed by an EndUser but can be claimed by a Hunter.
 3) Bounty - Only available for Hunter to be picked up for charging.
 4) Riding - An EndUser is currently using this vehicle; it can not be claimed by another user or Hunter.
 5) Collected - A Hunter has picked up a vehicle for charging.
 6) Dropped - A hunter has returned a vehicle after being charged.
 
*Not commisiooned for service, not claimable by either EndUsers nor Hunters.

 7) ServiceMode
 8) Terminated
 9) Unknown

# State transition scheme(with my comments)
<a href="http://tinypic.com?ref=2ln97ib" target="_blank"><img src="http://i64.tinypic.com/2ln97ib.png" border="0" alt="Image and video hosting by TinyPic"></a>
Here I modified this schene a bit to avoid some logical mistakes. In real world I would discuss that assumptions togather with team but I guess it's ok to do that for this kind of task.

 - [1] In requirements said that `A Hunter has picked up a vehicle for charging.` that means that battery level hsa to be increased to some level, for now we can always assume that it increases to 100%

 - [2] Since transition `BatteryLow -> Bounty` should happened automatically, Vehicle will never be in the BatteryLow state.

 - [3] Since transer `Ready -> Bounty` can be done automaticly after 9:30 PM then after the first day evening `Vehicle` will never be `Ready` again

 - [4] Transer `Ready -> Bounty` should be done only if the battery level of `Vehicle` is lower then 100%, in another case Hunter will be able to make cycle `Bounty -(Hunter)-> Collected -(Hunter)-> Dropped -(Hunter)-> Ready -(Automatically)-> Bounty` and abuse our system.

# Some tricky moments

 - Since the scheme of battery level change was not provided, for the sake of simplicity I will assume that every EndRide action will subtract 10% of battery level.

 - Since library has to be stateless, I can't run any timers to track when 48 hours will pass or local time, so any vehicle state manipulation will require pass localTime and can happened only when user do some actions with vehicle. Also app can run v.AutomaticStateChange() once at configurable period of time for each vehicle or GetState of ecah vehicle right after 21:30. Any of those cases are possible. 

# Test flows

 Flow#1

 1. User ride the vehicle until 21:30
 2. Vehicle automatically goes to the Bounty state
 3. Hunder do the chain: Collect -> ChargeAndDrop -> PrepareDropped 
 4. Admin terminate vehicle

  Flow#2

 1. User ride the vehicle until battery level goes lover than 20%
 2. Vehicle automatically goes to the Bounty state
 3. Hunder do the chain: Collect -> ChargeAndDrop -> PrepareDropped 
 4. Admin place vehicle to the service mode

# Usage
The best documentation are tests, so you can look at them. To tun tests go to the repo directory and run `go test`

