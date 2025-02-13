package main

import "fmt"

// 抽象类：饮料
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

// 模板方法
func MakeBeverage(b Beverage) {
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	b.AddCondiments()
}

// 具体类：茶
type Tea struct{}

func (t *Tea) BoilWater() {
	fmt.Println("烧水")
}

func (t *Tea) Brew() {
	fmt.Println("浸泡茶叶")
}

func (t *Tea) PourInCup() {
	fmt.Println("倒入杯子")
}

func (t *Tea) AddCondiments() {
	fmt.Println("加入柠檬")
}

// 具体类：咖啡
type Coffee struct{}

func (c *Coffee) BoilWater() {
	fmt.Println("烧水")
}

func (c *Coffee) Brew() {
	fmt.Println("用滤网滴滤咖啡")
}

func (c *Coffee) PourInCup() {
	fmt.Println("倒入杯子")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("加入糖和牛奶")
}

// 客户端代码
func main() {
	tea := &Tea{}
	coffee := &Coffee{}

	fmt.Println("制作茶...")
	MakeBeverage(tea)

	fmt.Println("\n制作咖啡...")
	MakeBeverage(coffee)
}
