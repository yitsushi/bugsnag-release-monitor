package main

import "time"

// Project represents the response from the list projects endpoint.
type Project struct {
	ID                    string    `json:"id"`
	Name                  string    `json:"name"`
	APIKey                string    `json:"api_key"`
	CollaboratorsCount    int       `json:"collaborators_count"`
	CreatedAt             time.Time `json:"created_at"`
	CustomEventFieldsUsed int       `json:"custom_event_fields_used"`
	DiscardedAppVersions  []string  `json:"discarded_app_versions"`
	DiscardedErrors       []string  `json:"discarded_errors"`
	ErrorsURL             string    `json:"errors_url"`
	EventsURL             string    `json:"events_url"`
	ForReviewErrorCount   int       `json:"for_review_error_count"`
	GlobalGrouping        []string  `json:"global_grouping"`
	HTMLURL               string    `json:"html_url"`
	IsFullView            bool      `json:"is_full_view"`
	Language              string    `json:"language"`
	LocationGrouping      []string  `json:"location_grouping"`
	OpenErrorCount        int       `json:"open_error_count"`
	ReleaseStages         []string  `json:"release_stages"`
	Slug                  string    `json:"slug"`
	Type                  string    `json:"type"`
	UpdatedAt             time.Time `json:"updated_at"`
	URL                   string    `json:"url"`
}
