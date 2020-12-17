package main

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"github.com/versent/saml2aws/v2/pkg/flags"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	//"github.com/versent/saml2aws/v2/pkg/awsconfig"
	"github.com/versent/saml2aws/v2/cmd/saml2aws/commands"
)

func main() {
	app := &cli.App{
		Name:  "saml2aws-refresh",
		Usage: "Automatically refresh AWS saml session",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "count",
				Usage:       "No. of times session has to be refreshed",
				Value:       1,
				DefaultText: "1",
			},
			&cli.StringFlag{
				Name:        "profile",
				Usage:       "AWS profile (partial match works if it matches exactly 1 profile)",
				Value:       "",
				DefaultText: "",
			},
		},
		Action: func(c *cli.Context) error {
			count := c.Int("count")
			inputProfile := c.String("profile")
			matchedProfiles := findProfile(inputProfile)

			fmt.Println("**********************************************")
			fmt.Println("Count:", count)
			fmt.Println("Input Profile:", inputProfile)
			fmt.Println("Matched Profiles:", matchedProfiles)
			fmt.Println("**********************************************")

			if len(matchedProfiles) == 0 {
				cli.Exit("no profiles matched!", 1)
			}

			loginWrapper := func() {
				login(matchedProfiles)
			}

			interval := 59 * time.Minute
			Schedule(loginWrapper, interval, count)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func login(profiles []string) error {
	for _, profile := range profiles {
		fmt.Printf("logging to %s ...\n", profile)
		loginExecFlags := flags.LoginExecFlags{
			CommonFlags: &flags.CommonFlags{
				IdpAccount:      profile,
				SkipPrompt:      true,
				Profile:         profile,
				DisableKeychain: false,
			},
			Force:       true,
			ExecProfile: "",
		}
		err := commands.Login(&loginExecFlags)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func findProfile(inputProfile string) []string {
	profiles := getAllProfiles()
	matchedProfiles := []string{}
	for _, profile := range profiles {
		if strings.Contains(profile, inputProfile) {
			matchedProfiles = append(matchedProfiles, profile)
		}
	}
	return matchedProfiles
}

func getAllProfiles() []string {
	homeDir, err := os.UserHomeDir()
	check(err)
	samlConfigPath := fmt.Sprintf("%s/.saml2aws", homeDir)
	samlConfigBytes, err := ioutil.ReadFile(samlConfigPath)
	check(err)
	samlConfig := string(samlConfigBytes)
	pattern := regexp.MustCompile(`\[(.*)\]`)
	matches := pattern.FindAllStringSubmatch(samlConfig, -1)
	profiles := []string{}
	for _, match := range matches {
		profiles = append(profiles, match[1])
	}
	return profiles

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Schedule calls function `f` for given number of times with given interval
func Schedule(f func(), interval time.Duration, times int) {
	// Receiving from a nil channel blocks forever
	t := time.NewTicker(interval)

	count := 0
	done := execute(f, count, times, t)
	count++
	if done {
		return
	}

	for {
		<-t.C
		done := execute(f, count, times, t)
		count++
		if done {
			return
		}
	}
}

func execute(f func(), count int, times int, t *time.Ticker) bool {
	fmt.Printf("\n=========> %s: Executing function... count: %d\n", time.Now(), count+1)
	f()
	if (count + 1) == times {
		t.Stop()
		return true
	}
	return false
}
