package maybe

func FromPointer[V any, VP *V /* only accept pointer */](v VP) T[V] {
	if v != nil {
		return Some(*v)
	}
	return None[V]()
}
