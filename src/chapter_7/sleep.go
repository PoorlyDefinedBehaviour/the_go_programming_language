package main

import (
	"flag"
	"fmt"
	"time"
)

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

type celsiusFlag struct {
	Celsius
}

// implement flag.Value interface
func (flag *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		flag.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		flag.Celsius = FahrenheitToCelsius(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	period := flag.Duration("period", 1*time.Second, "sleep period")
	temperature := CelsiusFlag("temperature", 20.0, "the temperature")

	flag.Parse()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	fmt.Println("temperature", *temperature)

}
