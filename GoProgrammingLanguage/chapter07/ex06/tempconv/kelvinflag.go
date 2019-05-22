package tempconv

import (
	"flag"
	"fmt"
)

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

type kelvinFlag struct{ Kelvin }

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Kelvin = CToK(Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Kelvin = Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
