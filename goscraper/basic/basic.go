package main

import (
	"fmt"
	"strings"
)

/* naked return */
func lenAndUpper(name string) (length int, uppercase string){
	
	/* defer */
	defer fmt.Println("I'm done")
	/* defer will be executed after function is finished */

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	/* normal one */
	// for i :=0; i< len(numbers); i++ {
	// }

	/* Iteration index, and value */
	// for index, number := range numbers {
	// }
	for _, number := range numbers {
		total += number
	}
	return total
}
/* if else */
func canIDrink(age int) bool {
	/* variable expression */
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

/* switch */ 
func canIDrinkSwitch(age int) bool {
	/* Can use variable expression as well */
	switch  {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}


func main () {
	// fmt.Println(lenAndUpper("zzid"))
	
	// result  := superAdd(1,2,3,4,5,6)
	// fmt.Println(result)

	// fmt.Println(canIDrinkSwitch(16))

	/* Array */
	// names := [5]string{"zzid", "Hwang", "DongYun"}
	
	/* Slice */
	// names := []string{"zzid", "Hwang", "DongYun"}
	// fmt.Println(names)

	// names2 := append(names, "Nadia") // >> this will "return new Slice"
	// fmt.Println(names)
	// fmt.Println(names2)
	
}