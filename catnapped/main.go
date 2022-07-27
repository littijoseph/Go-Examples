package main

import("fmt"
"math/rand"
"time")

func main()  {
	guestList:=[5]string{"Joseph","Joe","Josh","Johan","Jude"}
	catStorage:=[3]string{"Box","Basket","Toy chest"}
	fmt.Println("Guest List:",guestList)
	fmt.Println("Cat Storage:",catStorage)
	guest:=getRandomElement(guestList [:])
	storage:=getRandomElement(catStorage [:])
	fmt.Println(guest,"has napped the cat in",storage)
	
}

func getRandomElement(slice []string) string  {
	length:=len(slice)
	rand.Seed(time.Now().UnixNano())
	randomValue:=rand.Intn(length)
	return slice[randomValue]
}