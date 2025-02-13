package main

import "fmt"

// 策略接口：支付方式
type PaymentStrategy interface {
	Pay(amount float64) string
}

// 具体策略：信用卡支付
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用信用卡支付了 %.2f 元", amount)
}

// 具体策略：支付宝支付
type AlipayPayment struct{}

func (a *AlipayPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用支付宝支付了 %.2f 元", amount)
}

// 具体策略：微信支付
type WechatPayment struct{}

func (w *WechatPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用微信支付了 %.2f 元", amount)
}

// 上下文：支付服务
type PaymentService struct {
	strategy PaymentStrategy
}

func (p *PaymentService) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentService) Pay(amount float64) string {
	return p.strategy.Pay(amount)
}

// 客户端代码
func main() {
	// 创建支付服务
	paymentService := &PaymentService{}

	// 使用信用卡支付
	paymentService.SetStrategy(&CreditCardPayment{})
	fmt.Println(paymentService.Pay(100.50))

	// 使用支付宝支付
	paymentService.SetStrategy(&AlipayPayment{})
	fmt.Println(paymentService.Pay(200.75))

	// 使用微信支付
	paymentService.SetStrategy(&WechatPayment{})
	fmt.Println(paymentService.Pay(300.25))
}
