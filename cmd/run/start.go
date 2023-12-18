package run

import (
	"evm-inscriptions/app"
	"evm-inscriptions/utils/config"
	"github.com/bitxx/load-config/source/file"
	"github.com/spf13/cobra"
	"log"
)

var (
	configPath string
	StartCmd   *cobra.Command
)

const (
	name           = "name"
	chainType      = "chain-type"
	chainUrl       = "chain-url"
	chainTimeout   = "chain-timeout"
	mintTimes      = "mint-times"
	mintDelay      = "mint-delay"
	mintPrivateKey = "mint-private-key"
	mintGasPrice   = "mint-gas-price"
	mintGasLimit   = "mint-gas-limit"
	mintData       = "mint-data"
	logPath        = "log-path"
	logLevel       = "log-level"
	logStdout      = "log-stdout"
	logType        = "log-type"
	logCap         = "log-cap"
	rethDifficulty = "reth-difficulty"
)

func init() {
	StartCmd = &cobra.Command{
		Use:          "start",
		Short:        "run the evm inscriptions",
		Example:      "mint start -c settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			//后执行
			if configPath != "" {
				config.Setup(
					file.NewSource(file.WithPath(configPath)),
				)
			}

			flag := cmd.PersistentFlags()

			//优先命令行覆盖配置
			if name, _ := flag.GetString(name); name != "" {
				config.ApplicationConfig.Name = name
			}
			if chainType, _ := flag.GetString(chainType); chainType != "" {
				config.ChainConfig.ChainType = chainType
			}
			if chainUrl, _ := flag.GetString(chainUrl); chainUrl != "" {
				config.ChainConfig.Url = chainUrl
			}
			if chainTimeout, _ := flag.GetInt64(chainTimeout); chainTimeout > 0 {
				config.ChainConfig.Timeout = chainTimeout
			}
			if mintTimes, _ := flag.GetInt(mintTimes); mintTimes > 0 {
				config.MintConfig.Times = mintTimes
			}
			if mintDelay, _ := flag.GetInt(mintDelay); mintDelay > 0 {
				config.MintConfig.Delay = mintDelay
			}
			if mintPrivateKey, _ := flag.GetString(mintPrivateKey); mintPrivateKey != "" {
				config.MintConfig.PrivateKey = mintPrivateKey
			}
			if mintGasPrice, _ := flag.GetString(mintGasPrice); mintGasPrice != "" {
				config.MintConfig.GasPrice = mintGasPrice
			}
			if mintGasLimit, _ := flag.GetString(mintGasLimit); mintGasLimit != "" {
				config.MintConfig.GasLimit = mintGasLimit
			}
			if mintData, _ := flag.GetString(mintData); mintData != "" {
				config.MintConfig.Data = mintData
			}
			if logPath, _ := flag.GetString(logPath); logPath != "" {
				config.LoggerConfig.Path = logPath
			}
			if logLevel, _ := flag.GetString(logLevel); logLevel != "" {
				config.LoggerConfig.Level = logLevel
			}
			if logStdout, _ := flag.GetString(logStdout); logStdout != "" {
				config.LoggerConfig.Stdout = logStdout
			}
			if logType, _ := flag.GetString(logType); logType != "" {
				config.LoggerConfig.Type = logType
			}
			if logCap, _ := flag.GetUint(logCap); logCap > 0 {
				config.LoggerConfig.Cap = logCap
			}
			if rethDifficulty, _ := flag.GetString(rethDifficulty); rethDifficulty != "" {
				config.RethConfig.Difficulty = rethDifficulty
			}

			if config.ChainConfig.Url == "" {
				log.Fatalf("%s参数不得为空", chainUrl)
			}
			if config.MintConfig.PrivateKey == "" {
				log.Fatalf("%s参数不得为空", mintPrivateKey)
			}
			if config.MintConfig.GasPrice == "" {
				log.Fatalf("%s参数不得为空", mintGasPrice)
			}
			if config.MintConfig.GasLimit == "" {
				log.Fatalf("%s参数不得为空", mintGasLimit)
			}
			if config.MintConfig.Data == "" {
				log.Fatalf("%s参数不得为空", mintData)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return run()
		},
	}
	//此处先执行
	cmd := StartCmd.PersistentFlags()
	cmd.StringVarP(&configPath, "config", "c", "", "Start server with provided configuration file")
	cmd.String(name, "evm-inscriptions", "node name")
	cmd.String(chainType, "evm", "chain type. evm 通用evm的基本方式打mint，一般情况下，每个平台的mint都是支付手续费，data输入内容，即可mint\nreth 算力计算，获取结果，然后再去mint，此时不需要配置下方mint中的data")
	cmd.String(chainUrl, "", "链结点地址，提示：可前往此处 https://publicnode.com/ 查找满足你需要的免费evm结点url")
	cmd.Int64(chainTimeout, 60, "请求超时时间，如无特殊必要，默认60即可，单位：秒")
	cmd.Int64(mintTimes, 1, "铸造次数，通俗说，批量铸造多少个，就配置多少")
	cmd.Int64(mintDelay, 5, "每次mint间隔时间，单位秒")
	cmd.String(mintPrivateKey, "", "私钥")
	cmd.String(mintGasPrice, "", "配置gasPrice，单位是wei ； 备注：1Gwei = 1000000000 wei")
	cmd.String(mintGasLimit, "", "配置gasLimit")
	cmd.String(mintData, "", "铭文16进制字符串，比如mint马蹄链的anteater，则使用下面铭文（仅限示例，该铭文很有可能已经被mint完了）\n如果上面chainType配置了reth，则该data参数无需配置（配置了也没任何用处），内部会自动算出提交")
	cmd.String(rethDifficulty, "0x000077777777", "根据官方最新难度，自行进行调整，我目前这里默认给的是2023-12-14最新难度：0x000077777777")
	cmd.String(logPath, "", "log path")
	cmd.String(logLevel, "debug", "log level")
	cmd.String(logStdout, "default", "default,file")
	cmd.String(logType, "default", "default、zap、logrus")
	cmd.Uint(logCap, 50, "log cap")
}

func run() error {
	app.NewApp().Start()
	return nil
}
