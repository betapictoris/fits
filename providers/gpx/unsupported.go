package main

import (
	"errors"
	"time"

	"codeberg.org/betapictoris/fits/provider"
)

// --- Unsupported by the GPX provider --- //
func (p Provider) GetAllEvents() ([]provider.Event, error) {
	return nil, errors.New("unimplemented")
}

func (p Provider) GetEventsBetween(after time.Time, before time.Time) ([]provider.Event, error) {
	return nil, errors.New("unimplemented")
}

func (p Provider) GetEventByID(ID string) (*provider.Event, error) {
	return nil, errors.New("unimplemented")
}
