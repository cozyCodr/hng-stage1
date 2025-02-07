package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

// IsPrime Check if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPerfect Check if a number is a perfect number
func IsPerfect(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

// IsArmstrong Check if a number is an Armstrong number
func IsArmstrong(n int) bool {
	sum := 0
	temp := n
	digits := int(math.Log10(float64(n))) + 1
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// DigitSum Sum of digits of a number
func DigitSum(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// GetFunFact Fetch fun fact from Numbers API
func GetFunFact(n int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d/math", n))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
