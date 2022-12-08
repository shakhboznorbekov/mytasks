// package main

// import "fmt"

// type Os struct {
// 	Name string
// }

// type Linux struct{
// 	Os
// }

// //overiding
// func (l Linux) GetName() string{
// 	//Super call
// 	name:=l.Os.GetName()

// 	return "Linux" + name
// }

// type Windows struct {
// 	Os
// }

// func (o Os) GetName() string{
// 	return o.Name
// }

// type OSInterface {
// 	GetName()  string
// }

// func Rum(os OSInterface) {
// 	fmt.Println("Loading"+ os.GetName()+"...")
// }

// type MacOs struct{
// 	Os
// }

// func main(){
// 	Run(Linux{Os{"Ubuntu 20.04"}})
// 	Run(Linux{Os{"Xubuntu"}})
// 	Run(Windows{Os{"Windows 11"}})
// 	Run(MacOs{Os{"Macintosh M2"}})
// }