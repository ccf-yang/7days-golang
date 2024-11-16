package baseusage

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)



func VariableExamples() {
	// 基本类型示例
	integerExample := 42
	floatExample := 3.14
	stringExample := "你好，Go！"
	booleanExample := true

	fmt.Printf("整数: %d, 浮点数: %.2f, 字符串: %s, 布尔值: %t\n", integerExample, floatExample, stringExample, booleanExample)
		// 常用的格式化占位符示例
	// %d: 整数            // 用于整数
	// %f: 3.14            // 用于浮点数
	// %s: "Hello"         // 用于字符串
	// %t: true            // 用于布尔值
	// %x: 0xA             // 用于十六进制
	// %p: 0xc00008a018    // 用于指针
	// %v: [1 2 3]         // 用于任何类型的默认格式
	// %T: []int           // 用于打印类型
	// %%: 50%%            // 用于打印百分号本身


	fmt.Printf("integerExample 的类型: %T\n", integerExample)

	// 常用的判断类型的方法：
	// 1. 使用 reflect.TypeOf() 获取类型信息
	// 2. 使用类型断言 value.(Type)
	// 3. 使用 switch 语句进行类型选择
	// 4. 使用 reflect.ValueOf().Kind() 获取底层类型
	// 5. 使用 fmt.Sprintf("%T", value) 获取类型的字符串表示
	// 方法1：直接类型断言
	if _, ok := interface{}(integerExample).(int); ok {
		// 这行代码的目的是检查 integerExample 是否是 int 类型。让我们逐步拆解：
		// interface{}(integerExample): 将 integerExample 转换为空接口类型（interface{}）
		// 在 Go 中，空接口可以存储任意类型的值
		// 这是一种将具体类型转换为接口类型的方式
		// .(int)
		// 这是类型断言的语法
		// 尝试将空接口中的值断言为 int 类型
		// _, ok :=
		// 如果断言成功，ok 为 true
		// 第一个返回值（被丢弃的 _）是断言成功后的值
		fmt.Println("integerExample 是整型")
	}

	// 方法2：使用反射
	if reflect.TypeOf(integerExample).Kind() == reflect.Int {
		fmt.Println("integerExample 是整型")
	}


	// 复合类型示例
	arrayExample := [3]int{1, 2, 3}
	sliceExample := []string{"a", "b", "c"}
	mapExample := map[string]int{"键": 1}

	fmt.Printf("数组: %v, 切片: %v, 映射: %v\n", arrayExample, sliceExample, mapExample)

	// 结构体示例
	type PersonExample struct {
		Name string
		Age  int
	}
	structExample := PersonExample{Name: "爱丽丝", Age: 30}
	fmt.Printf("结构体: %+v\n", structExample)
	fmt.Printf("结构体: %v\n", structExample)   // 输出: {Alice 30}     // 默认格式，只显示值
	fmt.Printf("结构体: %+v\n", structExample)  // 输出: {Name:Alice Age:30}  // 显示字段名和值
	fmt.Printf("结构体: %#v\n", structExample)  // 输出: main.Person{Name:"Alice", Age:30}  // 显示完整类型信息

	// 零值示例
	var zeroIntExample int
	var zeroStringExample string
	fmt.Printf("零值 - 整数: %d, 字符串: '%s'\n", zeroIntExample, zeroStringExample)

	// 类型转换示例
	typeConversionExample := float64(integerExample)
	fmt.Printf("类型转换: %f\n", typeConversionExample)
	

	// 类型断言示例
	var interfaceExample interface{} = "字符串"
	assertedValue, ok := interfaceExample.(string)
	fmt.Printf("类型断言: 值=%s, 成功=%t\n", assertedValue, ok)

	// 指针示例 &获取变量地址，然后可以赋值给指针变量, 指针变量存储的是变量的内存地址，*指针变量可以获取变量的值
	pointerExample := &integerExample
	fmt.Printf("指针: %p, 值: %d\n", pointerExample, *pointerExample)
	pointerExample2 := new(int)
	fmt.Printf("指针2: %p, 值: %d\n", pointerExample2, *pointerExample2) // 指针2: 0xc00008a018, 值: 0, 初始new仅仅分配了内存，全给0值，make会分配内存结构，并初始化
	*pointerExample2 = 10
	fmt.Printf("指针2: %p, 值: %d\n", pointerExample2, *pointerExample2)

	// 通道示例
	channelExample := make(chan int, 1)
	go func() { channelExample <- 42 }()
	fmt.Printf("通道接收: %d\n", <-channelExample)

	// 函数类型示例
	functionExample := func(x int) int { return x * 2 }
	fmt.Printf("函数结果: %d\n", functionExample(5))
}

// 加法函数
func Add(a, b int) int {
	return a + b
}

// 多返回值函数
func MultiReturn(a int) (int, int) {
	return a, a * 2
}

