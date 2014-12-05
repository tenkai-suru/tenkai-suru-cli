package commands

import (
  "github.com/codegangsta/cli"
  "fmt"
)

var HealthCheck = cli.Command{
  Name:      "health-check",
  ShortName: "hc",
  Usage:     "[project name] get health status of the project",
  Action: healthCheck,
}

func healthCheck(c *cli.Context) {
  projectName := c.Args().First()
  fmt.Printf("Checking %s health.....\n", projectName)
  httpGet(projectName + "/health-check")
}
