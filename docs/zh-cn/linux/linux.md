### 文件、目录

- mkdir -p
- touch
- cat
- tac
- more
- less
- tail -f
- cd
- ls
- cp
- rm
- mv
- pwd
- chmod
- chown
- chgrp
- alias



### 磁盘

- 单位B、 KB、MB、GB、TB、PB
- df
- du


### 打包

- tar
- man tar

```

tar -czvf example.tar.gz /etc/example/Read.md

tar -zxvf example.tar.gz -C /etc

```


### Linux 文件系统


- /bin
- /dev
- /etc
- /home
- /lib
- /sbin
- /tmp


### 包管理


- brew install
- brew search
- brew list




----


### vi/vim

> 文本编辑器

#### 1. 模式

- 一般模式
- 编辑模式
- 命令行模式

> 增删改查


- vim filename1 filename2 filename3
- :sp :vsp

###### 一般模式

- h,j,k,l : 上下左右
- G、gg,nG,ngg 跳转
- 0, $   行
- x, dd, ndd
- u, p
- r
- yy,p ,nyy, P


###### 编辑模式

- esc


- a, A
- o, O
- i, I

###### 底部命令命令行

- :set nu
- :set nonu
- :set hlsearch
- :set nohlsearch
- :syntax on
- :syntax off
- n, N (next, Next)
- /key
- ?key
- n1,n2s/old/new/g
- n1,n2s/old/new/gc
- 1,$s/old/new/gc
- 1,$s/old/new/g


###### 块选择

Ctrl +v

- ~
- >
- <
- d
- y


###### 多个文件


- vi file1 file2
- :n
- :N
- :e file
- :sp 上下
- ctrl ww
- ctrl w h,j,k,l 箭头
- :ls




- :w
- :x
- ZZ
- :q
- :q!
- :wq


### shell

- 设置
- 环境变量
- 脚本


### sed

```
// 第一行

sed -n '1p' passwd1

// 第n,m 行

sed -n '1,3p' passwd1

// 过滤

sed -n '/a/p' passwd1

sed -n '/a/,/b/p' passwd1

// 1,5行，并打印行号
sed -n '1,5p' -e '=' passwd1

// 替换

sed -e 's/a/A/g' a.txt


// 删除 行

sed '1d' a.txt
sed '1,3d' a.txt

// 过滤删除行
sed '/a/d' a.txt


// 更新

sed -i '_bak' 's/a/A/g' a.txt

// 替换

sed 'y/old/new/' a.txt
```

### tr

```
cat passwd1 | tr '[a-z]' '[A-Z]'

cat passwd1 | tr '[:lower:]' '[:upper:]'

cat passwd1 | tr -d '@'

cat passwd1 | tr -d '\n'
```

### cut

```
cat passwd1 | cut -f 1 -d '@'
```

### sort

```
- r: 反向
- n: 数字
- t: 分割符
- k: 第几列
```

### grep



### awk








### 开发 三个客流报表

- 抓拍记录查询: done
- 去重客流查询：done
- 识别列表记录：时间：since , then // done

62c10ad7e78ac62de163abb27ef04cfa3e15637c
01962a6c730b7b5cf4d9ecc4484aff4f782a6c2d
4ea885624796c4b166cd0b3e07bae8e2a2f9749b


01962a6c730b7b5cf4d9ecc4484aff4f782a6c2d