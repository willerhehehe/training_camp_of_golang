package main

import (
	"fmt"
	"practice/repository"
)


func main() {
	id := 1
	person, err := repository.GetPerson(id)
	if err != nil {
		panic(err)
	}
	if person == nil{
		fmt.Printf("no person with id: %d\n", id)
	}else{
		fmt.Println(person)
	}
}

