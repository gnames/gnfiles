package ipfs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cheekybits/is"
	"github.com/gnames/gnfiles/io/ipfs"
	"github.com/gnames/gnsys"
)

const (
	shellURL = "localhost:5001"
	ipfsCID  = "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8"
	olegCID  = "QmZC4FytJmNQwXX2GyPaJ4gbKbPsDE7VHFgACAsb71iKkT"
	testDir  = "../../testdata"
)

var (
	ipfsDir      = filepath.Join(testDir, "test", "ipfs")
	olegFilePath = filepath.Join(ipfsDir, "oleg.txt")
)

func TestNew(t *testing.T) {
	is := is.New(t)
	efs := ipfs.NewExoFS(shellURL)
	err := efs.Connect("")
	is.Nil(err)

}

func TestAdd(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)

	id, err := efs.Add(olegFilePath)
	is.Nil(err)
	is.Equal(id, olegCID)

	ok, err := efs.PinExists(olegCID)
	is.Nil(err)
	is.True(ok)

	_, err = efs.Add("nodir")
	is.NotNil(err)
}

func TestGet(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)
	path := filepath.Join(testDir, "about")

	err := efs.Get(olegCID, path)
	is.Nil(err)
	ok, err := gnsys.FileExists(path)
	is.True(ok)

	err = os.Remove(path)
	is.Nil(err)
	ok, err = gnsys.FileExists(path)
	is.Nil(err)
	is.False(ok)
}

func TestUnpinPin(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)

	id, err := efs.Add(olegFilePath)
	is.Equal(id, olegCID)
	is.Nil(err)

	ok, err := efs.PinExists(olegCID)
	is.Nil(err)
	is.True(ok)

	err = efs.Unpin(olegCID)
	ok, err = efs.PinExists(olegCID)
	is.Nil(err)
	is.False(ok)

	err = efs.Pin(olegCID)
	ok, err = efs.PinExists(olegCID)
	is.Nil(err)
	is.True(ok)
}
