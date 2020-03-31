package terrafire

import "io"

type SopsClientMock struct {
	decryptFile func(input string, output io.Writer) error
}

func (c *SopsClientMock) DecryptFile(input string, output io.Writer) error {
	if c.decryptFile != nil {
		return c.decryptFile(input, output)
	}
	return nil
}
