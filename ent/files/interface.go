package files

type Files interface {
	SetMetaData() error
	PublishMetaData() (cid string, err error)
	Dump(force bool) error
	Update() error
}
