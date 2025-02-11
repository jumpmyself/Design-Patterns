package main

import "fmt"

// -----------抽象层--------------
type Fruit2 interface {
	Show() // 接口的方法
}

// -----------实现层--------------
type Apple struct {
	Fruit2
}

func (apple *Apple) Show() {
	fmt.Println("Apple")
}

type Banana struct {
	Fruit2
}

func (banana *Banana) Show() {
	fmt.Println("Banana")
}

type Pear struct {
	Fruit2
}

func (pear *Pear) Show() {
	fmt.Println("Pear")
}

// -----------工厂模块--------------
type Factory struct {
}

func (factory *Factory) CreateFruit(name string) Fruit2 {
	var fruit Fruit2

	if name == "Apple" {
		// apple构造初始化业务
		fruit = new(Apple) // 满足多态条件的赋值，父类指针指向子类对象
	} else if name == "Banana" {
		fruit = new(Banana)
	} else if name == "Pear" {
		fruit = new(Pear)
	}
	return fruit
}

// -------业务逻辑层-------
func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()

}
