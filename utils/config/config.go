package config

import (
	"fmt"
	loadconfig "github.com/bitxx/load-config"
	"github.com/bitxx/load-config/source"
	"log"
)

type Config struct {
	Application *Application `yaml:"application"`
	Logger      *Logger      `yaml:"logger"`
	Chain       *Chain       `yaml:"chain"`
	Mint        *Mint        `yaml:"mint"`
	Reth        *Reth        `yaml:"reth"`
	callbacks   []func()
}

func (e *Config) init() {
	e.Logger.Setup()
	e.runCallback()
}

func (e *Config) Init() {
	e.init()
	log.Println("!!! client config init")
}

func (e *Config) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Config) OnChange() {
	e.init()
	log.Println("!!! client config change and reload")
}

// Setup 载入配置文件
func Setup(s source.Source,
	fs ...func()) {
	_cfg := &Config{
		Application: ApplicationConfig,
		Chain:       ChainConfig,
		Mint:        MintConfig,
		Logger:      LoggerConfig,
		Reth:        RethConfig,
		callbacks:   fs,
	}
	var err error
	loadconfig.DefaultConfig, err = loadconfig.NewConfig(
		loadconfig.WithSource(s),
		loadconfig.WithEntity(_cfg),
	)
	if err != nil {
		log.Println(fmt.Sprintf("New client config object fail: %s, use default param to start", err.Error()))
		return
	}
	_cfg.Init()
}
