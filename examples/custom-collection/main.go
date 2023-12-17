package main

import (
	"fmt"

	ds "github.com/rubensdev/go-snippets/data-structures"
)

func main() {
	c := ds.MyCollection{Data: []int{1, 2, 3}}
	for it := c.Iterate(); it.Next(); {
		fmt.Println(it.Value())
	}
}
