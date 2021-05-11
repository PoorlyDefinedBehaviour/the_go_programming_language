package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.16
	FreezinhC     Celsius = 0
	BoilingC      Celsius = 100
)

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Types can have method's associated with them.
// These method's are calle type method's
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
