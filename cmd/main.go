package main

import (
	"log"

	"github.com/betapictoris/fits/provider"
)

func main() {
	m := provider.Meters(3200)

	log.Println(m.ToMiles().ToMeters())
}
