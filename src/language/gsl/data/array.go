package data

import "sort"

/*
数组与集合
	数组声明：
		var identifier [len]type  //e.g: var arr1 [5]int,索引范围从 0 到 len(arr1)-1
		var identifier = new([len]type)	//e.g: var arr1 = new([5]int)
	多维数组:
		e.g: var arr [3][5]int
	切片: #对数组一个连续片段的引用
		声明: var identifier []type  //（不需要说明长度,未初始化之前默认为 nil，长度为 0）
		初始化:
			var slice1 []type = arr1[start:end]
			var slice1 []type = make([]type, len, cap)
		长度: len()
		容量: cap()
		重组: slice1 = slice1[0:end] //e.g: var ar = [10]int{0,1,2,3,4,5,6,7,8,9};var a = ar[5:7];a = a[0:4]
		复制: copy(sl_to, sl_from)
		追加: append(sl3, 4, 5, 6)
	Map: #元素对（pair）的无序集合
		声明: var map1 map[keyType]valueType 	//e.g: var map1 map[string]int
		初始化:
			var map1 = make(map[keyType]valueType,cap)	// mapCreated := make(map[string]float32)
			var map1 = map[keyType]valueType{key1:value1,...}  // e.g: mapLit = map[string]int{"one": 1, "two": 2}
		赋值: map1[key1] = val1
		长度: len(map1)
		键值对存在检测: val1, isPresent = map1[key1] // isPresent为bool值，true表示键值对存在
		删除: delete(map1, key1) // 如果 key1 不存在，该操作不会产生错误。
 */

// 数组
// var arrAge = [5]int{18, 20, 15, 22, 16}
// var arrLazy = [...]int{5, 6, 7, 8, 22}
// var arrLazy = []int{5, 6, 7, 8, 22}
// var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
// var arrKeyValue = []string{3: "Chris", 4: "Ron"}
// var arr [3][5]int
// var arr [2][2][2]float64
// 切片
//

func Ints(a []int) { sort.Ints(a) } // 排序
func IntsAreSorted(a []int) bool { return sort.IntsAreSorted(a) } // 检测是否排序

