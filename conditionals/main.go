package main
import "fmt"

func main()  {
	/*IF Conditional
	Inside the main function,add an if statement that checks heistReady and if it’s true, 
	print the string "Let's Go!"*/
	var heistReady bool
	heistReady=true
	if(heistReady){
		fmt.Println("Let's Go!!")
	}
	/* If-Else Conditional 
	Add an else statement to the existing if statement. 
	The block for the else statement prints "Act normal."*/
	if(heistReady==false){
		fmt.Println("Lets Go!")
	}else {
		fmt.Println("Act Normal!")
	}

	/* comparison operators '==' '!=' create a if statement that checks if 
	lockCombo and robAttempt are the same.*/
	lockCombo:="123"
	robAttempt:="25-06-2011"
	if(lockCombo==robAttempt){
		fmt.Println("Unlock vault")
	}
	if(lockCombo!=robAttempt){
		fmt.Println("Failed Heist")
	}

	/* Comparison Operators "<,>,<=,>=" In main.go create an if statement that checks 
	if vaultAmt is at least $200,000 (200000)if it is, print out "We're going to need more bags.".*/
	vaultAmt:=23457
	if(vaultAmt>=200000){
		fmt.Println("We're going to need more bags.")
	}else{
		fmt.Println("No more bags needed")
	}	
}