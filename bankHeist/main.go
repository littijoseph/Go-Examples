package main
import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	//Declare the boolean variable isHeistOn with a value of true.
	isHeistOn:=true
	//Create another variable, eludedGuards and assign it a value of rand.Intn(100).
	eludedGuards:=rand.Intn(100)
	//if statement that checks if eludedGuards is at least 50. add a print statement that 
	//confirms that the heist continues
	if(eludedGuards>=50){
		fmt.Println("Looks like you've managed to make it past the guards.Good job, but remember, this is the first step.")
	}else {
		isHeistOn=false
		fmt.Println("Plan a better disguise next time?")
	}
	//create a variable openedVault and assign it a random number between 0 and 99
	openedVault:=rand.Intn(100)
	//Create an if statement that checks if isHeistOn is true and if openedVault is at least 70
	if(isHeistOn && openedVault>=70){
		fmt.Println("Vault has been opened. Grab and Go!!")
	}else if(isHeistOn){
		isHeistOn=false
		fmt.Println("Vault can't be opened.")
	}else{
		fmt.Println("What's the combo to this lock again??")
	}
	//create a variable leftSafely with value of rand.Intn(5) to account 5 different situations.
	leftSafely:=rand.Intn(5)
	//Create if statement that checks isHeistOn.add the code for previously mentioned 5 situations.
	if(isHeistOn){
		switch leftSafely {
		case 0:
			isHeistOn=false
			fmt.Println("Failed Heist")
		case 1:
			isHeistOn=false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn=false
			fmt.Println("When did they start raising dogs in vaults??")
		case 3:
			isHeistOn=false
			fmt.Println("Looks like this fingerprint scanner won’t accept any fingerprint…")
		default:
			fmt.Println("Success..Start the getaway car!")
			
		}
	}
	//add if statement that checks isHeistOn.create a variable amtStolen & 
	//assign it 10000 + rand.Intn(1000000)
	if(isHeistOn){
		amtStolen:=10000+rand.Intn(1000000)
		fmt.Println("Amount taken:",amtStolen)

	}
	fmt.Println("Is heist on:",isHeistOn)
}