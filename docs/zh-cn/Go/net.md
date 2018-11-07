# net/http


### 1. Request

- r.Method 表示请求方法
- r.ParseForm 开始解析参数

如何知道是哪种类型的参数？

- r.PostForm
- r.PostFormValue

> Content-Type: application/x-www-form-urlencoded

- r.Form
- r.FormValue

> query 参数


- r.URL.Path

> 获取路径中的参数
