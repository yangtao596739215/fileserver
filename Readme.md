## 功能说明
```
把执行命令的目录以http的方式暴露出来，可以通过浏览器下载目录上的文件
```

## 使用说明
go run file.go -addr=:9999

go build file.go && file -addr=:9999

go install github.com/yangtao596739215/fileserver && fileserver -addr :8080


两种命令行参数:
```
-addr :8089  
-addr=:8089 
```

## 注意事项
```
go 1.16以上才支持go install
go install 的路径是 GOBIN，如果没设置则是GOROOT/bin
需要把go的install目录加入到环境变量中，这样就可以在任何目录直接使用了
```