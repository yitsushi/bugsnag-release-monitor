package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var (
		targetOrgName       string
		targetProjectAPIKey string
		targetProjectName   string
		apiToken            string
		version             string
		compact             bool
	)

	flag.StringVar(&targetOrgName, "org-name", "", "Name of the Organization")
	flag.StringVar(&targetProjectAPIKey, "project-report-api-key", "", "Reporting API Key of the Project")
	flag.StringVar(&targetProjectName, "project-name", "", "Name of the Project")
	flag.StringVar(&apiToken, "api-token", "", "API Token (authentication)")
	flag.StringVar(&version, "release-version", "", "Release version")
	flag.BoolVar(&compact, "compact", false, "Compact view")

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

	orgID := findOrganizationID(bugsnagClient, targetOrgName)
	if orgID == "" {
		fmt.Println("Organization not found!")
		return
	}

	projectID := findProjectID(bugsnagClient, orgID, targetProjectAPIKey, targetProjectName)
	if projectID == "" {
		fmt.Println("Project not found!")
		return
	}

	errorList := bugsnagClient.ListErrorsForProject(projectID, createFilters(version))
	generateReport(errorList, compact)
}

func findOrganizationID(bugsnagClient *Client, target string) string {
	orgs := bugsnagClient.ListOrganizations()
	for _, org := range orgs {
		if org.Name == target {
			return org.ID
		}
	}

	return ""
}

func findProjectID(bugsnagClient *Client, orgID, targetKey, targetName string) string {
	projects := bugsnagClient.ListProjectsForOrganization(orgID)
	for _, project := range projects {
		if project.APIKey == targetKey || project.Name == targetName {
			return project.ID
		}
	}

	return ""
}

func createFilters(version string) *FilterParameter {
	filters := NewFilterParameter()
	filters.Add("app.release_stage", "eq", "production")
	filters.Add("release.seen_in", "eq", version)
	filters.Add("event.since", "eq", "1d")
	filters.Add("error.status", "eq", "open")

	return filters
}

func generateReport(errorList []*Error, compact bool) {
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
