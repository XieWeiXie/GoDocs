# 算法部署和迁移




## 部署: GPU

容器化部署:

- 基础镜像：nvidia-go
- 下载 DeepFaceServer 
- 生成 License 
- 启动

> 详见： aiface 项目 run.sh




GPU 运行： titian-api-go 、DeepServer、fastSearch


CPU 运行：业务项目 和 titan-api-go 交互



## 迁移

- 获取需要特征迁移的账号
- 根据账号获取得到组：vip 组、回头客组
- 运行脚本：迁移特征值



> old 版本

- 数据库连接



account 
```
18818687626 13818055668 18576687168 18349320426 13570303801 18058252400
```



go run main.go -src_ip='192.168.1.94' -db_port='5432' -db_user='root' -db_password='Read123@P2' -db_name='titanhq_production'