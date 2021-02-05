package fs

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gnames/gnfiles/ent/localfs"
	"github.com/gnames/gnfiles/ent/metadata"
	"github.com/gnames/gnfiles/ent/paths"
	"github.com/gnames/gnfmt"
)

type lfs struct {
	root string
}

func NewLocalFS(dir string) localfs.LocalFS {
	return &lfs{root: dir}
}

func (l *lfs) MetaData() (metadata.MetaFiles, error) {
	res := make(map[string]*metadata.MetaData)
	err := filepath.Walk(l.root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			path = l.trimPath(path)
			if info.IsDir() || strings.HasPrefix(path, "/_META") {
				return nil
			}
			name := filepath.Base(path)
			dir := filepath.Dir(path)
			res[path] = &metadata.MetaData{
				Name: name,
				Path: dir,
			}
			res[path].Size = info.Size()
			res[path].AddTime = time.Now()
			res[path].SHA, err = l.shaFile(path)
			return err
		})
	return metadata.MetaFiles(res), err
}

func (l *lfs) SaveMetaData(md metadata.MetaFiles) (err error) {
	var metaFile *os.File
	var jsonMeta []byte
	path := paths.MetaPath(l.root)

	enc := gnfmt.GNjson{Pretty: true}
	jsonMeta, err = enc.Encode(md)

	if err == nil {
		metaFile, err = os.Create(path)
		defer metaFile.Close()
	}

	if err == nil {
		_, err = metaFile.Write(jsonMeta)
	}
	return err
}

func (l *lfs) SaveKey(key string) (err error) {
	path := paths.KeyPath(l.root)
	keyFile, err := os.Create(path)
	defer keyFile.Close()

	if err == nil {
		_, err = keyFile.Write([]byte(key + "\n"))
	}

	return err
}

func (l *lfs) trimPath(path string) string {
	return path[len(l.root):]
}

func (l *lfs) shaFile(path string) (string, error) {
	path = filepath.Join(l.root, path)
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
