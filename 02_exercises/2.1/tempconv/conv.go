package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(Fahrenheit(c*9/5) + FreezingF)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(Kelvin(c) + FreezingK)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - FreezingF) * 5 / 9)
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(Kelvin((f-FreezingF)*5/9) + FreezingK)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - FreezingK)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(Fahrenheit((k-FreezingK)*9/5) + FreezingF)
}
