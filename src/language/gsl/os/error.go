package os

import "errors"

/*
go 错误

# go语言通常在函数和方法中返回错误对象作为它们的唯一或最后一个返回值——如果返回 nil，则没有错误发生
# panic(): 当发生像数组下标越界或类型断言失败这样的运行错误时，Go 运行时会触发运行时 panic() 程序。（类似于java catch 和 finally 里的错误处理）
# recover(): 取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil。panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。（类似于java throw new Exception）
# defer: 函数 return 后调用。（类似于java finally）

 */

type error interface {
	Error() string
}
func New(text string) error { return errors.New(text) }
