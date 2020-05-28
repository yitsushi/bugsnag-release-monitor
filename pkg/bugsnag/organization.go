package bugsnag

import "time"

// Creator is a nested element of an Organization.
type Creator struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Organization represents the response from the list organization endpoint.
type Organization struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	AutoUpgrade      bool      `json:"auto_upgrade"`
	CanStartProTrial bool      `json:"can_start_pro_trial"`
	CollaboratorsURL string    `json:"collaborators_url"`
	CreatedAt        time.Time `json:"created_at"`
	Creator          Creator   `json:"creator"`
	ProjectsURL      string    `json:"projects_url"`
	Slug             string    `json:"slug"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpgradeURL       string    `json:"upgrade_url"`
}
