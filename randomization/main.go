package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main()  {
	fmt.Println(rand.Intn(100))
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}