# zerorpc 

[![Build Status](https://travis-ci.org/ylxb23/zerorpc.svg?branch=main)](https://travis-ci.org/ylxb23/zerorpc)
[![codecov](https://codecov.io/gh/ylxb23/zerorpc/branch/main/graph/badge.svg)](https://codecov.io/gh/ylxb23/zerorpc)
![GoReport](https://goreportcard.com/badge/github.com/ylxb23/zerorpc)


## 使用方法
初始化并启动RPC server端：
```go
package main

import (
    "fmt"
    "github.com/sirupsen/logrus"
    "github.com/ylxb23/zerorpc"
)

const serverAddr = ":9999"

func main() {
    server := zerorpc.NewServer(serverAddr)
	logrus.Infof("rpc server start at%v", serverAddr)
	student := Student{
		Name: "Nick",
	}
	// registry method
	server.Register("sayHello", student.SayHello)
	// start rpc server
	server.Run()
}

type Student struct {
	Name string
}

func (s *Student) SayHello(word string) string {
	return fmt.Sprintf("Hello, %s, I'm %s.", word, s.Name)
}
```

启动RPC client端：
```go
package main

import (
    "github.com/sirupsen/logrus"
    "github.com/ylxb23/zerorpc"
)

const serverAddr = ":9999"
// func declare
var sayHello func(word string) string

func main() {
	// init client
	rpcClient := zerorpc.NewGrettyClient(serverAddr)
	rpcClient.ConfigRemoteMethod("sayHello", &sayHello)

	res := sayHello("word") // "Hello, word, I'm Nick."
	logrus.Infof("test res:%v", res)
	rpcClient.Destroy()
}
```

## 特性
go语言的简易rpc框架，通过反射实现远端服务的调用。压测结果支持24k+ qps：
```text
goos: windows
goarch: amd64
pkg: github.com/ylxb23/zerorpc
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkCallRpc
time="2021-07-25T20:23:04+08:00" level=info msg="benchmark call rpc finished. invoke times:1"
time="2021-07-25T20:23:04+08:00" level=info msg="benchmark call rpc finished. invoke times:100"
time="2021-07-25T20:23:04+08:00" level=info msg="benchmark call rpc finished. invoke times:10000"
time="2021-07-25T20:23:31+08:00" level=info msg="benchmark call rpc finished. invoke times:1000000"
time="2021-07-25T20:24:53+08:00" level=info msg="benchmark call rpc finished. invoke times:2708848"
BenchmarkCallRpc-8       2708848             30249 ns/op
PASS
ok      github.com/ylxb23/zerorpc       109.366s
```
