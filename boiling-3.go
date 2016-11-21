// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezingF %g°F = fToC(freezingF) %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("boilingF  %g°F = fToC(boilingF)  %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"

	fmt.Printf("BoilingC-FreezingC %g°C\n", BoilingC-FreezingC) // "100" °C
	BoilingF := CToF(BoilingC)
	fmt.Printf("BoilingF-CToF(FreezingC) %g°F\n", BoilingF-CToF(FreezingC)) // "180" °F
	fmt.Printf("FToC(BoilingF)-FreezingC %g°C\n", FToC(BoilingF)-FreezingC) // "100" °C

	var c Celsius
	var f Fahrenheit
	fmt.Println("c == 0", c == 0)                         // "true"
	fmt.Println("f >= 0", f >= 0)                         // "true"
	fmt.Println("Fahrenheit(c) == f", Fahrenheit(c) == f) // "true"
	fmt.Println("c == Celsius(f)", c == Celsius(f))       // "true"!

	ct := FToC(212.0)
	fmt.Println("ct.String()=",ct.String()) // "100°C"
	fmt.Printf("v ct=%v\n", ct)             // "100°C"; no need to call String explicitly
	fmt.Printf("s ct=%s\n", ct)             // "100°C"
	fmt.Println("ct=",ct)                   // "100°C"
	fmt.Printf("g ct=%g\n", ct)             // "100"; does not call String
	fmt.Println("float64(ct)=",float64(ct)) // "100"; does not call String

}

func (ct Celsius) String() string { return fmt.Sprintf("%g°C", ct) }

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
