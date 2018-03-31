package permutations

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
)

var permutationTests = []struct {
	name     string
	input    []string
	expected [][]string
}{
	{
		"size 0",
		[]string{},
		[][]string{},
	},
	{
		"size 1",
		[]string{"1"},
		[][]string{
			[]string{"1"},
		},
	},
	{
		"size 2",
		[]string{"1", "2"},
		[][]string{
			[]string{"1", "2"},
			[]string{"2", "1"},
		},
	},
	{
		"size 3",
		[]string{"1", "2", "3"},
		[][]string{
			[]string{"1", "2", "3"},
			[]string{"1", "3", "2"},
			[]string{"2", "1", "3"},
			[]string{"2", "3", "1"},
			[]string{"3", "1", "2"},
			[]string{"3", "2", "1"},
		},
	},
}

func TestAllPermutationsSimple(t *testing.T) {
	for _, test := range permutationTests {
		t.Run(test.name, func(t *testing.T) {
			checkUniqueness(t, test.expected) // Assert tests are alright
			res := AllPermutationsNaive(test.input)
			assert.DeepEqual(t, res, test.expected)
		})
	}
}

func TestAllPermutationsLexicographic(t *testing.T) {
	for _, test := range permutationTests {
		t.Run(test.name, func(t *testing.T) {
			res := AllPermutationsLexicographic(test.input)
			assert.DeepEqual(t, res, test.expected)
		})
	}
}

var permutationCountTests = []struct {
	numObjects int
	expected   int
}{
	{
		0,
		0,
	},
	{
		1,
		1,
	},
	{
		2,
		2,
	},
	{
		3,
		6,
	},
	{
		4,
		24,
	},
	{
		5,
		120,
	},
	{
		6,
		720,
	},
	{
		10,
		3628800,
	},
}

func checkUniqueness(t *testing.T, perms [][]string) {
	t.Helper()
	for _, perm := range perms {
		count := 0
		for _, permToCheck := range perms {
			if reflect.DeepEqual(perm, permToCheck) {
				count++
			}
		}
		assert.Equal(t, count, 1, fmt.Sprintf("permuation %v should be unique", perm))
	}

}

func TestAllPermutationsNaiveCount(t *testing.T) {
	for _, test := range permutationCountTests {
		t.Run(fmt.Sprintf("n=%d,r=%d", test.numObjects, test.numObjects), func(t *testing.T) {
			if test.numObjects > 10 {
				t.Skipf("naive implementation too slow")
			}
			array := []string{}
			for i := 0; i != test.numObjects; i++ {
				array = append(array, fmt.Sprintf("%d", i))
			}
			perms := AllPermutationsNaive(array)
			assert.Equal(t, len(perms), test.expected)
			if test.numObjects < 10 {
				// Too slow on high number of permutations
				checkUniqueness(t, perms)
			}

		})
	}
}

func TestAllPermutationsLexicographicCount(t *testing.T) {
	for _, test := range permutationCountTests {
		t.Run(fmt.Sprintf("n=%d,r=%d", test.numObjects, test.numObjects), func(t *testing.T) {
			array := []string{}
			for i := 0; i != test.numObjects; i++ {
				array = append(array, fmt.Sprintf("%d", i))
			}
			perms := AllPermutationsLexicographic(array)
			assert.Equal(t, len(perms), test.expected)
			if test.numObjects < 10 {
				// Too slow on high number of permutations
				checkUniqueness(t, perms)
			}
		})
	}
}
