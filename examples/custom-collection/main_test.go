package main

import (
	"testing"

	ds "github.com/rubensdev/go-snippets/data-structures"
)

func TestCollectionIteration(t *testing.T) {
	want := 6

	c := ds.MyCollection{Data: []int{1, 2, 3}}

	got := 0
	for it := c.Iterate(); it.Next(); {
		got += it.Value()
	}

	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
