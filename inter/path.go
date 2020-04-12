package inter

type BasePath interface {
	AppPath() string
	BasePath() string
	BootstrapPath() string
	ConfigPath() string
	DatabasePath() string
	LangPath() string
	PublicPath() string
	StoragePath() string
	ResourcePath() string
	EnvironmentFile() string
	EnvironmentTestingFile() string
}
