package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/exercism/cli/api"
	"github.com/exercism/cli/config"
	"github.com/exercism/cli/user"
)

// Restore returns a user's solved problems.
func Restore(ctx *cli.Context) {
	c, err := config.New(ctx.GlobalString("config"))
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewClient(c)

	problems, err := client.Restore()
	if err != nil {
		log.Fatal(err)
	}

	hw := user.NewHomework(problems, c)
	if err := hw.Save(); err != nil {
		log.Fatal(err)
	}

	hw.Summarize(user.HWNotSubmitted)

	if len(hw.Errors) != 0 {
		fmt.Println("There was errors saving files:")
		for _, err := range hw.Errors {
			fmt.Printf("- %s:\n", err)
			for _, filePath := range err.FilePaths {
				fmt.Printf("\t * %s\n", filePath)
			}
		}

		fmt.Printf("\nIf you wish to override these files re-run the restore command using the --force flag\n")
	}
}
