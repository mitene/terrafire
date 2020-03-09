package terrafire

type TerraformClientMock struct {
}

func (c *TerraformClientMock) Plan(dir string) error {

	return nil
}

func (c *TerraformClientMock) Apply(dir string) error {

	return nil
}
