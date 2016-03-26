package bandit

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var arm string

func TestBandit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// normal distribution, mean 0 with some randomness, so given enough iterations our value should be roughly 0.5 and 0.25 respectively
	reward := map[string]func() float64{
		"die":  func() float64 { return rand.NormFloat64() + 0.5 },
		"live": func() float64 { return rand.NormFloat64() + 0.25 },
	}

	b := NewNArmedBandit(250, []string{"die", "live"})

	// choose an actual, calculate the expected reward, update bandit with reward
	for i := 0; i < 100000; i++ {
		arm := b.Choose()
		reward := reward[arm]()
		b.Update(arm, reward)
	}

	assert.InDelta(t, 0.5, b.values["die"], 0.05)
	assert.InDelta(t, 0.25, b.values["live"], 0.05)
}

// Should be more than fast enough for anyone to care about
func BenchmarkEpsilonChoice(b *testing.B) {
	b.StopTimer()
	reward := map[string]func() float64{
		"die":  func() float64 { return rand.NormFloat64() + 0.5 },
		"live": func() float64 { return rand.NormFloat64() + 0.25 },
	}

	bandit := NewNArmedBandit(250, []string{"die", "live"})

	// choose an actual, calculate the expected reward, update bandit with reward
	for i := 0; i < 1000000; i++ {
		arm := bandit.Choose()
		r := reward[arm]()
		bandit.Update(arm, r)
	}

	b.StartTimer()

	// used to stop compiler removing the function call
	for i := 0; i < b.N; i++ {
		arm = bandit.Choose()
	}
}
