package main

import "fmt"

// 抽象的银行业务员
type AbstractBanker interface {
	DoWork() // 抽象的处理业务接口
}

// 存款的业务员
type SaveBanker struct {
	AbstractBanker
}

func (sb *SaveBanker) DoWork() {
	fmt.Println("进行了存款业务")
}

// 转账的业务员（这个是新增加的）
type TransferBanker struct {
	AbstractBanker
}

func (tb *TransferBanker) DoWork() {
	fmt.Println("进行了转账")
}

// 股票的业务员
type DepositBanker struct {
	AbstractBanker
}

func (db *DepositBanker) DoWork() {
	fmt.Println("股票业务")
}

//func main() {
//	// 存款业务
//	sb := SaveBanker{}
//	sb.DoWork()
//
//	// 转账业务
//	st := TransferBanker{}
//	st.DoWork()
//
//	// 股票业务
//	db := DepositBanker{}
//	db.DoWork()
//}

// 实现一个架构层（基于抽象层进行业务封装-针对interface接口进行封装）
func BankBusiness(banker AbstractBanker) {
	// 通过接口向下来调用（多态的现象）
	banker.DoWork()
}

func main() {
	// 存款业务
	BankBusiness(&SaveBanker{})
	// 转账业务
	BankBusiness(&TransferBanker{})
	// 股票业务
	BankBusiness(&DepositBanker{})
}
