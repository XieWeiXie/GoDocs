## Docker


- Dockerfile

> 编写单个镜像的文件

- Docker command

> 操作 镜像、容器、服务等命令

- Docker compose

> 编写复杂镜像的文件


```
docker ps -a
docker ps -l
docker images
docker attach | ctrl + P | ctrl + Q
docker run -it -d alpine /bin/sh
docker exec -it *** /bin/sh
docker stop
docker kill
exit
docker top dockername



```




### 常用的镜像的使用

- postgres

```
>> docker run -d -p 5432:5432 ******** -d postgres
>> docker exec -it 26bdbd43c30f /bin/sh
>> su postgres
>> psql -U postgres
>> createdb databaseName

```

- redis

```
>> docker run --name some-redis -d redis -p 6379:6379 redis-server --appendonly yes
>> docker exec -it ***** redis-cli
>> docker exec -it ****** /bin/sh
>> redis-cli

```

### 维护日常的容器的使用

> docker/src/mongo-docker/docker-compose.yaml： 镜像启动文件
> docker/src/mongo-docker/mongo-data : 挂载目录

### 如何学会新的镜像的使用

1. docker hub 查看官方的介绍和示例
2. 编写 Dockerfile 文件
3. 每个镜像的启动以 docker-compose 的方式运行，多个容器更应该这么做

### Dockerfile 指令

- FROM
- LABEL MAINTAINER="xiewei(1156143589@qq.com)"
- WORKDIR
- RUN
- COPY(ADD)
- CMD
- ENDPOINT
- ENV
- EXPOSE
- VOLUME


### Dockerfile golang 版模版

```$xslt
FROM golang:1.9.4
LABEL MAINTAINER="XieWei(1156143589@qq.com)"

WORKDIR ["/go/src/projectName"]

COPY vendor.json vendor/vender.json

COPY . .

RUN go get ...

RUN make install

CMD ["bash", "-c", "/go/src/projectName/..."]



```

```$xslt
docker build . -t=project -f path
```