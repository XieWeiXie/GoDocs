# GRPC

> 远程调用函数


- ProtocolBuffer: 语法
    - 包
    - 导入
    - 定义函数
    - 定义请求参数
    - 定义响应参数
    - 数据类型
- protoc（ Protocol Compiler） 命令行工具，根据 .proto 文件生成指定的编程语言的文件
    - `-I` == `--proto-path`
    - `--go_out=`
    - `*.proto`

- 编写服务端逻辑
    - 请求参数
    - 响应信息
    - 数据库交互
- 编写客户端逻辑
    - 构造请求参数
    - 调用函数
    - 得到结果



### ProtocolBuffer

- 数据类型
- 嵌套
- 子属性
- 枚举