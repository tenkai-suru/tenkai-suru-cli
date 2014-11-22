package commands

import (
  "github.com/codegangsta/cli"
)

var Deploy = cli.Command{
  Name:      "deploy",
  ShortName: "d",
  Usage:     "[project name] deploy the project",
  Action: deploy,
}

func deploy(c *cli.Context) {
  projectName := c.Args().First()
  println("Deploying ", projectName)
  httpPost(projectName + "/deploy")
}
