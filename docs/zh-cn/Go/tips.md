# Go tips


#### 匿名函数

处理微小的功能，使用匿名函数

### 文件的组织

- 自上而下
- 功能按文件区分

> 参考 go mod 命令行组织方式

### 接口

> 以 er 结尾命名

``` 
type Stringer interface {
	String() string
}
```