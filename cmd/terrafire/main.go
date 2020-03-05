package main

import "github.com/mitene/terrafire"

func main() {
	terrafire.GetSource("maychannel-dev", "terraform", "master")
}
