#go blog test！
仅后端接口，无前端展示
#部署文档（需要有go环境和go mod）
git clone https://github.com/James2333/gin-blog.git
go mod download
go mod tidy
修改models/models.go里的数据库文件
go run main.go

#更新了tokne验证，需要先请求auth，带账号密码，get返回的token值
请求中需要 带有token值才能修改文章。

#数据库表可配置自动生成
在models下的代码中都加入db.AutoMigrate(&结构体)
这样会自动创建表

