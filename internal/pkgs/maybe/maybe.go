package maybe

import "errors"

var (
	ErrUnwrapNoneValue = errors.New("unwrap none value")
)

type T[V any] struct {
	set   bool
	value V
}

func Some[V any](v V) T[V] {
	return T[V]{
		set:   true,
		value: v,
	}
}

func None[V any]() T[V] {
	return T[V]{
		set: false,
	}
}

func (op T[V]) IsSome() bool {
	return op.set
}

func (op T[V]) IsNone() bool {
	return !op.set
}

func (op T[V]) TryGetValue() (V, bool) {
	if op.set {
		return op.value, true
	}
	var zero V
	return zero, false
}

func (op T[V]) UnsafeUnwrap() V {
	if !op.set {
		panic(ErrUnwrapNoneValue)
	}
	return op.value
}

func (op T[V]) ToPointer() *V {
	if op.set {
		return &op.value
	}
	return nil
}
