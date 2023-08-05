# memo 备忘录
## 项目介绍
**此项目使用 `G0` + `Gin` + `Gorm` + `Redis` + `JWT` + `MySQL` ，基于 `RESTful API`实现的一个备忘录**。  

## 项目功能
* 用户注册登录
* 新增/删除/修改/查询/ 备忘录
* Redis 存储每条备忘录的浏览次数
* 分页搜索功能

## 项目主要依赖
* Go
* Gin               // 连接数据库时用到的依赖：go get -u github.com/gin-gonic/gin"
* Gorm              // 连接数据库时用到的依赖：go get -u github.com/jinzhu/gorm
* MySQL             // 连接数据库时用到的依赖（mysql数据库驱动）：go get -u github.com/jinzhu/gorm/dialects/mysql
* Redis
* ini               // 读取配置文件 config.ini 时用到的依赖： go get gopkg.in/ini.v1
* jwt-go

## 项目结构
```
memo
├── api                        // 用于定义接口函数
├── cache                      // 放置 redis 缓存
├── conf                       // 处理配置数据文件夹
│   ├── conf.go                // 读取配置数据
│   └── config.ini             // 配置文件
├── middleware                 // 应用中间件  jwt
├── model                      // 应用数据库模型
│   ├── init.go                // 连接数据库
├── pkg                        // 
│   ├── e                      // 封装错误码
│   ├── logging                // 日志打印
│   └── utils                  // 工具函数
├── routes                     // 路由逻辑处理
├── serializer                 // 将数据序列化为 json 的函数
├── serive                     // 接口函数的实现
├──                            // 配置文件
├── main.go
│   │   └── main()             // 程序入口
└── README.md
```

## 项目构建运行
本项目使用go mod进行管理，因此可以通过go mod tidy下载所需的依赖，通过go run main.go启动该项目。

## 理论准备（RESTful API —— 入门理解）
**什么是 REST？**  
由 Roy Thomas Fielding 在他2000年的博士论文中提出：  
在符合架构原理的前提下，理解和评估以网络为基础的应用软件的架构设计，得到一个功能强、性能好、适宜通信的架构。  
他对**互联网软件**的架构原则，定名为 **REST**（Representational State Transfer）。  

**如何理解 REST？**  
REST = Representational + State + Transfer 表现层状态转化  
Resource：资源  
Representational：某种表现形式，比如用 JSON、XML、JPEG 等  
State Transfer：状态转化，通过 HTTP 动词实现

**什么是 API？**  
**API**(Application Programming Interface应用程序编程接口) 是一些预先定义好的函数，目的是提供应用程序以及开发人员基于某软件或硬件得以访问一组例程的能力。

**为什么会有 RESTful API？**  
前端设备层出不穷（手机、平板、电脑、其他专用设备……），促使前后端分离，方便不同的前端设备与后端进行通信，导致 API 架构的流行。  

RESTful 架构，因结构清晰、符合标准、易于理解、扩展方便，是目前流行的一种互联网软件架构。  

RESTful API 是目前比较成熟的一套互联网应用程序的 API 设计理论。

**什么是 RESTful 架构？**  
如果一个架构符合 REST 原则，就称它为 `RESTful` 架构。  
RESTful 架构，是面向资源的架构：
1. 每一个 URI 代表一种资源；
2. 客户端和服务器之间，传递这种资源的某种表现层；
3. 客户端通过 HTTP 动词，对服务器端资源进行操作，实现“表现层状态转化”。  

**什么是 RESTful API？**  
符合 REST 架构设计的 API ，就是 RESTful API 。

特征：URL 用于定位资源，用 HTTP 动词描述操作。  

常见举例：  
查询编号为1的图书：  
[GET] http://127.0.0.1:8080/v1/books/1

删除编号为1的图书：  
[DELETE] http://127.0.0.1:8080/v1/books/1

**URI 设计规则**  
**规则1**：资源表示一种实体，URI 使用名词表示，不应该包含动词；一般来说，数据库中的表都是同种记录的“集合”（collection），所以 URI 中的名词应该使用复数。  

正确举例：  
GET/zoos：列出所有动物园  
POST/zoos：新建一个动物园  
GET/zoos/ID：获取某个指定动物园的信息

错误举例：
GET/zoos/show/1  
正确写法：
GET/zoos/1  

