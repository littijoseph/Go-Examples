package main
import "fmt"

func main()  {
	step1:="Breath in..."
	step2:="Breath out..."
	sprint:=fmt.Sprint(step1,step2)
	sprintln:=fmt.Sprintln(step1,step2)
	fmt.Println(sprint)
	fmt.Println(sprintln)

	template:="I wish i had a %v"
	pet:="dog"
	wish:=fmt.Sprintf(template,pet)
	fmt.Println(wish)

	correctAns := "A"
	answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)
	 
	fmt.Print(answer)
}