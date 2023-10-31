package main

func main() {
	EvenOrOdd(55)
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}

}
