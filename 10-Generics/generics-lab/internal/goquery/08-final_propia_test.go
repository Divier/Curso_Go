package goquery

import "testing"

func TestPartitionOwn(test *testing.T) {
	// Pases un slice de numeros del 1 al 7
	// Obtener los 2 slices ya separados
	// Comparar los tamaños de cada slice que cumplan la condicion
	// Dividir 2 slices en numeros pares e impares
	numeros := []int{1, 2, 3, 4, 5, 6, 7}
	pares, impares := PartitionOwn(numeros, func(n int) bool {
		return n%2 == 0
	})
	if len(pares) != 3 {
		test.Fatalf("Se esperaban 3 numeros pares, pero se obtuvieron %d", len(pares))
	}
	if len(impares) != 4 {
		test.Fatalf("Se esperaban 4 numeros pares, pero se obtuvieron %d", len(impares))
	}
}
