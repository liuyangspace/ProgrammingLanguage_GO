package os

import "time"

/*

协程
	go关键字: 协程通过使用关键字 go 调用（执行）一个函数或者方法来实现的（也可以是匿名或者 lambda 函数）
	format:
		//go 关键字放在方法调用前新建一个 goroutine 并执行方法体
		go GetThingDone(param1, param2);
		//新建一个匿名方法并执行
		go func(param1, param2) {
		}(val1, val2)
		//直接新建一个 goroutine 并在 goroutine 中执行代码块
		go {
			//do someting...
		}
	配置:
		runtime.GOMAXPROCS() // 配置并行执行比较适合适合于CPU密集型、并行度比较高的情景，如果是IO密集型使用多核的化会增加cpu切换带来的性能损失。
	chan关键字: 通常可以使用如下格式来声明通道：
		# 未初始化的通道的值是nil。
		var identifier chan dataType	// 无缓冲同步通道,e.g: var ch1 chan string
		var send_only chan<- int 		// channel can only receive data
		var recv_only <-chan int		// channel can only send data
		var identifier = make(chan dataType)	//无缓冲同步通道，e.g: ch1 := make(chan string),funcChan := make(chan func())
		var identifier = make(chan dataType, bufferSize) //带缓冲的异步通道,e.g: ch1 := make(chan string, 100)
		# 阻塞: 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束,在接收者准备好之前是阻塞的。
		# 可用close(ch)关闭通道，可用cap()检测通道容量
		# 检测通道是否关闭：v, ok := <-ch   // ok is true if v received value
	通信操作符 <-:
		流向通道（发送）: ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）,
		从通道流出（接收）:
			int2 = <- ch 或 int2 := <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）
			if <- ch != 1000{...} 表示：单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证
	select关键字: # 从多个通道中获取值，用法类似 switch-case-default
			select {
				case u:= <- ch1: ...
				case v:= <- ch2: ...
				default: ...// no value ready to be received
			}
		1) 如果通道都阻塞了，会等待直到其中一个可以处理
		2) 如果多个通道可以处理，随机选择一个
		3) 如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。


# linux gc 编译器下,可设置 GOMAXPROCS 为一个大于默认值 1 的数值来允许运行时支持使用多于 1 个的操作系统线程
# 正在执行的协程可以可以使用 runtime.Goexit() 来停止。
# Go 协程（goroutines）和协程（coroutines）对比:
	Go 协程意味着并行（或者可以以并行的方式部署），协程一般来说不是这样的
	Go 协程通过通道（channel）来通信；协程通过让出和恢复操作来通信
 */

type Ticker struct { // 循环间隔某段时间向通道C发送时间
	time.Ticker
}
type Timer struct { // 延迟某段时间向通道C发送时间
	time.Timer
}
func NewTicker(d time.Duration) *time.Ticker { return time.NewTicker(d) } //
