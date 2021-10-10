// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilinC       Celsius    = 100
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilinF       Fahrenheit = 212
	AbsoluteZeroK Kelvin     = 0
	FreezingK     Kelvin     = 273.15
	BoilinK       Kelvin     = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", c)
}

func (c Kelvin) String() string {
	return fmt.Sprintf("%g°K", c)
}
