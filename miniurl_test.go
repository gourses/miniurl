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

//go playground (selain suorittaa go:ta)
//Esimerkki dokumentti ointiin

func ExampleHash() {
	const input = "https://github/heppu/miniurl"
	output := miniurl.Hash(input)
	fmt.Println(output)
	//output
	//
}

func BenchmarkHash(b *testing.B) {
	// benchmark ( go test -bench=. )
	// -run  komennolla voi rajata mitä testiä ajetaan
	// go test -bench=. -run='^S' -benchmem
	//  go test -bench=. -run='^S'
	const input = "https://github/heppu/miniurl"
	for n := 0; n < b.N; n++ {
		miniurl.Hash(input)
	}
}

func FuzzHash(f *testing.F) {
	//Fuzzy testing, protocol library testing
	f.Add("some strings")
	f.Fuzz(func(t *testing.T, input string) {
		output1 := miniurl.Hash(input)
		output2 := miniurl.Hash(input)
		assert.Equal(t, output1, output2)
		assert.Len(t, output1, 32)
		assert.Len(t, output2, 32)
	})
}
