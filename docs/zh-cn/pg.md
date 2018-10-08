


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

## 9. 常用操作






- [官方文档](https://www.postgresql.org/)
