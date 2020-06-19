package runner

import "github.com/stretchr/testify/mock"

type TerraformMock struct {
	mock.Mock
}

func NewTerraformMock() *TerraformMock {
	return &TerraformMock{}
}

func (m *TerraformMock) Plan(option TerraformOption, workspace string, vars []string, varfiles []string, destroy bool) ([]byte, error) {
	args := m.Called(option, workspace, vars, varfiles, destroy)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *TerraformMock) Apply(option TerraformOption, destroy bool) error {
	return m.Called(option, destroy).Error(0)
}
