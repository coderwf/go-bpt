package bound

import (
	"encoding/xml"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)


var docs = generateList(1000)
var findNum = find("Go", docs)


func read(doc string) ([]item, error) {
	time.Sleep(time.Millisecond) // Simulate blocking disk read.
	var d document
	if err := xml.Unmarshal([]byte(file), &d); err != nil {
		return nil, err
	}
	return d.Channel.Items, nil
}

func find(topic string, docs []string) int {
	var found int
	for _, doc := range docs {
		items, err := read(doc)
		if err != nil {
			continue
		}
		for _, item := range items {
			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}
	return found
}

func findConcurrent(goroutines int, topic string, docs []string) int {
	var found int64

	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			var lFound int64
			for doc := range ch {
				items, err := read(doc)
				if err != nil {
					continue
				}
				for _, item := range items {
					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
			atomic.AddInt64(&found, lFound)
			wg.Done()
		}()
	}

	wg.Wait()

	return int(found)
}


func BenchmarkSingleThreadIo(b *testing.B){
	for i:= 0; i<b.N; i++{
		n := find("Go", docs)
		if n != findNum{
			b.Errorf("single thred n(%d) != findNum(%d)", n, findNum)
		}
	}
}

func BenchmarkConcurrentCpuNumIo(b *testing.B){
	for i:= 0; i<b.N; i++{
		n := findConcurrent(runtime.NumCPU(),"Go", docs)
		if n != findNum{
			b.Errorf("concurrent n(%d) != findNum(%d)", n, findNum)
		}
	}
}


/*
****************************************

GOGC=off go test -cpu 1 -run none -bench . -benchtime 5s -count=3

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThreadIo                5        1115279867 ns/op
BenchmarkSingleThreadIo                5        1114765554 ns/op
BenchmarkSingleThreadIo                5        1111183179 ns/op
BenchmarkConcurrentCpuNumIo           20         284160463 ns/op
BenchmarkConcurrentCpuNumIo           20         279193069 ns/op
BenchmarkConcurrentCpuNumIo           20         281765347 ns/op
PASS
ok      go-bpt/sync/bound       39.017s
****************************************



****************************************

GOGC=off go test -cpu 2 -run none -bench . -benchtime 5s -count=3

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThreadIo-2                      5        1113257701 ns/op
BenchmarkSingleThreadIo-2                      5        1112071081 ns/op
BenchmarkSingleThreadIo-2                      5        1111645330 ns/op
BenchmarkConcurrentCpuNumIo-2                 20         281253711 ns/op
BenchmarkConcurrentCpuNumIo-2                 20         280499506 ns/op
BenchmarkConcurrentCpuNumIo-2                 20         280125911 ns/op
PASS
ok      go-bpt/sync/bound       38.942s
****************************************



****************************************

GOGC=off go test -cpu 3 -run none -bench . -benchtime 5s -count=3

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThreadIo-3                      5        1113007330 ns/op
BenchmarkSingleThreadIo-3                      5        1112950285 ns/op
BenchmarkSingleThreadIo-3                      5        1111904997 ns/op
BenchmarkConcurrentCpuNumIo-3                 20         280593565 ns/op
BenchmarkConcurrentCpuNumIo-3                 20         279817299 ns/op
BenchmarkConcurrentCpuNumIo-3                 20         279114416 ns/op
PASS
ok      go-bpt/sync/bound       38.893s
****************************************



****************************************

GOGC=off go test -cpu 4 -run none -bench . -benchtime 5s -count=3

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThreadIo-4                      5        1115158585 ns/op
BenchmarkSingleThreadIo-4                      5        1112329311 ns/op
BenchmarkSingleThreadIo-4                      5        1108657300 ns/op
BenchmarkConcurrentCpuNumIo-4                 20         281697669 ns/op
BenchmarkConcurrentCpuNumIo-4                 20         279351443 ns/op
BenchmarkConcurrentCpuNumIo-4                 20         279056014 ns/op
PASS
ok      go-bpt/sync/bound       38.896s
****************************************

由此可见对于io密集型操作,开启多携程能够明显的提升性能
但是多携程情况下增加cpu核数并不会有太明显的性能提升

*/