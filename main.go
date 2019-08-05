package main // import "moul.io/cryptoguess"

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"moul.io/cryptoguess/cryptoguess"
)

func main() {
	if err := guess(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func guess() error {
	flag.Parse()
	var data []byte
	var err error
	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("input must be from stdin or file")
	}

	question := cryptoguess.New(data)
	fmt.Println(question)
	return nil
}
