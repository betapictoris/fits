package provider

type Meters float32

func (d Meters) ToMiles() Miles {
	return Miles(d / 1609)
}

type Miles float32

func (d Miles) ToMeters() Meters {
	return Meters(d * 1609)
}

type Point struct {
	Latitude  float32
	Longitude float32
	Elevation Meters
}
