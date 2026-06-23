package goquery

func PartitionOwn[T any](items []T, funcion func(T) bool) (yes []T, no []T) {
	// Creacion de 2 slices vacios
	yes = make([]T, 0, len(items))
	no = make([]T, 0, len(items))
	for _, item := range items {
		if funcion(item) {
			yes = append(yes, item)
		} else {
			no = append(no, item)
		}
	}
	return yes, no
}
