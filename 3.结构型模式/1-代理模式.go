package main

import "fmt"

type Goods struct {
	Kind string // 商品总类
	Fact bool   // 商品真伪
}

// ---------抽象层-----------
type Shopping interface {
	Buy(goods *Goods) // 某任务
}

// ---------实现层-----------
type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物、买了", goods.Kind)
}

type AmericaShopping struct {
}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物、买了", goods.Kind)
}

type AfrikaShopping struct {
}

func (afs *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲进行了购物、买了", goods.Kind)
}

// ---------海外代理-----------
type OverSeasProxy struct {
	Shopping Shopping // 代理某个主题、这是抽象的数据类型
}

func NewOverSeasProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{shopping}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	// 1 先辨别真伪
	if op.distinguish(goods) == true {
		// 2 调用具体要被代理的购物方式的Buy()方法
		op.Shopping.Buy(goods)
		// 3 海关安检
		op.check(goods)
	}
}

// 辨别真伪的能力
func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪。")
	if goods.Fact == true {
		fmt.Println("一眼顶真", goods.Kind, ",可以购买。")
	} else {
		fmt.Println("假货", goods.Kind, "别买")
	}
	return goods.Fact
}

// 案件流程
func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "]进行了海关检查，成功带回物品。")
}

// ---------业务逻辑层-----------
func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}
	g2 := Goods{
		Kind: "GET证书",
		Fact: false,
	}

	//var shopping Shopping
	//shopping = new(KoreaShopping)
	//
	//// 1.先辨别真伪
	//if g1.Fact == true {
	//	//2.购买
	//	shopping.Buy(&g1)
	//	//3.海关安检
	//	fmt.Println("带回物品")
	//}

	var kShopping Shopping
	kShopping = new(KoreaShopping)

	var kProxy Shopping
	kProxy = NewOverSeasProxy(kShopping)
	kProxy.Buy(&g1)

	var amShopping Shopping
	amShopping = new(AmericaShopping)

	var am_proxy Shopping
	am_proxy = NewOverSeasProxy(amShopping)
	am_proxy.Buy(&g2)
}
