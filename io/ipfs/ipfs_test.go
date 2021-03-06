package ipfs_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/gnfiles/ent/paths"
	"github.com/gnames/gnfiles/io/ipfs"
	"github.com/gnames/gnsys"
	"github.com/matryer/is"
)

const (
	keyName     = "self"
	shellURL    = "localhost:5001"
	olegCID     = "QmZC4FytJmNQwXX2GyPaJ4gbKbPsDE7VHFgACAsb71iKkT"
	testDataDir = "../../testdata"
)

var (
	testDir      = filepath.Join(testDataDir, "test")
	newDir       = filepath.Join(testDataDir, "new")
	ipfsDir      = filepath.Join(testDataDir, "test", "ipfs")
	olegFilePath = filepath.Join(ipfsDir, "oleg.txt")
	metaSrc      = filepath.Join(testDataDir, "_META.json")
)

func TestGetMetaData(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)

	metaDest := paths.MetaPath(testDir)

	md, err := efs.GetMetaData(metaSrc, metaDest)
	is.NoErr(err)        // get metadata returns error
	is.True(len(md) > 0) // metadata should not be empty
}

func TestAdd(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)

	id, err := efs.Add(olegFilePath)
	is.NoErr(err)         // cannot add file
	is.Equal(id, olegCID) // file has wrong CID

	ok, err := efs.PinExists(olegCID)
	is.NoErr(err) // pin check failed
	is.True(ok)   // file is not preserved by pin

	_, err = efs.Add("nodir")
	is.True(err != nil) // should not add unknown dir
}

func TestGet(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)
	path := filepath.Join(testDataDir, "about")

	err := efs.Get(olegCID, path)
	is.NoErr(err) // download from IPFS failed
	ok, err := gnsys.FileExists(path)
	is.NoErr(err) // file check failed
	is.True(ok)   // file doe not exist

	err = os.Remove(path)
	is.NoErr(err) // could not remove file
	ok, err = gnsys.FileExists(path)
	is.NoErr(err) // file check failed
	is.True(!ok)  // file still exists after "removal"
}

func TestUnpinPin(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)

	id, err := efs.Add(olegFilePath)
	is.Equal(id, olegCID) // could not add file to IPRS correctly
	is.NoErr(err)         // error while adding file to IPFS

	ok, err := efs.PinExists(olegCID)
	is.NoErr(err) // pin check failed
	is.True(ok)   // file is not preserved by pin

	err = efs.Unpin(olegCID)
	ok, err = efs.PinExists(olegCID)
	is.NoErr(err) // unpinning of a file produced an error
	is.True(!ok)  // file is still pinned

	err = efs.Pin(olegCID)
	ok, err = efs.PinExists(olegCID)
	is.NoErr(err) // pinning file produced an error
	is.True(ok)   // file is not preserved by pin
}

func TestPublish(t *testing.T) {
	// this is a long test, skip it unless
	// something is wrong with publishing.
	if true {
		return
	}
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)
	key, err := efs.Publish("self", olegCID)
	is.NoErr(err)                         // cannot publish
	is.True(strings.HasPrefix(key, "k5")) // did not return correct key
}

func TestKeyIPNS(t *testing.T) {
	efs := ipfs.NewExoFS(shellURL)
	is := is.New(t)

	key, err := efs.KeyIPNS(keyName)
	is.NoErr(err)                            // cannot find ipns key
	is.True(strings.HasPrefix(key.Id, "k5")) // wrong ipns key id

	key, err = efs.KeyIPNS("nokey")
	is.True(err != nil) // bad key should return error
	is.Equal(key, nil)  // bad key should be nil
}
