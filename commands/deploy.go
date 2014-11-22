package commands

import (
  "github.com/codegangsta/cli"
  "net/http"
)

var DeployCommand = cli.Command{
  Name:      "deploy",
  ShortName: "d",
  Usage:     "[project name] deploy the project",
  Action: deploy,
}

func deploy(c *cli.Context) {
  projectName := c.Args().First()
  println("Deploying ", projectName)
  resp, _ := http.Get("https://tenkai-suru-api.herokuapp.com/" + projectName + "/deploy")
  println("Response body ", resp.Body)
  println(resp.Status)
}
