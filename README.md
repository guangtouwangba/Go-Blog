# Go-Blog

## 配置文件
项目启动会读取config目录下的`config.yaml`文件，配置文件内容如下：
```yaml
httpserver:
  port: 8080
  host: 主机地址，默认0.0.0.0
```

## 启动
直接运行`start.sh`或者`docker-compose up -d`即可启动项目
