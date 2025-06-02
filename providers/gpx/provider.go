package main

import (
	"encoding/xml"
	"regexp"
	"time"

	"codeberg.org/betapictoris/fits/provider"
)

type Provider struct{}

var _ provider.Provider = Provider{}

// GetMetadata returns a provider.Metadata object so the host system can read
// information about the plugin.
func (p Provider) GetMetadata() provider.Metadata {
	return provider.Metadata{
		DisplayName: "Fits GPX Parser",
		Version:     "1.0.0",
		Developer: provider.DeveloperMetadata{
			DisplayName: "Daniel Hall",
			Email:       "beta@hai.haus",
			HomepageURL: "https://www.hai.haus/",
		},

		HomepageURL:   "https://codeberg.org/betapictoris/fits",
		RepositoryURL: "https://codeberg.org/betapictoris/fits",
		IssuesURL:     "https://codeberg.org/betapictoris/fits",
	}
}

// CanParse matches all files that end with `.gpx`.
func (p Provider) CanParse(filename string) bool {
	can, err := regexp.MatchString("(.*).gpx", filename)

	if !can || err != nil {
		return false
	}
	return true
}

// ParseEvent takes a raw GPX file and reads the data from it.
func (p Provider) ParseEvent(raw []byte) (*provider.Event, error) {
	var gpx GPX
	if err := xml.Unmarshal(raw, &gpx); err != nil {
		return nil, err
	}

	// Create event object
	var event provider.Event

	// Find type
	eventType := provider.EventTypeRunning

	if gpx.Track.TrackType == "running" {
		eventType = provider.EventTypeRunning
	}

	// Add metadata
	event.DisplayName = gpx.Track.DisplayName
	event.StartTime = gpx.Metadata.Timestamp
	event.Type = eventType

	event.Points = make(map[time.Time]units.Point)
	event.HeartRate = make(map[time.Time]float32)
	event.Cadence = make(map[time.Time]float32)
	event.Power = make(map[time.Time]float32)

	// Parse track points
	for _, segment := range gpx.Track.Segments {
		var p units.Point

		p.Elevation = units.Meters(segment.Elevation)
		p.Longitude = segment.Longitude
		p.Latitude = segment.Latitude

		// Save data
		event.Power[segment.Timestamp] = segment.Power
		event.HeartRate[segment.Timestamp] = segment.HeartRate
		event.Cadence[segment.Timestamp] = segment.Cadence
		event.Points[segment.Timestamp] = p
	}

	return &event, nil
}

// Command line interface.
func main() {

}
