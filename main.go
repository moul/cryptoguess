package main // import "moul.io/cryptoguess"

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"moul.io/cryptoguess/cryptoguess"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	input := []byte{}
	for {
		chunk, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		input = append(input, chunk)
	}
	question := cryptoguess.New(input)
	fmt.Println(question)
}
