package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println(args)
	//client := terrafire.NewGithubClient()
	//err := client.GetSource("maychannel-dev", "terraform", "terrafire-test", "aws/test", "out")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//tc := terrafire.NewTerraformClient("out")
	//err = tc.Plan()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
