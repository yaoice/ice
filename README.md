# k8s api代理

> 这是一个封装了原生k8s api的代理.

### 使用说明

1. 把k8s认证的.kube/config文件放到app/config位置
2. 安装godep: `go get github.com/tools/godep`
3. 安装go依赖: `go restore`
4. 运行: `bee run -gendoc=true -downdoc=true`
5. 界面查看接口: `http://<server_ip>:8080/swagger/`
