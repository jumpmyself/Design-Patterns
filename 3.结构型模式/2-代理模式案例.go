package main

import "fmt"

// 抽象主题
type BeautyWoman interface {
	// 对男人抛媚眼
	MakeEyesWithMan()
	// 和男人共度美好时光
	HappyWithMan()
}

// 具体主题
type Panjinlian struct {
}

// 对男人抛媚眼
func (p *Panjinlian) MakeEyesWithMan() {
	fmt.Println("潘金莲抛了个媚眼")
}

// 和男人共度美好时光
func (p *Panjinlian) HappyWithMan() {
	fmt.Println("潘金莲和男人一起玩的很开心")
}

// 中间的代理人
type WangPo struct {
	woman BeautyWoman
}

func (w WangPo) MakeEyesWithMan() {
	w.woman.MakeEyesWithMan()
}

func (w WangPo) HappyWithMan() {
	w.woman.HappyWithMan()
}

func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman}
}

// 业务逻辑

func main() {
	// 大官人想找金莲、让王婆安排
	wangPo := NewProxy(new(Panjinlian))
	// 王婆命令金莲抛媚眼
	wangPo.MakeEyesWithMan()
	// 王婆命令金莲和大官人玩
	wangPo.HappyWithMan()
}
