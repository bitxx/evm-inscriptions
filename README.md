# evm-inscriptions
evm生态铭文脚本  
`说明：`  
用第三方平台打铭文，得提供自己的密钥，总是不太放心，为此自己开发了这个脚本。  
代码很简单，没后门，为了方便使用，我编译了几个常用平台。不放心的可以自行审阅代码并重新编译  

# 功能
1. 理论上目前支持所有evm生态相关公链的铭文mint，目前我只测试了`https://evm.ink/` 中的三条公链`ethereum、bnb、马蹄`,一次性mint千八百张没什么问题，只要你gas足够多
2. 支持[reth](https://reth.cc/list) 的mint，貌似进度已经90%多，还剩不到1万张。这个需要cpu计算，然后才能mint，这整个过程我也整合了。  

# 编译
```shell
# 直接编译
go build -o mint main.go

# 交叉编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mint.exe main.go

# 交叉编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o mint main.go
```
可使用该项目根目录的`build.sh`脚本，可快速各平台打包  
`mint`二进制和配置文件`settings.yml`均在同一目录下，配置文件参数设置好，命令行进入当前目录，直接执行`mint`即可  

## 针对windows的使用方式
主要针对行外人  
使用linux的用户，基本都会用命令，可参考下方windows版的使用方式：  
1. 解压后，得到两个文件：  
二进制程序`mint`  
配置文件`settings.yml`（根据需要调整里面的配置参数，具体参数作用，在里面的备注上已经做了详细解释）  
2. 这两个文件放在相同目录下，
3. windows版，打开cmd命令框，并进入上述文件所在目录(该过程可以去网上找，到处都有讲解步骤)  
4. 然后输入：`./mint.exe` 回车，即可开始执行