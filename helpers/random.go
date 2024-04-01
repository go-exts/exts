package helpers

import (
	"math/rand/v2"
)

const (
	randCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*-_=+"
)

const (
	RandomNumber = iota + 1                                                // RandomNumber define the number random
	RandomLower  = RandomNumber << iota                                    // RandomLower define the lowercase random
	RandomUpper  = RandomNumber << iota                                    // RandomUpper define the uppercase random
	RandomSymbol = RandomNumber << iota                                    // RandomSymbol define the symbol random
	RandomAll    = RandomNumber | RandomLower | RandomUpper | RandomSymbol // RandomAll define the all random
)

const (
	// randomNone   = 0
	randomNumber = iota
	randomLower
	randomUpper
	randomSymbol
	randomAll
	randomMax
)

var randIndex = [randomMax][2]int{
	// randomNone:   {0, 0},
	randomNumber: {0, 10},
	randomLower:  {10, 36},
	randomUpper:  {36, 62},
	randomSymbol: {62, 74},
	randomAll:    {0, 74},
}

type Rand struct {
	randType int
	charset  string
}

// Rand returns a random string
func (r Rand) Rand(size int) string {
	return string(r.RandBytes(size))
}

// RandBytes returns a random byte slice
func (r Rand) RandBytes(size int) []byte {
	return randomByte(r.charset, size)
}

var globalRand = NewRand(RandomAll)

func loadCharset(rand int) string {
	var charset string
	if rand == RandomAll {
		return randCharset
	}
	var sta, end int
	for i := 0; i < randomMax-1; i++ {
		if rand&(1<<uint(i)) != 0 {
			sta, end = randomToIndex(i)
			charset += randCharset[sta:end]
		}
	}
	return charset
}

func randomToIndex(rand int) (int, int) {
	return randIndex[rand][0], randIndex[rand][1]
}

func NewRand(rnd int) *Rand {
	return &Rand{
		randType: rnd,
		charset:  loadCharset(rnd),
	}
}

// InitRand initializes the global random generator
func InitRand(rnd int) {
	globalRand = NewRand(rnd)
}

// RandString generates a random string of given size
func RandString(size int) string {
	return globalRand.Rand(size)
}

// RandBytes generates a random byte slice of given size
func RandBytes(size int) []byte {
	return globalRand.RandBytes(size)
}

func randomByte(charset string, size int) (r []byte) {
	ret := make([]byte, 0, size)
	ls := len(charset)
	for ; size > 0; size-- {
		ret[size-1] = charset[rand.IntN(ls)]
	}
	return ret
}
