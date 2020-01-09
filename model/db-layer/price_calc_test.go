package dblayer


import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	var raito1 float64 = 0.01
	var raito2 float64 = 0.02
	var cp = 100
	fmt.Println(uint64(float64(cp) * raito1 + 0.5))
	fmt.Println(uint64(float64(cp) * raito2 + 0.5))
	
	cp = 1000
	fmt.Println(uint64(float64(cp) * raito1 + 0.5))
	fmt.Println(uint64(float64(cp) * raito2 + 0.5))
}



