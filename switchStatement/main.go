package main
import "fmt"

func main()  {
	clothingChoice := "sweater"
	switch clothingChoice {
	case "shirt":fmt.Println("We have shirts in S and M only.")
	case "polos":fmt.Println("We have polos in L and XL only")
	case "sweater":fmt.Println("We have sweaters in S,M,L and XL")
	case "jacket":fmt.Println("We have jackets in all sizes")
	default:fmt.Println("Sorry we dont carry that")
		
	}
}