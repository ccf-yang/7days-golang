package main

import (
	"fmt"
	"go-base/baseusage"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
	"math"
)



func main() {
	// 注意： Go 需要在 go.mod 文件所在的目录下运行，以正确识别模块和包的导入路径，如果没有包依赖，可以用code-runner运行
	fmt.Println("=== 变量示例 ===")
	baseusage.VariableExamples()

	fmt.Println("\n=== 类型转换示例 ===")
	baseusage.TypeChangeExample()

	fmt.Println("\n=== string 用法示例 ===")
	baseusage.StringExamples()

	fmt.Println("\n=== slice 用法示例 ===")
	baseusage.SliceExamples()

	fmt.Println("\n=== 函数示例 ===")
	fmt.Println("加法结果:", baseusage.Add(5, 3))
	a, b := baseusage.MultiReturn(10)
	fmt.Printf("多返回值结果: %d, %d\n", a, b)
	fmt.Println("可变参数函数和:", baseusage.VariadicFunc(1, 2, 3, 4, 5))
	fmt.Println("5的阶乘:", baseusage.Factorial(5))
	fmt.Println("应用函数结果:", baseusage.ApplyFunc(func(x int) int { return x * x }, 4))

	rect := baseusage.Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形面积: %.2f\n", rect.Area())

	fmt.Println("\n=== 控制流示例 ===")
	baseusage.ControlFlowExamples()

	fmt.Println("\n=== 接口和多态示例 ===")
	circle := baseusage.Circle{Radius: 2.5}
	printArea(circle)
	printArea(rect)

	fmt.Println("\n=== 并发示例 ===")
	baseusage.ConcurrencyExamples()

	fmt.Println("\n=== 反射示例 ===")
	baseusage.ReflectionExample("你好，反射！")

	fmt.Println("\n=== 标准库使用 ===")
	fmt.Printf("格式化字符串: %s\n", "示例")
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(time.Now())
	fmt.Println(math.Sqrt(16))

	jsonData, _ := json.Marshal(map[string]string{"键": "值"})
	fmt.Println(string(jsonData))

	file, _ := os.Create("example.txt")
	defer file.Close()
	io.WriteString(file, "你好，文件！")

	fmt.Println("\n=== 文件和JSON示例 ===")
	baseusage.FileAndJSONExamples()

	fmt.Println("\n=== 错误处理示例 ===")
	if err := baseusage.SomeFunction(); err != nil {
		fmt.Println("错误:", err)
	}
}

func printArea(s baseusage.Shape) {
	fmt.Printf("形状面积: %.2f\n", s.Area())
}