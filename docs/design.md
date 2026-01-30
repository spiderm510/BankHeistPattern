# Monopoly Go – Bank Heist Pattern Recognition Service

## Design & Implementation Overview

## 1. Overall Service Approach

The service is designed as a **stateful pattern-recognition system** that assists players in the Monopoly Go Bank Heist mini-game by recommending the next vault door to open with the highest probability of yielding a **Diamond Ring (Bankrupt outcome)**.

At system startup, the service **initializes all 27,720 possible outcome patterns** of the 12-door grid. Each pattern represents a complete game configuration and is assigned an initial equal frequency (weight). These patterns form the **probabilistic knowledge base** of the system.

The service exposes two main REST APIs:

**Prediction API** (/api/predict)
Uses partial game information to estimate Diamond Ring probabilities for unrevealed doors and recommends the optimal next move.

**Pattern-Frequency Update API** (/api/update)
Learns from completed games by increasing the frequency of the matching pattern, allowing the system to improve future predictions.

All data is persisted using **file-based JSON storage**, ensuring learning persists across service restarts.

## 2. Structure of the Patterns

Patterns are represented using a fixed structure:
```go
type Pattern struct {
    Doors     [12]int `json:"doors"`     // index-based door outcomes
    Frequency int     `json:"frequency"` // weight
}
```
Pattern representation details:
* The 12 doors are encoded in a row-major, index-based format:

  ```go
  index = (row - 1) * 4 + (col - 1)
  ```
* Each door value is:
  * 1 → Diamond Ring
  * 2 → Cash Stacks
  * 3 → Silver Coins

* Frequency represents how often this exact configuration has occurred in completed games.

This structure is intentionally simple and fixed to ensure:
* Fast comparisons
* Easy persistence
* Deterministic learning behavior

## 3. Prediction and Recommendation Logic
### 3.1 Pattern Matching

When a prediction request is received:

  1. The service converts the user’s revealed doors into index-based constraints.

  2. All stored patterns are filtered to retain only those consistent with the revealed outcomes.

A pattern is considered valid if, for every revealed door, the pattern contains the same outcome at the corresponding index.

### 3.2 Probability Calculation

For each unrevealed door:
1. The service iterates over all matched patterns.
2. Outcome frequencies are accumulated using the pattern’s Frequency as a weight.
3. The Diamond Ring probability is calculated as:

```sql
P(Diamond at door i) =
  sum(frequency of patterns where door i = Diamond)
  --------------------------------------------------
        sum(frequency of all matched patterns)
```


This produces a weighted probability estimate that reflects learned behavior.

### 3.3 Recommendation Strategy

The recommendation engine selects the next door based on:
1. **Highest Diamond Ring probability**
2. **Lowest Silver Coin probability** (tie-breaker)
3. **Pattern frequency weighting**

This strategy directly supports the service goal of maximizing the chance of achieving a Bankrupt (100%) steal.

## 4. Influence of Pattern Frequencies on Predictions

Pattern frequencies act as adaptive weights in the probability model.
* Initially, all patterns have equal influence.
* When a game is completed, the exact matching pattern’s frequency is incremented via the **/api/update** endpoint.

Over time, patterns that occur more frequently:
* Contribute more weight during probability calculations
* Bias predictions toward historically successful configurations

This mechanism allows the system to:
* Learn from real gameplay
* Adapt recommendations dynamically
* Improve prediction accuracy without external dependencies

The learning process is incremental, deterministic, and persistent, ensuring long-term consistency and explainability.

## 5. Summary

The service combines:

* Exhaustive pattern initialization (27,720 patterns)
* Deterministic pattern matching
* Frequency-weighted probability estimation
* Persistent learning from completed games

Together, these components form a robust and extensible pattern-recognition system that fulfills all core requirements of the assignment while remaining efficient, transparent, and easy to maintain.