package fs_test

import (
	"testing"

	"github.com/gnames/gnfiles/io/fs"
	"github.com/matryer/is"
)

const (
	testDir = "../../testdata/test"
)

func TestMetaData(t *testing.T) {
	is := is.New(t)
	l := fs.NewLocalFS(testDir)
	md, err := l.MetaData()
	is.NoErr(err)
	is.True(len(md) > 0) // empty metadata
}
