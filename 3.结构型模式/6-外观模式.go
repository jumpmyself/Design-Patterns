package main

import "fmt"

// 子系统：CPU
type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU 冻结")
}

func (c *CPU) Jump(position int) {
	fmt.Printf("CPU 跳转到位置 %d\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU 开始执行")
}

// 子系统：内存
type Memory struct{}

func (m *Memory) Load(position int, data string) {
	fmt.Printf("内存加载数据 '%s' 到位置 %d\n", data, position)
}

// 子系统：硬盘
type HardDrive struct{}

func (h *HardDrive) Read(lba int, size int) string {
	return fmt.Sprintf("硬盘从 LBA %d 读取数据，大小为 %d", lba, size)
}

// 外观类：计算机
type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputer() *Computer {
	return &Computer{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *Computer) Start() {
	fmt.Println("计算机正在启动...")
	c.cpu.Freeze()
	c.memory.Load(0, "启动扇区")
	c.cpu.Jump(0)
	c.cpu.Execute()
	fmt.Println("计算机启动完成")
}

func (c *Computer) Shutdown() {
	fmt.Println("计算机正在关闭...")
	fmt.Println("计算机已关闭")
}

func main() {
	computer := NewComputer()
	computer.Start()
	computer.Shutdown()
}
