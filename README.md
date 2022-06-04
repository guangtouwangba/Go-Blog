# Go-Blog

## 配置文件
项目启动会读取config目录下的`config.yaml`文件，配置文件内容如下：
```yaml
httpserver:
  port: 8080
  host: 主机地址，默认0.0.0.0
```

## 启动
方法1：
修改配置文件config.yaml中database的地址为`db:3306`运行`start.sh`或者`docker-compose up -d`即可启动项目

方法2：
本地启动 (适用于本地开发)：
```shell
# 启动mysql
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
make dev
```

## 数据库字段更新
该项目基于`gorm`的automigration的功能，在项目启动时会自动创建未存在的数据库，后续如需更新数据库字段，可以将`config.yaml`中的
`update`设为true，然后重新启动项目即可。
弊端：`update`在字段无需更新时需手动设为`false`，否则会造成数据丢失。


TODO List:

基础设计&设施：
- [ ] 依赖倒置，重新设计依赖管理，使用依赖倒置的方法注入外部依赖
- [ ] k8s脚本准备
- [ ] 前端低代码框架替换

技术方案：
- [ ] 完善jwt+redis过期认证过期接口


## 需求目录：
### 用户部分
- [ ] 用户点赞
- [ ] 用户评论
- [ ] 用户创建个人信息
- [ ] 用户编辑个人信息
- [ ] 用户上传头像

### 内容展示部分
- [ ] 发布评论
- [ ] 发布博客/文章
- [ ] 编辑已经发布的文章
- [ ] 根据关键字搜索文章
- [ ] 根据内容进行搜索文章
- [ ] 展示热搜关键字