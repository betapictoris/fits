package main

import (
	"log"

	"codeberg.org/betapictoris/fits/provider"
)

func main() {
	m := provider.Meters(3200)

	log.Println(m.ToMiles().ToMeters())
}
