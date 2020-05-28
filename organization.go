package main

import "time"

type Organization struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	AutoUpgrade      bool        `json:"auto_upgrade"`
	CanStartProTrial bool        `json:"can_start_pro_trial"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CreatedAt        time.Time   `json:"created_at"`
	Creator          interface{} `json:"creator"`
	ProTrialEndsAt   interface{} `json:"pro_trial_ends_at"`
	ProTrialFeature  bool        `json:"pro_trial_feature"`
	ProjectsURL      string      `json:"projects_url"`
	Slug             string      `json:"slug"`
	UpdatedAt        time.Time   `json:"updated_at"`
	UpgradeURL       string      `json:"upgrade_url"`
}