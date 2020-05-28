package main

import "time"

type CreatedIssue struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type GroupingFields struct {
	ErrorClass string `json:"errorClass"`
	File       string `json:"file"`
	LineNumber int    `json:"lineNumber"`
}

type Error struct {
	AssignedCollaboratorID     string         `json:"assigned_collaborator_id"`
	CommentCount               int            `json:"comment_count"`
	Context                    string         `json:"context"`
	CreatedIssue               CreatedIssue   `json:"created_issue"`
	ErrorClass                 string         `json:"error_class"`
	Events                     int            `json:"events"`
	EventsURL                  string         `json:"events_url"`
	FirstSeen                  time.Time      `json:"first_seen"`
	FirstSeenUnfiltered        time.Time      `json:"first_seen_unfiltered"`
	GroupingFields             GroupingFields `json:"grouping_fields"`
	GroupingReason             string         `json:"grouping_reason"`
	ID                         string         `json:"id"`
	LastSeen                   time.Time      `json:"last_seen"`
	LastSeenUnfiltered         time.Time      `json:"last_seen_unfiltered"`
	Message                    string         `json:"message"`
	MissingDsyms               []string       `json:"missing_dsyms"`
	OriginalSeverity           string         `json:"original_severity"`
	OverriddenSeverity         interface{}    `json:"overridden_severity"`
	ProjectID                  string         `json:"project_id"`
	ProjectURL                 string         `json:"project_url"`
	ReleaseStages              []string       `json:"release_stages"`
	ReopenRules                interface{}    `json:"reopen_rules"`
	Severity                   string         `json:"severity"`
	Status                     string         `json:"status"`
	UnthrottledOccurrenceCount int            `json:"unthrottled_occurrence_count"`
	URL                        string         `json:"url"`
	Users                      int            `json:"users"`
}
