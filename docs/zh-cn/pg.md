# PostgreSQL


> 如何学、按类型分类、按关系分类



## 1. 安装

```
brew install postgresql

```

## 2. 初始化

```
init /usr/local/var/postgres

```

## 3. 启动

```

pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start

```

## 4. 查看状态

```

pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log status

```

## 5. 停止

```

pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log stop -s -m fast

```

## 6. 创建数据库

```

createdb -O xiewei dbname

```

## 7. 删除数据库

```

dropdb -U xiewei  dbname

```

## 8. 增删改查


根据数据类型：

- 数值类型
  - 比较操作
  - 函数统计操作
- 文本类型
  - 拼接
  - index
  - 切片
- 时间类型
  - 比较操作
  - 年月日时分秒：单位


根据表：

- 单表
- 多表：连接查询：根据表之间的共同字段等

## 9. 常用操作






- [官方文档](https://www.postgresql.org/)
