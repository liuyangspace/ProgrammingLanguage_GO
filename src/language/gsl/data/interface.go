package data

/*
interface 接口

#接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
定义：
	type interfaceName interface {
		otherInterfaceName
		MethodName(param_list) return_type
		...
	}
类型断言: interfaceVar.(Type)	// 测试在某个时刻 interfaceVar 是否包含类型 Type(struct类型须前置*) 的值,e.g:
	v, ok := varI.(*T);	//如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false
# Go 语言中接口可以包含一个或多个其他的接口,作用相当于直接将这些内嵌接口的方法列举在外层接口中一样
# Go 语言中接口有一个接口类型的变量 var ai interfaceName，
# 多个类型可以实现同一个接口，一个类型可以实现多个接口。
# (多态:)不论接口和类型定义在不同的包中，只要类型实现了接口中的方法，它就实现了此接口(可以两种类型变量间相互赋值)。e.g:
	type Shaper interface { Area() float32 }
	type Square struct { side float32 }
	func (sq *Square) Area() float32 { return sq.side * sq.side }
	func main() {
		sq1 := Square{5}
		var areaIntf Shaper
		areaIntf = sq1
		// shorter,without separate declaration:
		// areaIntf := Shaper(sq1)
		// or even:
		// areaIntf := sq1
		fmt.Printf("The square has area: %f\n", areaIntf.Area())
	}
# Go 语言规范定义了接口方法集的调用规则：
	类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	类型 T 的可调用方法集包含接受者为 T 的所有方法,不包含接受者为 *T 的方法
空接口:
	interface {}	//类似 Java/C# 中所有类的基类： Object
	# 可以给一个空接口类型的变量 var val interface {} 赋任何类型的值。
 */
