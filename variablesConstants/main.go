package main
import "fmt"

func main()  {
	var stationName string
	var trainTime int8
	var isUpDownTrain bool


	stationName="Thrissur"
	trainTime=10
	isUpDownTrain=false

	fmt.Println("Current Station: ",stationName)
	fmt.Println("Arrival time: ",trainTime)
	fmt.Println("Is train up down: ", isUpDownTrain)

	stationName="Ernamkulam"
	trainTime=12
	isUpDownTrain=true

	fmt.Println("Current Station: ",stationName)
	fmt.Println("Arrival time: ",trainTime)
	fmt.Println("Is train up down: ", isUpDownTrain)

	fmt.Println("sum of the numbers 2 & 10 is: ", 2+10)
	fmt.Println(23*1000)


	const earthsGravity=9.80665
	fmt.Println(earthsGravity)

	var favoriteSnack string
	favoriteSnack="momo"
	fmt.Println("My favouite snack is "+favoriteSnack)

}