package nofmtprintf

import (
	"fmt"
	. "fmt"
	format "fmt"
)

func dostuff() {
	fmt.Printf("hello")    // want `Don't use fmt.Printf`
	format.Printf("hello") // want `Don't use fmt.Printf`
	Printf("hello")        // want `Don't use fmt.Printf`
	fmt.Println("hey")

	fmt := printer{}
	fmt.Printf("fake")
	fmt.Println("fake")
}

type printer struct{}

func (printer) Printf(string)  {}
func (printer) Println(string) {}
