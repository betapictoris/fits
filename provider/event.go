package provider

import (
	"time"
)

type EventType int

const (
	EventTypeRunning EventType = iota
	EventTypeCycling
	EventTypeHiking
)

// Event is a single workout or race.
type Event struct {
	DisplayName string
	Type        EventType

	StartTime time.Time

	Points    map[time.Time]Point
	HeartRate map[time.Time]float32
	Cadence   map[time.Time]float32
	Power     map[time.Time]float32
}
