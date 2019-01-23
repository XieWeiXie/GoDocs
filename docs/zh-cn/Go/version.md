# Go 版本管理


### govendor

```
govendor init
govendor add +external
govendor fetch
```

### dep


```
dep init
dep ensure

```

### go mod

> https://docs.gomods.io/

```
// on 开启、off 关闭、auto
export GO111MODULE=on

go mod init
go mod tidy
go mod vendor
go list -m

```

example:
```

go mod init github.com/wuxiaoxiaoshen/go-emoji
go mod tidy
go mod vendor 

```
