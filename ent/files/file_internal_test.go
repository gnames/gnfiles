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

func TestPublishMetaData(t *testing.T) {
	var err error
	var cid string
	is := is.New(t)
	fcfg := Config{
		Root:    testDir,
		KeyName: "self",
		ID:      "k51qzi5uqu5dhjz3b0j9bhjp3uqblb04iy2v7pxcm7dnj07zkr1im675hu7o1x",
	}
	f := New(fcfg, exo, local)
	err = f.SetMetaData()
	is.NoErr(err) // cannot set metadata
	err = f.Dump(false)
	is.NoErr(err) // cannot dump files
	err = f.Update()
	is.NoErr(err) // cannot upload files
	cid, err = f.PublishMetaData()
	is.NoErr(err)          // cannot publish metadata
	is.True(len(cid) > 40) // did not generate CID
}
