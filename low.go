package intset

// Go has at least 32 bits in an int. We should really use int32 instead.
var bits_per_int int = 4 * 8
