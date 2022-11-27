基于 [double-entry-generator](https://github.com/deb-sig/double-entry-generator) 的改造。

# 安装
```shell
go get -u github.com/miaogaolin/beancount-go
```

# 用法
```shell
beancount-go translate --config example.config example.csv
```

# 支付宝
* 过滤掉交易关闭
* 过滤掉配置文件中进出账户相同的交易
* 过滤掉交易金额为0的