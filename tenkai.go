package main

import (
  "github.com/codegangsta/cli"

  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = "Tenkai Suru"
  app.Usage = "To Deploy"
  app.Version = "0.0.1"

  app.Run(os.Args)
}
