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
