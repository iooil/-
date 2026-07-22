# 校园选课评价系统

基于 **Go + Gin + GORM + MySQL** 与 **微信小程序（WXML / WXSS / JavaScript）** 开发的前后端分离校园选课评价系统。

本项目面向高校学生选课与教学评价场景，实现学生注册登录、课程浏览与筛选、在线选课与退课、教师信息查询、课程评价、教师评价及个人评价记录查询等功能。

---

## 项目架构

```text
微信小程序前端
WXML + WXSS + JavaScript
        │
        │ HTTP / JSON
        ▼
Go 后端服务
Gin + JWT + bcrypt + GORM
        │
        ▼
MySQL 8.0
```

主要业务流程：

```text
用户注册/登录
      ↓
课程浏览与筛选
      ↓
在线选课 / 退课
      ↓
查看教师信息
      ↓
评价课程 / 评价教师
      ↓
查看“我的评价”
```

## 技术栈

### 前端
- 微信小程序原生框架
- WXML
- WXSS
- JavaScript
- wx.request
- 本地缓存 Token

### 后端
- Go
- Gin
- GORM
- JWT
- bcrypt

### 数据库
- MySQL 8.0

## 已实现功能

### 用户模块
- 学生注册与登录
- bcrypt 密码加密存储
- JWT Token 身份认证
- 用户信息查询
- 首页数据统计

### 课程模块
- 课程列表展示
- 课程搜索与筛选
- 在线选课
- 防止重复选课
- 课程容量校验
- 我的课程查询
- 退课

### 教师模块
- 教师列表
- 教师搜索与筛选
- 教师详情
- 所授课程展示
- 综合评分展示

### 评价模块
- 课程星级评价
- 教师星级评价
- 文字评价
- 评价资格校验
- 防止重复评价
- 我的课程评价记录
- 我的教师评价记录

## 项目目录

```text
Life/
├─ miniapp/
│  ├─ images/
│  ├─ pages/
│  │  ├─ login/
│  │  ├─ register/
│  │  ├─ home/
│  │  ├─ courses/
│  │  ├─ teachers/
│  │  ├─ teacher-detail/
│  │  ├─ my-courses/
│  │  ├─ course-review/
│  │  ├─ teacher-review/
│  │  └─ my-reviews/
│  ├─ utils/
│  │  └─ request.js
│  ├─ app.js
│  ├─ app.json
│  └─ app.wxss
│
└─ campus-course-server/
   ├─ config/
   │  └─ database.go
   ├─ controller/
   │  ├─ auth_controller.go
   │  ├─ user_controller.go
   │  ├─ course_controller.go
   │  ├─ teacher_controller.go
   │  ├─ my_course_controller.go
   │  ├─ review_controller.go
   │  └─ my_review_controller.go
   ├─ middleware/
   │  └─ auth.go
   ├─ model/
   │  ├─ user.go
   │  ├─ course.go
   │  ├─ teacher.go
   │  └─ review.go
   ├─ utils/
   ├─ go.mod
   └─ main.go
```

## 数据库设计

系统包含 6 张核心数据表：

| 数据表 | 说明 |
|---|---|
| users | 学生用户信息 |
| teachers | 教师基础信息 |
| courses | 课程基础信息 |
| course_selections | 学生选课记录 |
| course_reviews | 课程评价记录 |
| teacher_reviews | 教师评价记录 |

## 主要接口

### 用户认证
| 方法 | 接口 | 功能 |
|---|---|---|
| POST | /api/auth/register | 学生注册 |
| POST | /api/auth/login | 学生登录 |

### 用户信息
| 方法 | 接口 | 功能 |
|---|---|---|
| GET | /api/user/profile | 获取学生信息 |
| GET | /api/user/dashboard | 获取首页统计数据 |

### 课程
| 方法 | 接口 | 功能 |
|---|---|---|
| GET | /api/courses | 获取课程列表 |
| GET | /api/courses/filters | 获取课程筛选项 |
| POST | /api/courses/:id/select | 选择课程 |
| GET | /api/my/courses | 获取我的课程 |
| POST | /api/my/courses/:id/drop | 退课 |

### 教师
| 方法 | 接口 | 功能 |
|---|---|---|
| GET | /api/teachers | 获取教师列表 |
| GET | /api/teachers/filters | 获取教师筛选项 |
| GET | /api/teachers/:id | 获取教师详情 |

### 评价
| 方法 | 接口 | 功能 |
|---|---|---|
| POST | /api/reviews/course | 提交课程评价 |
| POST | /api/reviews/teacher | 提交教师评价 |
| GET | /api/my/reviews | 获取我的评价 |

## 本地运行

### 环境要求
- Go 1.20+
- MySQL 8.0+
- 微信开发者工具
- Git

### 创建数据库

```sql
CREATE DATABASE campus_course
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;
```

### 配置数据库连接

修改：

```text
campus-course-server/config/database.go
```

示例：

```go
dsn := "root:你的密码@tcp(127.0.0.1:3306)/campus_course?charset=utf8mb4&parseTime=True&loc=Local"
```

> 不要将真实数据库密码提交到公开仓库，推荐使用环境变量。

### 安装依赖并启动后端

```bash
cd campus-course-server
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
go run .
```

启动成功后测试：

```text
http://127.0.0.1:8080/api/hello
```

### 配置小程序请求地址

修改：

```text
miniapp/utils/request.js
```

模拟器调试：

```javascript
const BASE_URL = 'http://127.0.0.1:8080'
```

真机调试时改为电脑局域网 IPv4，例如：

```javascript
const BASE_URL = 'http://192.168.1.100:8080'
```

并确保：
- 手机与电脑连接同一网络
- Go 后端监听 :8080
- Windows 防火墙放行 8080 端口
- 开发阶段启用“不校验合法域名”

## 测试账号

```text
学号：2023040100
密码：123456
```

> 公开仓库只应保留测试账号，不要上传真实用户信息。

## 项目截图

建议新建：

```text
docs/images/
```

并加入：

```text
login.png
home.png
courses.png
teachers.png
my-courses.png
my-reviews.png
```

示例：

```markdown
![登录页面](docs/images/login.png)
![学生首页](docs/images/home.png)
![课程中心](docs/images/courses.png)
```

## 开发过程中解决的问题

- Gin 依赖下载失败：配置 GOPROXY 并执行 go mod tidy
- GORM 扫描 gin.H 报错：改用专用结构体接收查询结果
- 小程序页面跳转失败：在 app.json 中统一注册页面路径
- 教师评价缺少关联字段：补充 teacher_id 的前后端传递
- 真机无法访问 127.0.0.1：改用电脑局域网 IPv4
- 评价重复提交：后端增加用户、课程和教师组合校验

## 后续计划

- 增加管理员后台
- 实现课程与教师信息 CRUD
- 增加评价审核与删除功能
- 完善评分统计
- 增加课程推荐与排行榜
- 部署到云服务器并配置 HTTPS
- 优化移动端适配和异常提示

## 项目说明

本项目为生产实习课程设计项目，主要用于学习和展示：

- Go Web 后端开发
- Gin 路由与中间件
- GORM 数据库操作
- JWT 身份认证
- 微信小程序开发
- 前后端接口联调
- MySQL 数据库设计

## License

本项目仅用于学习与课程设计展示。
