package data

/*
数据类型:
	预定义类型:
		# 布尔类型
		bool		//  true 或者 false
		# 数字
		int			//
		uint8		// 无符号 8 位整型 (0 到 255)
		uint16		// 无符号 16 位整型 (0 到 65535)
		uint32		// 无符号 32 位整型 (0 到 4294967295)
		uint64		// 无符号 64 位整型 (0 到 18446744073709551615)
		int8		// 有符号 8 位整型 (-128 到 127)
		int16		// 有符号 16 位整型 (-32768 到 32767)
		int32		// 有符号 32 位整型 (-2147483648 到 2147483647)
		int64		// 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
		float32		// IEEE-754 32位浮点型数
		float64		// IEEE-754 64位浮点型数
		complex64	// 32 位实数和虚数
		complex128	// 64 位实数和虚数
		# 字符串
		string
		# 其他
		void		//
		byte		// 类似 uint8
		rune		// 类似 int32
		uint		// 32 或 64 位
		int			// 与 uint 一样大小
		uintptr		// 无符号整型，用于存放一个指针
		pointer		// 指针，初始值 nil
	复合类型:
		数组:
			声明: var variableName [SIZE][SIZE2]... variableType
			初始化: var variableName = [SIZE|...][SIZE2]...variableType{ {value1,...} , ...}
			使用: variableName[index1][index2]...	// 数组下标从 0 开始
			# slice，map
		指针:
			声明:
				var pointerName *varType
				var pointerName **varType		// 指向指针的指针
				var pointerName [SIZE|*varType	// 指针数组
			使用:
				pointerName = &varName			// & 取变量地址
				varName = *pointerName			// * 取指针指向变量的值
		结构:
			定义: type structName struct{ varName varType,... }
			声明:
				var varStructName structName
				var varStructPointer *structName	// 结构体指针
			初始化:
				varStructName := structName{ [varName:]value,... }
				varStructPointer = &varStructName
			访问成员:
				varStructName.varName
				varStructPointer.varName
	接口: type interfaceName interface{ ... }							// 接口的声明
	别名: type aliasName typeName										// 类型别名， e.g: type IZ int
# 类型判断:
	格式化输出:
		func typeof(v interface{}) string {
			return fmt.Sprintf("%T", v)
		}
	反射:
		func typeof(v interface{}) string {
			return reflect.TypeOf(v).String()
		}
	type-switch断言:
		func typeof(v interface{}) string {
			switch t := v.(type) {
			case int:
				return "int"
			case float64:
				return "float64"
			//... etc
			default:
				_ = t
				return "unknown"
			}
		}
	断言:
		if sv, ok := v.(Stringer); ok {
			fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
		}
*/
