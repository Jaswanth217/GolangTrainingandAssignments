package main

import "fmt"

func main() {
	maxPrime := 100

	fmt.Println(maxPrime)
	fmt.Println()
	primes := getPrimes(maxPrime)
	fmt.Println(primes)
}

func getPrimes(maxNum int) []int {
	primes := make([]int, 0, maxNum)

	for i := 2; i <= maxNum; i++ {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			primes = append(primes, i)
		}
	}
	return primes
}
