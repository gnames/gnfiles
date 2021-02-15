package files

type Files interface {
	SetMetaData() error
	MetaDownload() error
	MetaUpload() error
	PublishMetaData() (cid string, err error)
	Download() error
	Upload() error
}
