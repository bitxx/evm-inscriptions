# evm-inscriptions
evm生态铭文脚本  
`说明：`  
用第三方平台打铭文，得提供自己的密钥，总是不太放心，为此自己开发了这个脚本。  
代码很简单，没后门，为了方便使用，我编译了几个常用平台。不放心的可以自行审阅代码并重新编译

**声明：本项目完全免费，只是出于个人爱好而分享。没有任何克扣gas和盗取私钥行为，如果不放心，请自行审阅代码并编译。使用期间，出现的任何问题，本项目及作者概不负责**

**声明：针对有可能被第三方盗用本项目并改造成收费的，我只想说，吃相不要太难看，差不多就行了。我也不想花精力去纠结这些行为，只希望使用者能擦亮眼，都是自己的辛苦钱，别乱砸**

![示例](/example.jpg)

# 功能
1. 理论上目前支持所有evm生态相关公链的铭文mint，目前我测试了`https://evm.ink/` 中的三条公链`ethereum、bnb、马蹄`,还有AVAX也测试通过，一次性mint百八十万张没什么问题，只要你gas足够多

## 下载入口
不想编译或者不会编译的，可以直接使用我编译好的二进制直接操作：[evm-inscriptions已编译程序下载入口](https://github.com/bitxx/evm-inscriptions/releases)  
`再次强调`：该项目没有任何后门。如实在不放心，请自行审阅和编译。使用期间，出现的任何问题，本项目及作者概不负责。  

## 针对windows的使用方式
主要针对行外人  
使用linux的用户，基本都会用命令，可参考下方windows版的使用方式：
1. 解压后，得到两个文件：  
   二进制程序`mint`  
   配置文件`settings.yml`（根据需要调整里面的配置参数，具体参数作用，在里面的备注上已经做了详细解释）
2. 这两个文件放在相同目录下，
3. windows版，打开cmd命令框，并进入上述文件所在目录(该过程可以去网上找，到处都有讲解步骤)
4. 然后输入：`./mint.exe` 回车，即可开始执行

# 编译
需要有golang环境，`1.21.5`及以上的版本

```shell
# 项目根目录添加依赖包
go mod tidy

# 直接编译当前平台
go build -o mint main.go

# 交叉编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mint.exe main.go

# 交叉编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o mint main.go
```
可使用该项目根目录的`build.sh`脚本，可快速各平台打包  
`mint`二进制和配置文件`settings.yml`均在同一目录下，配置文件参数设置好，命令行进入当前目录，直接执行`mint`即可


