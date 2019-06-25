# AiFace


> 算法服务容器化


## 项目目录

```
- config: 配置文件，用于认证权限 license
- DeepFaceServer: 算法服务，二进制服务，运行在 GPU 上
- deployments: dockerfile 文件
- log: 日志文件
- web_server: 健康检测， 运行一个检测图片的二进制文件，是否返回正确的结果
- run.sh: 
    - 下载 deepfaceserver.bin
    - 拷贝配置文件
    - 

```