package first

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"slices"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name      string
		integer8  int8
		integer10 int64
		integer16 int16
		real      float64
		str       string
		boolean   bool
		complex   complex64
	}{
		{
			name:      "001",
			integer8:  052,
			integer10: 42,
			integer16: 0x2A,
			real:      3.14,
			str:       "Golang",
			boolean:   true,
			complex:   1 + 2i,
		},
		{
			name:      "002",
			integer8:  053,
			integer10: 40,
			integer16: 0x2C,
			real:      0.16,
			str:       "evil golang",
			boolean:   true,
			complex:   1 + 2i,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var input []any = []any{test.integer8, test.integer10, test.integer16, test.real, test.str, test.boolean, test.complex}

			hasher := sha256.New()
			fmt.Printf("%v", input...)
			_, err := hasher.Write([]byte(fmt.Sprintf("%v", input...)))
			if err != nil {
				t.Fatal(err)
			}
			got := currentAssignment(io.Discard, input...)
			want := hex.EncodeToString(hasher.Sum(nil))
			if got != want {
				t.Errorf("got: %v; want: %v", got, want)
			}
		})
	}

}

func TestToSingleString(t *testing.T) {
	tests := []struct {
		name string
		vals any
		want string
	}{
		{
			name: "001",
			vals: []any{052, 42, 0x2A, 3.14, "Golang", true, 1 + 2i},
			want: "4242423.14Golangtrue(1+2i)",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := toSingleString(test.vals)
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
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
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
			val:  int8(053),
			want: "значение 053 принадлежит к типу int8\n",
		},
		{
			val:  int64(40),
			want: "значение 40 принадлежит к типу int64\n",
		},
		{
			val:  int16(0x2C),
			want: "значение 0x2C принадлежит к типу int16\n",
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

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			buffer := bytes.Buffer{}
			getTypes(&buffer, test.val)
			got := buffer.String()
			if got != test.want {
				t.Errorf("got: %s; want: %s", got, test.want)
			}
		})
	}

}
