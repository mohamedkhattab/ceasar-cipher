package bruteforce

import (
	"bufio"
	"fmt"
	"os"
)

var store = make([]string, 1000)

func init() {
	f, err := os.Open("words.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		store = append(store, scanner.Text())
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
	}
}
