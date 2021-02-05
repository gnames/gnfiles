package files

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/gnames/gnfiles/ent/exofs"
	"github.com/gnames/gnfiles/ent/localfs"
	"github.com/gnames/gnfiles/ent/metadata"
	"github.com/gnames/gnfiles/ent/paths"
	"github.com/gnames/gnsys"
	api "github.com/ipfs/go-ipfs-api"
)

type files struct {
	root      string
	key       api.Key
	exo       exofs.ExoFS
	local     localfs.LocalFS
	exometa   metadata.MetaFiles
	localmeta metadata.MetaFiles
}

type Config struct {
	Dir string
}

func New(
	dir string,
	key api.Key,
	exofs exofs.ExoFS,
	lfs localfs.LocalFS,
) Files {

	return &files{
		root:  dir,
		key:   key,
		exo:   exofs,
		local: lfs,
	}
}

func (f *files) SetMetaData() (err error) {
	log.Print("Getting metadata")
	ipnsPath := paths.IPNSPath(f.key.Id)
	metaPath := paths.MetaPath(f.root)

	f.localmeta, err = f.local.MetaData()

	if err != nil {
		return err
	}

	f.exometa, err = f.exo.MetaData(ipnsPath, metaPath)
	if err != nil {
		return err
	}

	if len(f.localmeta)+len(f.exometa) == 0 {
		return errors.New("No remote or local files exist")
	}

	if len(f.localmeta) == 0 {
		for k, v := range f.exometa {
			f.localmeta[k] = v
			f.localmeta[k].Action = metadata.Download
		}
		return nil
	}

	if len(f.exometa) == 0 {
		for k := range f.localmeta {
			f.localmeta[k].Action = metadata.Upload
		}
		return nil
	}

	f.localmeta = f.localmeta.Sync(f.exometa)
	return nil
}

func (f *files) PublishMetaData() (err error) {
	log.Print("Publishing metadata")
	var cid, key string
	err = f.local.SaveMetaData(f.localmeta)
	if err == nil {
		cid, err = f.exo.Add(paths.MetaPath(f.root))
	}
	if err == nil {
		key, err = f.exo.Publish(f.key.Name, cid)
	}
	if err == nil {
		f.local.SaveKey(key)
	}
	return err
}

func (f *files) Dump() error {
	log.Print("Downloading files")
	for k, v := range f.localmeta {
		if v.Action != metadata.Download || v.ID == "" {
			continue
		}
		path := paths.RootPath(f.root, k)
		gnsys.MakeDir(filepath.Dir(path))
		err := f.exo.Get(paths.IPFSPath(v.ID), "./"+path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *files) Update() error {
	log.Print("Updating files")
	for k, v := range f.localmeta {
		if v.Action != metadata.Upload {
			continue
		}
		path := paths.RootPath(f.root, k)
		id, err := f.exo.Add(path)
		if err != nil {
			return err
		}
		f.localmeta[k].Info.ID = id
	}
	return f.PublishMetaData()
}
