package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestMult(t *testing.T) {
	total := Mult(3, 4)
	if total != 12 {
		t.Error("Mult was incorrect, got: %d, want: %d", total, 12)
	}
}

func TestDiv(t *testing.T) {
	total, _ := Div(8, 4)
	if total != 2 {
		t.Error("Div was incorrect, got: %d, want: %d", total, 2)
	}
}

func TestDivZero(t *testing.T) {
	Expected := "can't div with zero"
	_, e := Div(8, 0)
	if e.Error() != Expected {
		t.Errorf("Error got: %v, want: %v.", e, Expected)
	}
}

func TestSumTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestMultTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{2, 2, 4},
		{5, 2, 10},
	}

	for _, table := range tables {
		total := Mult(table.x, table.y)
		if total != table.n {
			t.Errorf("Mult of (%d*%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestDivTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
		e string
	}{
		{1, 1, 1, ""},
		{4, 2, 2, ""},
		{4, 0, -1, "can't div with zero"},
		{0, 0, -1, "can't div with zero"},
	}

	for _, table := range tables {
		total, e := Div(table.x, table.y)
		if e != nil && e.Error() != table.e {
			t.Errorf("Error got: %v, want: %v.", e, table.e)
		}
		if total != table.n {
			t.Errorf("Div of (%d/%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}