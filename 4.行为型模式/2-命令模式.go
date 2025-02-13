package main

import "fmt"

// 核心计算模块、命令的接收者
type Doctor struct {
}

func (doctor *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (doctor *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象的命令
type Command interface {
	Treat()
}

// 治疗眼睛的病单
type CommandTreateEye struct {
	doctor *Doctor
}

func (command *CommandTreateEye) Treat() {
	command.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreateNose struct {
	doctor *Doctor
}

func (command *CommandTreateNose) Treat() {
	command.doctor.treatNose()
}

// 护士、命令的调用者
type Nurse struct {
	CmdList []Command // 收集的命令集合
}

// 发送病单、发送命令的方法
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat() // 多态现象、调用具体的命令、就会调用病单已经绑定的医生的诊断方法
	}
}

// 病人
func main() {
	//doctor := new(Doctor)
	//doctor.treatEye()
	//doctor.treatNose()

	// 依赖病单、通过填写病单让医生看病
	doctor := new(Doctor)
	cmdEye := CommandTreateEye{doctor}
	cmdEye.Treat()

	cmdNose := CommandTreateNose{doctor}
	cmdNose.Treat()

	// 护士
	nurse := new(Nurse)
	// 收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	// 执行病单指令
	nurse.Notify()
}
