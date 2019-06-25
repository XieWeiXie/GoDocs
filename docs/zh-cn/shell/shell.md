# Shell 指南


大纲：

- 变量
- 循环
    - for ((;;))
    - for ... in ...
    - while
- 判断
    - if ... fi
    - case ...in  esac
- 函数
- 数据类型
    - 字符串
        - 拼接
        - 长度
        - 截取
        - 分割
        - 替换
    - 列表
        - 长度
        - 索引
- 文件包含
    - .
    - source



## 删除除某文件之外的文件

- rm -f !(某文件)

## 数值运算

- let
- expr
- ()
- []
- bc // 通过管道
    - scale
    - ibase
    - obase
    - sqrt

```
sed -r 's/.*"(.+)".*/\1/'
```