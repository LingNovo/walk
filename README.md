# walk
面试题

请仔细看清题目的要求，仔细看清题，仔细看清题！！一个生成目录树哈希的小工具

## 小工具需求说明
1. 用golang/js开发，代码放到github上，用github进行问题跟踪
2. 对整个目录下的所有文件进行遍历，获取所有文件的大小和计算文件的sha1哈希值，记录在一个文件里面
3. 结果文件格式：每一行一个文件，用逗号隔开，前面是文件名称，后面是哈希值，文件大小
4. 需要可以指定忽略哪些目录、文件，需要支持通配符
5. 代码实现简洁，运行性能高得分高
6. 要求通过测试代码自我证明代码能够可靠运行并正确实现上述功能
7. 要求写安全稳定可靠的代码
8. 要求支持go get, go run, go build命令
9. 不限时间，做好了就提交github，告诉我链接

# 描述

## 目录清单

```
- walk (根目录)
-- core（业务包）
---  cover.out（覆盖率测试结果文件）
---  file_test.go（文件信息测试）
---  file.go（文件信息）
---  logic_test.go（逻辑处理测试）
---  logic.go（逻辑处理）
---  pool_test.go(线程池测试)
---  pool.go(线程池)
---  record.go(记录器)
-- init.go（程序初始化）
-- mian.go（程序入口）
```
## 程序设计
程序主要采用面向对象思想进行业务处理
### 抽象出业务结构fileInfoEx 文件信息对象，扩展结构体函数：
1. format() 格式化文件信息fileInfoEx实例为字符串
2. NewFileInfoEx() 构造函数，实例化一个fileInfoEx文件信息

### 设计线程池goPool 对象，扩展结构体函数：
1. Init() 初始化线程池，指定线程数量及任务总数
2. Start() 启动线程池，并行执行任务
3. Stop() 终止所有任务，释放资源
4. Start() 启动线程池，并行执行任务
5. AddTask() 添加一个任务
6. SetFinshCallback() 设置回调函数

### 抽象recorder 文件信息记录器对象，扩展结构体函数：
1. newRecorder() 记录器构造函数，实例化一个recorder 对象
2. createFile() 创建记录文件，初始化记录器
3. Write() 写入数据
4. Close() 关闭记录器，释放资源

### 业务处理函数：
1. Walk() 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。

### 业务处理思路
1. 初始化程序，接收控制台参数。
2. 遍历目录获取文件信息，eg：名称、哈希值、大小。
3. 初始化线程池，添加并行任务，将文件信息追加到指定文件。

### 遇到的问题
1. linux 环境下打开文件数量上限默认为１０２４．超出时程序报“too many open files”异常，解决办法，线程数设小一点，或者更改系统设置。

## 运行实例：

获取目录"/home/bo/develop/"下的文件信息，忽略目录名称包含"go"字符 的目录，忽略后缀为".go"的文件。

```
cd core
go build
go install
cd ..
go run init.go main.go -root="/home/bo/develop/" -ignore="go" -suffix="*.md"
```

## 测试报告

对应用程序分别进行功能测试、基准测试、覆盖率测试，以下是测试结果

### 功能测试
1. 测试目的：确保程序稳定运行，功能性代码可以运行通过
2. 测试结果：
```
go test -v
=== RUN   Test_newFileInfoEx
--- PASS: Test_newFileInfoEx (0.00s)
=== RUN   Test_Walk
--- PASS: Test_Walk (0.84s)
=== RUN   Test_Pool
--- PASS: Test_Pool (0.00s)
=== RUN   Test_Write
--- PASS: Test_Write (0.00s)
PASS
ok  	walk/core	0.851s

```

### 基准测试
1. 测试目的：测试程序性能
2. 测试结果：
```
$ go test -bench="." -v
=== RUN   Test_newFileInfoEx
--- PASS: Test_newFileInfoEx (0.00s)
=== RUN   Test_Walk
--- PASS: Test_Walk (0.79s)
=== RUN   Test_Pool
--- PASS: Test_Pool (0.00s)
=== RUN   Test_Write
--- PASS: Test_Write (0.00s)
Benchmark_newFileInfoEx-4   	2000000000	         0.00 ns/op
Benchmark_Walk-4            	2000000000	         0.40 ns/op
Benchmark_Pool-4            	2000000000	         0.00 ns/op
Benchmark_Write-4           	2000000000	         0.00 ns/op
PASS
ok  	walk/core	23.402s
```

### 覆盖率测试
1. 测试目的：测试代码完整性
2. 测试结果：
```
$ go test -cover
PASS
coverage: 77.5% of statements
ok  	walk/core	0.838s

```
### 综合分析
1. 性能有待优化
2. 可扩展性、维护性较强
3. 程序结构复杂圈数较低
4. 耦合性较低
