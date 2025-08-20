package sdk

type Meta struct {
	PluginName    string
	PluginVersion string
	SDKVersion    string
	GoVersion     string
}

type Compat interface {
	SDKCompat() string
}
