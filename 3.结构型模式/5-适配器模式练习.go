package main

import "fmt"

type Attack interface {
	Fight()
}

type Dabaojian struct {
}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了大宝剑技能")
}

type Hero struct {
	Name   string
	attack Attack // 攻击方式
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能")
	h.attack.Fight() // 使用具体的战斗方式
}

// 适配者
type PowerOff struct {
}

func (p *PowerOff) FangYu() {
	fmt.Println("敌方减少了100滴血")
}

// 适配器
type Adopt struct {
	poweroff PowerOff
}

func (a *Adopt) Fight() {
	a.poweroff.FangYu()
}

func NewAdopt(p *PowerOff) *Adopt {
	return &Adopt{*p}
}

func main() {
	//gaiLun := Hero{Name: "盖伦", attack: new(Dabaojian)}
	//gaiLun.Skill()

	gaiLun := Hero{
		Name:   "盖伦",
		attack: NewAdopt(new(PowerOff)),
	}
	gaiLun.Skill()
}
