package miniurl_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	miniurl "github.com/RichieRock/gourses-miniurl"
)

func TestHashLength(t *testing.T) {
	const (
		input = "https://github.com/RichieRock/gourses-miniurl"
		expectedLength = 32
	)

	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func TestHashIsDeterministic(t *testing.T) {
	const (
		input = "https://github.com/RichieRock/gourses-miniurl"
	)

	output1 := miniurl.Hash((input))
	output2 := miniurl.Hash((input))

	assert.Equal(t, output1, output2)
}

func ExampleHash() {
	const (
		input = "https://github.com/RichieRock/gourses-miniurl"
	)

	output := miniurl.Hash((input))

	fmt.Println(output)
	// output:
	// 7dd0152e5a839fde93952b98be0152ca
}

func BenchmarkHash(b *testing.B) {
	const input = "https://github.com/RichieRock/gourses-miniurl"

	for n := 0; n < b.N; n++ {
		miniurl.Hash(input)
	}
}

func FuzzHash(f *testing.F) {
	f.Add("some string here")

	f.Fuzz(func(t *testing.T, input string) {
		output1 := miniurl.Hash(input)
		output2 := miniurl.Hash(input)
		assert.Equal(t, output1, output2)
		assert.Len(t, output1, 32)
		assert.Len(t, output2, 32)
	})
}