# redis


### 1. 基本的数据类型

- string
- list
- sort
- zsort
- hash

针对各种数据类型的基本操作：

- C
- U
- R
- D

### 2. 典型应用场景

- TOP N


### 3. Tips

批量删除 key
```
redis-cli keys "*idiom:hash*" | xargs redis-cli del

```
