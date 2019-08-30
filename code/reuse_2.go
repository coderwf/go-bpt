package code

/*
代码复用的第二种方式
//抽象出公共方法组成新的struct并使用

*/

import "fmt"

type sayer interface {
	say(msg string)
}

type Animal struct {
	sayer
}

func (a *Animal) sayHello(){
	fmt.Println("--------do something before say----------")
	a.say("hello")
	fmt.Println("--------do something after say----------")
}

type dogSayer struct {
}

func (d *dogSayer) say(msg string){
	fmt.Printf("dog say: %s\n", msg)
}

type fishSayer struct {
}

func (f *fishSayer) say(msg string){
	fmt.Printf("fish say: %s\n", msg)
}


func main(){
	d := Animal{&dogSayer{}}
	d.sayHello()
	f := Animal{&fishSayer{}}
	f.sayHello()
}
