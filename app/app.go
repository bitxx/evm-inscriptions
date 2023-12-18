package app

import (
	"encoding/hex"
	"evm-inscriptions/utils/config"
	"evm-inscriptions/utils/log"
	"fmt"
	"github.com/bitxx/ethutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const (
	EvmTypeReth = "reth"
)

type App struct {
	client *ethutil.EthClient
}

func NewApp() *App {

	return &App{
		client: ethutil.NewEthClient(config.ChainConfig.Url, config.ChainConfig.Timeout),
	}
}

func (a *App) Start() {
	var err error
	defer func() {
		if err != nil {
			log.Errorf("mint失败，程序终止，失败原因：%s", err)
			return
		}
		log.Info("所有mint任务完成")
		time.Sleep(3 * time.Second) //3秒种后，窗口关闭
	}()
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

		balanceStr, er := a.TokenBalanceOf()
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

		data := config.MintConfig.Data
		if config.ChainConfig.ChainType == EvmTypeReth {
			data = a.RethCalc(config.RethConfig.Difficulty)
		}

		hash, er := a.Mint(data)
		err = er
		if err != nil {
			log.Errorf("第%d张mint异常，原因：%s", i, err)
			continue
		}
		log.Infof("第%d张mint成功，hash：%s", i, hash)
	}
}

// TokenBalanceOf
//
//	@Description: 查询余额
//	@receiver a
//	@return balance
//	@return err
func (a *App) TokenBalanceOf() (balance string, err error) {
	account, err := a.client.AccountWithPrivateKey(config.MintConfig.PrivateKey)
	if err != nil {
		return "", err
	}
	return a.client.TokenBalanceOf(account.Address)
}

// Mint
//
//	@Description: mint
//	@receiver a
//	@return hash
//	@return err
func (a *App) Mint(data string) (hash string, err error) {
	account, err := a.client.AccountWithPrivateKey(config.MintConfig.PrivateKey)
	if err != nil {
		return "", err
	}
	return a.client.TokenTransfer(config.MintConfig.PrivateKey, config.MintConfig.GasPrice, config.MintConfig.GasLimit, "", "0", account.Address, data)
}

// RethCalc
//
//	@Description: reth计算难度结果
//	@receiver a
//	@param difficulty
//	@return string
func (a *App) RethCalc(difficulty string) string {

	type Result struct {
		TotalTime int
		Result    string
	}
	log.Info(fmt.Sprintf("reth当前难度：%s，计算中，请稍后...", config.RethConfig.Difficulty))
	begin := time.Now()
	find := make(chan Result)
	temp := `data:application/json,{"p":"rerc-20","op":"mint","tick":"rETH","id":"%s","amt":"10000"}`
	challenge, _ := hex.DecodeString("7245544800000000000000000000000000000000000000000000000000000000")
	for i := 0; i < runtime.NumCPU()-1; i++ {
		time.Sleep(time.Millisecond * 10)
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		go func() {
			for {
				answer := make([]byte, 32)
				rd.Read(answer)
				payload := append(answer, challenge...)
				hashed := crypto.Keccak256Hash(payload)
				hashedHex := hashed.Hex()
				if strings.HasPrefix(hashedHex, difficulty) {
					endTime := time.Now()
					result := Result{
						TotalTime: endTime.Second() - begin.Second(),
						Result:    fmt.Sprintf(temp, "0x"+hex.EncodeToString(answer)),
					}
					find <- result
					break
				}
			}
		}()
	}

	for {
		select {
		case v := <-find:
			result := hexutil.Encode([]byte(v.Result))
			log.Infof("计算成功，耗时：%ds,  16进制计算结果：%s", v.TotalTime, result)
			return result
		}
	}
}
