package puzzle

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anttikivi/garns/util"
)

func TestParse(t *testing.T) {
	want := Puzzle{
		{6, 0, 0, 2, 0, 9, 4, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 2, 0},
		{0, 9, 0, 0, 0, 4, 0, 1, 6},
		{0, 0, 0, 0, 0, 8, 0, 3, 1},
		{4, 0, 3, 0, 1, 0, 2, 0, 8},
		{1, 7, 0, 5, 0, 0, 0, 0, 0},
		{2, 5, 0, 4, 0, 0, 0, 8, 0},
		{0, 3, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 6, 8, 0, 3, 0, 0, 4},
	}
	lines, _ := util.ReadLines("test1.txt")
	got, _ := Parse(lines)
	if len(want) != len(got) {
		t.Errorf("Parse() output length mismatch, expected: %d, got: %d", len(want), len(got))
	}
	if !reflect.DeepEqual(want, got) {
		for i, a := range want {
			fmt.Println("Row", i, "want:", a)
			fmt.Println("Row", i, "got: ", got[i])
		}
		t.Errorf("Parse() expected: %v, got: %v", want, got)
	}

	want = Puzzle{
		{5, 0, 0, 0, 3, 0, 8, 0, 0},
		{8, 0, 0, 0, 1, 0, 0, 3, 0},
		{0, 1, 0, 0, 0, 0, 6, 0, 7},
		{6, 0, 0, 2, 0, 0, 9, 0, 5},
		{0, 0, 0, 9, 0, 3, 0, 0, 0},
		{4, 0, 2, 0, 0, 8, 0, 0, 6},
		{2, 0, 1, 0, 0, 0, 0, 5, 0},
		{0, 6, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 9, 0, 4, 0, 0, 0, 2},
	}
	lines, _ = util.ReadLines("test2.txt")
	got, _ = Parse(lines)
	if len(want) != len(got) {
		t.Errorf("Parse() output length mismatch, expected: %d, got: %d", len(want), len(got))
	}
	if !reflect.DeepEqual(want, got) {
		for i, a := range want {
			fmt.Println("Row", i, "want:", a)
			fmt.Println("Row", i, "got: ", got[i])
		}
		t.Errorf("Parse() expected: %v, got: %v", want, got)
	}
}
