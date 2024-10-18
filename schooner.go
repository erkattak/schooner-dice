package schooner

import (
	"sort"
)

// Score returns the score of a dice roll for the given category.
func Score(category Category, diceRoll []int) int {
	r := parseRoll(diceRoll)
	switch category {
	case Ones, Twos, Threes, Fours, Fives, Sixes, Sevens, Eights:
		// sum of selected number
		return r.sumKind(category.toInt())
	case ThreeOfAKind, FourOfAKind:
		return r.scoreMultipleOfKind(category.toInt())
	case FullHouse:
		return r.scoreFullHouse()
	case SmallStraight:
		return r.scoreSmallStraight()
	case AllDifferent:
		return r.scoreAllDifferent()
	case LargeStraight:
		return r.scoreLargeStraight()
	case Schooner:
		return r.scoreSchooner()
	case Chance:
		return sumAll(diceRoll)
	default:
		return 0
	}
}

// TopCategories returns the best scoring category of all qualifying categories.
// There could be multiple if there is a tie for the top category.
func TopCategories(diceRoll []int) []Category {
	scores := make(map[Category]int, len(allCategories))

	var topScore int
	for _, cat := range allCategories {
		score := Score(cat, diceRoll)
		scores[cat] = Score(cat, diceRoll)
		if score > topScore {
			topScore = score
		}
	}
	var topCats []Category
	for cat, score := range scores {
		if score == topScore {
			topCats = append(topCats, cat)
		}
	}

	return topCats
}

// Category is a scoring category that a dice roll may satisfy.
type Category string

func (c Category) toInt() int {
	if v, ok := values[c]; ok {
		return v
	}
	return 0
}

var values = map[Category]int{
	Ones:   1,
	Twos:   2,
	Threes: 3,
	Fours:  4,
	Fives:  5,
	Sixes:  6,
	Sevens: 7,
	Eights: 8,

	ThreeOfAKind: 3,
	FourOfAKind:  4,
}

const (
	Ones          Category = "ONES"
	Twos                   = "TWOS"
	Threes                 = "THREES"
	Fours                  = "FOURS"
	Fives                  = "FIVES"
	Sixes                  = "SIXES"
	Sevens                 = "SEVENS"
	Eights                 = "EIGHTS"
	ThreeOfAKind           = "THREE_OF_A_KIND"
	FourOfAKind            = "FOUR_OF_A_KIND"
	FullHouse              = "FULL_HOUSE"
	SmallStraight          = "SMALL_STRAIGHT"
	AllDifferent           = "ALL_DIFFERENT"
	LargeStraight          = "LARGE_STRAIGHT"
	Schooner               = "SCHOONER"
	Chance                 = "CHANCE"
)

var allCategories = []Category{
	Ones,
	Twos,
	Threes,
	Fours,
	Fives,
	Sixes,
	Sevens,
	Eights,
	ThreeOfAKind,
	FourOfAKind,
	FullHouse,
	SmallStraight,
	AllDifferent,
	LargeStraight,
	Schooner,
	Chance,
}

var (
	smallStraights = [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{4, 5, 5, 6},
		{4, 5, 6, 7},
		{5, 6, 7, 8},
	}
)

type roll map[int]int

func parseRoll(dice []int) roll {
	o := make(roll)
	for _, v := range dice {
		o[v]++
	}
	return o
}

func (r roll) keys() []int {
	keys := make([]int, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (r roll) sumKind(kind int) int {
	if c, ok := r[kind]; ok && c > 1 {
		return c * kind
	}
	return 0
}

func (r roll) sum() int {
	var res int
	for v, c := range r {
		res += v * c
	}
	return res
}

func (r roll) scoreMultipleOfKind(multiple int) int {
	for _, o := range r {
		if o >= multiple {
			return r.sum()
		}
	}
	return 0
}

func (r roll) isFullHouse() bool {
	return len(r) == 2
}

func (r roll) scoreFullHouse() int {
	if r.isFullHouse() {
		return 25
	}
	return 0
}

func (r roll) isSmallStraight() bool {
	if len(r) < 4 {
		return false
	}

	var lastKey int
	var skipped bool
	for _, k := range r.keys() {
		if lastKey == 0 {
			lastKey = k
			continue
		}
		if lastKey+1 != k && !skipped {
			skipped = true
		} else if lastKey+1 != k && skipped {
			return false
		}
		lastKey = k
	}

	return true
}

func (r roll) scoreSmallStraight() int {
	if r.isSmallStraight() {
		return 30
	}
	return 0
}

func (r roll) isAllDifferent() bool {
	return len(r) == 5
}

func (r roll) scoreAllDifferent() int {
	if r.isAllDifferent() {
		return 35
	}
	return 0
}

func (r roll) isLargeStraight() bool {
	if len(r) != 5 {
		return false
	}

	var lastKey int
	for _, k := range r.keys() {
		if lastKey == 0 {
			lastKey = k
			continue
		}
		if lastKey+1 != k {
			return false
		}
		lastKey = k
	}
	return true
}

func (r roll) scoreLargeStraight() int {
	if r.isLargeStraight() {
		return 40
	}
	return 0
}

func (r roll) isSchooner() bool {
	return len(r) == 1
}

func (r roll) scoreSchooner() int {
	if r.isSchooner() {
		return 50
	}
	return 0
}

func sumAll(dice []int) int {
	var res int
	for _, v := range dice {
		res += v
	}
	return res
}

func sumAllOfKind(t int, dice []int) int {
	var res int
	for _, v := range dice {
		if v == t {
			res += v
		}
	}
	return res
}
