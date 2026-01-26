package loops

import "fmt"

func Run() {

	// There's only one loop in go 

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// The init and post statements are optional. and then it become while loop 
	// sum := 1
	// for sum < 10 {
	// 	sum += sum
	// 	fmt.Println(sum)
	// }
	
	// if we remove the condition we get a infinite loop like this 
	sum := 1
	for {
		if (sum > 100) { // can add checks to terminate infinite loop
			break
		}
		sum += sum
	}
	fmt.Println(sum)


	
	// for j := range 10 {
	// 	fmt.Println(j)
	// }


}