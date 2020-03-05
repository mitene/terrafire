package main

import (
	"github.com/mitene/terrafire"
	"log"
)

func main() {
	client := terrafire.NewGithubClient()
	buf, err := client.GetSource("maychannel-dev", "terraform", "master")
	if err != nil {
		log.Fatal(err)
	}
	terrafire.Unzip(buf, "out")
}
