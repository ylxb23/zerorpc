package zerorpc

import "github.com/sirupsen/logrus"

var rpcClient *GrettyClient

func init() {
	// start rpc server
	server := NewServer(serverAddr)
	logrus.Infof("rpc server start at%v", serverAddr)
	student := Student{
		Name: "Nick",
	}
	// registry method
	server.Register("sayHello", student.SayHello)
	go server.Run()

	// init client
	rpcClient = NewGrettyClient(serverAddr)
	rpcClient.ConfigRemoteMethod("sayHello", &sayHello)
}
