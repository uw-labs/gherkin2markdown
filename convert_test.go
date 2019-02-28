package g2md_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uw-labs/gherkin2markdown"
)

func TestConvertFileToString(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature: Foo"))
	assert.NoError(t, err)

	_, err = g2md.ConvertFileToString(f.Name())
	assert.NoError(t, err)
}

func TestConvertFileToString_Error(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature"))
	assert.NoError(t, err)

	_, err = g2md.ConvertFileToString(f.Name())
	assert.Error(t, err)
}

func TestConvertFiles_NonReadableSourceDir(t *testing.T) {
	d, err := ioutil.TempDir("", "")
	assert.Nil(t, err)
	defer os.RemoveAll(d)

	assert.Error(t, g2md.ConvertFiles("foo", d))
}
