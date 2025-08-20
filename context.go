package sdk

type Context interface {
	Log() Logger
	Bus() Bus
	Config() map[string]interface{}
}
