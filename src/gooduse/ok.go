package gooduse

/*

常见 value,ok = MyFunc() 模式

1) 错误处理:
		if value, err := pack1.Func1(param1); err != nil {
			// do err
			return err
		}
		Process(value)// 函数Func1没有错误:

2) 检测映射中是否存在一个键值 ：key1在映射map1中是否有值？
	if value, isPresent = map1[key1]; isPresent {
			Process(value)
	}
	// key1不存在

3) 检测一个接口类型变量varI是否包含了类型T：类型断言：
	if value, ok := varI.(T); ok {
		Process(value)
	}
	// 接口类型varI没有包含类型T

4) 检测一个通道ch是否关闭：
		for input := range ch {
			Process(input)
		}
	或者:
		for {
			if input, open := <-ch; !open {
				break // 通道是关闭的
			}
			Process(input)
		}
 */