package main
import "fmt"

//Write a function called fuelGauge() which takes an int parameter, fuel.
func fuelGauge(fuel int){
	fmt.Println("You have,",fuel,"gallons of fuel left!")
}
/*function that calculates the amount of fuel needed to go to a specific planet.
that has a single parameter, a string called planet and returns an int*/
func calculateFuel(planet string) int  {
	var fuel int
	/*Create a switch or a series of if & else statements to assign fuel in the following cases:
	Venus: 300000
	Mercury: 500000
	Mars: 700000
	If it’s not listed above: 0*/
	
	switch planet{
	case "Venus":
		fuel=300000
	case "Mercury":
		fuel=500000
	case "Mars":
		fuel=700000
	default:
		fuel=0
	}
	return fuel
}

//greetPlanet(), this function should take a single parameter, planet, a string.
func greetPlanet(planet string){
	fmt.Println("Welcome to planet ",planet)
}

//create a function called cantFly() doesn’t have any parameters and doesn’t return anything.
func cantFly(){
	fmt.Println("We do not have the available fuel to fly there.")
}

//create a function called flyToPlanet() that uses the previous functions to coordinate our flight.
func flyToPlanet(planet string,fuel int) int{
	//Inside flyToPlanet() create the following variables fuelRemaining and fuelCost, both ints.
	var fuelRemaining,fuelCost int
	fuelRemaining=fuel
	fuelCost=calculateFuel(planet)
	if(fuelRemaining>=fuelCost){
		greetPlanet(planet)
		fuelRemaining-=fuelCost
	}else {
		cantFly()
	}
	return fuelRemaining
}
//Create a function that returns your spaceship back to your home planet.
func returnHome(planet string,fuel int)int{
	var fuelRemaining,fuelCost int
	homePlanet:="Earth"
	fuelRemaining=fuel
	fuelCost=calculateFuel(planet)
	if(fuelRemaining>=fuelCost){
		fuelRemaining-=fuelCost
		fmt.Println("Welcome back to ",homePlanet)
		fuelGauge(fuelRemaining)
	}else {
		cantFly()
	}
	return fuelRemaining
}
func main()  {
	fuel:=1000000
	var planetChoice string
	fmt.Println("Which Planet you want to visit?")
	fmt.Scan(&planetChoice)
	fmt.Printf("I wish to visit planet %s \n",planetChoice)
	fmt.Println("Your spaceship is on ",planetChoice)
	fuel=flyToPlanet(planetChoice,fuel)
	fuelGauge(fuel)
	fuel=returnHome(planetChoice,fuel)
	
}