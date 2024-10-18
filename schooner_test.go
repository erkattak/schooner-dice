package schooner_test

import (
	"fmt"
	"reflect"
	"schooner"
	"sort"
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		name      string
		category  schooner.Category
		diceRolls [][]int
		expected  []int
	}{
		{
			name:     "ones",
			category: schooner.Ones,
			diceRolls: [][]int{
				{2, 3, 4, 5, 6},
				{5, 8, 3, 1, 7},
				{1, 5, 7, 2, 1},
				{1, 1, 1, 2, 2},
				{6, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			expected: []int{0, 0, 2, 3, 4, 5},
		},
		{
			name:     "twos",
			category: schooner.Twos,
			diceRolls: [][]int{
				{1, 3, 4, 5, 6},
				{5, 2, 3, 1, 7},
				{1, 2, 7, 2, 1},
				{1, 1, 2, 2, 2},
				{6, 2, 2, 2, 2},
				{2, 2, 2, 2, 2},
			},
			expected: []int{0, 0, 4, 6, 8, 10},
		},
		{
			name:     "threes",
			category: schooner.Threes,
			diceRolls: [][]int{
				{2, 1, 4, 5, 6},
				{5, 8, 3, 1, 7},
				{1, 3, 3, 2, 1},
				{1, 3, 3, 3, 2},
				{3, 3, 3, 1, 3},
				{3, 3, 3, 3, 3},
			},
			expected: []int{0, 0, 6, 9, 12, 15},
		},
		{
			name:     "fours",
			category: schooner.Fours,
			diceRolls: [][]int{
				{2, 3, 2, 5, 6},
				{5, 8, 4, 1, 7},
				{1, 5, 4, 4, 1},
				{1, 4, 4, 4, 2},
				{4, 4, 4, 4, 1},
				{4, 4, 4, 4, 4},
			},
			expected: []int{0, 0, 8, 12, 16, 20},
		},
		{
			name:     "fives",
			category: schooner.Fives,
			diceRolls: [][]int{
				{2, 3, 4, 1, 6},
				{5, 8, 3, 1, 7},
				{1, 5, 7, 2, 5},
				{5, 5, 1, 5, 2},
				{5, 5, 5, 5, 1},
				{5, 5, 5, 5, 5},
			},
			expected: []int{0, 0, 10, 15, 20, 25},
		},
		{
			name:     "sixes",
			category: schooner.Sixes,
			diceRolls: [][]int{
				{2, 3, 4, 5, 1},
				{5, 8, 6, 1, 7},
				{1, 6, 7, 6, 1},
				{6, 1, 6, 2, 6},
				{6, 1, 6, 6, 6},
				{6, 6, 6, 6, 6},
			},
			expected: []int{0, 0, 12, 18, 24, 30},
		},
		{
			name:     "sevens",
			category: schooner.Sevens,
			diceRolls: [][]int{
				{2, 3, 4, 5, 6},
				{5, 8, 3, 1, 7},
				{7, 2, 6, 8, 7},
				{1, 5, 7, 7, 7},
				{7, 7, 7, 1, 7},
				{7, 7, 7, 7, 7},
			},
			expected: []int{0, 0, 14, 21, 28, 35},
		},
		{
			name:     "eights",
			category: schooner.Eights,
			diceRolls: [][]int{
				{2, 3, 4, 5, 6},
				{5, 8, 3, 1, 7},
				{1, 8, 8, 2, 1},
				{1, 8, 1, 8, 8},
				{8, 8, 8, 1, 8},
				{8, 8, 8, 8, 8},
			},
			expected: []int{0, 0, 16, 24, 32, 40},
		},
		{
			name:     "three_of_a_kind",
			category: schooner.ThreeOfAKind,
			diceRolls: [][]int{
				{1, 2, 3, 4, 5},
				{1, 1, 2, 3, 4},
				{1, 1, 1, 3, 4},
				{5, 6, 5, 6, 6},
			},
			expected: []int{0, 0, 10, 28},
		},
		{
			name:     "four_of_a_kind",
			category: schooner.FourOfAKind,
			diceRolls: [][]int{
				{1, 2, 3, 4, 5},
				{1, 1, 3, 4, 5},
				{1, 1, 1, 4, 5},
				{1, 1, 1, 1, 5},
			},
			expected: []int{0, 0, 0, 9},
		},
		{
			name:     "full_house",
			category: schooner.FullHouse,
			diceRolls: [][]int{
				{1, 1, 4, 4, 5},
				{1, 2, 3, 4, 5},
				{8, 8, 6, 6, 6},
			},
			expected: []int{0, 0, 25},
		},
		{
			name:     "small_straight",
			category: schooner.SmallStraight,
			diceRolls: [][]int{
				{4, 4, 4, 4, 4},
				{1, 1, 3, 3, 3},
				{4, 5, 1, 7, 4},
				{1, 8, 2, 3, 4},
				{2, 2, 3, 4, 5},
				{2, 3, 4, 5, 6},
				{1, 4, 5, 6, 7},
				{8, 7, 1, 6, 5},
			},
			expected: []int{0, 0, 0, 30, 30, 30, 30, 30},
		},
		{
			name:     "all_different",
			category: schooner.AllDifferent,
			diceRolls: [][]int{
				{1, 1, 3, 4, 5},
				{1, 2, 3, 4, 5},
				{4, 5, 6, 1, 3},
			},
			expected: []int{0, 35, 35},
		},
		{
			name:     "large_straight",
			category: schooner.LargeStraight,
			diceRolls: [][]int{
				{2, 5, 7, 4, 8},
				{1, 2, 3, 4, 5},
				{2, 3, 4, 5, 6},
				{3, 4, 5, 6, 7},
				{8, 6, 7, 4, 5},
			},
			expected: []int{0, 40, 40, 40, 40},
		},
		{
			name:     "schooner",
			category: schooner.Schooner,
			diceRolls: [][]int{
				{1, 1, 1, 3, 1},
				{1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2},
				{3, 3, 3, 3, 3},
				{4, 4, 4, 4, 4},
				{5, 5, 5, 5, 5},
				{6, 6, 6, 6, 6},
				{7, 7, 7, 7, 7},
				{8, 8, 8, 8, 8},
			},
			expected: []int{0, 50, 50, 50, 50, 50, 50, 50, 50},
		},
		{
			name:     "chance",
			category: schooner.Chance,
			diceRolls: [][]int{
				{2, 3, 4, 5, 6},
				{5, 8, 3, 1, 7},
				{1, 8, 8, 2, 1},
				{1, 8, 1, 8, 8},
				{8, 8, 8, 1, 8},
				{8, 8, 8, 8, 8},
			},
			expected: []int{20, 24, 20, 26, 33, 40},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.diceRolls) != len(tc.expected) {
				t.Fatalf("expected diceRolls (%d) and expected (%d) to be same length", len(tc.diceRolls), len(tc.expected))
			}
			for i, diceRoll := range tc.diceRolls {
				t.Run(fmt.Sprintf("roll_%d", i), func(t *testing.T) {
					score := schooner.Score(tc.category, diceRoll)
					if score != tc.expected[i] {
						t.Errorf("expected %d, got %d", tc.expected[i], score)
					}
				})
			}
		})
	}
}

func TestTopCategories(t *testing.T) {
	tests := []struct {
		name     string
		diceRoll []int
		expected []schooner.Category
	}{
		{
			name:     "all_different",
			diceRoll: []int{1, 5, 7, 3, 8},
			expected: []schooner.Category{
				schooner.AllDifferent,
			},
		},
		{
			name:     "full_house",
			diceRoll: []int{1, 2, 1, 2, 2},
			expected: []schooner.Category{
				schooner.FullHouse,
			},
		},
		{
			name:     "three_of_a_kind_chance",
			diceRoll: []int{3, 3, 3, 6, 7},
			expected: []schooner.Category{
				schooner.ThreeOfAKind,
				schooner.Chance,
			},
		},
		{
			name:     "large_straight",
			diceRoll: []int{8, 5, 7, 6, 4},
			expected: []schooner.Category{
				schooner.LargeStraight,
			},
		},
		{
			name:     "idk",
			diceRoll: []int{1, 1, 1, 1, 1},
			expected: []schooner.Category{
				schooner.Schooner,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := schooner.TopCategories(tc.diceRoll)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			sort.Slice(tc.expected, func(i, j int) bool {
				return tc.expected[i] < tc.expected[j]
			})
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
