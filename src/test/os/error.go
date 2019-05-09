package os

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func TestPanic()  {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

func main() {
	fmt.Printf("error: %v", errNotFound)
	TestPanic()
}