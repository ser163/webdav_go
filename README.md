# webday_go

  一个简易的go版 webdav 服务程序。


## 编译安装

### 下载源码
```
  git clone https://github.com/ser163/webdav_go
```
### 进入目录
```
   cd webdav
```
### 编译源码
```
   go mod download
   go build -ldflags="-s -w"
```
### 运行程序
```
   ./webDav_go .
```


## 使用帮助

```
Usage of webDav_go:
  -a string
        地址
  -p string
        共享路径 (default ".")
  -port int
        端口 (default 8080)
```


### 1.直接输入路径
```
  webDav_go c:\ or webDav_go . (当前目录)
```
### 2.使用IP和端口号
```
   webDav_go -a "192.168.0.11" -port 8081
```
### 3.使用path定义路径
```
   webDav_go -p "/opt/data/webdav" -a "192.168.0.11" -port 8081
```

### 4.开启日志
```
   webDav_go -log true -F "/var/log/webdav.log"
```
### 5.只读模式
```
   webDav_go -R true
```
### 6.用户验证
```
   webDav_go -user "admin" -pass "123"
```
### 7.https模式
```
   webDav_go -ssl true -ssl-key "key.pem" --ssl-cert "cert.pem" -prot 443
```
### 8.Linux下在后台运行
```
   nohup ./webDav &
```

## 挂载方法
### windows下挂载
  打开命令提示工具,从运行中输入cmd命令
```
   net use Y: http://127.0.0.1:8080/ 123  /user:admin /persistent:YES
```
  123 为服务器密码. http://127.0.0.1:8080/ 为服务器路径 persistent 为保存链接状态

> 从Windows Vista起，微软就禁用了http形式的基本WebDAV验证形式（KB841215），必须使用https连接
  HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\WebClient\Parameters
找到BasicAuthLevel把这个值从1改为2，然后进控制面板，服务，把WebClient服务重启（没有启动的就启动它）。

### Linux下挂载
```
   sudo apt install davfs2
   sudo mount.davfs http://locaohost:8080/dav/ /opt/Backup/
```


## 参考

此程序参考: [Golang 实现简单WebDAV系统](https://www.cnblogs.com/singinger/p/13433780.html)


