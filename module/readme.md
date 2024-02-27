# 模块应用示例


## 工具使用

生成Service

`./gf gen service -s './module/hello/logic' -d './module/hello/service`


## 注意事项

在写单例测试的时候，要注意引用相关的包，否则会出现错误

例如：task_test.go 文件的引用，需要引用如下包
```go
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "sagooiot/module/simcard/logic/sim"

```


