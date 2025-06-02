package main

import (
	"encoding/xml"
	"time"
)

type GPX struct {
	XMLName  xml.Name `xml:"gpx"`
	Metadata struct {
		Timestamp time.Time `xml:"time"`
	} `xml:"metadata"`
	Track struct {
		DisplayName string `xml:"name"`
		TrackType   string `xml:"type"`
		Segments    []struct {
			Latitude  float32   `xml:"lat,attr"`
			Longitude float32   `xml:"lon,attr"`
			Elevation float32   `xml:"ele"`
			Timestamp time.Time `xml:"time"`
			Power     float32   `xml:"extensions>power"`
			HeartRate float32   `xml:"extensions>gpx:TrackPointExtension>gpxtpx:hr"`
			Cadence   float32   `xml:"extensions>gpx:TrackPointExtension>gpxtpx:cad"`
		} `xml:"trkseg>trkpt"`
	} `xml:"trk"`
}
