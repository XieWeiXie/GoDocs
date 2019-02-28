# Golang


## 环境搭建


- GOROOT: 源代码的安装目录
- GOPATH: 项目的工程目录
  - src: 源代码
  - bin：编译后的文件
  - pkg：编译过程中生成的一些文件
## 基础概念


- package： 包里面的函数是如何调用的
- import: 路径的形式
- main



- 变量
  - 其一个好名字
  - 指定数据类型：整型、浮点型、数组、切片...
- 数据类型
- 控制流程
- 循环
- 函数


- 先判断正向的逻辑
- 先判断简单的
- 先判断有趣的，有缺陷的


函数：

- 输入：参数，指明类型
- 输出： 返回值，指明类型


结构体：

- 属性的公有私有，首字母的大小写
- 方法的编写
- 如何“继承” ： 组合
- 访问属性和方法：n.Name , n.GetID()
- 传值还是传指针



接口:

- 定义: 方法的集合
- 内置 error


## 内置库的使用


- 字符串： strings
- 字符串： strconv
- 时间： time
- 文件操作： file,os



---

1. 安装
  - 源代码
  - 设置环境变量：GOROOT, GOPATH
2. IDE
  - vscode
  - goland：jetbrains
  - 快捷键
3. 第一个程序
  - 整体结构
    - 包名：package
    - 导入：库：内置、第三方的: 路径的方式
    - 导入未使用，重命名，排序
    - main
    - fmt

4. 基本的语法
  - 变量：命名：见闻知意，名词
  - 数据类型：
    - 数值型：整型、浮点型...
    - 字符串：
      - 索引index
      - 长度
      - 大写
      - 小写
      - 分隔： “a,v,b”
      - 转换：“1”， 1， “a,v,c” ["a", "v", "c"]
      - strings
    - 数组
      - [3]int
      - [4]float
      - [2]string
      - 长度
      - 取值
    - 切片
      - []int
      - []float
      - []string
      - []interface{}
    - map
      - key
      - value
      - {"name": "xiewei"}
      - 遍历： for ... range ...
    - 结构体
      - 定义： type Name struct {}
      - 方法： Name.func
    - 函数
      - 功能：不用关注细节
      - 复用
      - 好的命名：多是动词开头 toString, doServer, doLoop
      - 私有、公有
    - 可理解的代码

----

### 内置库

- strings


### 并发

- 如何启动并发：go 关键字
- 并发如何通信：channel
- 两个方向: 通道内存数据，或者通道内取数据，使用并发

```go
go func() {

  c <- 1
}

<-c

```

```go
go func(){

  <-c
}
c <- 1
```

- 实现同步

```go
var w sync.WaitGroup

w.Add(delta int)

go func (){
  defer w.Done()

}

w.Wait()

```

- 缓冲的通道


```

- runtime.GOMAXPROCS(0)：线程的个数，0 表示获取，不为 0 表示设置
- runtime.NumGoroutine() : 协程的个数
- runtime.NumCPU


```

- 锁
> 解决共享内存的问题，在写操作的过程中加锁

```go

mux sync.Mutex

mux.Lock
mux.Unlock

```

- 并发版的爬虫思路

```
1. 两个 通道：一个通道用来存链接，一个通道用来存结果
2. 不断的往存链接的通道存链接，同时协程获取结果，不断的把结构存另一个通道
3. 不断的从结果通道内取结果，交给 redis  存储
4. 构建API 从 redis 内获取结果

```


### 版本包管理

- govendor

```go
>> govendor init
>> govendor add +external
>> govendor install
>> govendor fetch [address]

```

- dep

```go

>> dep init // 初始化
>> dep ensure // 添加版本
>> dep status // 各个库的版本的信息
>> dep check // 检查状态

```

- mod

```go

go version 1.11

export GO111MODULE on

>> go mod init github.com/wuxiaoxiaoshen/project_name
>> go mod tidy
>> go mod vendor


```

### Jsonb

```
如何转变

```


###  指定浮点型的精度

```$xslt
strconv.FormatFloat(float64(i.Count)/float64(length)*100, 'f', 1, 32)

```

### sql 中字符串带引号

 ```
fmt.Sprintf("'%#v'", "data")
 ```


 ### jsonb  的处理

1. struct --> jsonb

 ```
 func (rule PostRules) JsonbLowHandler() postgres.Jsonb {
 	var values postgres.Jsonb
 	values.RawMessage, _ = json.Marshal(rule.LowRule)
 	return values
 }



 ```

2. jsonb -- > struct

```
func (rule FrequentCustomerRule) GetHighFrequency() rulePairs {
	var highRulePairs rulePairs
	if err := json.Unmarshal(rule.HighFrequency.RawMessage, &highRulePairs); err != nil {
		return nil
	}
	return highRulePairs
}
```


### 对切片指定长度和 append 不混合使用 


### 编译时用ldflags设置变量的值

``` 
package main

import "fmt"

var (
    VERSION    string
    BUILD_TIME string
    GO_VERSION string
)

func main() {
    fmt.Printf("%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)
}

```

``` 
go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"

// 1.0.0
   2019年 2月28日 星期四 09时44分39秒 CST
   go version go1.11 darwin/amd64
```


### go 交叉编译



- GOOS 目标操作系统: darwin(mac)、linux、windows
- GOARCH 目标处理器架构: 386、amd64、arm

Mac to Linux
``` 

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go


``` 

Linux to mac

``` 
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```