package os
/*
reflect 反射

方法和类型的反射,e.g:
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))			// type: float64
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)						// value: 3.4
	fmt.Println("type:", v.Type())					// type: float64
	fmt.Println("kind:", v.Kind())					// kind: float64
	fmt.Println("value:", v.Float())				// value: 3.4
	fmt.Println(v.Interface())						// 3.4
	fmt.Printf("value is %5.2e\n", v.Interface())	// value is 3.40e+00
	y := v.Interface().(float64)
	fmt.Println(y)									// 3.4
通过反射修改(设置)值,e.g:
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.SetFloat(3.1415) 							// Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())	// settability of v: false
	v = reflect.ValueOf(&x) 						// Note: take the address of x.
	fmt.Println("type of v:", v.Type())				// type of v: *float64
	fmt.Println("settability of v:", v.CanSet())	// settability of v: false
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)			// The Elem of v is:  <float64 Value>
	fmt.Println("settability of v:", v.CanSet())	// settability of v: true
	v.SetFloat(3.1415) 								// this works!
	fmt.Println(v.Interface())						// 3.1415
	fmt.Println(v)									// <float64 Value>
反射结构:
	type NotknownType struct { s1, s2, s3 string }
	func (n NotknownType) String() string { return n.s1 + " - " + n.s2 + " - " + n.s3 }
	// variable to investigate:
	var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}
	// main
	value := reflect.ValueOf(secret) 				// <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    				// main.NotknownType
	// alternative:
	//typ := value.Type()  							// main.NotknownType
	fmt.Println(typ)								// main.NotknownType
	knd := value.Kind() 							// struct
	fmt.Println(knd)								// struct
	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		// value.Field(i).SetString("C#")
	}
	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) 							// [Ada - Go - Oberon]


参见:
	"reflect" 包
	reflect.TypeOf(x)
	reflect.ValueOf(x)
 */
