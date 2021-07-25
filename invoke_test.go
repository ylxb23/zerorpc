package zerorpc

import (
	"fmt"
	"testing"
)

const serverAddr = ":9999"

// func declare
var sayHello func(word string) string

type Student struct {
	Name string
}

func (s *Student) SayHello(word string) string {
	return fmt.Sprintf("Hello, %s, I'm %s.", word, s.Name)
}

func TestServerRun(t *testing.T) {
	res := sayHello("word")
	t.Logf("test res:%v", res)
}
