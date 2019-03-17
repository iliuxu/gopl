package tempconv

import "fmt"

func Example_one() {
	fmt.Printf("%g\n", BoilingC - FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC))
	fmt.Printf("%g\n", AbsoluteZeroC)
	fmt.Printf("%g\n", FToC(Fahrenheit(FreezingC)))
}