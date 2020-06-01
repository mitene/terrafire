package runner

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"path"
	"strings"
)

type BlobS3 struct {
	bucket string
	prefix string
}

func NewBlobS3(bucket string, prefix string) Blob {
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	return &BlobS3{
		bucket: bucket,
		prefix: prefix,
	}
}

func (b *BlobS3) Get(project string, workspace string) (io.ReadCloser, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.key(project, workspace)),
	})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (b *BlobS3) Put(project string, workspace string, source io.ReadSeeker) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := s3.New(sess)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:   source,
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.key(project, workspace)),
	})
	return err
}

func (b *BlobS3) key(project string, workspace string) string {
	return b.prefix + path.Join(project, workspace, "artifact")
}
