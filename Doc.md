
## 部署环境要求：

- Golang 环境
- Node 环境
## git仓库：
http://gitlab.hulian120.cn:1980/jkhh50/ai-assistant.git

## 编译前端资源
分别在frontendAdmin、frontendCustomer文件夹安装依赖和编译
```
npm i
npm run build
```
## 编译后端二进制文件
在主目录执行
```
go build ai-admin
```
执行上面命令会自动下载依赖并编译生成ai-admin二进制文件

## 运行
可以直接命令行运行或者使用pm2之类的进程管理器
```
./ai-admin  -port=8088
```

## 参数说明

- port string
监听端口号，默认8088 (default "8088")

## 访问后台
使用路径   `你的域名/admin` 访问，然后注册用户，用户名不可重复
如果注册用户名为 `Admin`会自动成为管理员，并且是启用状态，否则都是未启用用户





## 交叉编译
在win 11 cmd或者powershell执行以下命令

set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
然后正常go build即可

再切换回windows

go env -w GOARCH=amd64
go env -w GOOS=windows
