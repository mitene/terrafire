package terrafire

type TerraformClientMock struct {
	plan  func(dir string, params *ConfigTerraformDeployParams) (string, error)
	apply func(dir string, params *ConfigTerraformDeployParams) error
}

func (c *TerraformClientMock) Plan(dir string, params *ConfigTerraformDeployParams) (string, error) {
	if c.plan != nil {
		return c.plan(dir, params)
	}
	return "", nil
}

func (c *TerraformClientMock) Apply(dir string, params *ConfigTerraformDeployParams) error {
	if c.apply != nil {
		return c.apply(dir, params)
	}
	return nil
}
