package math_problems

import (
	"fmt"
)

func Main() {
	romanToInt("LVIII")
}

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := range answer {
		n = i + 1
		ans := ""
		found := false
		if n%3 == 0 {
			found = true
			ans += "Fizz"
		}
		if n%5 == 0 {
			found = true
			ans += "Buzz"
		}
		if !found {
			ans = fmt.Sprint(n)
		}
		answer[i] = ans
	}
	return answer
}

func countPrimes(n int) int {
	isPrime := func(x int) bool {
		if x < 2 {
			return false
		}
		for i := 2; i <= x/2; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	count := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for i := 0; ; i++ {
		m := 1
		for j := 0; j <= i; j++ {
			m *= 3
		}
		if m == n {
			return true
		}
		if m > n {
			return false
		}
	}
}

func romanToInt(s string) int {
	repr := make(map[byte]int16)
	repr["I"[0]] = 1
	repr["V"[0]] = 5
	repr["X"[0]] = 10
	repr["L"[0]] = 50
	repr["C"[0]] = 100
	repr["D"[0]] = 500
	repr["M"[0]] = 1000
	prev := "f"[0]
	repr[prev] = 1<<15 - 1

	var res int16
	for _, v_ := range s {
		v := byte(v_)
		if repr[v] > repr[prev] {
			res = res - 2*repr[prev] + repr[v]
		} else {
			res += repr[v]
		}
		prev = v
	}
	return int(res)
}
