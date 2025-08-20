package sdk

type Plugin interface {
	Init(Context) error
	Stop() error
}
