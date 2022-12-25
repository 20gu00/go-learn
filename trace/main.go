package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	defer Trace()()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t2()
		wg.Done()
	}()
	foo()
	wg.Wait()
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
	t2()
}

func t1() {
	defer Trace()
}

func t2() {
	defer Trace()
	t1()
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	// caller 函数
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	// groutineID go开发团队建议不要依赖groutine的id 生产环境不要用
	gid := curGoroutineID()
	mu.Lock()
	indents := m[gid]    // 获取当前gid对应的缩进层次
	m[gid] = indents + 1 // 缩进层次+1后存入map
	mu.Unlock()
	printTrace(gid, name, "->", indents+1)
	return func() {
		mu.Lock()
		indents := m[gid]    // 获取当前gid对应的缩进层次
		m[gid] = indents - 1 // 缩进层次-1后存入map
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}
