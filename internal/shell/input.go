package shell

import (
	"bufio"
	"fmt"
	"os"
)

func (sh *Shell) ParseInput(ch chan string) {
	defer close(ch)

	// open input file
	f, err := os.Open(sh.input)
	if err != nil {
		panic(fmt.Errorf("error reading input file: %s", err))
	}

	// creating scanner on \n
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ch <- scanner.Text()
	}
}
