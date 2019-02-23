# travis


[Travis](https://docs.travis-ci.com)

> 语法

- install 安装依赖
- script 运行脚本

``` 
install: true // 直接跳过安装依赖
script: true  // 直接跳过运行脚本

```

### 具体语法


#### 单个命令

``` 
install: command one
script: command one 

```

#### 多个命令

``` 
install:
    - command one
    - command two
script:
    - command one
    - command two
    
```

``` 
install: command one && command two
script: command one && command two

```

#### 存在默认执行命令

比如： go 


#### notification


#### deployment


#### 钩子

- before
- after


#### 环境变量

- env

#### 整体生命周期

``` 

before_install
install
before_script
script
aftersuccess or afterfailure
[OPTIONAL] before_deploy
[OPTIONAL] deploy
[OPTIONAL] after_deploy
after_script
```