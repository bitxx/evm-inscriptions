package app

import (
	"evm-inscriptions/utils/config"
	"github.com/bitxx/ethutil"
)

type App struct {
	client *ethutil.EthClient
}

func NewApp() *App {

	return &App{
		client: ethutil.NewEthClient(config.ChainConfig.Url, config.ChainConfig.Timeout),
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
//	@Description: reth计算难度结果。这个是以前reth项目计算难度的方法，该项目已mint结束，注销此方法
//	@receiver a
//	@param difficulty
//	@return string
/*func (a *App) RethCalc(difficulty string) string {

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
}*/
