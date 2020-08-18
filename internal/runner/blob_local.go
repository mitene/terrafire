package runner

import (
	"github.com/mitene/terrafire/internal/utils"
	"io"
	"os"
	"path/filepath"
)

type BlobLocal struct {
	root string
}

func NewBlobLocal(root string) Blob {
	return &BlobLocal{
		root: root,
	}
}

func (b *BlobLocal) Get(project string, workspace string) (io.ReadCloser, error) {
	return os.Open(b.path(project, workspace))
}

func (b *BlobLocal) Put(project string, workspace string, source io.ReadSeeker) error {
	path := b.path(project, workspace)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer utils.LogDefer(f.Close)

	_, err = io.Copy(f, source)
	if err != nil {
		return err
	}

	return nil
}

func (b *BlobLocal) path(project string, workspace string) string {
	return filepath.Join(b.root, project, workspace, "artifact")
}
