package main // import "moul.io/cryptoguess"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v2"
	"moul.io/cryptoguess/cryptoguess"
)

func main() {
	app := cli.App{
		Name: "cryptoguess",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug", Aliases: []string{"D"}},
			&cli.BoolFlag{Name: "list", Aliases: []string{"l"}},
		},
		Action: guess,
	}
	if err := app.Run(os.Args); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func guess(c *cli.Context) error {
	if c.Bool("list") {
		fmt.Println("Available experiments:")
		for _, guesser := range cryptoguess.AvailableExperiments {
			tmp := guesser(nil)
			fmt.Printf("- %s\n", tmp.Name())
		}
		return nil
	}
	files := []string{}
	for _, arg := range c.Args().Slice() {
		if arg == "-" {
			arg = "/dev/stdin"
		}
		files = append(files, arg)
	}
	if len(files) == 0 {
		files = append(files, "/dev/stdin")
	}
	longest := 0
	for _, file := range files {
		if len(file) > longest {
			longest = len(file)
		}
	}

	for _, file := range files {
		var data []byte
		var err error
		switch {
		case file == "/dev/stdin":
			data, err = ioutil.ReadAll(os.Stdin)
		default:
			data, err = ioutil.ReadFile(file)
		}
		left := file + ":" + strings.Repeat(" ", longest-len(file))
		if err != nil {
			fmt.Printf("%s err: %v\n", left, err)
			continue
		}

		question := cryptoguess.New(data)
		fmt.Printf("%s %s\n", left, question.Short())
		if c.Bool("debug") {
			for _, experiment := range question.Experiments {
				fmt.Printf("- %s\n", experiment)
			}
		}
	}
	return nil
}
