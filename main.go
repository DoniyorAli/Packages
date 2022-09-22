package main

import (
	"fmt"
	numbox "numbox/Working"
)
func main() {

	one := numbox.Work{Value: "12345"} 
	two := numbox.Work{Value: "12345"}
	fmt.Println("Two numbers were added:",numbox.Add(one, two))

	three := numbox.Work{Value: "12345"} 
	four := numbox.Work{Value: "345"}
	fmt.Println("Two numbers were Subtractioned:", numbox.Sub(three, four))

	five := numbox.Work{Value: "11"} 
	six := numbox.Work{Value: "10"}
	fmt.Println("Two numbers were multiplied:", numbox.Multiples(five, six))

	seven := numbox.Work{Value: "11"} 
	eight := numbox.Work{Value: "10"}
	fmt.Println("Two numbers were divisioned:", numbox.Division(seven, eight))

	nine := numbox.Work{Value: "123"} 
	ten := numbox.Work{Value: "10"}
	fmt.Println("Two numbers were modded:", numbox.Mod(nine, ten))

	// onee := numbox.Work{Value: "123"} 
	// fmt.Println("Two numbers were modded:", numbox.Abs(onee))


}

