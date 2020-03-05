package main

import (
	"fmt"
	"log"

	"github.com/mitene/terrafire"
)

func main() {
	client := terrafire.NewGithubClient()
	buf, err := client.GetSource("maychannel-dev", "terraform", "master")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(buf)
}
