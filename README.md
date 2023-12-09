# evm-inscriptions
evm生态铭文脚本

# 编译
```shell
# 直接编译
go build -o mint main.go

# 交叉编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mint.exe main.go

# 交叉编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o mint main.go
```

## 使用方式
使用linux的用户，基本都会用命令，这里就不介绍linux版的使用方式了，只介绍windows版的：  
1. 解压后，得到两个文件：  
二进制程序`mint`  
配置文件`settings.yml`（根据需要调整里面的配置参数，具体参数作用，在里面的备注上已经做了详细解释）  
2. 这两个文件放在相同目录下，
3. windows版，打开cmd命令框，并进入上述文件所在目录(该过程可以去网上找，到处都有讲解步骤)  
4. 然后输入：`./mint.exe` 回车，即可开始执行