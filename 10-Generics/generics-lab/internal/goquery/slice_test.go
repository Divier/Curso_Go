package goquery

import "testing"

func TestFilterOwn(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5, 6}
	got := Filter(numbers, func(n int) bool { return n%2 == 0 })
	want := []int{2, 4, 6}

	if len(got) != len(want) {
		t.Fatalf("got %d even numbers, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%v, want=%v", i, got[i], want[i])
		}
	}
}

type UserOwn struct {
	Id   int
	Name string
}

func TestMapToNamesOwn(t *testing.T) {
	users := []UserOwn{{Id: 1, Name: "Alice"}, {Id: 2, Name: "Bob"}, {Id: 3, Name: "Charlie"}}
	got := Map(users, func(u UserOwn) string { return u.Name })
	want := []string{"Alice", "Bob", "Charlie"}

	if len(got) != len(want) {
		t.Fatalf("got %d names, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%v, want=%v", i, got[i], want[i])
		}
	}
}
