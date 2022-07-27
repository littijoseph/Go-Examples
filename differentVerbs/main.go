//Familiarize with different verbs

package main
import "fmt"

func main()  {
	guess:="C"
	fmt.Printf("Is your final answer %v \n",guess)

	floatValue:=1.75
	fmt.Printf("This value's type is %T \n",floatValue)
	fmt.Printf("This value's type is %T \n",guess)

	yearsOfExp:=3
	reqYearsOfExp:=15
	fmt.Printf("I have %d years of experience in GoLang but this job requires %d years of experience \n",yearsOfExp,reqYearsOfExp)

	stockPrice:= 3.5598
	fmt.Printf("Each share of Luna is %.3f \n",stockPrice)
	fmt.Printf("Each share of Luna is %f \n",stockPrice)

}