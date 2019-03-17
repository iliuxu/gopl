package tempconv

import "fmt"

type Celsius float64

type Fahrenheit float64

const (
	AbsoluteZeorC Celsius = -273.15
	FreezingC = 0
	BoilingC = 	100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}