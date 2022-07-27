package main
import "fmt"
func main()  {
	//in addition to checking rightTime, use the && operator to check if rightPlace is also true.
	rightTime:=true
	rightPlace:=true

	if(rightTime&&rightPlace){
		fmt.Println("We're outta here!")
		} else {
			fmt.Println("Be patient...")
		}
	//use the || operator to check for enoughBags.
	enoughRobbers := false
	enoughBags := true
	if(enoughRobbers||enoughBags){
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
	
	readyToGo := true

	if (!readyToGo) {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}

	
