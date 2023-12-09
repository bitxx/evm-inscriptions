package main

import (
	"errors"
	"evm-inscriptions/app"
	"evm-inscriptions/utils/config"
	"evm-inscriptions/utils/log"
	"fmt"
	"github.com/bitxx/load-config/source/file"
	"github.com/shopspring/decimal"
	"time"
)

const (
	//configPath = "settings.dev.yml" //测试专用
	configPath = "settings.yml" //测试专用
)

const (
	EvmTypeReth = "reth"

	RethDifficulty = "0x007777777" //当前reth难度
)

func init() {
	config.Setup(
		file.NewSource(file.WithPath(configPath)),
	)
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Errorf("mint失败，程序终止，失败原因：%s", err)
			return
		}
		log.Info("所有mint任务完成")
		time.Sleep(3 * time.Second) //3秒种后，窗口关闭
	}()
	evmApp := app.NewApp()
	//转eth单位精度
	accuracyEth, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		return
	}
	//转Gwei单位精度
	accuracyGWei, err := decimal.NewFromString("1000000000")
	if err != nil {
		return
	}
	gasPrice, err := decimal.NewFromString(config.MintConfig.GasPrice)
	if err != nil {
		return
	}
	log.Infof("欢迎使用 %s ", config.ApplicationConfig.Name)
	log.Infof("当前节点地址：%s", config.ChainConfig.Url)
	log.Infof("总计mint张数：%d", config.MintConfig.Times)
	log.Infof("单张gas price：%sGwei", gasPrice.DivRound(accuracyGWei, 18))
	log.Infof("单张gas limit：%s\n\n", config.MintConfig.GasLimit)

	//开始mint
	for i := 1; i <= config.MintConfig.Times; i++ {
		balanceStr, er := evmApp.TokenBalanceOf()
		err = er
		if err != nil {
			return
		}
		balance, er := decimal.NewFromString(balanceStr)
		err = er
		if err != nil {
			return
		}
		log.Infof("当前账户余额：%s", balance.DivRound(accuracyEth, 4))

		data := config.MintConfig.Data
		if config.ChainConfig.ChainType == EvmTypeReth {
			data = evmApp.RethCalc(RethDifficulty)
		}

		hash, er := evmApp.Mint(data)
		err = er
		if err != nil {
			err = errors.New(fmt.Sprintf("第%d张mint异常，原因：%s", i, err))
			return
		}
		log.Infof("第%d张mint成功，hash：%s", i, hash)
		time.Sleep(3 * time.Second)
	}
}
