package main
import "fmt"
func main()  {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ",name,"!!")

	var age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("My name is %s and I am %d years old \n",name,age)

	newMessage:=fmt.Sprintf("This is %s and i am of %d years",name,age)
	fmt.Println(newMessage)
}