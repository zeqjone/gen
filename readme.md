# mysql to go model

根据数据库表结构生成 go 模型的工具

## 使用方法

- 下载源代码
- 如果本地没有 goimports, 则请下载 ```go get -u golang.org/x/tools/cmd/goimports```
- 执行 ```go build -o gen main.go``` 构建可执行文件
- ```gen -h``` 查看指令集
- 所有的指令集类似 git 指令格式

## 版本规划

- [x] 根据配置文件，支持主流的 ORM 库，默认 gorm
- [ ] 根据配置文件，支持主流的 orm 库，beego
- [x] go format 仅格式化生成的文件
- [ ] 支持达梦数据库，ongoing

## 联系

mailto: zeq_jone@163.com
