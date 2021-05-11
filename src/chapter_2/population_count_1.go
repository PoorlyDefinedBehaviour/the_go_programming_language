package populationcount

// The number of bits set to 1 in a uint64 value is called population count

var pc [256]byte

// a function named init() is called automatically
// when the package is initialized
func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}

func PopulationCount(x uint64) int {
	count := 0

	for i := 0; i < 8; i += 1 {
		count += pc[byte(x>>(i*8))]
	}

	return count
}
