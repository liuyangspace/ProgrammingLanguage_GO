package data

/*
struct
定义:
	type identifier struct {
		field1 type1 ["tag1"]
		field2 type2 ["tag2"]
		...
	}
	// e.g: type T struct {a, b int "b"}
	# Go语音允许结构体内嵌，即结构体内字段亦允许为结构体类型
匿名字段:
	结构体可以包含一个或多个 仅有类型没有显示字段名的字段，此时类型就是字段的名字。e.g:
		type T struct { int }; t := T(1); t.int = 2
初始化:
	声明赋值: var structVar structType; structVar.fieldName = value ...
	new内存分配: structVar := new(structType)
	&符赋值: structVar := &structType{[fieldName.]value...} 或 structVar := structType{[fieldName.]value...}
		//据说语音内部亦使用new分配内存,e.g:
			type Interval struct { start , end int }
			intr := Interval{0, 3}            //值必须以字段在结构体定义时的顺序给出
			intr := Interval{end:5, start:1}  //字段名加一个冒号放在值的前面，值的顺序不必一致
			intr := Interval{end:5}           //字段名加一个冒号放在值的前面，某些字段可以被忽略掉
选择器（selector）: Go 语言中使用 structVar.fieldName 格式引用 结构体类型 或 结构体类型指针 的字段
方法: 参考 function.go 方法函数
继承: 结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法。e.g:
	type Point struct { x, y float64 }
	func (p *Point) Abs() float64 { return math.Sqrt(p.x*p.x + p.y*p.y) }
	type NamedPoint struct { Point, name string }
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印5
	// 内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法
格式化输出: 如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
格式化描述符:
	%v	//字段输出
	%T 	//给出类型的完全规格
	%#v	//出实例的完整输出，包括它的字段

 */

