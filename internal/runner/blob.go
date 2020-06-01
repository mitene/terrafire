package runner

import "io"

type Blob interface {
	Get(project string, workspace string) (io.ReadCloser, error)
	Put(project string, workspace string, source io.ReadSeeker) error
}
