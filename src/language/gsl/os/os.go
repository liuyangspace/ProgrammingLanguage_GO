package os

import (
	"flag"
	"os"
	"os/exec"
)

/*


#os 包中有一个 string 类型的切片变量 os.Args，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。


 */


func Bool(name string, value bool, usage string) *bool { return flag.Bool(name,value,usage) }
func Exit(code int) { os.Exit(code) }

func Command(name string, arg ...string) *exec.Cmd { return exec.Command(name) }
func StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error) { return os.StartProcess(name,argv,attr) }
