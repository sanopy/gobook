// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

type StringSet map[string]bool

func newStringSet(elms ...string) StringSet {
	s := make(StringSet)
	for _, e := range elms {
		s[e] = true
	}
	return s
}

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]StringSet{
	"algorithms": newStringSet("data structures"),
	"calculus":   newStringSet("linear algebra"),

	"compilers": newStringSet(
		"data structures",
		"formal languages",
		"computer organization",
	),

	"data structures":       newStringSet("discrete math"),
	"databases":             newStringSet("data structures"),
	"discrete math":         newStringSet("intro to programming"),
	"formal languages":      newStringSet("discrete math"),
	"networks":              newStringSet("operating systems"),
	"operating systems":     newStringSet("data structures", "computer organization"),
	"programming languages": newStringSet("data structures", "computer organization"),
}

func main() {
	v := topoSort(prereqs)
	for i, course := range v {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Printf("is sorted: %v\n", isSorted(prereqs, v))
}

func topoSort(g map[string]StringSet) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items StringSet)

	visitAll = func(items StringSet) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(g[item])
				order = append(order, item)
			}
		}
	}

	items := make(StringSet)
	for key := range g {
		items[key] = true
	}

	visitAll(items)
	return order
}

func isSorted(g map[string]StringSet, v []string) bool {
	order := make(map[string]int)
	for i, key := range v {
		order[key] = i
	}

	for _, c1 := range v {
		for c2 := range g[c1] {
			if order[c1] < order[c2] {
				return false
			}
		}
	}

	return true
}
