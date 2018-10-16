# Gorm


- 索引
- 联合索引
- 唯一索引


- 链式操作：scope
- 错误处理：bool、error、notFound
- 钩子hook: 函数集合



1. 表设计
2. 单表
3. 多表


- 搜索
- 更新
- 查询
- 删除


## 表定义

- 表名
- 列名
- 类型
- 默认
- 唯一键
- 索引
- 联合索引
- 主键
- 联合主键
- 一对多
- 多对多

> 根据结构体的 Struct tag 来操作

```

'gorm:"type:varchar;default:'';column:shop_name;index"'

```
## 表操作

- 基于对象操作
- 原生 SQL 语句

- 行、列、记录、搜索、查询、删除、更新...


> 1. 注意 SQL 语句的熟练度即可解决问题
