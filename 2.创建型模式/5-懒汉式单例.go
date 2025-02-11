package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
三个要点：
	一是某个类只能有一个实例；
	二是它/须自行创建这个实例；
	三是它，须自行向整个系统提供这个实例。保证一个类永远只能有一个对象
*/

// 保证一个类永远只能有一个对象、这个对象还能被系统其他模块使用

// 1.如果这个类是首字母大写的话，这个类就可以被外界用来创建对象、可能会创建多个对象、所以这个类应该是非共有访问、首字母需要小写
type singelton1 struct {
}

// 2.还要有一个指针指向这个唯一对象、但是这个指针永远不能改变方法
// golang中没有常指针概念、因此只能通过将这个指针私有化不然外部模块访问
var (
	lock        sync.Mutex
	instance1   *singelton1
	initialized uint32
)

// 如果全部私有化、那么外部模块就无法访问到这个对象、所以需要对外提供一个方法来获取这个对象
func GetInstance1() *singelton1 {
	// 为了线程安全加个互斥锁
	//lock.Lock()
	//defer lock.Unlock()

	if atomic.LoadUint32(&initialized) == 1 {
		return instance1
	}
	//如果不存在、加锁申请
	lock.Lock()
	defer lock.Unlock()
	if instance1 == nil {
		instance1 = new(singelton1)
		// 设计标记为1
		atomic.StoreUint32(&initialized, 1)
	}
	return instance1
}

func (s *singelton1) Echo1() {
	fmt.Println("懒汉式单例的方法")
}

func main() {
	s1 := GetInstance1()

	s2 := GetInstance1()

	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}

// 延申1.对于懒汉式单例，如果高并发的情况下有可能创建多个对象、需要考虑线程安全、可以加个锁
// 延申2.如果每次获取这个单例的时候都加个锁、会比较慢、性能不好、可以使用原子标记
