package config

type Reth struct {
	Difficulty string //难度，根据官方调整
}

var RethConfig = new(Reth)
