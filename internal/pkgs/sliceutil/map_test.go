package sliceutil

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMap_Empty(t *testing.T) {
	var (
		source   = []int{}
		expected = []string{}
	)

	got := Map(source, func(x int) string {
		return strconv.FormatInt(int64(x), 10)
	})

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("Map() mismatch (-want, +got):\n%s", diff)
	}
}

func TestMap_NotEmpty(t *testing.T) {
	var (
		source   = []int{1, -1, 2, -2, 3, -3}
		expected = []string{"1", "-1", "2", "-2", "3", "-3"}
	)

	got := Map(source, func(x int) string {
		return strconv.FormatInt(int64(x), 10)
	})

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("Map() mismatch (-want, +got):\n%s", diff)
	}
}

func TestMapWithIndex_Empty(t *testing.T) {
	var (
		source   = []int{}
		expected = []string{}
	)

	got := MapWithIndex(source, func(index uint, x int) string {
		return fmt.Sprintf("%d:%d", index, x)
	})

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("MapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestMapWithIndex_NotEmpty(t *testing.T) {
	var (
		source   = []int{1, -1, 2, -2, 3, -3}
		expected = []string{"0:1", "1:-1", "2:2", "3:-2", "4:3", "5:-3"}
	)

	got := MapWithIndex(source, func(index uint, x int) string {
		return fmt.Sprintf("%d:%d", index, x)
	})

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("MapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestTryMap_Empty(t *testing.T) {
	var (
		source   = []int{}
		expected = []string{}
	)

	got, err := TryMap(source, func(x int) (string, error) {
		return fmt.Sprintf("%d", x), nil
	})

	if err != nil {
		t.Errorf("unexpected error = %v", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("TryMapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestTryMap_NotEmpty(t *testing.T) {
	var (
		source   = []int{1, -1, 2, -2, 3, -3}
		expected = []string{"1", "-1", "2", "-2", "3", "-3"}
	)

	got, err := TryMap(source, func(x int) (string, error) {
		return fmt.Sprintf("%d", x), nil
	})

	if err != nil {
		t.Errorf("unexpected error = %v", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("TryMapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestTryMap_Error(t *testing.T) {
	var (
		source         = []int{1, -1, 2, -2, 3, -3}
		expected       = ([]string)(nil)
		unhandledError = errors.New("unhandled error")
	)

	got, err := TryMap(source, func(x int) (string, error) {
		if x == 2 {
			return "", unhandledError
		}
		return fmt.Sprintf("%d", x), nil
	})

	if err != unhandledError {
		t.Errorf("expected unhandled error, got = %v", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("TryMapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestTryMapWithIndex_Empty(t *testing.T) {
	var (
		source   = []int{}
		expected = []string{}
	)

	got, err := TryMapWithIndex(source, func(index uint, x int) (string, error) {
		return fmt.Sprintf("%d:%d", index, x), nil
	})

	if err != nil {
		t.Errorf("unexpected error = %v", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("TryMapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestTryMapWithIndex_NotEmpty(t *testing.T) {
	var (
		source   = []int{1, -1, 2, -2, 3, -3}
		expected = []string{"0:1", "1:-1", "2:2", "3:-2", "4:3", "5:-3"}
	)

	got, err := TryMapWithIndex(source, func(index uint, x int) (string, error) {
		return fmt.Sprintf("%d:%d", index, x), nil
	})

	if err != nil {
		t.Errorf("unexpected error = %v", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("TryMapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestTryMapWithIndex_Error(t *testing.T) {
	var (
		source         = []int{1, -1, 2, -2, 3, -3}
		expected       = ([]string)(nil)
		unhandledError = errors.New("unhandled error")
	)

	got, err := TryMapWithIndex(source, func(index uint, x int) (string, error) {
		if index == 2 {
			return "", unhandledError
		}
		return fmt.Sprintf("%d:%d", index, x), nil
	})

	if err != unhandledError {
		t.Errorf("expected unhandled error, got = %v", err)
	}

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("TryMapWithIndex() mismatch (-want, +got):\n%s", diff)
	}
}

func TestListIntToString(t *testing.T) {
	var (
		source   = []int64{1, -1, 2, -2, 3, -3}
		expected = "[1,-1,2,-2,3,-3]"
	)

	got := ListIntToString(source)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("ListIntToString() mismatch (-want, +got):\n%s", diff)
	}
}
