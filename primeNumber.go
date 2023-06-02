package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var num int
	fmt.Print("Bir sayı girin: ")
	_, err := fmt.Scanf("%d\n", &num)
	if err != nil {
		fmt.Println("Hatalı bir giriş yaptınız. Bir sayı girin.")
		return
	}
	primes := findPrimes(num)
	fmt.Printf("%d'e kadar olan asal sayılar: %v\n", num, primes)
}
