package files

import (
	"testing"

	"github.com/gnames/gnfiles/io/fs"
	"github.com/gnames/gnfiles/io/ipfs"
	"github.com/matryer/is"
)

const (
	testDir  = "../../testdata/test"
	shellURL = "localhost:5001"
)

var (
	exo   = ipfs.NewExoFS(shellURL)
	local = fs.NewLocalFS(testDir)
)

func TestSetMetaData(t *testing.T) {
	is := is.New(t)

	key, err := exo.KeyIPNS("self")
	is.NoErr(err) // cannot find key

	f := New(testDir, key, exo, local)
	_ = f.SetMetaData()
	// is.NoErr(err) // cannot set metadata
}

func TestPublishMetaData(t *testing.T) {
	is := is.New(t)
	key, err := exo.KeyIPNS("self")
	is.NoErr(err) // cannot find key

	f := New(testDir, key, exo, local)
	_ = f.SetMetaData()
	err = f.PublishMetaData()

	is.NoErr(err) // cannot publish metadata
}
