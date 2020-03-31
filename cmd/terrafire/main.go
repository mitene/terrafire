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
		terrafire.NewSopsClient(),
	)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	action := args[0]
	switch action {
	case "plan":
		var reportType terrafire.ReportType
		switch t := os.Getenv("TERRAFIRE_REPORT"); t {
		case "github":
			reportType = terrafire.ReportTypeGithub
		case "":
			reportType = terrafire.ReportTypeNone
		default:
			log.Fatalf("invalid report type: %s", t)
		}
		err := runner.Plan(cwd, reportType)
		if err != nil {
			log.Fatalln(err)
		}
	default:
	}
}
