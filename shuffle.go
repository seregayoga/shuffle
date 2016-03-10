package shuffle

import (
	"math/rand"
	"time"
)

// Shuffle mixes randomly the input slice
func Shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())

	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
