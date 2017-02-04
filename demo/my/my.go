package my

func Fib(n int) uint64 {
	var (
		previousNumber uint64 = 1
		number         uint64 = 1
	)

	if n < 3 {
		return number
	}

	for i := 2; i < n; i++ {
		previousNumber, number = number, previousNumber+number
	}

	return number
}

type rankMask uint16

func (d rankMask) findLowBit() int {
	return 1
}

func demo() {
	var ranks rankMask
	ranks = 0xf
	lowBit := ranks.findLowBit()
	_ = lowBit
}
