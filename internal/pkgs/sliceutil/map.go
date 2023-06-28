package sliceutil

import (
	"fmt"
	"strings"
)

func Map[T any, U any](ts []T, f func(t T) U) []U {
	return MapWithIndex(ts, func(_ uint, t T) U {
		return f(t)
	})
}

func MapWithIndex[T any, U any](ts []T, f func(index uint, t T) U) []U {
	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(uint(i), t)
	}
	return us
}

func TryMap[T any, U any](ts []T, f func(t T) (U, error)) ([]U, error) {
	return TryMapWithIndex(ts, func(_ uint, t T) (U, error) {
		return f(t)
	})
}

func TryMapWithIndex[T any, U any](ts []T, f func(index uint, t T) (U, error)) ([]U, error) {
	us := make([]U, len(ts))
	for i, t := range ts {
		u, err := f(uint(i), t)
		if err != nil {
			return nil, err
		}
		us[i] = u
	}
	return us, nil
}

func ListIntToString(ts []int64) string {
	strSlice := make([]string, len(ts))
	for i, num := range ts {
		strSlice[i] = fmt.Sprintf("%d", num)
	}
	return "[" + strings.Join(strSlice, ",") + "]"
}
