package tempconv

// CToF Convert  C -> F
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Convert F -> C
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
