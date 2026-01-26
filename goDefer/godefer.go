package godefer

import "fmt"

func Run() {
	fmt.Println("------- DEFER -------")
	// So bacially what defer is , in go when we write something with defer some expression or any function it exexute when it's sourrinding function return
	// basically all expression or function with defer moved to a stack call defer stack
	// and when the function return, go execute all the things from defer stack in LIFO order (in reverse order)
	
	defer fmt.Println("Hello")
	defer fmt.Println("HOla")
	fmt.Println("World")

	fmt.Println("------- DEFER WITH LOOPS -------")

	for i := range 10 {
		defer fmt.Println(i)
	}
	

	// WHY WE USE DEFER ? 
	// and this is very good thing to use i think , as a human there's chance that we forgot many things 
	// for example if we are working with files , we might forgot to close that file 
	// me might frogot to close db connection while working with databases 
	// int that cases we can use defer 

	// -> CLOSE FILES
	// file, _ := os.Open("a.txt")
	// defer file.Close()

	// -> CLOSE DB CONNECTIONS 
	// defer db.Close()

	// used for cleanup + safety + avoiding resource leaks

}