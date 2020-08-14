package main

import "fmt"

// 多态实现

type MobileStorager interface {
	Read()
	Write()
}

// 移动硬盘
type MobileDisk struct {
}

func (m *MobileDisk) Read() {
	fmt.Println("移动硬盘在读取数据")
}

func (m *MobileDisk) Write() {
	fmt.Println("移动硬盘在写数据")
}

// u盘
type UDisk struct {
}

func (u *UDisk) Read() {
	fmt.Println("U盘在读取数据")
}

func (u *UDisk) Write() {
	fmt.Println("U盘在写入数据")
}

// 计算机
type Computer struct {
}

func (c *Computer) CpuRead(i MobileStorager) {
	i.Read()
}

func (c *Computer) CpuWrite(i MobileStorager) {
	i.Write()
}

func main() {
	m := &MobileDisk{}
	c := &Computer{}

	c.CpuRead(m)
	c.CpuWrite(m)
}
