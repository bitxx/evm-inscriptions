package main

import (
	"evm-inscriptions/app"
	"evm-inscriptions/utils/config"
	"evm-inscriptions/utils/log"
	"github.com/bitxx/load-config/source/file"
	"github.com/shopspring/decimal"
	"time"
)

const (
	configPath = "settings.dev.yml" //测试专用
	//configPath = "settings.yml" //正式专用
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
		time.Sleep(time.Duration(config.MintConfig.Delay) * time.Second) //防止多次异常，导致连续mint同一个nonce

		balanceStr, er := evmApp.TokenBalanceOf()
		err = er
		if err != nil {
			log.Errorf("第%d张读取账户余额异常，原因：%s", i, err)
			continue
		}
		balance, er := decimal.NewFromString(balanceStr)
		err = er
		if err != nil {
			log.Errorf("第%d张解析账户余额异常，原因：%s", i, err)
			continue
		}
		log.Infof("当前账户余额：%s", balance.DivRound(accuracyEth, 4))

		hash, er := evmApp.Mint(config.MintConfig.Data)
		err = er
		if err != nil {
			log.Errorf("第%d张mint异常，原因：%s", i, err)
			continue
		}
		log.Infof("第%d张mint成功，hash：%s", i, hash)
	}
}
