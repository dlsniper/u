package main

import "sort"


type demo struct{}

func (demo) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
}
