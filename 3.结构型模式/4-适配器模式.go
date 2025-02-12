package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

// 被适配的角色、适配者
type v220 struct {
}

func (v v220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 适配器类
type Adapter struct {
	v220 *v220
}

func (a *Adapter) Use5V() {
	fmt.Println("phone进行了充电..")
	a.v220.Use220V()
}

func NewAdapter(v220 *v220) *Adapter {
	return &Adapter{v220: v220}
}

// 业务类
type Phone1 struct {
	v V5
}

func NewPhone(v V5) *Phone1 {
	return &Phone1{v: v}
}

// 普通的业务功能
func (p *Phone1) Charge() {
	fmt.Println("phone 进行了充电")
	p.v.Use5V()
}

func main() {
	phone := NewPhone(NewAdapter(&v220{}))
	phone.Charge()
}
