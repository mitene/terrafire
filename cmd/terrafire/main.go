package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mitene/terrafire"
)

func main() {
	applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)
	autoApprove := applyCmd.Bool("auto-approve", false, "Skip interactive approval of plan before applying.")

	if len(os.Args) < 2 {
		log.Fatalln("error!!!!")
	}

	fmt.Println(os.Args[1])

	runner := terrafire.NewRunner(
		terrafire.NewGithubClient(),
		terrafire.NewTerraformClient(),
	)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
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
	case "apply":
		applyCmd.Parse(os.Args[2:])
		err := runner.Apply(cwd, *autoApprove)
		if err != nil {
			log.Fatalln(err)
		}
	default:
	}
}
