package main
import "fmt"
//Multiple Variable Declaration
func main()  {
	//On a single line, declare two int32 variables: one named magicNum and the other powerLevel.
	var magicNum,powerLevel int32
	//Assign magicNum a value of 2048 and powerLevel a value of 9001.
	magicNum=2048
	powerLevel=9001
	fmt.Println("MagicNum is :",magicNum,"Power Level is :",powerLevel)
	/*In a single line, declare and initialize two variables:
	 amount with a value of 10. unit with a value of "doll hairs".*/
	 amount,unit:=10,"doll hairs"
	fmt.Println(amount,unit," Thats expensive!")
}