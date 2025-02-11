package main

import "fmt"

// 类的改动是通过增加代码进行的、而不是修改源代码
type Banker struct {
}

// 存款业务
func (b Banker) Save() {
	fmt.Println("进行了 存款业务。。。")
}

// 转账业务
func (b Banker) Transfer() {
	fmt.Println("进行了 转账业务。。。")
}

// 支付业务
func (b Banker) Pay() {
	fmt.Println("进行了 支付业务。。。")
}

// 增加一个股票业务(这个功能是新添加的)
func (b Banker) Deposit() {
	fmt.Println("进行了 股票业务。。。")
}

func main() {
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()
	banker.Deposit()
}

// 由于前面的功能已经很多了、这时候新增加一个股票业务、可能会前面的功能互相影响、耦合度较高
