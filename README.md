# FISCO-BCOS Go SDK 演示

> 获取 SDK 国密 TLS 分支 `git clone -b feat/gmtls https://github.com/channingduan/go-sdk.git`
> 配置 go mod replace 到上面的项目地址.
> 修改配置文件.
> 运行,国密: go run main.go --isGM=true. 非国密: go run main.go --isGM=false

## 配置文件说明
> 非国密的版本,请直接用[官方配置文件](https://github.com/FISCO-BCOS/go-sdk#%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6%E8%AF%B4%E6%98%8Econfigtoml)
> 国密需要添加加密证书,配置文件如下:

```toml
[Network]
Type="channel"
CAFile="./configs/sdk/gm/gmca.crt"
Cert="./configs/sdk/gm/gmsdk.crt"
Key="./configs/sdk/gm/gmsdk.key"

# 国密 TLS 加密证书
EnCert="./configs/sdk/gm/gmensdk.crt"
EnKey="./configs/sdk/gm/gmensdk.key"

[[Network.Connection]]
NodeURL="bcos.gm:20200"
GroupID=1

[Account]
KeyFile="./configs/sdk/gm/user.pem"

[Chain]
ChainID=1
# 开始国密 TLS 支持
SMCrypto=true

[Log]
Path="./"
Level="debug"
```