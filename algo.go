package permutations

import (
	"sort"
)

// AllPermutationsNaive find all permutations of the values.
func AllPermutationsNaive(values []string) [][]string {
	if len(values) == 1 {
		return [][]string{values}
	}
	if len(values) == 2 {
		return [][]string{
			[]string{values[0], values[1]},
			[]string{values[1], values[0]},
		}
	}

	allPerms := [][]string{}
	for val1Idx, val1 := range values {
		allValues := append([]string{}, values[:val1Idx]...)
		allValues = append(allValues, values[val1Idx+1:]...)
		perms := AllPermutationsNaive(allValues)

		for _, perm := range perms {
			newPerms := append([]string{val1}, perm...)
			allPerms = append(allPerms, newPerms)
		}
	}
	return allPerms
}

func factorial(n int) uint64 {
	var factval uint64 = 1
	if n >= 0 {
		for i := 1; i <= n; i++ {
			factval *= uint64(i)
		}
	}
	return factval
}

func reverse(slice []string) {
	for idx := len(slice)/2 - 1; idx >= 0; idx-- {
		opp := len(slice) - 1 - idx
		slice[idx], slice[opp] = slice[opp], slice[idx]
	}

}

// AllPermutationsLexicographic find all permutations of the values.
// Resources:
// - https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
// - https://www.quora.com/How-would-you-explain-an-algorithm-that-generates-permutations-using-lexicographic-ordering
func AllPermutationsLexicographic(values []string) [][]string {
	numValues := len(values)
	if numValues == 0 {
		return [][]string{}
	}
	numPerms := factorial(numValues)
	allPerms := make([][]string, numPerms)
	sort.Strings(values)

	allPerms[0] = values

	// fmt.Println("start:", values, "numPerms:", numPerms)

	for permNum := uint64(1); permNum != numPerms; permNum++ {
		nextPerm := make([]string, numValues)
		copy(nextPerm, values)

		var k int
		for k = numValues - 2; k >= 0; k-- {
			if values[k] < values[k+1] {
				break
			}
		}
		l := k
		for i := k + 1; i < numValues; i++ {
			if values[i] > values[k] {
				l = i
			}
		}

		// Swap
		nextPerm[k], nextPerm[l] = nextPerm[l], nextPerm[k]

		// Reverse
		reverse(nextPerm[k+1:])

		allPerms[permNum] = nextPerm
		// fmt.Printf("perm #%d - %v - k=%d l=%d - %v\n", permNum, values, k, l, nextPerm)
		values = nextPerm
	}

	return allPerms
}
