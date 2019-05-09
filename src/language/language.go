package language
/**
1,基础语法:
	注释:
		// 单行注释
		\/* 多行注释 *\/
	结束符:
		'\n'(换行符)或 ';'(分号,一行多个语句必须用分号),
	执行函数:
		func init() { } // 每个包完成初始化后自动执行，并且执行优先级比 main 函数高
		func main() { }	// main包内的main函数作为程序入口点启动
	基础组成:
		包
		标识量(常量,变量)
		主函数(init,main)/函数
2,Keyword:
	break		default			func		interface		select
	case		defer			go			map				struct
	chan		else			goto		package			switch
	const		fallthrough		if			range			type
	continue	for				import		return			var
	# 预定义标识符
	append	bool	byte	cap			close	complex	complex64	complex128	uint16
	copy	false	float32	float64		imag	int		int8		int16		uint32
	int32	int64	iota	len			make	new		nil			panic		uint64
	print	println	real	recover		string	true	uint		uint8		uintptr
3,标识量
	命名: 字母(A~Z和a~z)数字(0~9)、下划线_组成, 不能以数字开头, 区分大小写。
	变量声明:
		var varName[, varName2] [type][= value[, value2]] 	// e.g: var age int
																	var vname1, vname2, vname3 type
																	var vname1, vname2 = v1, v2
		varName[, varName2] := value[, value2]				// e.g: c := 10
																	vname1, vname2 := v1, v2
		var (												//
			varName1 type
			varName2 type2
		)
		# Go 语言中全局变量与局部变量同名时，局部变量优先
	常量声明:
		const vatName [type] = value; // e.g: const PI = 3.14;
		const varName[, varName2] := value[, value2]
	特殊标识量:
		_ 			// 空白占位符，只写变量(可用作函数传参)
		iota		// 每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
		nil			// 空指针
4, 运算符
	算术运算符:
		+	-	*	/	%	++	--
	关系运算符:
		==	!=	>	<	>=	<=
	逻辑运算符:
		&&	||	!
	位运算符:
		&	|	^	<<	>>
	赋值运算符:
		=	+=	-=	*=	/=	%=	<<=	>>=	&=	^=	|=
	其他运算符:
		&	返回变量存储地址
		*	指针变量
		:=	声明并赋值
5, 语句结构:
	分支:
		if [initialization;] condition { } [ [ else if condition2 { } ] else { } ] // e.g: if val := 10; val>10 {...} else if val>5 {...} else {...}
		switch var1 { case val1: ... default: ... }
		select { case communication clause : ... default: ... }
	循环:
		for init; condition; post { } // e.g: for i:=0,j:=0; i<5,j<5; i++,j++ { ... }
		for condition { } // 条件循环，e.g: for i >= 0 { .. }
		for ix, val := range coll { } // 遍历循环(类似foreach), e.g: str := "Go is a language!"; for index, char := range str { ... }
		for { }
	跳转:
		break;
		continue;
		goto label; ...; label: statement;
		return ...;
	回调:
		defer ...
*,标准库(包)
	unsafe			包含了一些打破 Go 语言“类型安全”的命令，一般的程序中不会被使用，可用在 C/C++ 程序的调用中。
	syscall			底层的外部包，提供了操作系统底层调用的基本接口。
	fmt				提供了格式化输入输出功能。
	os				提供给我们一个平台无关性的操作系统功能接口，采用类Unix设计，隐藏了不同操作系统间差异，让不同的文件系统和操作系统对象表现一致。
	os/exec			提供我们运行外部操作系统命令和程序的方式。
	archive/tar 	压缩(解压缩)文件功能。
	archive/zip-compress：压缩(解压缩)文件功能。
	io				提供了基本输入输出功能，大多数是围绕系统功能的封装。
	bufio 			缓冲输入输出功能的封装。
	path/filepath 	用来操作在当前系统中的目标文件名路径。
	flag 			对命令行参数的操作。
	math 			基本的数学函数。
	math/cmath 		对复数的操作。
	math/rand 		伪随机数生成。
	sort 			为数组排序和自定义集合。
	math/big 		大数的实现和计算。 　　
	time			日期和时间的基本操作。
	log 			记录程序运行时产生的日志。
	strings 		提供对字符串的操作。
	strconv 		提供将字符串转换为基础类型的功能。
	unicode 		为 unicode 型的字符串提供特殊的功能。
	regexp: 		正则表达式功能。
	bytes 			供对字符型分片的操作。
	index/suffixarray 子字符串快速查询。
	container/list  双链表。
	container/ring  环形链表。
	encoding/json 	读取并解码和写入并编码 JSON 数据。
	encoding/xml 	简单的 XML1.0 解析器,有关 JSON 和 XML 的实例请查阅第 12.9/10 章节。
	text/template 	生成像 HTML 一样的数据与文本混合的数据驱动模板（参见第 15.7 节）。
	net 			网络数据的基本操作。
	http 			提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
	html 			HTML5 解析器。
	runtime 		Go 程序运行时的交互操作，例如垃圾回收和协程创建。
	reflect 		实现通过程序运行时反射，让程序操作任意类型的变量。
 */

 /*

 FAQ:
 	1) 发生错误时使用defer关闭一个文件
 		defer仅在函数返回时才会执行，在循环的结尾或其他一些有限范围的代码内不会执行。
 	2) 何时使用new()和make()
 		- 切片、映射和通道，使用make
		- 数组、结构体和所有的值类型，使用new


  */