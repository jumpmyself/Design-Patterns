package main

import "fmt"

// 司机张三、李四、汽车、宝马、奔驰
// -张三 开奔驰
// -李四 开宝马

// 奔驰汽车
type Benz struct {
}

func (b *Benz) run() {
	fmt.Println("Benz is running")
}

// 宝马汽车
type BWM struct {
}

func (b *BWM) run() {
	fmt.Println("BWM is running")
}

// 司机张三
type zhangsan struct {
}

func (zs *zhangsan) DriveBenz(benz *Benz) {
	benz.run()
	fmt.Println("zhangsan drive benz is done")
}

// 新加一个方法
func (zs *zhangsan) DriveBwm(benz *BWM) {
	benz.run()
	fmt.Println("zhangsan drive bwm is done")

}

// 司机李四
type lisi struct {
}

func (l *lisi) DriveBWM(bwm *BWM) {
	bwm.run()
	fmt.Println("lisi drive bwm is done")
}

// 新加一个方法
func (l *lisi) DriveBenz(benz *Benz) {
	benz.run()
	fmt.Println("lisi drive benz is done")
}

func main() {
	// 张三开奔驰
	benz := &Benz{}
	zhangsan := &zhangsan{}
	zhangsan.DriveBenz(benz)

	// 李四开宝马
	bwm := &BWM{}
	lisi := &lisi{}
	lisi.DriveBWM(bwm)
}
