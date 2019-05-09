package io

import (
	"os"
	"io/ioutil"
	"io"
	"fmt"
	"bufio"
)

/*
文件输入输出

# Go 语言中，文件使用指向 os.File 类型的指针来表示

 */

type File struct { os.File }
func Open(name string) (*os.File, error) { return os.Open(name) }
func OpenFile(name string, flag int, perm os.FileMode) (*File, error) { return os.OpenFile(name,flag,perm) }
func (file *File) Close() error { return file.Close() }
func ReadFile(filename string) ([]byte, error) { return ioutil.ReadFile(filename) }
func WriteFile(filename string, data []byte, perm os.FileMode) error { return ioutil.WriteFile(filename,data,perm) }
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) { return fmt.Fscanf(r,format,a) }

func Copy(dst io.Writer, src io.Reader) (written int64, err error) { return io.Copy(dst,src)}

type Writer struct {
	bufio.Writer
}
func (b *Writer) WriteString(s string) (int, error) { return b.WriteString(s) }
func (b *Writer) Flush() error { return b.Flush() }
