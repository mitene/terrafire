package main

import (
	"log"

	"github.com/mitene/terrafire"
)

func main() {
	client := terrafire.NewGithubClient()
	err := client.GetSource("maychannel-dev", "terraform", "master", "aws/main", "out")
	if err != nil {
		log.Fatal(err)
	}
}
