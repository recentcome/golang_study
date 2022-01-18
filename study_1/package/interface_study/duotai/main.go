package main

import "fmt"

//定义一个USB接口
type Usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p *Phone) start() {
	fmt.Println("手机接口", p.Name, "开始运行")
}

func (p *Phone) stop() {
	fmt.Println("手机接口", p.Name, "停止运行")
}

type Camera struct {
	Name string
}

func (c *Camera) start() {
	fmt.Println("相机接口", c.Name, "开始运行")
}

func (c *Camera) stop() {
	fmt.Println("相机接口", c.Name, "停止运行")
}

func main() {
	//定义一个USB接口类型切片
	var usb = [3]Usb{
		&Phone{
			"小米",
		},
		&Phone{
			"华为",
		},
		&Camera{
			"尼康",
		},
	}
	usb[0].start()
}
