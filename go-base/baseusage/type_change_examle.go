package baseusage

import (
	"fmt"
	"strconv"
)

func TypeChangeExample() {
	// 1. 基本类型转换
	// int 到 float64
	intValue := 42
	floatValue := float64(intValue)
	fmt.Printf("int 转 float64: %d -> %.2f\n", intValue, floatValue)

	// float64 到 int（会截断小数部分）
	var floatValue64 float64 = 3.14
	floatToIntValue := int(floatValue64)
	fmt.Printf("float64 转 int: 3.14 -> %d\n", floatToIntValue)

	// int 到 string
	intToStringValue := strconv.Itoa(intValue)
	fmt.Printf("int 转 string: %d -> %s\n", intValue, intToStringValue)

	// string 到 int
	stringToIntValue, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("string 转 int 错误:", err)
	} else {
		fmt.Printf("string 转 int: \"123\" -> %d\n", stringToIntValue)
	}

	// 2. 字符串与其他类型转换
	// float64 到 string
		// strconv.FormatFloat 函数用于将浮点数转换为字符串
		// 参数说明：
		// 1. 3.14: 要转换的浮点数
		// 2. 'f': 格式标记，'f'表示固定精度小数格式
		// 3. 2: 小数点后的位数
		// 4. 64: 表示输入的浮点数是float64类型
	floatToStringValue := strconv.FormatFloat(3.14, 'f', 2, 64)
	fmt.Printf("float64 转 string: 3.14 -> %s\n", floatToStringValue)

	// bool 到 string
	boolValue := true
	boolToStringValue := strconv.FormatBool(boolValue)
	fmt.Printf("bool 转 string: %t -> %s\n", boolValue, boolToStringValue)

	// 3. 接口类型转换和类型断言
	var interfaceValue interface{} = "Hello, Go!"
	
	// 安全的类型断言
	if strValue, ok := interfaceValue.(string); ok {
		fmt.Printf("接口类型断言成功: %s\n", strValue)
	}

	// 4. 类型转换的边界情况
	// 注意：超出范围的转换可能会导致溢出
	var smallInt int8 = 127
	var largeInt int16 = int16(smallInt)
	fmt.Printf("int8 转 int16: %d -> %d\n", smallInt, largeInt)

	// 5. 自定义类型转换
	type MyInt int
	var myIntValue MyInt = 100
	normalIntValue := int(myIntValue)
	fmt.Printf("自定义类型转换: MyInt(%d) -> int(%d)\n", myIntValue, normalIntValue)
}

// IntToString converts an integer to a string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}
