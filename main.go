package main

import (
	"fmt"
	"log"

	"./accounts"
	"./mydict"
)

func main() {
	account := accounts.NewAccount("miso")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(100)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)

	dictionary := mydict.Dictionary{"first": "First word"}

	word := "hello"
	def := "World"
	err2 := dictionary.Add(word, def)
	if err2 != nil {
		log.Fatalln(err2)
	}

	definition, _ := dictionary.Search(word)
	fmt.Println(definition)

	err3 := dictionary.Update(word, "second")
	if err3 != nil {
		log.Fatalln(err3)
	}
	definition2, _ := dictionary.Search(word)
	fmt.Println(definition2)

	err4 := dictionary.Delete(word)
	if err4 != nil {
		log.Fatalln(err4)
	}
	definition3, _ := dictionary.Search(word)
	fmt.Println(definition3)
}
