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
