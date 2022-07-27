package main
import "fmt"

func main()  {
	coolSneakers := 65.99
  	niceNecklace := 45.50
	//Create a variable called taxCalculation, make it an float64
	var taxCalculation float64
	//Add the value of coolSneakers to taxCalculation using the += operator.
	taxCalculation += coolSneakers
	//Add the value of niceNecklace to taxCalculation using the += operator.
	taxCalculation += niceNecklace
	// Compute the NYC sales tax 8.875% of the purchase here:
	taxCalculation *= 0.08875
	fmt.Println("You made a purchase of ",coolSneakers+niceNecklace,"with a tax amount of:",taxCalculation,"total amount payable is:",niceNecklace+coolSneakers+taxCalculation)
}