package main

import "fmt"

type Cat struct {
}

func (cat *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫添加一个可以睡觉的方法（使用继承实现 ）
type Cat2 struct {
	Cat
}

func (cat2 *Cat2) Sleep() {
	fmt.Println("小猫睡觉觉")
}

// 给小猫添加一个可以睡觉的fangfa(通过组合实现)
type Cat3 struct {
	C *Cat
}

func (cat3 *Cat3) Eat() {
	// 通过组合进来的原有类的功能
	cat3.C.Eat()
}

// 更轻量的方法是直接把cat传到cat3的方法里面，这样的话。结构体cat3与cat甚至没有依赖
//func (cat3 *Cat3) Eat(cat *Cat) {
//	cat.Eat()
//}

func (cat3 *Cat3) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := &Cat{}
	c.Eat()
	fmt.Println("--------------------")

	c2 := &Cat2{}
	c2.Eat()
	c2.Sleep()
	fmt.Println("--------------------")

	c3 := &Cat3{}
	c3.Eat()
	c3.Sleep()

}
