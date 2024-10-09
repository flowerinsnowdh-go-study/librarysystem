# librarysystem
这是一个学习语言的项目，无实际意义

项目内容是模拟一个图书馆系统，使用 SQLite 储存数据

# 系统介绍
每名学生可以同时借阅 1 本书

每本书可以同时被一名玩家借阅

# 编译
```shell
go build
```

# 添加学生
```sqlite
INSERT INTO `student` (`name`) VALUES ('书名')
```

# 添加书
```sqlite
INSERT INTO `book` (`name`) VALUES ('书名')
```

# 下一步
准备写成 HTTP 网页端服务器
