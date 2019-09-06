package syncpool

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
并发使用
*/

type Payload struct {
	length int
	header int
	body int
}

var P = new(Payload)

var poolSync = sync.Pool{
	New: func() interface{} {
		return new(Payload)
	},
}

func doWithNoPool(){
	var wait sync.WaitGroup
	wait.Add(100)
	for i:= 0;i < 100; i++{
		go func() {
			for j:= 0;j < 1000; j++{
				time.Sleep(1 * time.Millisecond)
				p := new(Payload)
				P = p
			}//for
			wait.Done()
		}()
	}
	wait.Wait()
}

func doWithPool(){
	var wait sync.WaitGroup
	wait.Add(100)
	for i:= 0;i < 100; i++{
		go func() {
			for j:= 0; j< 1000; j++{
				time.Sleep(1 * time.Millisecond)
				p := poolSync.Get()
				pay := p.(*Payload)
				P = pay
				poolSync.Put(p)
			}
			wait.Done()
		}()
	}
	wait.Wait()
}

func BenchmarkCUsePool(b *testing.B){
	for i := 0;i< b.N; i++{
		runtime.GC()
		doWithPool()
	}//for
}

func BenchmarkCNo(b *testing.B){
	for i := 0;i< b.N; i++{
		runtime.GC()
		doWithNoPool()
	}//for
}


func BenchmarkGc(b *testing.B){
	for i:= 0;i < b.N; i++{
		runtime.GC()
	}
}


/*
goos: darwin
goarch: amd64
pkg: go-bpt/sync/syncpool
BenchmarkCUsePool             10        1364522359 ns/op            6432 B/op        101 allocs/op
BenchmarkCNo                  10        1381827245 ns/op         3206416 B/op     100101 allocs/op

*/