// 可变参数函数
func VariadicFunc(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 递归函数：阶乘
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// 函数作为参数
func ApplyFunc(f func(int) int, x int) int {
	return f(x)
}

// 结构体定义
type Rectangle struct {
	Width, Height float64
}

// 结构体方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 控制流示例
func ControlFlowExamples() {
	// if-else 示例
	if x := 10; x > 5 {
		fmt.Println("x 大于 5")
	} else {
		fmt.Println("x 不大于 5")
	}

	// switch 示例
	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday:
		fmt.Println("是周末")
	default:
		fmt.Println("是工作日")
	}

	// for 循环示例
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// range 循环示例，range用于遍历切片，数组，map，通道等，返回索引和值
	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}

	// defer 示例
	defer fmt.Println("这将最后打印")

	// 恢复 panic 示例
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("从 panic 中恢复:", r)
		}
	}()
	// 注释掉 panic 以允许程序继续运行
	panic("出错了！") //panic触发会让recover收到信号，然后可以恢复
}

// 接口定义
type Shape interface {
	Area() float64
}

// 圆形结构体
type Circle struct {
	Radius float64
}

// 圆形的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 多态函数
func PrintArea(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
}

// 并发示例
func ConcurrencyExamples() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("这是一个 goroutine")
	}()

	ch := make(chan int, 1)
	go func() {
		ch <- 42
	}()
	fmt.Println("从通道接收:", <-ch)

	select {
	case msg := <-ch:
		fmt.Println("接收到:", msg)
	default:
		fmt.Println("没有接收到消息")
	}

	var mu sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println("计数器:", counter)

	wg.Wait()
}

// 反射示例
func ReflectionExample(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("类型: %v, 值: %v\n", t, v)
}

// 自定义错误类型
type MyError struct {
	message string
}

func (e MyError) Error() string {
	return e.message
}

// 返回错误的函数
func SomeFunction() error {
	return MyError{"这是一个自定义错误"}
}

// 变量示例函数
func VariableExamples2() {
	// 基本类型变量声明
	var intVar int = 42
	var floatVar float64 = 3.14
	var stringVar string = "Hello, Go!"
	var boolVar bool = true

	// 简短声明
	sliceVar := []int{1, 2, 3, 4, 5}
	mapVar := map[string]int{"apple": 1, "banana": 2}

	fmt.Printf("整数: %d\n浮点数: %f\n字符串: %s\n布尔值: %v\n切片: %v\n映射: %v\n", 
		intVar, floatVar, stringVar, boolVar, sliceVar, mapVar)
}

// 基本函数示例
func Add2(a, b int) int {
	return a + b
}

// 多返回值函数
func MultiReturn2(x int) (int, int) {
	return x * 2, x * 3
}

// 可变参数函数
func VariadicFunc2(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 递归函数 - 阶乘
func Factorial2(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial2(n-1)
}

// 高阶函数
func ApplyFunc2(f func(int) int, x int) int {
	return f(x)
}

// 结构体和方法
type Rectangle2 struct {
	Width, Height float64
}

func (r Rectangle2) Area2() float64 {
	return r.Width * r.Height
}

// 控制流示例
func ControlFlowExamples2() {
	// if-else
	x := 10
	if x > 5 {
		fmt.Println("x大于5")
	} else {
		fmt.Println("x小于或等于5")
	}

	// switch-case
	switch x {
	case 10:
		fmt.Println("x是10")
	case 5:
		fmt.Println("x是5")
	default:
		fmt.Println("x是其他值")
	}

	// for循环
	for i := 0; i < 3; i++ {
		fmt.Printf("循环: %d\n", i)
	}
}

// 接口和多态
type Shape2 interface {
	Area2() float64
}

type Circle2 struct {
	Radius float64
}

func (c Circle2) Area2() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 并发示例
func ConcurrencyExamples2() {
	// goroutine
	go func() {
		fmt.Println("异步goroutine")
	}()

	// 通道
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println("接收通道值:", <-ch)

	// WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2")
	}()
	wg.Wait()
}

// 反射示例
func ReflectionExample2(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("类型: %v, 值: %v\n", t, v)
}

// 错误处理示例
func SomeFunction2() error {
	if runtime.GOOS == "darwin" {
		return fmt.Errorf("macOS系统不支持此操作")
	}
	return nil
}

// 文件和JSON处理示例
func FileAndJSONExamples() {
	// JSON编码
	data := map[string]string{"name": "Go语言", "version": "1.23"}
	jsonBytes, _ := json.Marshal(data)
	fmt.Println("JSON编码:", string(jsonBytes))

	// 文件操作
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("文件创建错误:", err)
		return
	}
	defer file.Close()

	// 写入文件
	_, err = io.WriteString(file, "Hello, Go语言文件操作!")
	if err != nil {
		fmt.Println("写入文件错误:", err)
	}

	// 字符串处理
	fmt.Println("大写转换:", strings.ToUpper("hello go"))
	fmt.Println("字符串包含:", strings.Contains("Go语言", "语言"))
}

