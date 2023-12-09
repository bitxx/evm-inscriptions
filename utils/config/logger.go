package config

import (
	"evm-inscriptions/utils/log"
)

type Logger struct {
	Type   string
	Path   string
	Level  string
	Stdout string
	Cap    uint
}

func (e Logger) Setup() {

	log.Init(log.LoggerConf{
		Type:   e.Type,
		Path:   e.Path,
		Level:  e.Level,
		Stdout: e.Stdout,
		Cap:    e.Cap,
	})
}

var LoggerConfig = new(Logger)
