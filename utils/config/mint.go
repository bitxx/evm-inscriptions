package config

type Mint struct {
	Times      int    //次数
	Delay      int    //间隔时间
	PrivateKey string //私钥
	GasPrice   string
	GasLimit   string
	Data       string //铭文未编码信息
	//MaxPriorityFeePerGas string //EIP1559需要配置，为后续可能的扩展使用
}

var MintConfig = new(Mint)
