# 函数式编程概论
## 背景
硬件性能的提升以及编译技术和虚拟机技术的改进，一些曾被性能问题所限制的动态语言开始受到关注，Python、Ruby 和 Lua 等语言都开始在应用中崭露头角。
伴随动态语言的流行，函数式编程也再次进入了我们的视野。
函数式编程是一种编程模型，他将计算机运算看做是数学中函数的计算，并且避免了状态以及变量的概念

## 函数
一等公民：
函数作为变量对待。也就说，函数与变量没有差别，它们是一样的，变量出现的地方都可以替换成函数，并且编译也是可以通过的，没有任何语法问题。

## 函数使用
1. 函数可以定义函数类型
2. 函数可以赋值给变量
3. 高阶函数---可以作为入参也可以作为返回值
4. 动态创建函数
5. 匿名函数
6. 闭包

### 1.定义函数类型：
```go
type Operation func(a,b int) int
// -----Operation :type name类型名称
//-----func(a,b int) int:signature函数签名
func Add func(a,b int) int{
    return a+b
}
//符合函数签名的函数
```

### 2.声明函数类型的变量和为变量赋值：
```go
var op Operation
op = Add
fmt.Println(op(1,2))
```
变量op是Operation类型的，可以把Add作为值赋值给变量op，执行op等价于执行Add。

### 3.函数作为其他函数入参
```go
type Calculator struct {
    v int
}
func (c Calculator)Do(op Operation,a int){
    c.v = op(c.v,a)
}
func main(){
   var calc Calculator
   calc.Do(add,1)
}
```


### 4. 函数作为返回值+动态创建
```go
type Operation func(b int)int
func Add(b int) Operation{
   addB := func(a int)int{
      return a + b
    }
   return addB
}

type Calculator struct {
    v int
}
func (c Calculator)Do(op Operation){
    c.v = op(c.v)
}
func main(){
   var calc Calculator
   calc.Do(add(1)) //c.v = 1
}
```


### 5. 匿名函数
```go
func(a int)int{}
   func Add(b int) Operation{
   return func(a int)int{
       return a + b
   }
}
```


### 6. 闭包
#### 定义
闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)   
#### 解析
闭包只是在形式和表现上像函数，但实际上不是函数。函数是一些可执行的代码，这些代码在函数被定义后就确定了，不会在执行时发生变化。
所以一个函数只有一个实例。闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。
所谓引用环境是指在程序执行中的某个点所有处于活跃状态的约束所组成的集合。
不同的引用环境和相同的函数组合可以产生不同的实例。

```go
type Operation func(b int)int
func Add(b int) Operation{
   addB := func(a int)int{
       return a + b
   }
   return addB
}
```

比如匿名函数里直接使用了变量b，该匿名函数也是闭包函数。
Note:一个函数可以是匿名函数，但不是闭包函数，因为闭包有时是有副作用的。
