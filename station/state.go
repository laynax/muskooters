package station

import "muskooters/user"

type State int

type Transitions struct {
	From  State
	To    State
	Roles []user.Role
}

const (
	Ready      State = iota
	BatteryLow
	Bounty
	Riding
	Collected
	Dropped
)

var (
	csRoles = []user.Role{user.Client, user.Scooter}
	sRoles  = []user.Role{user.Scooter}
	hRoles  = []user.Role{user.Hunter}
)

var graph = []Transitions{
	{From: Ready, To: Riding, Roles: csRoles},
	{From: Ready, To: Bounty, Roles: sRoles},
	{From: Riding, To: Ready, Roles: csRoles},
	{From: Riding, To: BatteryLow, Roles: sRoles},
	{From: BatteryLow, To: Bounty, Roles: sRoles},
	{From: Bounty, To: Collected, Roles: hRoles},
	{From: Collected, To: Dropped, Roles: hRoles},
	{From: Dropped, To: Ready, Roles: hRoles},
}
