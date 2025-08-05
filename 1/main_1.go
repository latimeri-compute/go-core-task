package main

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

	getTypes(os.Stdout, vals...)
	oneString := toSingleString(vals...)
	fmt.Println(oneString)
	runes := stringToRunes(oneString)
	fmt.Println(runes)
	runes = addGo2024(runes)
	fmt.Println(runes)

	fmt.Println(hashRunes(runes...))
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
		resString += stringToAdd
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

func stringToRunes(strings ...string) []rune {
	var runes []rune
	for _, s := range strings {
		runes = append(runes, []rune(s)...)
	}
	return runes
}

func hashRunes(runes ...rune) string {
	hasher := sha256.New()
	hasher.Write([]byte(string(runes)))
	return hex.EncodeToString(hasher.Sum(nil))
}
