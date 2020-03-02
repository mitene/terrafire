package main

import (
	"fmt"
	"github.com/mitene/terrafire/terrafire"
)

func main() {
	v, _ := terrafire.DecodeFile("sample/main.hcl")
	fmt.Printf("%#v\n", v)
	fmt.Println(v.Terrafire.Backend.Name)
	fmt.Println(v.Terrafire.Backend.Bucket)
	fmt.Println(v.Terrafire.Backend.Key)
}
