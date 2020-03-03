package main

import (
	"fmt"
	"log"

	"github.com/mitene/terrafire/terrafire"
)

func main() {
	v, err := terrafire.LoadConfig("./sample")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", v)
	fmt.Println(v.Terrafire.Backend.Name)
	fmt.Println(v.Terrafire.Backend.Bucket)
	fmt.Println(v.Terrafire.Backend.Key)

	fmt.Println(v.TerraformDeploy[0].Name)
	fmt.Println(v.TerraformDeploy[0].Source.Owner)
	fmt.Println(v.TerraformDeploy[0].Vars)

	fmt.Println(v.TerraformDeploy[1].Name)
	fmt.Println(v.TerraformDeploy[1].Source.Owner)
	fmt.Println(v.TerraformDeploy[1].Vars)
}
