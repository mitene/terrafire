package terrafire

type TerraformClientMock struct {
	plan func(dir string) error
	apply func(dir string) error
}

func (c *TerraformClientMock) Plan(dir string) error {
	if c.plan != nil {
		return c.plan(dir)
	}
	return nil
}

func (c *TerraformClientMock) Apply(dir string) error {
	if c.apply != nil {
		return c.apply(dir)
	}
	return nil
}
