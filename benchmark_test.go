package zerorpc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync/atomic"
	"testing"
)

func BenchmarkCallRpc(b *testing.B) {
	b.ResetTimer()
	b.SetParallelism(runtime.NumCPU())
	b.StartTimer()
	var count uint32 = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddUint32(&count, 1)
			logrus.Debugf("call result:%s", sayHello(fmt.Sprintf("Got num: %d", count)))
		}
	})
	b.StopTimer()
	logrus.Infof("benchmark call rpc finished. invoke times:%d", count)
}
