package config

type Application struct {
	Name    string
	Version string
}

var ApplicationConfig = new(Application)
