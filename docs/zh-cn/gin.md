# Gin 教程


- web 框架
- 优势
- 基本使用：路由、资源
- 实践：管理系统




# API 操作


- 请求参数: json、form、query、path
- 请求方法: PUT、POST、GET、DELETE
- 多表操作: Proload、Association
- 数据库增删改查





## HTTP

### 实体

```
content-type: text/plain 纯文本
content-type: application/json; charset=utf-8 json
content-type: multipart/form-data 表单，上传媒体文件
application/x-www-form-urlencoded 表单，标准的编码格式
allow: 运行的方法
```


### 请求参数

- 默认:
- 限定 eq=|eq=
- 是否必须: binding:required
- 类型: int、string、multipart.Form
- 举例：example



表单:

```

type PostParam struct {

    Name string `form:"name,default=xiewei" binding:"gt=0,max=50"`
    Gender int `form:"gender" binding:"eq=0|eq=1"`
    Avatar multipart.Form `form:"avatar" binding:"required"`

}


```


json:

```

type PostParam struct {

  Name string  `json:"name"`
  Age int `json:"age"`

}


```


### 响应

- json
- string
- xml
- html
- file
- attachment


### HTTP POST GOLNAG

```
m := map[string]interface{}{
  "name": "backy",
  "species": "dog",
}
mJson, _ := json.Marshal(m)
contentReader := bytes.NewReader(mJson)
req, _ := http.NewRequest("POST", "http://example.com", contentReader)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Notes","GoRequest is coming!")
client := &http.Client{}
resp, _ := client.Do(req)

```

### 请求参数，可以组合形式

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type A struct {
	Name string `form:"name" json:"name"`
}

type B struct {
	A
	Age int `form:"age" json:"age"`
}


func main() {
	g := gin.Default()
	g.POST("/v1/post", func(context *gin.Context) {
		var params B
		if err := context.ShouldBind(&params);err!=nil{
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": params.Name,
			"age": params.Age,
		})
	})
	g.Run(":1111")
}



```
