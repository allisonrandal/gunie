package main

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

func main() {
	var sourcetext string
	if len(os.Args) < 2 {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		sourcetext = string(buf)
	} else {
		buf, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Perl script \"%v\" doesn't seem to exist.\n", os.Args[1])
			log.Fatal(err)
		}
		sourcetext = string(buf)
	}

	gparser := &Gunie{Buffer: sourcetext}
	gparser.Init()
	err := gparser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	gparser.PrintSyntaxTree()

}
