// Reference: http://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf
package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
)

// precision parameter: required space is 2^precision bytes,
// error rate is 1.04/sqrt(2^precision)
const precision uint8 = 11

func main() {
	var n int

	for {
		fmt.Printf("Enter number of random values: ")
		fmt.Scanln(&n)
		estimate := HLL(n)
		fmt.Printf("The estimated number if %f (err=%f%%)\n\n", estimate, 100*math.Abs(estimate-float64(n))/float64(n))
	}
}

// HLL implements HyperLogLog algorithm
// n is the number of random values generated
func HLL(n int) float64 {
	b := precision
	m := uint32(math.Pow(2, float64(b)))

	// Create mask for computing w
	var mask uint32 = (0xFFFFFFFF << b) >> b

	// Initialize a collection of m registers, M[0], ..., M[m-1], to 0;
	M := make([]uint8, m)

	// Generate random inputs and compute M[]
	for i := 0; i < n; i++ {
		x := genRandomHash()
		idx := x >> (32 - b)
		w := x & mask

		// Compute number of leading zeroes plus one
		rho := 32 - b - uint8(math.Log2(float64(w)))
		if M[idx] < rho {
			M[idx] = rho
		}
	}

	// Compute raw estimate
	// Let V be the number of registers equal to 0;
	V := 0
	var temp float64
	var j uint32
	for j = 0; j <= m-1; j++ {
		//fmt.Printf("M[j]=%d %f\n", M[j], math.Pow(2, -float64(M[j])))
		temp += math.Pow(2, -float64(M[j]))
		if M[j] == 0 {
			V++
		}
	}
	E := computeAlpha(m) * float64(m*m) / temp

	if E <= 2.5*float64(m) {
		// Small range correction, use LinearCount for better estimation
		if V != 0 {
			E = float64(m) * math.Log(float64(m)/float64(V))
		}
	} else if E > float64(1<<32)/30 {
		// Large range correction
		E = -float64(1<<32) * math.Log(1-E/float64(1<<32))

	}

	return E
}

func computeAlpha(m uint32) float64 {
	switch m {
	case 16:
		return 0.673
	case 32:
		return 0.697
	case 64:
		return 0.709
	default:
		return 0.7213 / (1 + 1.079/float64(m))
	}
}

func genRandomHash() uint32 {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)
	h := md5.New()
	md5sum := h.Sum(randBytes)
	return binary.BigEndian.Uint32(md5sum[:4])
}
