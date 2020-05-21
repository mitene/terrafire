package executor

import (
	"os"
	"path/filepath"
)

type LocalBlob struct {
	dir string
}

func NewLocalBlob(dir string) *LocalBlob {
	return &LocalBlob{
		dir: dir,
	}
}

func (b *LocalBlob) New(project string, workspace string) (string, error) {
	d, err := b.Get(project, workspace)
	if err != nil {
		return "", err
	}

	err = os.RemoveAll(d)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(d, 0755)
	if err != nil {
		return "", err
	}

	return d, nil
}

func (b *LocalBlob) Get(project string, workspace string) (string, error) {
	return filepath.Join(b.dir, project, workspace), nil
}

func (*LocalBlob) Put(string, string) error {
	// noop
	return nil
}
