package main

type Agent struct {
	ID        int
	AccountID int
	Token     string // token is empty at first until someone requests it to be registered, once it is set, it must
	// be reset from the control panel
	LastIP     string
	Provider   string // calculate ISP ahead of time or client side? or on backend with maxmind lookup?
	LastUpdate int64
	Location   string
	TypeID     int
	GroupID    int
}

type AgentLocation struct {
	ID        int
	AccountID int
	Name      string
	Latitude  string
	Longitude string
}

type AgentGroup struct {
	ID        int
	AccountID int
	Name      string
}

type AgentType struct {
	ID        int
	AccountID int
	Name      string
}
