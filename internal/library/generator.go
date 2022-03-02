package library

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandomNumber type alias for generated random numbers
type RandomNumber int

func (r RandomNumber) String() string {
	return fmt.Sprintf("%d", r)
}

func generateRandomInteger(min, max int) RandomNumber {
	v, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	v.Add(v, big.NewInt(int64(min)))
	return RandomNumber(v.Int64())
}
