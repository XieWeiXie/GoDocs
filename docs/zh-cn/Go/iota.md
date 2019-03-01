# Iota 枚举

> 遇到下一个 const 恢复为 0 值，关注第一行的表达式

## 1. 每行一个值



## 2. 每行多个值


``` 

package main

import "fmt"

const  (
	A = iota
	B
	C
)
// Output:
// 0,1,2

const (
	AA = iota + 10
	BB
	CC
)
// Output:
// 10,11,12

const (
	AAA = 1<<iota
	BBB
	CCC
)

// Output:
// 1,2,4
const (
	AAAA, BBBB = iota,iota
	CCCC,DDDD
	EEEE,
	FFFF
)
// Output:
// 0,0,1,1,2,2

func main(){
	fmt.Println(A,B,C)
	fmt.Println(AA,BB,CC)
	fmt.Println(AAA,BBB,CCC)
	fmt.Println(AAAA,BBBB,CCCC,DDDD,EEEE,FFFF)
}
```