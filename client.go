package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type client struct {
	*http.Client
	APIKey string
}

func NewClient(key string) *client {
	return &client{&http.Client{Timeout: 10 * time.Second}, key}
}

func (c *client) ListErrorsForProject(projectID string, filters *FilterParameter) []*Error {
	var errorList []*Error

	path := fmt.Sprintf(
		"/projects/%s/errors?%s",
		projectID, filters.ToURLParams())
	c.Get(path, &errorList)

	return errorList
}

func (c *client) ListProjectsForOrganization(organizationID string) []*Project {
	var projects []*Project

	path := fmt.Sprintf("/organizations/%s/projects", organizationID)
	c.Get(path, &projects)

	return projects
}

func (c *client) ListOrganizations() []*Organization {
	var orgs []*Organization
	c.Get("/user/organizations", &orgs)

	return orgs
}

func (c *client) Send(req *http.Request) *http.Response {
	req.Header.Set("Authorization", "token " + c.APIKey)
	req.Header.Set("X-Version", "2")
	response, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	return response
}

func (c *client) Get(path string, parseTo interface{}) {
	baseURL := "https://api.bugsnag.com"
	url := fmt.Sprintf("%s%s", baseURL, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	response := c.Send(req)
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(parseTo)
}

func (c *client) Post(path string, payload []byte, parseTo interface{}) {
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

	json.NewDecoder(response.Body).Decode(parseTo)
}
