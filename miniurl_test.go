package miniurl_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"miniurl"
	"testing"
)

//testataan ulkoisesti miniurl, kun package on eri
// go:ssa kansio ja pakettirakenne implisiittisesti parempi, mutta jos eri, niin lisä tyyppiä
// go:ssa on sisäänrakennettu refaktorointityökalu

func TestHashLength(t *testing.T) {
	const (
		input          = "https://github/heppu/miniurl"
		expectedLength = 32
	)
	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func TestHashDeterministic(t *testing.T) {
	const (
		input          = "https://github/heppu/miniurl"
		expectedLength = 32
	)
	output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)
	assert.Equal(t, output1, output2)
}

//go playground
//Esimerkki koodia

func ExampleHash() {
	const input = "https://github/heppu/miniurl"
	output := miniurl.Hash(input)
	fmt.Println(output)
	//output
	//
}
