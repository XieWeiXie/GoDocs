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



### Raw, Scan

定义的字段需要和表中的列名一致，tag 不起作用。


### 记录

#### 1. has one

```go
package models

import "github.com/jinzhu/gorm"

type ExampleOne struct {
	gorm.Model
	FirstName      string       `gorm:"varchar" json:"first_name"`
	ExampleTwo     []ExampleTwo `gorm:"many2many:example_one_two"`
	ExampleThreeID uint
	ExampleThree   ExampleThree
}

type ExampleTwo struct {
	gorm.Model
	SecondName string       `gorm:"varchar" json:"second_name"`
	ExampleOne []ExampleOne `gorm:"many2many:example_one_two"`
}

type ExampleThree struct {
	gorm.Model
	FullName string `gorm:"varchar" json:"full_name"`
}
func insert() {
	var one models.ExampleOne
	var two models.ExampleTwo
	var three models.ExampleThree
	three = models.ExampleThree{
		FullName: "XieWei",
	}
	if dbError := DB.Save(&three).Error; dbError != nil {
		return
	}
	one = models.ExampleOne{
		FirstName:      "Xie",
		ExampleThreeID: three.ID,
	}
	two = models.ExampleTwo{
		SecondName: "Wei",
	}

	if dbError := DB.Save(&one).Error; dbError != nil {
		return
	}
	if dbError := DB.Save(&two).Error; dbError != nil {
		return
	}

	DB.Model(&one).Association("ExampleTwo").Append(&two)
	DB.Model(&one).Association("ExampleThree").Append(&three)
	fmt.Println(one)
}

```


#### 2. has many

#### 3. many2many
