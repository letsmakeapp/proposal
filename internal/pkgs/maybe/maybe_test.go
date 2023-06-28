package maybe

import "testing"

func TestMaybe_IsSome(t *testing.T) {
	tests := []struct {
		name     string
		create   func(t *testing.T) T[int]
		expected bool
	}{
		{
			name: "should return true",
			create: func(t *testing.T) T[int] {
				return Some(1)
			},
			expected: true,
		},
		{
			name: "should return false",
			create: func(t *testing.T) T[int] {
				return None[int]()
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.create(t)
			got := m.IsSome()
			if got != tt.expected {
				t.Errorf("expect IsSome() = %v but got %v", tt.expected, got)
			}
		})
	}
}

func TestMaybe_IsNone(t *testing.T) {
	tests := []struct {
		name     string
		create   func(t *testing.T) T[int]
		expected bool
	}{
		{
			name: "should return true",
			create: func(t *testing.T) T[int] {
				return Some(1)
			},
			expected: false,
		},
		{
			name: "should return false",
			create: func(t *testing.T) T[int] {
				return None[int]()
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.create(t)
			got := m.IsNone()
			if got != tt.expected {
				t.Errorf("expect IsNone() = %v but got %v", tt.expected, got)
			}
		})
	}
}

func TestMaybe_TryGetValue(t *testing.T) {
	tests := []struct {
		name   string
		create func(t *testing.T) T[int]
		value  int
		ok     bool
	}{
		{
			name: "should return true",
			create: func(t *testing.T) T[int] {
				return Some(1)
			},
			value: 1,
			ok:    true,
		},
		{
			name: "should return false",
			create: func(t *testing.T) T[int] {
				return None[int]()
			},
			ok: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.create(t)
			value, ok := m.TryGetValue()

			if ok != tt.ok {
				t.Errorf("expect ok = %v but got %v", tt.ok, ok)
			}

			if value != tt.value {
				t.Errorf("expect value = %v but got %v", tt.value, value)
			}
		})
	}
}
