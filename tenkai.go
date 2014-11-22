package main

import (
  "github.com/codegangsta/cli"
  "os"
  "github.com/tenkai-suru/tenkai-suru-cli/commands"
)

func main() {
  app := cli.NewApp()
  app.Name = "Tenkai Suru"
  app.Usage = "To Deploy"
  app.Version = "0.0.1"

  app.Commands = []cli.Command{
    commands.DeployCommand,
  }

  app.Run(os.Args)
}
