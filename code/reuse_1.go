package code

import "fmt"


//代码复用

/*
将公共代码抽象出来
非公共代码为struct所有
*/

//type Dog struct {
//
//}
//
//func (d *Dog) say(msg string){
//	fmt.Println("--------do something before say----------")
//	fmt.Printf("dog say: %s\n", msg)
//	fmt.Println("--------do something after say----------")
//}
//
//func (d *Dog) sayHello() {
//	d.say("hello")
//}
//
//type Fish struct {
//
//}
//
//func (f *Fish) say(msg string){
//	fmt.Printf("fish say: %s\n", msg)
//}
//
//func (f *Fish) sayHello() {
//	fmt.Println("--------do something before say----------")
//	f.say("hello")
//	fmt.Println("--------do something after say----------")
//}
//
//sayHello为共同的方法,可以抽象出来,但是go的没有继承机制,无法通过go类似的继承机制实现对子类的动态调用
//say为各个struct自己的方法
//通过继承的方式则sayHello中调用say不会动态调用而是静态调用继承的say,而不是自己重写的say
//只能将sayHello抽象为一个公共的模块

type Sayer interface {
	say(string)
	name() string
}

type Dog struct {}
type Fish struct{}

func (d *Dog) say(msg string){
	fmt.Printf("%s say: %s\n", d.name(), msg)
}

func (d *Dog) name() string{
	return "dog"
}

func (f *Fish) say(msg string){
	fmt.Printf("%s say: %s\n", f.name(), msg)
}

func (f *Fish) name() string{
	return "fish"
}

func SayHello(sayer Sayer){
	fmt.Println("--------do something before say----------")
	sayer.say("hello")
	fmt.Println("--------do something before say----------")
}

func (d *Dog) sayHello(){
	SayHello(d)
}

func (f *Fish) sayHello(){
	SayHello(f)
}

func main(){
	d := &Dog{}
	f := &Fish{}
	d.sayHello()
	f.sayHello()
}
