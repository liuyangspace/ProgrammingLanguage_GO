package os

import "runtime"

/*
调试:

#测试程序必须属于被测试的包，并且文件名满足这种形式 *_test.go
参考:
	go test 命令
 */

func Caller(skip int) {
	runtime.Caller(skip)
}

func log()  {
	println(log)
}