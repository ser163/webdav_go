# webday_go

  一个简易的go版 webdav 服务程序。


## 未来需要支持的功能

- [ ] https支持
 
- [ ] 加入http参数验证 


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

