package main

import "fmt"

// -----> 抽象层 <-----
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// -----> 实现层 <-----
type BenZ struct {
	//.....
}

func (b BenZ) Run() {
	fmt.Println("Benz is running")
}

type Bmw struct {
}

func (b Bmw) Run() {
	fmt.Println("Bmw is running")
}

type Zhangsan struct {
}

func (z Zhangsan) Drive(car Car) {
	fmt.Println("Zhangsan is drive")
	car.Run()
}

type LiSi struct {
}

func (l LiSi) Drive(car Car) {
	fmt.Println("Lisi is drive")
	car.Run()
}

type WangWu struct {
}

func (w WangWu) Drive(car Car) {
	fmt.Println("Wangwu is drive")
	car.Run()
}

// -----> 业务逻辑层 <-----
func main() {
	// 张三开奔驰
	var benz Car
	benz = new(BenZ)
	var zhangsan Driver
	zhangsan = new(Zhangsan)
	zhangsan.Drive(benz)

	// 李四开宝马
	var bmw Car
	bmw = new(Bmw)
	var lisi Driver
	lisi = new(LiSi)
	lisi.Drive(bmw)

	// 张三开宝马
	zhangsan.Drive(bmw)

	// 王五开奔驰
	var wangwu Driver
	wangwu = new(WangWu)
	wangwu.Drive(bmw)
	wangwu.Drive(benz)

}
