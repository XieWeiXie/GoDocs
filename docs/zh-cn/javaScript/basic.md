# 基本知识



## 语言特性

- 函数: 匿名函数、方法、常规函数
- 变量: 字符串、列表、对象、map、set、数值型
- 循环： for, while
- 分支: if..else, break、continue
- 对象

```

var print = function(message){
  console.log(message)
}

var intExample = 123
print(intExample)

var floatExample = 1.23
print(floatExample)

var strExample = 'how can i learn'
print(strExample)

var listExample = [1,2,"3", '4']
print(listExample)

var objectExample = {
  "1": "one",
  "2": "two",
  "3": listExample,
  "4": function(){
    return 4321
  }
}

print(objectExample)
print(objectExample["4"]())

function example(message){
  return message
}

print(example("normal function"))


function isBool(a, b){
  if (a > b) {
    return true
  }else{
    return false

  }
}

print(1,2)

```
## 浏览器属性

> 访问属性或者方法以 . 的形式
**windows**

- windows.InnerWidth
- windows.InnerHeight

**navigator**

- navigator.appName
- navigator.appVersion
- navigator.userAgent
- navigator.language
- navigator.platform
