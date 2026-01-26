package conditionals

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func Run() {
	fmt.Println("-------Go Conditionals-------")
	fmt.Println("------- IF ELSE -------")
	fmt.Println(sqrt(2),sqrt(-4))
	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))
	fmt.Println("------- SWITCH -------")
	trySwitch()
	findDay()
}

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
// Switch without a condition is the same as switch true.
func findDay(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
		case today + 0:
			fmt.Println("Today")
		case today + 1:
			fmt.Println("Tomorrow") 
		case today + 2:
			fmt.Println("In Two Days")
		default:
			fmt.Println("Too Far away")
	}
}
// A switch statement is a shorter way to write a sequence of if - else statements.
// the break statement that is needed at the end of each case in those languages is provided automatically in Go.
// Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers. 
func trySwitch(){
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS. ")
	case "linux":
		fmt.Println("Linux. ")
	default:
		fmt.Printf("%s \n",os)
	}
}

// the decalred variable v has local scope , only can access in that block . this is something similar to ternary operator 
// It has 2 parts:
// Short statement (runs first)
// Condition (checked after)
// v is only available inside this if block (and its else, if present).
// Outside the if, v does not exist.
func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x,n);  v > lim { // If with a short statement 
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func sqrt(x float64) string {
	if x < 0 { 
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}