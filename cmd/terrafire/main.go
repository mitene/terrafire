package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mitene/terrafire"
)

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println(args)
	if len(args) < 1 {
		log.Fatalln("error!!!!")
	}

	runner := terrafire.NewRunner(
		terrafire.NewGithubClient(),
		terrafire.NewTerraformClient(),
	)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	action := args[0]
	switch action {
	case "plan":
		err := runner.Plan(cwd, terrafire.ReportTypeGithub)
		if err != nil {
			log.Fatalln(err)
		}
	default:
	}
}
