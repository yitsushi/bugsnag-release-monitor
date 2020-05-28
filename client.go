package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client holds things together.
type Client struct {
	*http.Client
	APIKey string
}

const (
	timeout = 10
)

// NewClient creates a new client with given API Token.
func NewClient(key string) *Client {
	return &Client{&http.Client{Timeout: timeout * time.Second}, key}
}

// ListErrorsForProject lists all errors for a project with defined filters.
func (c *Client) ListErrorsForProject(projectID string, filters *FilterParameter) []*Error {
	var errorList []*Error

	path := fmt.Sprintf(
		"/projects/%s/errors?%s",
		projectID, filters.ToURLParams())
	c.Get(path, &errorList)

	return errorList
}

// ListProjectsForOrganization lists all projects for an organization.
func (c *Client) ListProjectsForOrganization(organizationID string) []*Project {
	var projects []*Project

	path := fmt.Sprintf("/organizations/%s/projects", organizationID)
	c.Get(path, &projects)

	return projects
}

// ListOrganizations lists all organizations for the user (API Token).
func (c *Client) ListOrganizations() []*Organization {
	var orgs []*Organization

	c.Get("/user/organizations", &orgs)

	return orgs
}

// Send appds some extra required header information on and sends out the request.
func (c *Client) Send(req *http.Request) *http.Response {
	req.Header.Set("Authorization", "token "+c.APIKey)
	req.Header.Set("X-Version", "2")
	response, err := c.Do(req)

	if err != nil {
		panic(err)
	}

	return response
}

// Get repares and sends a GET request.
func (c *Client) Get(path string, parseTo interface{}) {
	baseURL := "https://api.bugsnag.com"
	url := fmt.Sprintf("%s%s", baseURL, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	response := c.Send(req)
	defer response.Body.Close()

	_ = json.NewDecoder(response.Body).Decode(parseTo)
}

// Post repares and sends a POST request.
//   Caution: Maybe works, maybe not.
func (c *Client) Post(path string, payload []byte, parseTo interface{}) {
	baseURL := "https://api.bugsnag.com"
	url := fmt.Sprintf("%s%s", baseURL, path)

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	response := c.Send(req)

	fmt.Println(response)
	defer response.Body.Close()

	_ = json.NewDecoder(response.Body).Decode(parseTo)
}
