package data

/*
函数
	声明:
		# Go语音不允许函数重载(函数重载（function overloading）指的是可以编写多个不同的形参或者不同的返回值同名函数)。
		# Go语音支持函数递归调用
		func functionName( parameter [type][,...] ) [returnTypes] { ... } // 左大括号 { 必须与方法名声明放在同一行
		type functionName func(type,...) type,...	// 函数可以以申明的方式被使用，作为一个函数类型,e.g: type binOp func(int, int) int
	传参:
		# go 语言基本数据类型(bool、int、...)传参默认传值,复合类型(slice、map、interface、channel、...)默认隐式引用传递
		多返回值: func functionName( parameter [type][,...] ) returnType1,returnType2,... { ...; return returnValue1,returnValue2 }
				// e.g: func swap(x, y string) (string, string) { return y, x }
		命名的返回值: func functionName( parameter [type][,...] ) (return1 returnType1,return2 returnType2,...) { ... }
				// e.g: func getX2AndX3_2(input int) (x2 int, x3 int) {
				//    x2 = 2 * input; x3 = 3 * input
				//    return // return x2, x3  }
		引用传参: func functionName( parameter *type [,...] ) [returnTypes] { ... }	// 通过 * 标识指针,e.g: func swap(x *int, y *int) { ... }
		数组传参:
			func functionName( arrayName [arraySize]type [,...] ) [returnTypes] { ... } // e.g: void myFunction(param [10]int){ ... }
			func functionName( arrayName []type [,...] ) [returnTypes] { ... }
	    变长参数: func functionName(a, b, arg ...int) { if len(s)>0 { min := s[0] ... } }
		函数作为参数: func functionName(y int, f func(int, int)) { f(y, 2) }
		任意类型参数: a interface{} // 空接口表示任意类型，func Println(a ...interface{})
	使用:
		varName[,varName2...] = [packageName.]functionName(parameterName1,parameterName2 ...)	// e.g: a, b := swap("Mahesh", "Kumar")
		varName = [packageName.]functionName(&parameterName1,&parameterName2 ...)	// 通过 & 取地址，e.g: swap(&a, &b)
		# Go语音函数可以赋值给变量，e.g: type binOp func(int); add := binOp 。
	变量(匿名)函数:
		# 将函数赋给指定变量
		声明:
			getSquare := func(x float64) float64 { return math.Sqrt(x) }
			getSquareValue := func(x float64) float64 { return math.Sqrt(x) }(value)
		使用: getSquareRoot(9)
	闭包函数:
		# 闭包函数返回值为函数，且 返回函数可调用闭包函数的局部变量
		声明: func getSequence() func() int { i:=0; return func() int { return i+1  } }
		使用: nextNumber := getSequence() ; nextNumber()
	方法函数:
		# 方法函数 指明 返回值 接受变量
		# 接受变量类型和作用在它上面定义的方法必须在同一个包里定义()
		声明: func (varName varType) methodName(parameterName...) [returnTypes]{ ... }
		使用: varName.methodName(parameterName...)
	defer追踪:
		defer 关键字 声明函数结束后回调的语句，e.g:
			func function1() {
				defer Printf(" defer! ")
				fmt.Printf(" bottom! ")
			}
			// 运行输出: bottom!  defer!
	内置函数:
		名称				说明
		close				用于管道通信
		len、cap			len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
		new、make			new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作new() 是一个函数，
		copy、append		用于复制和连接切片
		panic、recover		两者均用于错误处理机制
		print、println		底层打印函数，在部署环境中建议使用 fmt 包
		complex、real imag	用于创建和操作复数
 */
//
