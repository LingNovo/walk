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
## 程序设计
程序主要采用面向对象思想进行业务处理
### 抽象出业务结构FileInfoEx 文件信息对象，扩展结构体函数：
1. format() 格式化文件信息FileInfoEx实例为字符串
2. getHash() 获取文件哈希值
3. writeToFile() 将格式化后的FileInfoEx 信息写入指定文件
4. NewFileInfoEx() 构造函数，实例化一个FileInfoEx文件信息

### 业务处理函数：
1. Walk() 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。

### 业务处理思路
1. 初始化程序，接收控制台参数。
2. 遍历目录获取文件信息，eg：名称、哈希值、大小。
3. 将文件信息追加到指定文件。

### 遇到的问题
1. 向指定文件追加内容时无法换行，解决办法：内容末尾加“\r\n”。

## 运行实例：
```
cd core
go build
cd ..
go run init.go main.go -root="d:\\develop\\res"
```

## 测试报告

对应用程序分别进行功能测试、基准测试、覆盖率测试，以下是测试结果

### 功能测试
1. 测试目的：确保程序稳定运行，功能性代码可以运行通过
2. 测试结果：
```
$ go test -v
=== RUN   Test_NewFileInfoEx
--- PASS: Test_NewFileInfoEx (0.22s)
=== RUN   Test_Walk
--- PASS: Test_Walk (0.40s)
PASS
ok      walk/core       0.702s
```

### 基准测试
1. 测试目的：测试程序性能
2. 测试结果：
```
$ go test -bench="." -v
=== RUN   Test_NewFileInfoEx
--- PASS: Test_NewFileInfoEx (0.20s)
=== RUN   Test_Walk
--- PASS: Test_Walk (0.35s)
Benchmark_NewFileInfoEx-4       2000000000               0.13 ns/op
Benchmark_Walk-4                2000000000               0.17 ns/op
PASS
ok      walk/core       7.792s
```

### 覆盖率测试
1. 测试目的：测试代码完整性
2. 测试结果：
```
$ go test -cover
PASS
coverage: 56.4% of statements
ok      walk/core       0.564s
```
### 综合分析
1. 性能有待优化
2. 可扩展性、维护性较强
3. 程序结构复杂圈数较低
4. 耦合性较低
