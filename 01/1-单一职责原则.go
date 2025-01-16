package main

import "fmt"

//type Clothes struct {
//}
//
//// 工作穿衣服
//func (c *Clothes) Onwork() {
//	fmt.Println("工作时穿的衣服")
//}
//
//// 逛街穿衣服
//func (c *Clothes) Onwalk() {
//	fmt.Println("逛街时穿的衣服")
//}
//
//func main() {
//	c := Clothes{}
//	// 工作的业务
//	fmt.Println("在工作")
//	c.Onwork()
//
//	// 逛街的业务
//	fmt.Println("在逛街")
//	c.Onwalk()
//}

// 因为clothes这个类可以做很多事情、所以可能会造成理解障碍、这个时候可以使用单一职责原则
// 类的职责单一、对外只提供一种功能、引起类变化的原因应该只有一个
type ClothesWalk struct {
}

func (cs *ClothesWalk) Style() {
	fmt.Println("逛街时穿的衣服")
}

type ClothesWork struct {
}

func (cs *ClothesWork) Style() {
	fmt.Println("工作时穿的衣服")
}

func main() {
	// 工作的业务
	cwork := ClothesWork{}
	cwork.Style()

	// 逛街的业务
	cwalk := ClothesWalk{}
	cwalk.Style()
}
