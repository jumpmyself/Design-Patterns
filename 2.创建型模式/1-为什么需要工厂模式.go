package main

import "fmt"

// 水果类
type Fruit struct {
	// 可能有很多水果
	// ....,
	// .....
	// ....,
	// .....
}

// 创建一个Fruit对象
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
		// 创建apple的逻辑
	} else if name == "banana" {
		// 创建香蕉的逻辑
	} else if name == "cherry" {
		// 创建樱桃的逻辑
	}

	return fruit
}

// 水果显示名称方法
// 如果要增加新类型的水果、必须修改Fruit结构体的的构造函数和其他相关代码、很麻烦、违反了开闭原则
func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "cherry" {
		fmt.Println("我是樱桃")
	}

}

// 业务逻辑
func main() {
	apples := NewFruit("apple")
	apples.Show("apple")

	bananas := NewFruit("banana")
	bananas.Show("bananas")
}
