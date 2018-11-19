# net/http

- 路径
- 参数
- 方法
- 头部信息
- 请求参数
- 响应
- 服务
- 客户端

> 看结构体是如何定义的，以及存在什么方法


两大类：

- 客户端：client  相关， 请求的是服务端
- 服务端： server 相关，完成的是服务端内容的构建




### 1. Request

- r.Method 表示请求方法
- r.ParseForm  / r.ParseMultipartForm 开始解析参数

如何知道是哪种类型的参数？

- r.PostForm
- r.PostFormValue

> Content-Type: application/x-www-form-urlencoded

- r.Form  和 PostForm 的区别就是，它把 URL 内的请求参数也包括在内
- r.FormValue

> query 参数


- r.URL.Path

> 获取路径中的参数


### 2. 状态码

- 1：请求已接受、等待响应
- 2：正常
- 3：重定向
- 4：客户端
- 5：服务端
