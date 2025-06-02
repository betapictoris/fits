package provider

import "time"

type Provider interface {
	GetMetadata() Metadata

	GetAllEvents() ([]Event, error)
	GetEventsBetween(after time.Time, before time.Time) ([]Event, error)
	GetEventByID(ID string) (*Event, error)

	CanParse(filename string) bool
	ParseEvent(raw []byte) (*Event, error)
}

// Metadata are details about the provider plugin to the host system.
type Metadata struct {
	DisplayName string
	Version     string
	Developer   DeveloperMetadata

	HomepageURL   string
	RepositoryURL string
	IssuesURL     string
}

// DeveloperMetadata is the data about the creator of the plugin.
type DeveloperMetadata struct {
	DisplayName string
	Email       string
	HomepageURL string
}
