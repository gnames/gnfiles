package localfs

import "github.com/gnames/gnfiles/ent/metadata"

type LocalFS interface {
	CreateMetaData() (metadata.MetaFiles, error)
	SaveMetaData(metadata metadata.MetaFiles) error
}
