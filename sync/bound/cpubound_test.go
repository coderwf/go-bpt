package bound

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var nums = genNumbers(1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 *20)

var res = addNum(nums)

func addNum(nums []int) int64{
    var num int64 = 0
    for _, n := range nums{
    	num += int64(n)
	}

    return num
}


func addNumConcurrent(nums []int, goroutines int) int64{
	var wait sync.WaitGroup

	var num int64 = 0

	var sliceLen = len(nums) / goroutines

	wait.Add(goroutines)

	for g := 0; g < goroutines; g++{
		go func(g int) {
			end := (g + 1) * sliceLen
			if end > len(nums){
				end = len(nums)
			}
			var pNum int64 = 0
			for _, n := range nums[g*sliceLen: end]{
				pNum += int64(n)
			}

			atomic.AddInt64(&num, pNum)
			wait.Done()
		}(g)
	}

	wait.Wait()
    return num
}

func BenchmarkSingleThread(b *testing.B){
    for i:= 0;i< b.N; i++{
    	n := addNum(nums)
		if n != res{
			b.Errorf("not equal %d != %d\n", res, n)
		}
	}
}

func BenchmarkConcurrentCpuNums(b *testing.B){
	for i:= 0;i <b.N; i++{
		n := addNumConcurrent(nums, runtime.NumCPU())
		if n != res{
			b.Errorf("not equal %d != %d\n", res, n)
		}
	}
}

/*
电脑cpu为2核4线程

运行python脚本
# !/bin/bash

# -*- coding:utf-8 -*-

import os

for i in range(1, 5):
    command = "GOGC=off go test -cpu %d -run none -bench . -benchtime 3s -count=5" % i

    os.system("echo '%s\n'" % ("*" * 40))
    os.system("echo '%s\n'" % command)
    os.system(command)
    os.system("echo '%s\n\n\n'" % ("*" * 40))

输出为下
*/
/*
****************************************

GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s -count=5

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThread              10000            464572 ns/op
BenchmarkSingleThread              10000            460184 ns/op
BenchmarkSingleThread              10000            451535 ns/op
BenchmarkSingleThread              10000            469571 ns/op
BenchmarkSingleThread              10000            464509 ns/op
BenchmarkConcurrentCpuNums         10000            475715 ns/op
BenchmarkConcurrentCpuNums         10000            474927 ns/op
BenchmarkConcurrentCpuNums         10000            475232 ns/op
BenchmarkConcurrentCpuNums         10000            468485 ns/op
BenchmarkConcurrentCpuNums         10000            476807 ns/op
PASS
ok      go-bpt/sync/bound       47.340s
****************************************



****************************************

GOGC=off go test -cpu 2 -run none -bench . -benchtime 3s -count=5

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThread-2            10000            478792 ns/op
BenchmarkSingleThread-2            10000            468516 ns/op
BenchmarkSingleThread-2            10000            470702 ns/op
BenchmarkSingleThread-2            10000            461592 ns/op
BenchmarkSingleThread-2            10000            459559 ns/op
BenchmarkConcurrentCpuNums-2       20000            256830 ns/op
BenchmarkConcurrentCpuNums-2       20000            255030 ns/op
BenchmarkConcurrentCpuNums-2       20000            254381 ns/op
BenchmarkConcurrentCpuNums-2       20000            251218 ns/op
BenchmarkConcurrentCpuNums-2       20000            253542 ns/op
PASS
ok      go-bpt/sync/bound       61.875s
****************************************



****************************************

GOGC=off go test -cpu 3 -run none -bench . -benchtime 3s -count=5

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThread-3            10000            466316 ns/op
BenchmarkSingleThread-3            10000            459336 ns/op
BenchmarkSingleThread-3            10000            469697 ns/op
BenchmarkSingleThread-3            10000            451563 ns/op
BenchmarkSingleThread-3            10000            459187 ns/op
BenchmarkConcurrentCpuNums-3       20000            237053 ns/op
BenchmarkConcurrentCpuNums-3       20000            236486 ns/op
BenchmarkConcurrentCpuNums-3       20000            236201 ns/op
BenchmarkConcurrentCpuNums-3       20000            235062 ns/op
BenchmarkConcurrentCpuNums-3       20000            237510 ns/op
PASS
ok      go-bpt/sync/bound       58.904s
****************************************



****************************************

GOGC=off go test -cpu 4 -run none -bench . -benchtime 3s -count=5

goos: linux
goarch: amd64
pkg: go-bpt/sync/bound
BenchmarkSingleThread-4            10000            467062 ns/op
BenchmarkSingleThread-4            10000            454313 ns/op
BenchmarkSingleThread-4            10000            491002 ns/op
BenchmarkSingleThread-4            10000            475302 ns/op
BenchmarkSingleThread-4            10000            462579 ns/op
BenchmarkConcurrentCpuNums-4       30000            157921 ns/op
BenchmarkConcurrentCpuNums-4       30000            183034 ns/op
BenchmarkConcurrentCpuNums-4       30000            165301 ns/op
BenchmarkConcurrentCpuNums-4       30000            159792 ns/op
BenchmarkConcurrentCpuNums-4       30000            162317 ns/op
PASS
ok      go-bpt/sync/bound       56.766s
****************************************

*/


/*
取平均值
GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s -count=5
BenchmarkSingleThread              10000            462074.2 ns/op
BenchmarkConcurrentCpuNums         10000            474233.2 ns/op

GOGC=off go test -cpu 2 -run none -bench . -benchtime 3s -count=5
BenchmarkSingleThread-2            10000            467832.2 ns/op
BenchmarkConcurrentCpuNums-2       20000            254200.2 ns/op

GOGC=off go test -cpu 3 -run none -bench . -benchtime 3s -count=5
BenchmarkSingleThread-3            10000            461219.8 ns/op
BenchmarkConcurrentCpuNums-3       20000            236462.4 ns/op

GOGC=off go test -cpu 4 -run none -bench . -benchtime 3s -count=5
BenchmarkSingleThread-4            10000            470051.6 ns/op
BenchmarkConcurrentCpuNums-4       30000            165673.0 ns/op


由此可见对于cpu密集型操作 需要分配的携程数刚好等于cpu可用线程数是最好的,再多分配对于性能提升不大

而且在cpu单线程情况下是无法实现真正意义上的并行的,所以反而会由于多携程不断切换上下文而导致计算更长的时间

*/