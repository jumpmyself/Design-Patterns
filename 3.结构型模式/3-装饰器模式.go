package main

import "fmt"

// -------抽象层--------
type Phone interface {
	Show() // 构件的功能
}

// 抽象的装饰器、装饰器的基础类（该类本应该是interface、但是golang的interface语法不可以有成员属性）
type Decorator struct {
	phone Phone // 组合进来抽象的phone
}

func (d *Decorator) Show() {
	if d != nil {
		d.phone.Show() // 调用被装饰的构件的原方法
	}

}

// 具体的构件
type HuaWei struct {
}

func (h *HuaWei) Show() {
	fmt.Println("展示了华为手机")
}

type Xiaomi struct {
}

func (x *Xiaomi) Show() {
	fmt.Println("展示了小米手机")
}

// 具体的装饰器
type MoDecorator struct {
	Decorator // 继承基础的装饰器类（主要是继承phone的成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show() // 调用被装饰的构件的原方法

	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

type KeDecorator struct {
	Decorator // 继承基础的装饰器类（主要是继承phone的成员属性）
}

func (ke *KeDecorator) Show() {
	ke.phone.Show() // 调用被装饰的构件的原方法

	fmt.Println("带壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

//-------实现层--------

// -------业务逻辑层--------
func main() {
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()

	fmt.Println("----------")
	// 用贴膜的装饰器、得到一个贴膜的华为手机

	var moHuaWei Phone
	moHuaWei = NewMoDecorator(huawei)
	moHuaWei.Show()
	fmt.Println("----------")

	var xiaomi Phone
	xiaomi = new(Xiaomi)
	var keXiaomi Phone
	keXiaomi = NewKeDecorator(xiaomi)
	keXiaomi.Show()

}
