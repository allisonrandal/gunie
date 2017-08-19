package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "version string")
	flag.Parse()
	if version {
		fmt.Println("Header: perly.c,v 1.0 87/12/18 15:53:31 root Exp\nPatch level: 0")
		os.Exit(0)
	}

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
	        gparser.PrintSyntaxTree()
		log.Fatal(err)
	}
	gparser.PrintSyntaxTree()

}
