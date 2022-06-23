package coverage

import (
	"errors"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeopleLen(t *testing.T) {
	var tData = []struct {
		fName     string
		structure People
		expected  int
	}{
		{"empty",
			People{},
			0},
		{"filled",
			People{
				Person{"a", "b", time.Unix(1405110000, 0)},
				Person{"a", "c", time.Unix(1405220000, 0)},
				Person{"e", "f", time.Unix(1405330000, 0)},
				Person{"g", "h", time.Unix(1405440000, 0)}},
			4},
	}
	for _, tcases := range tData {
		t.Run(tcases.fName, func(t *testing.T) {
			got := tcases.structure.Len()
			if got != tcases.expected {
				t.Errorf("got %d, expected %d", got, tcases.expected)
			}
		})
	}
}

func TestPeopleLess(t *testing.T) {
	var tData = []struct {
		fName     string
		structure People
		expected  bool
	}{
		{
			"birthday",
			People{
				Person{"a", "b", time.Unix(1405300000, 0)},
				Person{"a", "b", time.Unix(1405200000, 0)},
			},
			true,
		},
		{
			"first name",
			People{
				Person{"a", "b", time.Unix(1405300000, 0)},
				Person{"b", "b", time.Unix(1405300000, 0)},
			},
			true,
		},
		{
			"last name",
			People{
				Person{"a", "a", time.Unix(1405300000, 0)},
				Person{"a", "b", time.Unix(1405300000, 0)},
			},
			true,
		},
	}
	for _, tcases := range tData {
		t.Run(tcases.fName, func(t *testing.T) {
			got := tcases.structure.Less(0, 1)
			if got != tcases.expected {
				t.Errorf("got %t, expected %t", got, tcases.expected)
			}
		})
	}
}

func TestPeopleSwap(t *testing.T) {
	structure := People{
		Person{"a", "b", time.Unix(1405300000, 0)},
		Person{"c", "d", time.Unix(1405400000, 0)},
	}
	structure.Swap(0, 1)
	s := structure[0]
	if s.firstName != "c" && s.lastName != "d" && s.birthDay != time.Unix(1405400000, 0) {
		t.Errorf("swap failed")
	}
}

func TestNewM(t *testing.T) {
	var tData = []struct {
		fName     string
		input     string
		expectedM *Matrix
		expectedE error
	}{
		{
			"normal case",
			"1 3\n5 7",
			&Matrix{rows: 2, cols: 2},
			nil,
		},
		{
			"invalid element",
			"one",
			&Matrix{rows: 1, cols: 1},
			errors.New("strconv.Atoi: parsing \"one\": invalid syntax"),
		},
		{
			"single element",
			"5",
			&Matrix{rows: 1, cols: 1},
			nil,
		},
		{
			"invalid matrix",
			"1\n2 3",
			nil,
			errors.New("Rows need to be the same length"),
		},
	}
	for _, tcases := range tData {
		t.Run(tcases.fName, func(t *testing.T) {
			matrix, err := New(tcases.input)
			if err != nil {
				if tcases.expectedE.Error() != err.Error() {
					t.Errorf("got error `%v`, expected error `%v`", err, tcases.expectedE)
				}
			} else {
				if matrix.rows != tcases.expectedM.rows || matrix.cols != tcases.expectedM.cols {
					t.Errorf("got matrix %v, expected matrix %v", matrix, tcases.expectedM)
				}
			}
		})
	}
}

func TestRowsM(t *testing.T) {
	matrix, _ := New("1 3 5\n7 9 11\n13 15 17")
	rows := matrix.Rows()
	if len(rows) != 3 {
		t.Errorf("got matrix rows %d, expected 3", len(rows))
	}
	if len(rows[0]) != 3 {
		t.Errorf("got matrix rows number %d, expected 3", len(rows[0]))
	}
	if rows[0][0] != 1 {
		t.Errorf("got first element %d, expected 1", rows[0][0])
	}
}

func TestColsM(t *testing.T) {
	matrix, _ := New("1 3 5\n7 9 11\n13 15 17")
	columns := matrix.Cols()
	if len(columns) != 3 {
		t.Errorf("got matrix columns %d, expected 3", len(columns))
	}
	if len(columns[0]) != 2 {
		t.Errorf("got matrix columns number %d, expected 3", len(columns[0]))
	}
	if columns[0][0] != 1 {
		t.Errorf("got first element %d, expected 1", columns[0][0])
	}
}

func TestSetM(t *testing.T) {
	matrix, _ := New("1 3 5\n7 9 11\n13 15 17")
	var tData = []struct {
		fName    string
		row      int
		column   int
		value    int
		expected bool
	}{
		{
			"negative row",
			-1,
			0,
			100,
			false,
		},
		{
			"negative column",
			0,
			-1,
			100,
			false,
		},
		{
			"row number out is of index",
			5,
			1,
			100,
			false,
		},
		{
			"column number is out of index",
			1,
			5,
			100,
			false,
		},
		{
			"normal case",
			2,
			2,
			100,
			true,
		},
	}
	for _, tcases := range tData {
		t.Run(tcases.fName, func(t *testing.T) {
			changed := matrix.Set(tcases.row, tcases.column, tcases.value)
			switch changed {
			case tcases.expected:
				if matrix.Rows()[tcases.row][tcases.column] != tcases.value {
					t.Errorf("got %d, expected %d", matrix.Rows()[tcases.row][tcases.column], tcases.value)
				}
			default:
				t.Errorf("got %t, expected %t", changed, tcases.expected)
			}
		})
	}
}
