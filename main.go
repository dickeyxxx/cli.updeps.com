package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "up"
	app.Usage = "updeps.com command line tool"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:        "doc",
			Usage:       "gets the documentation for a dep",
			ShortName:   "d",
			Description: "foobar",
			Action: func(c *cli.Context) {
				pkg := c.Args().First()
				if pkg == "" {
					cli.ShowCommandHelp(c, "doc")
				} else {
					println("documentation for:", pkg)
				}
			},
		},
	}

	app.Run(os.Args)
}
