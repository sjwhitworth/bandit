package bandit

import (
	"math"
	"math/rand"
)

type Bandit interface {
	Choose() string
	Update(arm string, reward float64)
	Epsilon() float64
}

type NArmedBandit struct {
	Decay   int
	Choices []string

	values map[string]float64
	counts map[string]int
}

// func NewNArmedBandit returns an implementation of the multi armed bandit algorithm, with decaying epsilon
func NewNArmedBandit(decay int, choices []string) *NArmedBandit {
	bandit := &NArmedBandit{
		Decay:   decay,
		Choices: choices,
		values:  make(map[string]float64),
		counts:  make(map[string]int),
	}

	bandit.initValuesAndCounts()
	return bandit
}

func (n *NArmedBandit) initValuesAndCounts() {
	for _, v := range n.Choices {
		n.values[v] = 0.0
		n.counts[v] = 0
	}
}

func (n *NArmedBandit) Choose() string {
	ep := n.Epsilon()
	if rand.Float64() > ep {
		return n.highestReward()
	}

	return n.Choices[rand.Intn(len(n.Choices))]
}

func (n *NArmedBandit) Update(choice string, reward float64) {
	n.counts[choice]++
	count := n.counts[choice]
	value := n.values[choice]
	newValue := ((float64(count)-1)/float64(float64(count)))*value + (1/float64(float64(count)))*reward
	n.values[choice] = newValue
}

func (n *NArmedBandit) Epsilon() float64 {
	sum := 0
	for _, v := range n.counts {
		sum += v
	}

	return float64(n.Decay) / (float64(sum) + float64(n.Decay))
}

func (n *NArmedBandit) highestReward() string {
	maxStr := ""
	maxV := math.Inf(-1)
	for k, v := range n.values {
		if v > maxV {
			maxStr = k
			maxV = v
		}
	}
	return maxStr
}
