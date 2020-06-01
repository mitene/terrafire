package runner

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Zip_Unzip(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	err := Zip(buf, ".")
	assert.NoError(t, err)

	tmp, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer func() { _ = os.RemoveAll(tmp) }()

	err = Unzip(bytes.NewReader(buf.Bytes()), tmp)
	assert.NoError(t, err)

	d1, err := ioutil.ReadDir(".")
	assert.NoError(t, err)
	f1 := make([]string, len(d1))
	for i, d := range d1 {
		f1[i] = d.Name()
	}

	d2, err := ioutil.ReadDir(tmp)
	assert.NoError(t, err)
	f2 := make([]string, len(d2))
	for i, d := range d1 {
		f2[i] = d.Name()
	}

	assert.Equal(t, f1, f2)
}
