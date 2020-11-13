package main

import "fmt"

func main() {
	/*
	接口和结构体之间的联系和使用规范
	*/
	//接口：是一套标椎，适合于定义共性，抽离并定义一套标准
	//结构体：实体的描述和定义

	/*
	接口：有统一的共性进行判断时，往往使用的是接口

	*/
	person1:=NewChinese()
	age:=person1.Age()
	if age<20 {
		fmt.Println("在年龄上符合择偶标准")
	}else {
		fmt.Println("年龄不符合择偶标准")
	}
	//person1.IsMajiang

	person2:=NewJapanese()
	person2.Age()
	//person2.Php
}

type Person interface {
	Shanliang() bool
	WeiRenChuShi()
	Height() int
	Weight() int
	Age() int
	Salary() int
}

type Chinese struct {
	Name string
	Sex string
	IsShanliang bool
	High int
	Wei int
	AgeNum int
    Money int
	lsMajiang bool
}

func NewChinese() *Chinese {
	c:=&Chinese{
		Name:        "文豪",
		Sex:         "男",
		IsShanliang: true,
		High:        170,
		Wei:         140,
		AgeNum:      21,
		Money:       300000,
		lsMajiang:   true,
	}
	return c
}

func (c *Chinese) Shanliang() bool  {
	return c.IsShanliang
}

func (c *Chinese) WeiRenChuShi()  {
	fmt.Println(c.Name+"为人处世能力很好")
}

func (c *Chinese) Height() int {
	return c.High
}

func (c *Chinese) Weight() int {
	return c.Wei
}

func (c *Chinese) Age() int {
	return c.AgeNum
}

func (c *Chinese) Salary() int {
	return c.Money
}

func (c *Chinese) Majiang() {
	if c.lsMajiang {
		fmt.Println(c.Name+"会打麻将")
	}else {
		fmt.Println(c.Name+"不会打麻将")
	}
}

type Japanese struct {
	Name string
	AgeNum int
	lsShanliang bool
	High int
	Wei int
	Money int
	Php bool
}

func NewJapanese() *Japanese  {
	j:=&Japanese{
		Name:        "名媛",
		AgeNum:      23,
		lsShanliang: true,
		High:        165,
		Wei:         90,
		Money:       2000000,
		Php:         true,
	}
	return j
}

func (j *Japanese) Shanliang() bool {
	return j.lsShanliang
}

func (j *Japanese) WeiRenChuShi() {
	fmt.Println(j.Name+"为人处理能力很好")
}

func (j *Japanese) Height() int {
	return j.High
}

func (j *Japanese) Weight() int  {
  return j.Wei
}

func (j *Japanese) Age() int {
	return j.AgeNum
}

func (j *Japanese) Salary() int {
	return j.Money
}

func (j *Japanese) Eat(food string) {
	fmt.Println(j.Name+"喜欢吃：",food)
}

func (j *Japanese) Phps() bool {
	return j.Phps()
}