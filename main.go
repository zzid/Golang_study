package main

import (
	"fmt"

	"github.com/zzid/mydict"
	// "log"
	// "github.com/zzid/account"
)

func main() {
	/* 
	==========================

		Account 

	==========================
	*/
	// account := account.NewAccount("zzid")
	// account.Deposit(10)
	// // fmt.Println(account.Balance())

	// /* error control */
	// // err := account.Withdraw(20)
	// // if err != nil {
	// // 	// log.Fatalln(err) // this kills program
	// // 	fmt.Println(err)
	// // }
	// // fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account)

	/* 
	==========================

		Dictionary 

	==========================
	*/

	/* Search */
	// dictionary := mydict.Dictionary{"first":"First word"}
	// // fmt.Println(dictionary["first"])
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	/* Add */
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, err := dictionary.Search("hello")
	// fmt.Println(hello)
	
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	/* Update */
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// dictionary.Add(word, "First")
	// err := dictionary.Update(word, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// newdef, _ := dictionary.Search(word)
	// fmt.Println(newdef)

	/* Delete */
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// dictionary.Add(word, "First")
	// def, _ := dictionary.Search(word)
	// dictionary.Delete(word)
	// def, err := dictionary.Search(word)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(def)
	// }

}
