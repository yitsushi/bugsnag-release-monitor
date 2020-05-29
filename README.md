# Bugsnag Release Monitor

List errors for given version for development.


```
❯ bugsnag-release-monitor -help
Usage of bugsnag-release-monitor:
  -api-token string
        API Token (authentication)
  -compact
        Compact view
  -org-name string
        Name of the Organization
  -project-name string
        Name of the Project
  -project-report-api-key string
        Reporting API Key of the Project
  -release-version string
        Release version
  -since string
        Report since... (default "3h")
```

or with Docker:

```
docker run --rm yitsushi/bugsnag-release-monitor -help
❯ docker run --rm yitsushi/bugsnag-release-monitor -help
Usage of /github-pr-creator:
  -api-token string
        API Token (authentication)
  -compact
        Compact view
  -org-name string
        Name of the Organization
  -project-name string
        Name of the Project
  -project-report-api-key string
        Reporting API Key of the Project
  -release-version string
        Release version
  -since string
        Report since... (default "3h")
```

## How to use as a lib?

```go
package main

import "github.com/yitsushi/bugsnag-release-monitor/pkg/bugsnag"

func main() {
      bugsnagClient := bugsnag.NewClient(apiToken)

      filters := bugsnag.NewFilterParameter()
      filters.Add("app.release_stage", "eq", "production")
      filters.Add("release.seen_in", "eq", "myVersion")
      filters.Add("event.since", "eq", "3h")
      filters.Add("error.status", "eq", "open")

      errorList := bugsnagClient.ListErrorsForProject(projectID, filters)

      // do whatever you want with this list
}
```
