package first

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var integer8 int8 = 052
	var integer10 int64 = 42
	var integer16 int16 = 0x2A
	var real float64 = 3.14
	var str string = "Golang"
	var boolean bool = true
	var complex complex128 = 1 + 2i
	vals := []any{integer8, integer10, integer16, real, str, boolean, complex}

	fmt.Print(currentAssignment(os.Stdout, vals...))
}

func currentAssignment(w io.Writer, vals ...any) string {
	getTypes(w, vals...)
	runes := []rune(toSingleString(vals...))
	println(string(runes))
	runes = addGo2024(runes)
	println(string(runes))

	return hashRunes(runes...)
}

func getTypes(w io.Writer, vals ...any) {
	for _, val := range vals {
		fmt.Fprintf(w, "значение %v принадлежит к типу %s\n", val, reflect.TypeOf(val))
	}
}

func toSingleString(vals ...any) string {
	var resString string
	for _, val := range vals {
		stringToAdd := fmt.Sprintf("%v", val)
		resString += stringToAdd[1:]
	}
	fmt.Printf("%s", resString)
	return resString
}

func addGo2024(vals []rune) []rune {
	toAdd := []rune("go-2024")
	length := len(vals)
	res := append(append(vals[:length/2], toAdd...), vals[length/2:]...)

	return res
}

func hashRunes(runes ...rune) string {
	hasher := sha256.New()
	hasher.Write([]byte(string(runes)))
	return hex.EncodeToString(hasher.Sum(nil))
}
