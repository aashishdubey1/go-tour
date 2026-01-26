package pointers

import "fmt"

type User struct {
	Name string
	Age  int
}

func Run() {

	fmt.Println("-------Go POINTERS-------")

	// SO WHAT IS POINTERS ?

	// In computers we store all varibales in some sort of memory like RAM
	// this is like memory block with a value to store and a address
	// wHen we create a varible it get store on some memory block that also has an address
	// we can access those address using pointers

	// So Pointers is just a variable that store a memory address
	// The & operator generates a pointer to its operand.
	// The * operator denotes the pointer's underlying value.
	// Unlike C, Go has no pointer arithmetic. => WHAT DOES IT MEANS ?

	i := 45 // created a varible , that will get store on memory block

	pI := &i // we use & to get memory address , now pI holds the memory address of I

	fmt.Println(pI) // if we print that we get some random gibrish value , the value of memory address

	fmt.Println(*pI) //(dereference) to get the actual value that store at that address we put * before pointer variable

	*pI = 455 // that will set the value at pI to 455 , try logging value of i, also as we can see the type of pI is *int means it does not hold int value but the value at those address holds the int data

	fmt.Println(i)

	fmt.Println("-------POINTERS EXERCISES-------")

	fmt.Println("-------SWAP-------")
	x := 10
	y := 20
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)

	fmt.Println("-------STRUCT-------")

	ram := User{
		Name: "Ram",
		Age:  45,
	}

	fmt.Println(ram)
	fmt.Println(ram.Age)
	birthday(&ram)
	fmt.Println(ram.Age)

}

func birthday(u *User) {
	u.Age = u.Age + 1 // here we don't need star (*), because Go automatically dereferences struct pointers when you access fields.
	// (*u).Age = (*u).Age + 1 // we don't do something like this , This works only for struct field access (and methods), to make code cleaner.
}

func swap(x, y *int) {
	c := *x
	*x = *y
	*y = c
}
