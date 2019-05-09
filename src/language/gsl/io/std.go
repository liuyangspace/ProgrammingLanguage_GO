package io

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

/*
std 标准输入输出
go 语音 标准输入:
	var (
		Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	)

参考:
	fmt 包
 */
//std input
func Scanln(a ...interface{}) (n int, err error) { return fmt.Fscanln(os.Stdin, a...) } // 扫描来自标准输入的文本,e.g:fmt.Scanln(&firstName, &lastName)
func Scanf(format string, a ...interface{}) (n int, err error) { return fmt.Fscanf(os.Stdin, format, a...) } //
func Sscanf(str string, format string, a ...interface{}) (n int, err error) { return fmt.Sscanf(str,format,a) } //
//std output
func Println(a ...interface{}) (n int, err error) { return fmt.Fprintln(os.Stdout, a...) }
func Printf(format string, a ...interface{}) (n int, err error) { return fmt.Fprintf(os.Stdout, format, a...) }
//buffered 缓冲
type Reader struct { bufio.Reader }
func NewReader(rd io.Reader) *bufio.Reader { return bufio.NewReader(rd) }
func (b *Reader) ReadString(delim byte) (string, error) { return b.ReadString(delim) }
func (b *Reader) Read(p []byte) (n int, err error) { return b.Read(p) }

