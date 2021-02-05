package localfs

import "github.com/gnames/gnfiles/ent/metadata"

type LocalFS interface {
	MetaData() (metadata.MetaFiles, error)
	SaveMetaData(metadata metadata.MetaFiles) error
	SaveKey(key string) error
}
