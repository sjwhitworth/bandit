# bandit

An implementation of the N-Armed bandit algorithm. N armed, or multi armed bandits, are good at converging on picking the optimal value whilst still retaining some exploration of all values.

In practical terms, they are useful for:

* A/B testing: instead of naively picking an experiment at random, you can leave the experiment running, and a natural winner (based on a reward function) will emerge.
* Model selection: given a selection of predictive models, choose the most accurate model.

### Usage

```
// A decay constant, controlling the extent to which you explore (select a value at random to get an estimate of the true reward value) or exploit (choose the option with the highest reward). The higher the decay factor, the more you explore the probability space.
decayFactor := 50

// The different choices that you want to select between
choices := []string{"blue button", "yellow button"}

// Generate the bandit object.
bandit := NewNArmedBandit(250, choices)

// Returns a choice. Keep a map somewhere that you can index into, or a switch statement, based upon this value. In this case, we returned 'blue button'
arm := bandit.Choose()

// When you know the 'reward' for this choice (e.g. in an A/B test, if this resulted in a 'successful outcome', then 1 else 0), then you can update the bandit. In this case, we will be more likely to pick the blue button in future (cause everyone hates yellow, amirite?)
bandit.Update(arm, 1)
```
