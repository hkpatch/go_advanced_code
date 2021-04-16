package main

import "fmt"

func main() {
	arrayA := [2]int{100, 200}
	testArrayPoint1(&arrayA) // 1.传数组指针
	arrayB := arrayA[:]
	testArrayPoint2(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA) // arrayA : 0xc00000a0a0 , [100 400]
}

func testArrayPoint1(x *[2]int) {
	// 1.传数组指针，地址不变
	fmt.Printf("func Array1 : %p , %v\n", x, *x) // func Array1 : 0xc00000a0a0 , [100 200]

	// 增加100
	(*x)[1] += 100
}

func testArrayPoint2(x *[]int) {
	// 2.传指针切片，地址变化
	fmt.Printf("func Array2 : %p , %v\n", x, *x) // func Array2 : 0xc0000044c0 , [100 300]

	// 增加100
	(*x)[1] += 100
}
// 数组指针
//	优点：就算是传入10亿的数组，也只需要再栈上分配一个8个字节的内存给指针就可以了
// 缺点： 第一行和第三行指针地址都是同一个，万一原数组的指针指向更改了，那么函数里面的指针指向都会跟着更改

// 解决方法-----用切片指针
// 优点：用切片传数组参数，既可以达到节约内存的目的，也可以达到合理处理好共享内存的问题。
//	打印结果第二行就是切片，切片的指针和原来数组的指针是不同的。

/*
结论：
	把第一个大数组传递给函数会消耗很多内存，采用切片的方式传参可以避免上述问题。
	切片是引用传递，所以它们不需要使用额外的内存并且比使用数组更有效率
 */
