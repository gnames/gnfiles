package metadata

type MetaGetter interface {
	GetMetaData(source, targetPath string) (MetaFiles, error)
}
