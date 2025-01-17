package main

import (
	"fmt"
	"github.com/kumackey/kiriban/kiriban"
	"log"
	"os"
	"strconv"
)

func main() {
	prNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	c, err := kiriban.NewChecker()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(c.IsKiriban(prNumber))
}
