package syncpool

import (
	"sync"
	"testing"
)

var globalP *Packet

func usePacket(p *Packet){
	p.from = 1111
	p.to = 2222
	p.amount = 4.56
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Packet)
	},
}

type Packet struct {
	from int
	to int
	amount float64
}


func BenchmarkUsePool(b *testing.B){
    for i:= 0;i < b.N; i++{
    	for j:= 0; j < 10000; j++{
			p := pool.Get()
			usePacket(p.(*Packet))
			globalP = p.(*Packet)
			pool.Put(p)
		}
	}
}

func BenchmarkNo(b *testing.B){
    for i:= 0;i < b.N; i++{
    	for j:= 0;j< 10000; j++{
			p := new(Packet)
			usePacket(p)
			globalP = p
		}
	}
}


/*
1.每轮循环使用gc这种情况下缓存对象作用不大
func BenchmarkUsePool(b *testing.B){
    for i:= 0;i < b.N; i++{
    	for j:= 0; j < 10000; j++{
    		runtime.GC()
			p := pool.Get()
			usePacket(p.(*Packet))
			globalP = p.(*Packet)
			pool.Put(p)
		}
	}
}

go test -bench=. -benchmem -run=none -cpu=1 -benchtime=10s
结果
BenchmarkUsePool              20         641084512 ns/op         1680000 B/op      30000 allocs/op
BenchmarkNo               100000            191469 ns/op          320000 B/op      10000 allocs/op


2.不使用gc
goos: darwin
goarch: amd64
pkg: go-bpt/sync/syncpool
BenchmarkUsePool          100000            162737 ns/op               0 B/op          0 allocs/op
BenchmarkNo               100000            215103 ns/op          320000 B/op      10000 allocs/op
PASS
ok      go-bpt/sync/syncpool    41.459s


3.关闭gc
GOGC=off go test -bench=. -benchmem -run=none -cpu=1 -benchtime=10s
goos: darwin
goarch: amd64
pkg: go-bpt/sync/syncpool
BenchmarkUsePool          100000            155887 ns/op               0 B/op          0 allocs/op
BenchmarkNo                50000            425391 ns/op          320000 B/op      10000 allocs/op
PASS
ok      go-bpt/sync/syncpool    43.272s
*/