**规则2**：如果某些动作是 HTTP 动词表示不了的，应该把动作做成一种资源。  

错误举例：网上汇款，从账号1向账户2汇款500元  
POST/accounts/1/transfer/500/to/2  

正确的写法：把动词 transfer 改成 transaction，资源不能是动词，但是可以是一种服务：  
POST/transaction  
from=1&to=2&amount=500.00

比如登陆、退出：  
POST/sessions  
DELETE/sessions/{id}

**规则3**：参数的设计允许存在冗余，即允许 API 路径和 URL 参数偶尔有重复。  
比如，GET/zoos/ID/animals 与 /animals?zoo_id=ID 的含义是相同的。  
注意：对于 /zoos/ID/animals/ID/ 方式，注意关联层次不要太深，可以将关系通过参数方式表现：/animals?zoo_id=ID  

**规则4** 一些常见的参数：  
?limit=10：指定返回记录的数量  
?offset=10：指定返回记录的开始位置  
?page=2&per_page=10：指定第几页，以及每页的记录数  
?sortby=name&order=asc：指定返回结果按照哪个属性排序，以及排序顺序  

**HTTP 动词**  
7个 HTTP 动词：GET/POST/PUT/DELETE/PATCH/HEAD/OPTIONS  
GET(SELECT)：从服务器取出资源（一项或多项）  
POST(CREATE)：在服务器新建一个资源  
PUT(UPDATE)：在服务器更新资源（客户端提供改变后的完整资源）  
PATCH(UPDATE)：在服务器更新资源（客户端提供改变的属性）  
DELETE(DELETE)：从服务器删除资源  

HEAD：用于获取某个资源的元数据（metadata）  
OPTIONS：用于获取某个资源所支持的 Request 类型

**ContentType**  
用于指定请求和响应的 HTTP 内容类型。下面是几个常见的 Content-Type：  
1. text/html
2. text/plain
3. text/css
4. text/javascript
5. application/x-www-form-urlencoded
6. multipart/form-data
7. application/json
8. application/xml

前面几个是 html、css、javascript 的文件类型，后面四个是 POST 的发包方式。  

如果未指定 ContentType, GET 默认为 text/html 。  

**状态码**  
常见状态码：  
2xx 范围的状态码是保留给成功消息使用的。  
3xx 范围的状态码是保留给重定向用的。  
4xx 范围的状态码是保留给客户端错误用的，例如，客户端提供了一些错误的数据或请求了不存在的内容。这些请求应该是幂等的，不会改变任何服务器的状态。  
5xx 范围的状态码是保留给服务器端错误用的。这些错误常常是从底层的函数抛出来的，并且开发人员也通常没法处理。发送这类状态码的目的是确保客户端能得到一些响应。  

**版本化**  
一是在 URL 中包含版本信息：  
https://127.0.0.1:8080/v1/  

二是在请求头里面保存版本信息。  
Accept:vnd.example-com.foo+json;version=1.0  

相对来说，在请求头里面包含版本信息远没有放在 URL 里面来的容易。  

**返回结果**  
当使用不同的 HTTP 动词向服务器请求时，客户端需要在返回结果里面拿到一些列的信息。  
比如，当一个客户端创建一个资源时，客户端常常不知道新建资源的ID（也许还有其他属性，如创建和修改的时间戳等）。这些属性需要在随后的请求中返回，并且作为刚才 POST 请求的一个响应结果。  

下面是非常经典的 RESTful API 定义：  
GET/resources：返回一系列资源对象  
GET/resources/ID：返回单独的资源对象  
POST/resources：返回新创建的资源对象
PUT/resources：返回完整的资源对象  
PATCH/resources：返回完整的资源对象
DELETE/resources：返回一个空文档  

**认证**  
JWT（JSON Web Token）提供了一个轻量级的解决方案。  

**文档**  
必须要为 API 准备文档，否则没有人知道怎么使用它。  

不要截断文档示例中请求与响应的内容，要展示完整的东西。  

文档化每一个端点所预期的响应代码和可能的错误消息，和在什么情况下会产生这些的错误消息。

**总结**  
RESTful 架构  
是基于互联网环境的架构  
是架构风格而不是标准！  


https://gitee.com/tonghuaing/go_web_todo