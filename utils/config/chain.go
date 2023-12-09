package config

type Chain struct {
	Url       string
	Timeout   int64
	ChainType string
}

var ChainConfig = new(Chain)
