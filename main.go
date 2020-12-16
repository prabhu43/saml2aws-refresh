package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main() {
	app := &cli.App{
		Name: "saml2aws-refresh",
		Usage: "Automatically refresh AWS saml session",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "count",
				Usage:   "No. of times session has to be refreshed",
				Value: 1,
				DefaultText: "1",
			},
			&cli.StringFlag{
				Name:    "profile",
				Usage:   "AWS profile (partial match works if it matches exactly 1 profile)",
				Value: "",
				DefaultText: "",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			fmt.Println("count:", c.Int("count"))
			fmt.Println("profile:", c.String("profile"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
