package main

import (
	"fmt"
	"github.com/kumackey/kiriban"
	"log"
	"os"
	"strconv"
)

func main() {
	prNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	prNumber = 999

	c, err := kiriban.NewChecker()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("::set-output name=is_kiriban::%t\n", c.IsKiriban(prNumber))
}
