package models

/*

Agent entries will handle various types of data entries, we will need multiple tables for different types

*/

type SNMPEntry struct {
}

type SpeedTestEntry struct {
	// TODO
}

type TraceRouteEntry struct {
	// TODO
}

// general network test to generate graph of max jitter, max ping, avg ping, etc.

type GeneralEntry struct {
	ID         int
	AgentID    int
	Timestamp  int64
	MinLatency float64
	MaxLatency float64
	AvgLatency float64
	MinJitter  float64
	MaxJitter  float64
	AvgJitter  float64
}
