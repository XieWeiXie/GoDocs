# Makefile

> 项目编译、构建工具，类似于 shell 命令

``` 

target ... : prerequisites ...
    command
    ...
```

- target - 编译文件要生成的目标
- prerequisites - 编译文件需要的依赖
- command - 依赖生成目标所需要执行的命令（任意的 shell 命令），Makefile 中的命令必须以 [tab] 开头


``` 
BINARY=project

default:
    go build -o ${BINARY} -tags=jsoniter

.PHONY:  default

```