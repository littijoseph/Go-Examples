package main
import "fmt"

func main(){
	fastFoodMenu := []string{"Burgers","Nuggets","Fries"}
	fmt.Println("...**....Welcome to Fast Food Cafe....*...")
	fmt.Println("The menu contains:",fastFoodMenu)
	//While the order is not quit, use an indefinite loop to call the askOrder() over & over.
	var order string
	var total int
	for order != "Quit" {
		fmt.Println("inside loop")
		order=askOrder()
		if contains(fastFoodMenu,order){
			total += 4
		}else {
			fmt.Println("Item not in Menu..")
		}
	}
	for amount:=total;amount>=0;amount -=2{
		fmt.Println("Amount to be paid:",amount)
	}
	fmt.Println("Total amount to be paid:",total)
}
//Function to take orders from customers
func askOrder() string{
	var input string
	fmt.Println("What would you like to order?")
	fmt.Scan(&input)
	return input
}
//With you now taking orders, you can begin to check if its on the menu.
func contains(menu []string,order string) bool{
	//write a loop that uses the range keyword with the menu array.
	for _,item:=range menu{
		if order==item{
			return true
		}
	}
	return false

}
