package baseusage

import (
	"fmt"
	"strings"
)

func StringExamples() {
	s := "Hello, World!"

	// 字符串截断
	// 预期输出: 字符串截断: Hello, World! -> Hello
	fmt.Printf("字符串截断: %s -> %s\n", s, s[:5])

	// 字符串转slice
	// 预期输出: 字符串转slice: Hello, World! -> [72 101 108 108 111 44 32 87 111 114 108 100 33]
	bytes := []byte(s)
	fmt.Printf("字符串转slice: %s -> %v\n", s, bytes)

	// 字节slice转回字符串
	// 预期输出: 字节slice转回字符串: [72 101 108 108 111 44 32 87 111 114 108 100 33] -> Hello, World!
	backToString := string(bytes)
	fmt.Printf("字节slice转回字符串: %v -> %s\n", bytes, backToString)

	// 获取字符串长度
	// 预期输出: 字符串长度: Hello, World!  -> 13
	length := len(s)
	fmt.Printf("字符串长度: %s -> %d\n", s, length)


	// 字符串拼接
	// 预期输出: 字符串拼接: Hello, World -> Hello, World
	s1, s2 := "Hello", "World"
	result := s1 + ", " + s2
	fmt.Printf("字符串拼接: %s, %s -> %s\n", s1, s2, result)

	// 字符串查找
	// 预期输出: 字符串查找: 在 Hello, World! 中查找 'World' -> 索引 7
	index := strings.Index(s, "World")
	fmt.Printf("字符串查找: 在 %s 中查找 'World' -> 索引 %d\n", s, index)

	// 字符串替换
	// 预期输出: 字符串替换: Hello, World! -> Hello, Go!
	newS := strings.Replace(s, "World", "Go", 1)
	fmt.Printf("字符串替换: %s -> %s\n", s, newS)

	// 字符串分割
	// 预期输出: 字符串分割: Hello, World! -> [Hello World!]
	parts := strings.Split(s, ", ")
	fmt.Printf("字符串分割: %s -> %v\n", s, parts)

	// 字符串转大写
	// 预期输出: 字符串转大写: Hello, World! -> HELLO, WORLD!
	upper := strings.ToUpper(s)
	fmt.Printf("字符串转大写: %s -> %s\n", s, upper)

	// 字符串转小写
	// 预期输出: 字符串转小写: Hello, World! -> hello, world!
	lower := strings.ToLower(s)
	fmt.Printf("字符串转小写: %s -> %s\n", s, lower)

	// 去除首尾空白字符
	// 预期输出: 去除首尾空白字符: '  Hello, World!  ' -> 'Hello, World!'
	withSpaces := "  Hello, World!  "
	trimmed := strings.TrimSpace(withSpaces)
	fmt.Printf("去除首尾空白字符: '%s' -> '%s'\n", withSpaces, trimmed)

	// 判断字符串是否包含子串
	// 预期输出: 判断字符串是否包含子串: 在 Hello, World! 中查找 'World' -> true
	contains := strings.Contains(s, "World")
	fmt.Printf("判断字符串是否包含子串: 在 %s 中查找 'World' -> %t\n", s, contains)
}
