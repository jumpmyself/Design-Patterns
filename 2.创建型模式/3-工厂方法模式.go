package main

import "fmt"

// ----抽象层-------

// 水果类
type Fruit3 interface {
	Show()
}

// 工厂类（抽象的接口）
type AbstractFactory interface {
	CreateFruit() Fruit3 // 生产水果类（抽象）的生产链方法
}

// -----基础模块层-------

type Apple1 struct {
	Fruit3
}

func (apple *Apple1) Show() {
	fmt.Println("Apple")
}

type Banana1 struct {
	Fruit3
}

func (banana *Banana1) Show() {
	fmt.Println("Banana")
}

type Pear1 struct {
	Fruit3
}

func (pear *Pear1) Show() {
	fmt.Println("Pear")
}

// 新增加一个xigua
type Xigua struct {
	Fruit3
}

func (xigua *Xigua) Show() {
	fmt.Println("Xigua")
}

// ------基础的工厂模块-----

// 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (appleFactory *AppleFactory) CreateFruit() Fruit3 {
	var fruit Fruit3
	// 生产一个具体的苹果
	fruit = new(Apple1)
	return fruit
}

// 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (appleFactory *BananaFactory) CreateFruit() Fruit3 {
	var fruit Fruit3
	// 生产一个具体的苹果
	fruit = new(Apple1)
	return fruit
}

// 具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (appleFactory *PearFactory) CreateFruit() Fruit3 {
	var fruit Fruit3
	// 生产一个具体的苹果
	fruit = new(Apple1)
	return fruit
}

// 新增加一个xigua工厂
type XiguaFactory struct {
	AbstractFactory
}

func (xiguaFactory *XiguaFactory) CreateFruit() Fruit3 {
	var fruit Fruit3
	fruit = new(Apple1)
	return fruit
}

func main() {
	// 需求1：需要一个具体的苹果对象
	// 1.先要生成一个具体的苹果工厂
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)
	// 2.由苹果工厂生产具体的苹果对象
	var apple Fruit3
	apple = appleFactory.CreateFruit()
	apple.Show()

	// 需求2：需要一个具体的梨对象
	// 1.先要生成一个具体的梨工厂
	var pearFactory AbstractFactory
	pearFactory = new(AppleFactory)
	// 2.由梨工厂生产具体的苹果对象
	var pear Fruit3
	pear = pearFactory.CreateFruit()
	pear.Show()

	// 需求3：需要一个具体的梨对象
	// 1.先要生成一个具体的梨工厂
	pear1Factory := new(AppleFactory)
	pear1 := pear1Factory.CreateFruit()
	pear1.Show()

}
