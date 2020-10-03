package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "fuck-cli"
	app.Usage = "Unix time transform tool"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		fmt.Println(time.Now().Unix())
		fmt.Println(time.Now().Format("2006-1-2 15:04:00"))
		return nil
	}

	app.Run(os.Args)
}
