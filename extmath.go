package extmath

import "math"

const segment_threshold int = 1073741824 // 1GB


func fill_segment(segment_number, segment_size int) (segment []int) {

	var offset int

	segment = make([]int, segment_size, segment_size)
	offset = segment_number * segment_size
	for i := 0; i < segment_size; i++ {
		segment[i] = offset
		offset++
	}
	return segment
}


func sieve_of_eratosthenes(segment_size int) (primes []int) {

	var (
		prime             int = 2
		composite         int
		sqrt_segment_size int
		segment           []int
	)

	segment = fill_segment(0, segment_size)
	sqrt_segment_size = int(math.Sqrt(float64(segment_size)))
	for prime <= sqrt_segment_size {
		composite = int(math.Pow(float64(prime), 2))
		for composite < segment_size {
			segment[composite] = 0
			composite += prime
		}
		prime++
		for segment[prime] == 0 {
			prime++
		}
	}
	for i := 2; i < segment_size; i++ {
		if segment[i] != 0 {
			primes = append(primes, segment[i])
		}
	}
	return primes
}


func segmented_sieve(primes []int, segment_number, segment_size int) []int {

	var (
		prime, composite, limit, sqrt_limit, index, offset int
		segment                                            []int
	)

	segment = fill_segment(segment_number, segment_size)
	offset = segment[0]
	limit = segment[len(segment)-1]
	sqrt_limit = int(math.Sqrt(float64(limit)))
	index = 0
	prime = primes[index]
	for prime <= sqrt_limit {
		composite = int(math.Pow(float64(prime), 2))
		for composite < offset {
			composite += prime
		}
		for composite <= limit {
			segment[composite-offset] = 0
			composite += prime
		}
		index++
		if index < len(primes) {
			prime = primes[index]
		} else {
			break
		}
	}
	for i := 0; i < len(segment); i++ {
		if segment[i] != 0 {
			primes = append(primes, segment[i])
		}
	}
	return primes
}


func PrimeFactorization(number int) (factors []int) {

	var retry bool = true

	if number < 0 {
		return
	}
	// Start with a resonable amount of primes.
	primes := Primes(10000)
	result := number
	for retry {
		for i := 0; i < len(primes) && result != 1; {
			if result%primes[i] == 0 {
				factors = append(factors, primes[i])
				result /= primes[i]
			} else {
				i++
			}
		}
		if result != 1 {
			// We ran out of primes! Let's calculate them all.
			primes = Primes(number)
			result = number
			factors = factors[:0]
		} else {
			retry = false
		}
	}
	return
}


func Primes(upto int) (primes []int) {

	var segment_size, segment_total int

	if upto < 0 {
		return nil
	}
	if upto <= segment_threshold {
		primes = sieve_of_eratosthenes(upto)
		return primes
	}
	segment_size = int(math.Sqrt(float64(upto)))
	segment_total = upto / segment_size
	if upto%segment_size > 0 {
		segment_total++
	}
	primes = sieve_of_eratosthenes(segment_size)
	for i := 1; i < segment_total; i++ {
		primes = segmented_sieve(primes, i, segment_size)
	}
	return primes
}


// Greatest Common Divisor - Binary GCD algorithm (https://en.wikipedia.org/wiki/Binary_GCD_algorithm)
func Gcd(a, b uint) uint {

  if a == b {
    return a
  }
  if a == 0 {
    return b
  }
  if b == 0 {
    return a
  }
  // look for factors of 2
  if ^a & 1 == 1 {                    // a is even
    if b & 1 == 1 {                   // b is odd
      return Gcd(a >> 1, b)
    } else {                          // both a and b are even
      return Gcd(a >> 1, b >> 1) << 1 // add a power of two to the result
    }
  }
  if ^b & 1 == 1 {                    // a is odd, b is even
    return Gcd(a, b >> 1)
  }
  // reduce the larger argument
  if a > b {
    return Gcd((a - b) >> 1, b)
  }
  return Gcd((b - a) >> 1, a)
}


// Least Common Multiple - (https://en.wikipedia.org/wiki/Least_common_multiple)
func Lcm(a, b uint) uint {
  if a == 0 && b == 0 {
    return 0
  }
  if a > b {
    return (a / Gcd(a,b)) * b
  } else {
    return (b / Gcd(a,b)) * a
  }
}


// Factorial - n! = n.(n-1).(n-2)...1
func Factorial(x int) int {

  if x < 0 {
    return -1
  }
  if x == 0 {
    return 1
  }
  return x * Factorial(x-1)
}
