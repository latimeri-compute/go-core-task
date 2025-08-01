package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		initial StringIntMap
		key     string
		value   int
		want    StringIntMap
	}{
		{
			initial: StringIntMap{map[string]int{"a": 34, "hehee": 90}},
			key:     "5",
			value:   3231,
			want:    StringIntMap{map[string]int{"a": 34, "hehee": 90, "5": 3231}},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			test.initial.Add(test.key, test.value)
			EqualMap(t, test.initial, test.want)
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		initial StringIntMap
		key     string
		want    StringIntMap
	}{
		{
			name:    "удаление из карты существующей пары",
			initial: StringIntMap{map[string]int{"a": 34, "hehee": 90}},
			key:     "a",
			want:    StringIntMap{map[string]int{"hehee": 90}},
		},
		{
			name:    "удаление из карты несуществующей пары",
			initial: StringIntMap{map[string]int{"a": 34, "hehee": 90}},
			key:     "no",
			want:    StringIntMap{map[string]int{"a": 34, "hehee": 90}},
		},
		{
			name:    "удаление из пустой карты несуществующей пары",
			initial: StringIntMap{map[string]int{}},
			key:     "no",
			want:    StringIntMap{map[string]int{}},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			test.initial.Remove(test.key)
			EqualMap(t, test.initial, test.want)
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name    string
		initial StringIntMap
		key     string
		want    bool
	}{
		{
			name:    "существует",
			initial: StringIntMap{map[string]int{"a": 34, "hehee": 90}},
			key:     "a",
			want:    true,
		},
		{
			name:    "не существует",
			initial: StringIntMap{map[string]int{"a": 34, "hehee": 90}},
			key:     "noooo",
			want:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.initial.Exists(test.key)
			if got != test.want {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}
}

func EqualMap(t *testing.T, m1, m2 StringIntMap) {
	t.Helper()
	if !reflect.DeepEqual(m1, m2) {
		t.Errorf("got: %v; want: %v", m1, m2)
	}
}
