package main
import "fmt"

func addHundred(num *int){
*num+=100
}
func main(){
	address:="My first address location"
	// to find the address of a variable
	fmt.Println("My first address location is : ",&address)
	//Pointers: variables that store addresses.* operator signifies that the variable stores an address
	var variableAddress *string
	variableAddress=(&address)
	fmt.Println(variableAddress)
	//use pointer to access the address and change its value! This is called dereferencing/indirecting.
	fmt.Println(address)
	*variableAddress="String changes"
	fmt.Println(address)

	x:=1
	addHundred(&x)
	fmt.Println(x)
}	