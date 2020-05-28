package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var targetOrgName string
	var targetProjectAPIKey string
	var targetProjectName string
	var apiToken string
	var version string
	var compact bool

	var orgID string
	var projectID string

	flag.StringVar(&targetOrgName, "org-name", "","Name of the Organization")
	flag.StringVar(&targetProjectAPIKey, "project-report-api-key", "","Reporting API Key of the Project")
	flag.StringVar(&targetProjectName, "project-name", "","Name of the Project")
	flag.StringVar(&apiToken, "api-token", "","API Token (authentication)")
	flag.StringVar(&version, "release-version", "","Release version")
	flag.BoolVar(&compact, "compact", false,"Compact view")

	flag.Parse()

	if apiToken == "" {
		log.Fatalln("'api-token' is not defined.")
	}

	if version == "" {
		log.Fatalln("'release-version' is not defined.")
	}

	if targetOrgName == "" {
		log.Fatalln("'org-name' is not defined.")
	}

	if targetProjectAPIKey == "" && targetProjectName == "" {
		log.Fatalln("'project-name' or 'project-report-api-key' is not defined.")
	}

	bugsnagClient := NewClient(apiToken)

	orgs := bugsnagClient.ListOrganizations()
	for _, org := range orgs {
		if org.Name == targetOrgName {
			orgID = org.ID
			break
		}
	}
	if orgID == "" {
		fmt.Println("Organization not found!")
		return
	}

	projects := bugsnagClient.ListProjectsForOrganization(orgID)
	for _, project := range projects {
		if project.APIKey == targetProjectAPIKey {
			projectID = project.ID
			break
		}
	}
	if projectID == "" {
		fmt.Println("Project not found!")
		return
	}

	filters := NewFilterParameter()
	filters.Add("app.release_stage", "eq", "production")
	filters.Add("release.seen_in", "eq", version)
	filters.Add("event.since", "eq", "1d")
	filters.Add("error.status", "eq", "open")

	errorList := bugsnagClient.ListErrorsForProject(projectID, filters)

	for _, reported := range errorList {
		if compact {
			fmt.Printf("[%5d] %s: %s\n", reported.Events, reported.ErrorClass, reported.Context)
		} else {
			fmt.Printf(
				"[%5d] %s: %s\n--- BEGIN ---\n%s\n--- END ---\n\n",
				reported.Events, reported.ErrorClass,
				reported.Context, reported.Message)
		}
	}
}
