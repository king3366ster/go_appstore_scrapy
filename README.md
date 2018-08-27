# go_appstore_scrapy
go语言学习——获取腾讯应用商城APP，解包分析，并存入Sqlite数据库

## 工程入口 main.go

## 开发运行
- go run main.go

## 编译工程
- go build main.go

## 文件介绍
- app_info.sql => sqlite3 数据库表，需要实现建好
- requestApp.go 爬虫
- downloadApp.go 下载APP
- analyseApp.go 分析解包
- main.go 使用go routine并发、存入数据库