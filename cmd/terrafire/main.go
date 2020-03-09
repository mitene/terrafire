package main

import (
	"flag"
	"fmt"
	"github.com/mitene/terrafire"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println(args)
	if len(args) < 1 {
		log.Fatalln("error!!!!")
	}

	runner := terrafire.NewRunner()

	action := args[0]
	switch action {
	case "plan":
		err := runner.Plan()
		if err != nil {
			log.Fatalln(err)
		}
	default:
	}
}
