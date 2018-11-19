# 数据库

1. 数据库
2. 数据库管理系统

- 物理层
- 逻辑层
- 实体

> 相互映射


设计数据库管理系统的步骤：

1. 需求分析
2. 概念：ER 图 （实体 -- 联系）
3. 逻辑分析
4. 结构优化


> 属性  -- > 实体  -- >  联系


1. DDL
2. DML
3. DCL

范式：

- 原子性：属性不可分
- 主键、其他属性和主键唯一相关
-


1. 主键（超码）
2. 外键（外码）

## 表

### 1. 创建表


- 一些关键字
```
create table
primary key
not null
default
```

- 数据类型

```
- varchar 、char
- integer、numeric
- time 、timestamp
- jsonb、json
- text


```

- 约束

```
- not null
- primary key
- check
- foregin key
- default
- unique
```


### 2. 修改表

- 列
  - add column
  - drop column
- 约束
  - check
  - default
  - unique
  - not null
- 类型
- 重命名


```sql

alter table example_table add column fuck text;

alter table example_table drop column fuck;

alter table example_tale rename column fuck to fuckkk;

alter table example_table rename to example;

alter table example_table alter column fuckk set default '9.0';

alter table example_table alter column fuckk drop default;

alter table example_table alter column fuckk type varchar(50);




```




## 数据类型

- 数值型
- 字符串型
- 时间日期类型
- 布尔类型
- Json 类型、Jsonb 类型
- 数组：数值型数组、字符串型数

## 函数

### 1. 字符串

- 单个字符的操作
- 长度的操作
- 多个字符的操作

```
// 存储入数据库的需要时候就考虑的规则化

- 拼接： || concat
- 大小写转换: upper, lower
- 子串: substring
- 索引: position
- 前后字符的去除： trim
- 替换：overlay
- 长度： length, bit_length,char_length,octet_length, character_length

```

实例：

```sql
select ‘hello’ || ‘world’;
>> helloworld

select concat('hello', '', 'world');
>> helloworld

select upper('device'), lower('DE');
>> DEVICE de

select substring('hello' from 1 for 2);
>> he

select substring('Thomas' from '$"__o_a$"s' for '$');
>>

select position('md5' in 'hello md5');
>> 7

select trim(both '1' from '1232321212q1');
>> 232321212q

select overlay('Txxxxas' placing 'hom' from 2 for 4);
>>Thomas

select bit_length('H'), length('H2121'),octet_length('12'), char_length('helo'), character_length('12212');
>> 8            | 5        | 2              | 4             | 5

```
### 2. 数值型

> 各种类型的数值，说到底其实就是精度的不同


- 关系运算：大小判断
- 算术运算：加减乘除, 数学中的各种运算
- 统计运算（加减乘除）


> 1. 使用操作符 2. 使用数学函数


- 操作一个数值
- 两个或者多于两个



```

- 基本操作：加减乘除
- 三角运算

- 数学操作符
- 数学函数
- 三角函数
- 随机函数

```

示例：

```sql
select ceil(43.2);
>> 44
select floor(43.2);
>> 43
select round(43.2);
>> 43
select trunc(43.2);
>> 43
select 2 + 3 as add, 2-3 as sub, 2*3 as mul, 2/3 as dev, 2%3 as remain, 2^3 as power, 2! as m12,!!2 as m11,~2 as not2;
>>
+-------+-------+-------+-------+----------+---------+-------+-------+--------+
| add   | sub   | mul   | dev   | remain   | power   | m12   | m11   | not2   |
|-------+-------+-------+-------+----------+---------+-------+-------+--------|
| 5     | -1    | 6     | 0     | 2        | 8.0     | 2     | 2     | -3     |
+-------+-------+-------+-------+----------+---------+-------+-------+--------+


```

### 3. 布尔类型

- 逻辑运算
- 是非判断


> 多用于条件判断

```
- is
- is not null
- is true
- is false
- isnotnull
- isnull
- is not false
- is not true

```


### 4. 时间日期类型

> 1.  格式 2. 操作 3. 转换

- 时间单位
- 单个时间操作获取值
- 多个时间进行操作
- 时间类型和字符串类型之间转换

```
- timestamp: 日期 + 时间
- time： 时间
- date： 日期
- interval：

```

- 时间日期操作符

```
+
-
*
```

示例：

```sql
select current_time;
>> 12:07:39.058952+08:00
select current_date;
>> 2018-11-18
select current_timestamp;
>> 2018-11-18 12:08:20.137552+08
select localtime;
>> 12:08:48.654921
select localtimestamp;
>> 2018-11-18 12:10:26.081166

// 获取子域，即单个值
select * from example_time;
select date_part('year', n) from example_time;
select date_part('month', n) from example_time;
select date_part('week', n) from example_time;
select date_part('day',n) from example_time;
select date_part('min', n) from example_time;
select date_part('sec', n) from example_time;
select date_part('hour', n) from example_time;

// 截取，即表达式和原始的一致，但截取部分

select date_trunc('year',n) from example_time;
>> 2018-01-01 00:00:00
select date_trunc('month', n) from example_time;
>> 2018-11-01 00:00:00
select date_trunc('week', n) from example_time;
>> 2018-11-12 00:00:00
select date_trunc('hour', n) from example_time;
>> 2018-11-18 00:00:00
select date_trunc('day', n) from example_time;
>> 2018-11-18 00:00:00
select extract( hour from n) from example_time;
>>
select extract(year from n) from example_time;
>>
select extract(week from n) from example_time;
>>
select extract(min from n) from example_time;
>>
select extract(month from n ) from example_time;
>>
select extract(day from n) from example_time;
>>

extract == date_part
```

### 5. json, jsonb

> json

- 数值型
- 数组型
- 布尔型
- 字符串类型


```
- 操作符
  ->
  ->>
  #>

  @>
  <@


- 函数
  取元素


```


### 数组

- integer[]
- varchar[]
