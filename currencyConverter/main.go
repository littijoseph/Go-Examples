package main
import "fmt"

func main()  {
	currencies:=map[string]float32{"EUR":0.95,"JPY":130.2}
	fmt.Println(currencies)
	var dollarAmount float32
	var currency string
	fmt.Println("Enter the dollars to be converted:")
	fmt.Scan(&dollarAmount)
	if dollarAmount == 0{
		fmt.Println("Enter valid amount!")
	} else{
		fmt.Println("Current to be converted to:")
		fmt.Scan(&currency)
		rate,isValid:=currencies[currency]
		if isValid{
			targetCurrency:=rate*dollarAmount
			fmt.Println(targetCurrency)
		}else {
			fmt.Println("Error: currency not listed in map")
		}
	}
	
}