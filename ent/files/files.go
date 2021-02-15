package files

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/gnames/gnfiles/ent/exofs"
	"github.com/gnames/gnfiles/ent/localfs"
	"github.com/gnames/gnfiles/ent/metadata"
	"github.com/gnames/gnfiles/ent/paths"
	"github.com/gnames/gnsys"
	api "github.com/ipfs/go-ipfs-api"
)

type Config struct {
	Root       string
	KeyName    string
	Source     string
	WithUpload bool
}

type files struct {
	root      string
	source    string
	keyWrite  *api.Key
	exo       exofs.ExoFS
	local     localfs.LocalFS
	exometa   metadata.MetaFiles
	localmeta metadata.MetaFiles
}

func New(
	cfg Config,
	efs exofs.ExoFS,
	lfs localfs.LocalFS,
) Files {
	res := &files{
		root:   cfg.Root,
		source: cfg.Source,
		exo:    efs,
		local:  lfs,
	}
	var keyWrite *api.Key
	var err error
	if cfg.KeyName != "" {
		keyWrite, err = efs.KeyIPNS(cfg.KeyName)
		if err != nil {
			log.Printf("Cannot find key '%s'", cfg.KeyName)
		}
	}

	res.keyWrite = keyWrite
	return res
}

func (f *files) SetMetaData() (err error) {
	log.Print("Getting metadata")
	var metaPath string

	metaPath = paths.MetaPath(f.root)

	f.localmeta, err = f.local.CreateMetaData()

	if err == nil && f.source != "" {
		f.exometa, err = f.exo.GetMetaData(f.source, metaPath)
	}
	return err
}

func (f *files) MetaDownload() (err error) {
	if len(f.exometa) == 0 {
		err = errors.New("metadata from IPFS is empty")
	}

	if err == nil {
		for k, v := range f.exometa {
			f.localmeta[k] = v
			f.localmeta[k].Action = metadata.Download
		}
	}
	return err
}

func (f *files) MetaUpload() error {
	// return error, because no metadata can be created
	if len(f.localmeta)+len(f.exometa) == 0 {
		return errors.New("no remote or local files exist")
	}

	// if local dir is empty, download files from IPFS
	if len(f.localmeta) == 0 {
		for k, v := range f.exometa {
			f.localmeta[k] = v
			f.localmeta[k].Action = metadata.Download
		}
		return nil
	}

	// if IPFS metadata is empty, set all files for upload
	if len(f.exometa) == 0 {
		for k := range f.localmeta {
			f.localmeta[k].Action = metadata.Upload
		}
		return nil
	}

	// if both local and IPFS metadata exist, figure out what to upload
	f.localmeta = f.localmeta.Sync(f.exometa)
	return nil
}

func (f *files) PublishMetaData() (string, error) {
	log.Print("Publishing metadata")
	var cid, key string
	var err error
	err = f.local.SaveMetaData(f.localmeta)
	if err == nil {
		cid, err = f.exo.Add(paths.MetaPath(f.root))
	}
	if err == nil && f.keyWrite != nil {
		log.Printf("Publishing to permalink using '%s' key", f.keyWrite.Name)
		log.Printf("It will take a while...")
		key, err = f.exo.Publish(f.keyWrite.Name, cid)
	}
	if err == nil {
		log.Printf("The updated metadata CID: %s", cid)
		if key != "" {
			log.Printf("Key path: '%s'", paths.IPNSPath(key))
		}
	}
	return cid, err
}

func (f *files) Download() error {
	log.Print("Downloading files")
	for k, v := range f.localmeta {
		if v.ID == "" || v.Action != metadata.Download {
			continue
		}
		path := paths.RootPath(f.root, k)
		gnsys.MakeDir(filepath.Dir(path))
		fmt.Println(path)
		err := f.exo.Get(paths.IPFSPath(v.ID), path)
		if err != nil {
			log.Printf("cannot download '%s': %v", path, err)
		}
	}
	return nil
}

func (f *files) Upload() error {
	log.Print("Updating files")
	for k, v := range f.localmeta {
		if v.Action != metadata.Upload {
			continue
		}
		path := paths.RootPath(f.root, k)
		fmt.Println(path)
		id, err := f.exo.Add(path)
		if err != nil {
			return err
		}
		f.localmeta[k].Info.ID = id
	}
	_, err := f.PublishMetaData()
	return err
}
