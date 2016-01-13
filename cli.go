package main

import (
  "os"
  "runtime"
  "github.com/codegangsta/cli"
)

const (
  Version = "0.0.1"
)

var (
  app *cli.App
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  app = cli.NewApp()
  app.Name = "pochi_echo"
  app.Usage = "Echo message"
  app.Action = func(c *cli.Context) {
    println(c.Args()[0])
  }

  app.Run(os.Args)
}
