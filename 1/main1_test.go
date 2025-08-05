package main

import (
	"bytes"
	"fmt"
	"slices"
	"testing"
)

func TestToSingleString(t *testing.T) {
	tests := []struct {
		name string
		vals []any
		want string
	}{
		{
			name: "001",
			vals: []any{052, 42, 0x2A, 3.14, "Golang", true, 1 + 2i},
			want: "4242423.14Golangtrue(1+2i)",
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := toSingleString(test.vals...)
			if got != test.want {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}
}

func TestAddGo2024(t *testing.T) {
	tests := []struct {
		vals []rune
		want []rune
	}{
		{
			vals: []rune("Golang"),
			want: []rune("Golgo-2024ang"),
		},
		{
			vals: []rune("hmm"),
			want: []rune("hgo-2024mm"),
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := addGo2024(test.vals)
			res := slices.Compare(got, test.want)
			if res != 0 {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}
}

func TestGetTypes(t *testing.T) {
	tests := []struct {
		val  any
		want string
	}{
		{
			val:  int(053),
			want: "значение 43 принадлежит к типу int\n",
		},
		{
			val:  int(40),
			want: "значение 40 принадлежит к типу int\n",
		},
		{
			val:  int(0x2C),
			want: "значение 44 принадлежит к типу int\n",
		},
		{
			val:  0.16,
			want: "значение 0.16 принадлежит к типу float64\n",
		},
		{
			val:  "golang",
			want: "значение golang принадлежит к типу string\n",
		},
		{
			val:  true,
			want: "значение true принадлежит к типу bool\n",
		},
		{
			val:  1 + 2i,
			want: "значение (1+2i) принадлежит к типу complex128\n",
		},
	}

	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			buffer := bytes.Buffer{}
			getTypes(&buffer, test.val)
			got := buffer.String()
			if got != test.want {
				t.Errorf("got: %s; want: %s", got, test.want)
			}
		})
	}

}

func TestStringToRunes(t *testing.T) {
	tests := []struct {
		in   string
		want []rune
	}{
		{
			in:   "string",
			want: []rune{115, 116, 114, 105, 110, 103},
		},
		{
			in:   "",
			want: []rune{},
		},
		{
			in:   "go go go",
			want: []rune{103, 111, 32, 103, 111, 32, 103, 111},
		},
	}

	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := stringToRunes(test.in)

			if !slices.Equal(got, test.want) {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}
}

func TestHashRunes(t *testing.T) {
	tests := []struct {
		runes []rune
		want  string
	}{
		{
			runes: []rune{52, 50, 52, 50, 52, 50, 51, 46, 49, 52, 71, 111, 108, 103, 111, 45, 50, 48, 50, 52, 103, 111, 45, 50, 48, 50, 52, 40, 49, 43, 50, 105, 41},
			want:  "eb2cb5f348ea6438ef0bb0931a41f50c673fa02166464b01efb692fbd93a1621",
		},
		{
			runes: []rune{52, 50, 52, 50, 52, 50, 51, 46, 49, 52, 71, 111, 108, 97, 110, 103, 116, 114, 117, 101, 40, 49, 43, 50, 105, 41},
			want:  "2b3a78203af83942344fca645453cc2d4e7c1daa6aaa51114cb43bce89144e6b",
		},
	}

	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			got := hashRunes(test.runes...)
			if got != test.want {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}

}
