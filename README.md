# go-common

go 工程公共模块，抽取工具类、初始化配置等。

## 1.本地引入

```shell script
go get github.com/cuilan/go-common
```

### 项目中 go.mod 引入

```
module xxx

go 1.16

require (
    github.com/cuilan/go-common v0.0.1
    ...
)
```

版本号默认为git提交记录的版本号，也可以选择 `git tag` 版本号。

---

## 2.发布新版本

### 版本号

默认版本号为 git commit 版本号，不直观也不便于维护，使用 `git tag` 标签来命名版本号。

格式为：`vX.Y.Z`（Major.Minor.Patch），如：`v0.0.1`

* X: 表示主版本号，当兼容性变化，或重大功能升级时，X 需递增，0 为开发阶段。
* Y: 表示次版本号，当增加功能时(不影响兼容性)，Y 需递增。
* Z: 表示修订号，当做 Bug 修复时(不影响兼容性)，Z 需递增。

### 提交新版本

在自己的分支修改代码并提交后，发起合并至 **`master`** 分支，确保主分支代码为稳定版，再由 **`master`** 分支打 **`tag`** 并推送至远程仓库。

```shell
git checkout master

git pull origin master

git tag -a vx.y.z -m "vx.y.z release notes"

git push origin vx.y.z

```

---

## 3.模块说明

### entity

数据库表对应的实体。

### load

#### load_logger_conf.go

读取日志配置，并自动装配，依赖 `log.conf/app.conf` 配置文件，如下：

```
[log]
; 日志级别
log_level = 7
; 日志存放目录
log_path = logs/app.log
; 日志文件最大行数
maxlines = 3000
; 日志文件大小限制
maxsize = 10240000
```

代码中调用：
```go
import "go-common/load"

func init() {
    // 读取并初始化配置
    load.LoadLoggerConfig("../conf/log.conf")
}
```

#### load_mysql_conf.go

自动初始化、加载数据库配置，依赖 `mysql.conf/mysql.ini` 配置文件，如下：

```
[mysql]
host=localhost
user=root
password=123456
port=3306
db=weattech_dl
charset=utf8mb4
maxactive=10
maxidle=10

# 多数据配置，格式：[mysql-数据源名称]，同一项目中保证不重复即可
[mysql-xxx]
host=192.168.0.100
user=test
password=test123
port=3306
db=test_dl
charset=utf8mb4
maxactive=1
maxidle=100
```

代码中调用：
```go
import "go-common/load"

func init() {
    // 读取配置，并初始化数据源
    load.LoadMysqlConfig("./conf/mysql.conf")

    // 读取任意数据源配置
    source, maxactive, maxidle := load.LoadMysqlWithDataSource("../conf/mysql.conf", "mysql-haulers")
}
```

### orm

#### orm.go

初始化 MySQL 数据库连接信息。如果已经使用 `go-common/load` 包中的 `LoadMysqlConfig()` 方法，则可以省略此步骤。

### utils

#### common_util.go

基础工具方法

#### file_util.go

文件操作方法集合

#### generate_key_util.go

生成秘钥工具

#### minio_util.go

文件对象存储工具

#### notice.go

飞书机器人异常信息通知工具
