package terrafire

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Terrafire struct {
		Backend struct {
			Name   string `hcl:"name,label"`
			Bucket string `hcl:"bucket"`
			Key    string `hcl:"key"`
		} `hcl:"backend,block"`
	} `hcl:"terrafire,block"`
}

func DecodeFile(path string) (*Config, error) {
	var config Config
	err := hclsimple.DecodeFile(path, nil, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
