package main

import "fmt"

// 接口只有声明为平凡类型，不能声明为指针类型
type Name interface {
	printName()
}

type A struct {
	name string
}

// A 类型实现了接口的方法
func (a *A) printName()  {
	fmt.Println(a.name)
}

type B struct {
	a *A
}

type C struct {
	A
}

func getName(name Name)  {
	name.printName()
}


func main()  {
	a:= &A{"B"}
	// 所以 a 的实例的指针，可以被当作 a 类型来用。这里应该可以看作一个语法糖了。
	getName(a)

	b := B{a}
	// 因为不是匿名继承，所以这里必须显式地引用自己的成员变量来做这件事。
	b.a.printName()

	// 注意，即使使用了匿名的嵌入，这里也不能使用匿名初始化的语法，类似于CPP的初始化列表。
	c :=C{A: A{"name"}}

	// 但可以不经成员变量直接复用成员。
	c.printName()
}


