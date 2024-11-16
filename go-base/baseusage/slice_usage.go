package baseusage

import (
	"fmt"
	"sort"
)

// SliceExamples demonstrates various slice operations and techniques
func SliceExamples() {
	// 1. 创建和初始化切片
	// Creating and initializing slices
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("初始切片:", fruits)

	// 2. 使用 make 创建切片
	// Creating slice with make
	numbers := make([]int, 5, 10)  // 长度5，容量10
	fmt.Println("使用 make 创建的切片:", numbers)

	// 3. 切片追加元素
	// Appending elements to a slice
	fruits = append(fruits, "date", "elderberry") 
	fmt.Println("追加元素后:", fruits)

	// 4. 切片截取
	// Slice slicing
	subFruits := fruits[1:4]  // 从索引1到3的子切片
	fmt.Println("子切片:", subFruits)

	// 5. 切片长度和容量
	// Slice length and capacity
	fmt.Printf("切片长度: %d, 容量: %d\n", len(fruits), cap(fruits))

	// 6. 复制切片
	// Copying slices
	fruitsCopy := make([]string, len(fruits))
	copy(fruitsCopy, fruits)
	fmt.Println("复制的切片:", fruitsCopy)

	// 7. 删除切片元素
	// Removing elements from a slice
	fruits = append(fruits[:2], fruits[3:]...)
	fmt.Println("删除元素后:", fruits)

	// 8. 切片排序
	// Sorting slices
	intSlice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(intSlice)
	fmt.Println("排序后的整数切片:", intSlice)

	// 9. 多维切片
	// Multidimensional slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("多维切片:", matrix)

	// 10. 切片遍历
	// Slice iteration
	fmt.Print("遍历切片: ")
	for i, fruit := range fruits {
		fmt.Printf("%d:%s ", i, fruit)
	}
	fmt.Println()

	// 11. 空切片和 nil 切片
	// Empty and nil slices
	var emptySlice []int
	var nilSlice []int = nil
	fmt.Printf("空切片: %v, nil切片: %v\n", emptySlice, nilSlice)
	fmt.Printf("空切片长度: %d, nil切片长度: %d\n", len(emptySlice), len(nilSlice))

	// 12. 切片作为函数参数
	// Slice as function parameter
	modifySlice(fruits)
	fmt.Println("函数修改后:", fruits)
}

// 演示切片作为参数的函数
// Function demonstrating slice as parameter
func modifySlice(s []string) {
	if len(s) > 0 {
		s[0] = "modified"
	}
}
