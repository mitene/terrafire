package terrafire

import (
	"io"
	"os"
	"os/exec"
)

type SopsClient interface {
	DecryptFile(input string, output io.Writer) error
}

type SopsClientImpl struct {
}

func NewSopsClient() SopsClient {
	return &SopsClientImpl{}
}

func (s *SopsClientImpl) DecryptFile(input string, output io.Writer) error {

	cmd := exec.Command("sops", "-d", input)

	cmd.Stdout = output
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
